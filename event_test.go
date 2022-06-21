package cqrs_test

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/RamzesSoft/cqrs"
)

type SomeEvent struct {
	Item  string
	Count int
}

type SomeOtherEvent struct {
	OrderID string
}

func NewTestEventMessage(id string) *cqrs.EventDescriptor {
	ev := &SomeEvent{Item: uuid.NewString(), Count: rand.Intn(100)}
	return cqrs.NewEvent(id, ev)
}

func TestNewEventMessage(t *testing.T) {
	Convey("Check new Event", t, func() {
		id := uuid.NewString()
		ev := &SomeEvent{Item: "Some String", Count: 43}

		em := cqrs.NewEvent(id, ev)

		Convey("AggregateID equal prepared ID", func() {
			So(em.AggregateID(), ShouldEqual, id)
		})
		Convey("Event in EventMessage equal `em`", func() {
			So(em.Event(), ShouldResemble, ev)
		})
		Convey("EventType should be `cqrs_test.SomeEvent`", func() {
			So(em.EventType(), ShouldEqual, "cqrs_test.SomeEvent")
		})
	})
}

//func (s *EventSuite) TestShouldGetTypeOfAggregate(c *C) {
//em := &EventMessage{aggregate: &SomeAggregate{}}
//c.Assert(em.AggregateType(), Equals, "SomeAggregate")
//}
