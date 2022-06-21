package cqrs

//CommandDispatcher is the interface that should be implemented by command dispatcher
//
//The dispatcher is the mechanism through which command are distributed to
//the appropriate command handler.
//
//Command handlers are registered with the dispatcher for a given command type.
//It is good practice in CQRS to have only one command handler for a given command.
//When a command is passed to the dispatcher it will look for the registered command
//handler and call that handler's Handle method passing the command message as an
//argument.
//
//Commands contained in a CommandMessage envelope are passed to the CommandDispatcher via
//the dispatch method.
type CommandDispatcher interface {
	Dispatch(CommandMessage) error
	RegisterHandler(CommandHandler, ...interface{}) error
}
