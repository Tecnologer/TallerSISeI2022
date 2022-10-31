package messenger

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
)

type messenger struct {
	services         map[string]*service
	incomingMessages chan Message
}

type service struct {
	m         Messenger
	name      string
	isOpen    bool
	sleepTime time.Duration //time in seconds to wait for request new messages
}

func NewMessenger() *messenger {
	return &messenger{
		services: make(map[string]*service),
	}
}

func (m *messenger) Broadcast(msg string) (errs []error) {
	errs = make([]error, 0)
	for name := range m.services {
		e := m.SendMessage(name, "", msg)
		if e != nil {
			errs = append(errs, errors.Wrapf(e, "messenger.broadcast: service %s", name))
		}
	}

	fmt.Println("message broadcasted")

	return
}

func (m *messenger) BroadcastToOtherServices(service, msg string) (errs []error) {
	errs = make([]error, 0)
	for name := range m.services {
		if name == service {
			continue
		}

		e := m.SendMessage(name, "", msg)
		if e != nil {
			errs = append(errs, errors.Wrapf(e, "messenger.broadcast_to_other_services: service %s", name))
		}
	}

	fmt.Printf("message send to the services, except: %s\n", service)

	return
}

func (m *messenger) SendMessage(serviceName string, dst string, msg string) error {
	s, exists := m.services[serviceName]
	if !exists {
		return fmt.Errorf("the service %s doesn't exists", serviceName)
	}

	err := s.m.SendMessage(dst, msg)
	if err != nil {
		return errors.Wrap(err, "messenger.send_message")
	}

	return nil
}

func (m *messenger) AddService(name string, messenger Messenger) error {
	_, e := messenger.GetMe()
	if e != nil {
		return errors.Wrap(e, "messenger.add_service: test messenger connection")
	}

	m.services[name] = &service{
		name:      name,
		m:         messenger,
		isOpen:    false,
		sleepTime: time.Duration(5),
	}

	return nil
}

func (m *messenger) RegisterForAllMessages() (chan Message, error) {
	if m.incomingMessages != nil {
		return nil, fmt.Errorf("the service to recive messages is running")
	}

	m.incomingMessages = make(chan Message)
	for name := range m.services {
		go func(name string) {
			incomeMsg, err := m.RegisterForMessages(name)
			if err != nil {
				log.Println(err)
				return
			}

			for msg := range incomeMsg {
				m.incomingMessages <- msg
			}
		}(name)
	}

	return m.incomingMessages, nil
}

func (m *messenger) UnregisterForAllMessages() {
	if m.incomingMessages == nil {
		log.Println("there is not service to recieve messages")
		return
	}

	for name := range m.services {
		_ = m.UnregisterForMessages(name)
	}

	close(m.incomingMessages)
	m.incomingMessages = nil
}

func (m *messenger) RegisterForMessages(name string) (chan Message, error) {
	s, exists := m.services[name]
	if !exists {
		return nil, fmt.Errorf("the service %s doesn't exists", name)
	}

	messages := make(chan Message)
	s.isOpen = true

	go func(s *service, c chan<- Message) {
		for s.isOpen {
			messages, err := s.m.GetMessages()

			if err != nil {
				log.Printf("error getting messages from service %s. Err: %v", s.name, err)
				continue
			}

			for _, message := range messages {
				message.SetService(s.name)
				c <- message
			}
			time.Sleep(s.sleepTime * time.Second)
		}

		close(messages)
	}(s, messages)

	return messages, nil
}

func (m *messenger) UnregisterForMessages(name string) error {
	s, exists := m.services[name]
	if !exists {
		return fmt.Errorf("the service %s doesn't exists", name)
	}

	s.isOpen = false

	return nil
}

type Messenger interface {
	GetMessages() ([]Message, error)
	SendMessage(dest, msg string) error
	GetMe() (User, error)
}

type Message interface {
	SetService(string)
	GetService() string
	GetMsg() string
	From() User
}

type User interface {
	GetName() string
	GetID() string
	IsBot() bool
}
