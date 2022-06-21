package cqrs

//AggregateRoot is the interface that all aggregates should implement
type AggregateRoot interface {
	AggregateID() string
	OriginalVersion() int
	CurrentVersion() int
	RecordEvent(EventMessage)
	ReleaseEvents() []EventMessage
}

// AggregateBase is a type that can be embedded in an AggregateRoot
// implementation to handle common aggregate behaviour
//
// All required methods to implement an aggregate are here, to implement the
// Aggregate root interface your aggregate will need to implement the Apply
// method that will contain behaviour specific to your aggregate.
type AggregateBase struct {
	id      string
	version int
	changes []EventMessage
}

// NewAggregate constructs a new AggregateBase.
func NewAggregate(id string) AggregateRoot {
	return &AggregateBase{
		id:      id,
		changes: []EventMessage{},
		version: -1,
	}
}

// AggregateID returns the AggregateID
func (a *AggregateBase) AggregateID() string {
	return a.id
}

// OriginalVersion returns the version of the aggregate as it was when it was
// instantiated or loaded from the repository.
func (a *AggregateBase) OriginalVersion() int {
	return a.version
}

// CurrentVersion returns the version of the aggregate as it was when it was
// instantiated or loaded from the repository.
func (a *AggregateBase) CurrentVersion() int {
	return a.version + len(a.changes)
}

// RecordEvent stores the EventMessage in the change's collection.
//
// Changes are new, not saved event that have been applied to the aggregate.
func (a *AggregateBase) RecordEvent(event EventMessage) {
	a.changes = append(a.changes, event)
}

// ReleaseEvents returns the collection of new not saved event
//
// clear the collection of event
// and make OriginalVersion equal to CurrentVersion
func (a *AggregateBase) ReleaseEvents() []EventMessage {
	a.version = a.CurrentVersion()
	events := a.changes
	a.changes = []EventMessage{}
	return events
}
