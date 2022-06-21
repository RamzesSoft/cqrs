package cqrs_test

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/RamzesSoft/cqrs"
)

type SomeCommand struct {
	Item  string
	Count int
}

func NewSomeCommandMessage(id string) *cqrs.CommandDescriptor {
	ev := &SomeCommand{Item: uuid.NewString(), Count: rand.Intn(100)}
	return cqrs.NewCommandMessage(id, ev)
}

func TestNewCommandMessage(t *testing.T) {
	Convey("SomeCommand test", t, func() {
		id := uuid.NewString()
		cmd := &SomeCommand{Item: "Some String", Count: 43}

		cm := cqrs.NewCommandMessage(id, cmd)
		Convey("AggregateID should be equal prepared ID", func() {
			So(cm.AggregateID(), ShouldEqual, id)
		})
		Convey("Command should be equal prepared `command`", func() {
			So(cm.Command(), ShouldEqual, cmd)
		})
		Convey("CommandType should be equal to `cqrs_test.SomeCommand`", func() {
			So(cm.CommandType(), ShouldEqual, "cqrs_test.SomeCommand")
		})
	})
}

type SomeOtherCommand struct {
	OrderID string
}

func NewSomeOtherCommandMessage(id string) *cqrs.CommandDescriptor {
	ev := &SomeOtherCommand{id}
	return cqrs.NewCommandMessage(id, ev)
}

type ErrorCommand struct {
	Message string
}
