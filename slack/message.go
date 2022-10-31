package slack

import (
	"encoding/json"
	"strconv"

	"github.com/tecnologer/SISeI/messenger"
)

type messages struct {
	OK       bool       `json:"ok"`
	Latest   timestamp  `json:"latest"`
	Messages []*message `json:"messages"`
	Error    string     `json:"error,omitempty"`
}

type message struct {
	service   string    `json:"-"`
	Text      string    `json:"text"`
	User      string    `json:"user"`
	UserID    string    `json:"user_id"`
	BotID     string    `json:"bot_id"`
	Timestamp timestamp `json:"ts"`
}

func (m *message) GetMsg() string {
	if m == nil {
		return ""
	}

	return m.Text
}

func (m *message) GetService() string {
	if m == nil {
		return ""
	}

	return m.service
}

func (m *message) SetService(s string) {
	if m == nil {
		return
	}

	m.service = s
}

func (m *message) From() messenger.User {
	if m == nil {
		return nil
	}
	u := &User{
		ID:    m.UserID,
		Name:  m.User,
		BotID: m.BotID,
	}

	if u.ID == "" {
		u.ID = m.User
	}

	if m.BotID != "" {
		u.Name = m.BotID
	}

	return u
}

type User struct {
	ID    string `json:"user_id"`
	Name  string `json:"user"`
	BotID string `json:"bot_id"`
}

func (u *User) GetID() string {
	if u == nil {
		return ""
	}

	if u.BotID != "" {
		return u.BotID
	}

	return u.ID
}

func (u *User) GetName() string {
	if u == nil {
		return ""
	}

	if u.Name != "" {
		return u.Name
	}

	return u.ID
}

func (u *User) IsBot() bool {
	if u == nil {
		return false
	}

	return u.BotID != ""
}

type timestamp float64

func (t *timestamp) UnmarshalJSON(v []byte) error {
	var str string
	err := json.Unmarshal(v, &str)
	if err != nil {
		return err
	}

	fv, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*t = timestamp(fv)
	return nil
}
