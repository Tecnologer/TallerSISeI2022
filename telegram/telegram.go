package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/tecnologer/SISeI/messenger"
)

const urlBase string = "https://api.telegram.org/bot%s/%s"

type Telegram struct {
	token       string
	offset      int
	defaultDest int
}

func New(token string) *Telegram {
	return &Telegram{
		token:       token,
		defaultDest: 10244644,
	}
}

func (t *Telegram) GetMe() (messenger.User, error) {
	var user *User
	err := t.sendRequest("getMe", nil, &user)
	if err != nil {
		return nil, errors.Wrap(err, "telegram.get_me")
	}

	return user, nil
}

func (t *Telegram) GetMessages() ([]messenger.Message, error) {
	params := map[string]interface{}{
		"offset": t.offset + 1,
	}

	var result []*Update

	err := t.sendRequest("getUpdates", params, &result)
	if err != nil {
		return nil, errors.Wrap(err, "telegram.get_messages")
	}

	for _, update := range result {
		if t.defaultDest == 0 && update.Message != nil {
			t.defaultDest = update.Message.Chat.ID
		}

		if update.UpdateID <= t.offset {
			continue
		}

		t.offset = update.UpdateID
	}

	return parseMessages(result...), nil
}

func (t *Telegram) SendMessage(dest, msg string) error {
	if dest == "" {
		dest = fmt.Sprint(t.defaultDest)
	}

	params := map[string]interface{}{
		"chat_id": dest,
		"text":    msg,
	}

	var update *Message

	err := t.sendRequest("sendMessage", params, &update)
	if err != nil {
		return errors.Wrap(err, "telegram.send_message")
	}

	return nil
}

func (t *Telegram) getUrl(values ...interface{}) string {
	values = append([]interface{}{t.token}, values...)
	return fmt.Sprintf(urlBase, values...)
}

func (t *Telegram) sendRequest(method string, params map[string]interface{}, resDst interface{}) error {
	data := &response{
		Result: resDst,
	}

	urlReq, err := url.Parse(t.getUrl(method))
	if err != nil {
		return errors.Wrap(err, "telegram.send_request: build url")
	}

	values := urlReq.Query()
	for param, value := range params {
		values.Add(param, fmt.Sprint(value))
	}

	urlReq.RawQuery = values.Encode()

	response, err := http.Get(urlReq.String())
	if err != nil {
		return errors.Wrap(err, "telegram.send_request: http get")
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return errors.Wrap(err, "telegram.send_request: decode response body")
	}

	return nil
}

func parseMessages(inMessages ...*Update) []messenger.Message {
	var (
		messages = make([]messenger.Message, len(inMessages))
	)

	for i, m := range inMessages {
		messages[i] = m
	}

	return messages
}
