package cqrs

type EventHandler interface {
	Handle(EventMessage)
}
