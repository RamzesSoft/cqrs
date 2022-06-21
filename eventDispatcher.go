package cqrs

// EventDispatcher is the interface that an event bus must implement.
type EventDispatcher interface {
	PublishEvent(EventMessage)
	AddHandler(EventHandler, ...interface{})
}
