package cqrs_test

import (
	"testing"

	"github.com/google/uuid"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/RamzesSoft/cqrs"
	"github.com/RamzesSoft/cqrs/inMemory"
)

func TestNewHandlerEventDispatcher(t *testing.T) {
	Convey("Test InMemoryEventDispatcher", t, func() {
		bus := inMemory.NewEventDispatcher()
		Convey("EventDispatcher should not be nil", func() {
			So(bus, ShouldNotBeNil)
		})
	})
}

func TestEventBusPublishesEventsToHandlers(t *testing.T) {
	Convey("Test Publish One Event", t, func() {
		bus := inMemory.NewEventDispatcher()
		h := NewMockEventHandler()
		ev := NewTestEventMessage(uuid.NewString())
		bus.AddHandler(h, &SomeEvent{})
		bus.PublishEvent(ev)
		e := h.Events()
		So(e, ShouldResemble, []cqrs.EventMessage{ev})
	})
}

func TestRegisterMultipleEventsForHandler(t *testing.T) {
	Convey("Test Publish Multiply Event", t, func() {
		bus := inMemory.NewEventDispatcher()
		h := NewMockEventHandler()
		ev1 := cqrs.NewEvent(uuid.NewString(), &SomeEvent{Item: "Some Item", Count: 3456})
		ev2 := cqrs.NewEvent(uuid.NewString(), &SomeOtherEvent{OrderID: uuid.NewString()})

		bus.AddHandler(h, &SomeEvent{}, &SomeOtherEvent{})

		bus.PublishEvent(ev1)
		bus.PublishEvent(ev2)

		So(h.Events(), ShouldResemble, []cqrs.EventMessage{ev1, ev2})
	})
}

// Stubs
type MockEventBus struct {
	events []cqrs.EventMessage
}

func (m *MockEventBus) PublishEvent(event cqrs.EventMessage) {
	m.events = append(m.events, event)
}

func (m *MockEventBus) AddHandler(handler cqrs.EventHandler, event ...interface{}) {}
func (m *MockEventBus) AddLocalHandler(handler cqrs.EventHandler)                  {}
func (m *MockEventBus) AddGlobalHandler(handler cqrs.EventHandler)                 {}
