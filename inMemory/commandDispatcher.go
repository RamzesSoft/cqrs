package inMemory

import (
	"fmt"

	"github.com/RamzesSoft/cqrs"
)

//commandDispatcher provides a lightweight and performant in process commandDispatcher
type commandDispatcher struct {
	handlers map[string]cqrs.CommandHandler
	//isDebug bool
}

//NewCommandDispatcher constructs a new in memory commandDispatcher
func NewCommandDispatcher() cqrs.CommandDispatcher {
	return &commandDispatcher{handlers: make(map[string]cqrs.CommandHandler)}
}

//Dispatch passes the CommandMessage on to all registered command handlers.
func (b *commandDispatcher) Dispatch(command cqrs.CommandMessage) error {
	if handler, ok := b.handlers[command.CommandType()]; ok {
		return handler.Handle(command)
	}
	return fmt.Errorf("the command bus does not have a handler for command of type: %s", command.CommandType())
}

//RegisterHandler registers a command handler for the command types specified by the
//variadic command parameter.
func (b *commandDispatcher) RegisterHandler(handler cqrs.CommandHandler, commands ...interface{}) error {
	for _, command := range commands {
		typeName := cqrs.TypeWithPackage(command)
		if _, ok := b.handlers[typeName]; ok {
			return fmt.Errorf("duplicate command handler registration with command bus for command of type: %s", typeName)
		}
		b.handlers[typeName] = handler
	}
	return nil
}
