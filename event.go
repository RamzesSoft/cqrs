package cqrs

// EventMessage is the interface that a command must implement.
type EventMessage interface {
	// AggregateID returns the ID of the Aggregate that the event relates to
	AggregateID() string
	//Event Returns the actual event which is the payload of the event message.
	Event() interface{}
	// EventType returns a string descriptor of the command name
	EventType() string
}

// EventDescriptor is an implementation of the event message interface.
type EventDescriptor struct {
	id    string
	event interface{}
}

// NewEvent returns a new event descriptor
func NewEvent(aggregateID string, event interface{}) *EventDescriptor {
	return &EventDescriptor{
		id:    aggregateID,
		event: event,
	}
}

// EventType returns the name of the event type as a string.
func (c *EventDescriptor) EventType() string {
	return TypeWithPackage(c.event)
}

// AggregateID returns the ID of the Aggregate that the event relates to.
func (c *EventDescriptor) AggregateID() string {
	return c.id
}

// Event the event payload of the event message
func (c *EventDescriptor) Event() interface{} {
	return c.event
}
