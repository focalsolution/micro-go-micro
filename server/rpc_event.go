package server

import (
	"github.com/focalsolution/micro-go-micro/broker"
	"github.com/focalsolution/micro-go-micro/transport"
)

// event is a broker event we handle on the server transport
type event struct {
	message *broker.Message
}

func (e *event) Ack() error {
	// there is no ack support
	return nil
}

func (e *event) Message() *broker.Message {
	return e.message
}

func (e *event) Topic() string {
	return e.message.Header["Micro-Topic"]
}

func newEvent(msg transport.Message) *event {
	return &event{
		message: &broker.Message{
			Header: msg.Header,
			Body:   msg.Body,
		},
	}
}
