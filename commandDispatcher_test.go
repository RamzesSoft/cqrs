package cqrs_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/RamzesSoft/cqrs"
	"github.com/RamzesSoft/cqrs/inMemory"
)

func TestNewInternalCommandBus(t *testing.T) {
	Convey("NewCommandDispatcher created correct", t, func() {
		bus := inMemory.NewCommandDispatcher()
		So(bus, ShouldNotBeNil)
	})
}

func TestShouldHandleCommand(t *testing.T) {
	Convey("RegisterHandler test", t, func() {
		inMemoryDispatcher := inMemory.NewCommandDispatcher()
		err := inMemoryDispatcher.RegisterHandler(&TestCommandHandler{}, &SomeCommand{})
		So(err, ShouldBeNil)

		cmd := NewSomeCommandMessage(uuid.NewString())
		err = inMemoryDispatcher.Dispatch(cmd)
		So(err, ShouldBeNil)
		//So(inMemoryDispatcher.C, ShouldResemble, cmd)
	})
}

func TestShouldReturnErrorIfNoHandlerRegisteredForCommand(t *testing.T) {
	Convey("Error if command don't have handler", t, func() {
		cmd := NewSomeCommandMessage(uuid.NewString())
		inMemoryDispatcher := inMemory.NewCommandDispatcher()
		err := inMemoryDispatcher.Dispatch(cmd)
		So(err, ShouldBeError)
		So(err, ShouldResemble, fmt.Errorf("the command bus does not have a handler for command of type: %s", cmd.CommandType()))
	})
}

func TestDuplicateHandlerRegistrationReturnsAnError(t *testing.T) {
	Convey("RegisterHandler duplicate handler should return Error", t, func() {
		inMemoryDispatcher := inMemory.NewCommandDispatcher()
		err := inMemoryDispatcher.RegisterHandler(&TestCommandHandler{}, &SomeCommand{}, &SomeCommand{})
		So(
			err,
			ShouldResemble,
			fmt.Errorf(
				"duplicate command handler registration with command bus for command of type: %s",
				cqrs.TypeWithPackage(&SomeCommand{"", 0}),
			),
		)

	})
}

func TestCanRegisterMultipleCommandsForTheSameHandler(t *testing.T) {
	Convey("RegisterHandler multiple command for the same handler work correct", t, func() {
		inMemoryDispatcher := inMemory.NewCommandDispatcher()
		err := inMemoryDispatcher.RegisterHandler(&TestCommandHandler{}, &SomeCommand{}, &SomeOtherCommand{})
		So(err, ShouldBeNil)
	})
}
