package cqrs_test

import (
	"github.com/RamzesSoft/cqrs"
)

func NewMockEventHandler() *MockEventHandler {
	return &MockEventHandler{
		make([]cqrs.EventMessage, 0),
	}
}

type MockEventHandler struct {
	events []cqrs.EventMessage
}

func (m *MockEventHandler) Handle(event cqrs.EventMessage) {
	m.events = append(m.events, event)
}

func (m *MockEventHandler) Events() []cqrs.EventMessage {
	return m.events
}
