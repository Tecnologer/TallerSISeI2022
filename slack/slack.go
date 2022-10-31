package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/tecnologer/SISeI/messenger"
)

const urlBase = "https://slack.com/api/%s"

type Slack struct {
	myID        string
	token       string
	defaultDest string
	oldest      timestamp
}

func New(token string) *Slack {
	return &Slack{
		token:       token,
		defaultDest: "D0458TPQQ07",
		oldest:      1665344843.163739,
	}
}

func (t *Slack) GetMe() (messenger.User, error) {
	var user *User
	err := t.postRequest("auth.test", nil, &user)
	if err != nil {
		return nil, errors.Wrap(err, "slack.get_me")
	}

	t.myID = user.GetID()

	return user, nil
}

func (s *Slack) GetMessages() ([]messenger.Message, error) {
	//conversations.list
	var messages *messages

	params := map[string]interface{}{
		"channel": s.defaultDest,
	}

	if s.oldest > 0 {
		params["oldest"] = fmt.Sprintf("%f", s.oldest)
	}

	err := s.getRequest("conversations.history", params, &messages)
	if err != nil {
		return nil, errors.Wrap(err, "slack.get_messages")
	}

	for _, msg := range messages.Messages {
		if msg.Timestamp > s.oldest {
			s.oldest = msg.Timestamp
		}
	}

	return parseMessages(messages.Messages...), nil
}

func (t *Slack) SendMessage(dest, msg string) error {
	if dest == "" {
		dest = fmt.Sprint(t.defaultDest)
	}

	params := map[string]interface{}{
		"channel": dest,
		"text":    msg,
		"pretty":  "1",
	}

	var result map[string]interface{}
	err := t.postRequest("chat.postMessage", params, &result)
	if err != nil {
		return errors.Wrap(err, "slack.send_message")
	}

	// if !reflect.DeepEqual(result["ok"], interface{}(true)) {
	if result["ok"] != interface{}(true) {
		return errors.New(result["error"].(string))
	}

	return nil
}

func (t *Slack) getUrl(method string) string {
	return fmt.Sprintf(urlBase, method)
}

func (s *Slack) postRequest(action string, params map[string]interface{}, resDst interface{}) error {
	urlBase, err := url.Parse(s.getUrl(action))
	if err != nil {
		return errors.Wrap(err, "slack.post_request: build url")
	}

	var payload io.ReadWriter
	if len(params) > 0 {
		payload = &bytes.Buffer{}

		body, err := json.Marshal(params)
		if err != nil {
			return errors.Wrap(err, "slack.post_request: json marshal body")
		}

		_, err = payload.Write(body)
		if err != nil {
			return errors.Wrap(err, "slack.post_request: write json payload")
		}
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, urlBase.String(), payload)
	if err != nil {
		return errors.Wrap(err, "slack.post_request: new request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "slack.post_request: do request")
	}
	defer res.Body.Close()

	// result := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&resDst)
	if err != nil {
		return errors.Wrap(err, "slack.post_request: parse response body")
	}

	return nil
}

func (s *Slack) getRequest(action string, params map[string]interface{}, resDst interface{}) error {
	urlBase, err := url.Parse(s.getUrl(action))
	if err != nil {
		return errors.Wrap(err, "slack.get_request: build url")
	}

	method := "GET"

	values := urlBase.Query()

	for param, value := range params {
		values.Add(param, fmt.Sprint(value))
	}

	urlBase.RawQuery = values.Encode()

	client := &http.Client{}

	req, err := http.NewRequest(method, urlBase.String(), nil)
	if err != nil {
		return errors.Wrap(err, "slack.get_request: new request")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "slack.get_request: do request")
	}
	defer res.Body.Close()

	// var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resDst)
	if err != nil {
		return errors.Wrap(err, "slack.get_request: parse response body")
	}

	return nil
}

func parseMessages(inMessages ...*message) []messenger.Message {
	var (
		messages = make([]messenger.Message, len(inMessages))
	)

	for i, m := range inMessages {
		messages[i] = m
	}

	return messages
}
