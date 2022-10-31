package telegram

import (
	"fmt"

	"github.com/tecnologer/SISeI/messenger"
)

type Update struct {
	service  string   `json:"-"`
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type Message struct {
	MessageID  int    `json:"message_id"`
	From       *User  `json:"from"`
	SenderChat *Chat  `json:"sender_chat"`
	Date       int64  `json:"date"`
	Chat       *Chat  `json:"chat"`
	Text       string `json:"text"`
}

type User struct {
	ID        int    `json:"id"`
	IsBotFlag bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *Update) GetMsg() string {
	if u == nil || u.Message == nil {
		return ""
	}

	return u.Message.Text
}

func (u *Update) From() messenger.User {
	if u == nil || u.Message == nil {
		return nil
	}

	return u.Message.From
}

func (u *Update) GetService() string {
	if u == nil {
		return ""
	}

	return u.service
}

func (u *Update) SetService(s string) {
	if u == nil {
		return
	}

	u.service = s
}

func (c *User) GetID() string {
	if c == nil {
		return ""
	}

	return fmt.Sprint(c.ID)
}

func (c *User) GetName() string {
	if c == nil {
		return ""
	}

	return c.Username
}

func (u *User) IsBot() bool {
	if u == nil {
		return false
	}

	return u.IsBotFlag
}
