package cqrs

// CommandMessage is the interface that a command message must implement.
type CommandMessage interface {
	// AggregateID returns the ID of the Aggregate that the command relates to
	AggregateID() string
	// Command returns the actual command which is the payload of the command message.
	Command() interface{}
	// CommandType returns a string descriptor of the command name
	CommandType() string
}

// CommandDescriptor is an implementation of the command message interface.
type CommandDescriptor struct {
	id      string
	command interface{}
}

// NewCommandMessage returns a new command descriptor
func NewCommandMessage(aggregateID string, command interface{}) *CommandDescriptor {
	return &CommandDescriptor{
		id:      aggregateID,
		command: command,
	}
}

// CommandType returns the command type name as a string
func (c *CommandDescriptor) CommandType() string {
	return TypeWithPackage(c.command)
}

// AggregateID returns the ID of the aggregate that the command relates to.
func (c *CommandDescriptor) AggregateID() string {
	return c.id
}

// Command returns the actual command payload of the message.
func (c *CommandDescriptor) Command() interface{} {
	return c.command
}
