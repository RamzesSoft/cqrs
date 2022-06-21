package inMemory

import "github.com/RamzesSoft/cqrs"

// eventDispatcher provides a lightweight in process event bus
type eventDispatcher struct {
	eventHandlers map[string]map[cqrs.EventHandler]struct{}
}

// NewEventDispatcher constructs a new eventDispatcher
func NewEventDispatcher() cqrs.EventDispatcher {
	b := &eventDispatcher{
		eventHandlers: make(map[string]map[cqrs.EventHandler]struct{}),
	}
	return b
}

// PublishEvent publishes event to all registered event handlers
func (b *eventDispatcher) PublishEvent(event cqrs.EventMessage) {
	if handlers, ok := b.eventHandlers[event.EventType()]; ok {
		for handler := range handlers {
			handler.Handle(event)
		}
	}
}

// AddHandler registers an event handler for all the event specified in the
// variadic event parameter.
func (b *eventDispatcher) AddHandler(handler cqrs.EventHandler, events ...interface{}) {

	for _, event := range events {
		typeName := cqrs.TypeWithPackage(event)

		// There can be multiple handlers for any event.
		// Here we check that a map is initialized to hold these handlers
		// for a given type. If not we create one.
		if _, ok := b.eventHandlers[typeName]; !ok {
			b.eventHandlers[typeName] = make(map[cqrs.EventHandler]struct{})
		}

		// Add this handler to the collection of handlers for the type.
		b.eventHandlers[typeName][handler] = struct{}{}
	}
}
