package cqrs_test

import (
	"github.com/google/uuid"

	"github.com/RamzesSoft/cqrs"

	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAggregateBase(t *testing.T) {
	id := uuid.NewString()
	Convey("NewAggregate check create", t, func() {
		agg := cqrs.NewAggregate(id)
		Convey("Aggregate root not nil", func() {
			So(agg, ShouldNotBeNil)
		})
		Convey("AggregateID should be equal prepared ID", func() {
			So(agg.AggregateID(), ShouldEqual, id)
		})
		Convey("OriginalVersion should be equal `-1`", func() {
			So(agg.OriginalVersion(), ShouldEqual, -1)
		})
		Convey("CurrentVersion should be equal `-1`", func() {
			So(agg.CurrentVersion(), ShouldEqual, -1)
		})
	})
}

func TestTrackOneChange(t *testing.T) {
	Convey("Track one change", t, func() {
		ev := NewTestEventMessage(uuid.NewString())
		agg := NewSomeAggregate(ev.AggregateID())
		agg.RecordEvent(ev)
		So(agg.ReleaseEvents(), ShouldResemble, []cqrs.EventMessage{ev})
	})
}

func TestTrackMultipleChanges(t *testing.T) {
	Convey("Track multiple change", t, func() {
		agg := cqrs.NewAggregate(uuid.NewString())
		ev1 := NewTestEventMessage(agg.AggregateID())
		ev2 := NewTestEventMessage(agg.AggregateID())

		agg.RecordEvent(ev1)
		agg.RecordEvent(ev2)

		Convey("OriginalVersion should be equal `-1`", func() {
			So(agg.OriginalVersion(), ShouldEqual, -1)
		})
		Convey("CurrentVersion should be equal `1`", func() {
			So(agg.CurrentVersion(), ShouldEqual, 1)
		})

		events := agg.ReleaseEvents()
		Convey("ReleaseEvents should return both Events", func() {
			So(events, ShouldResemble, []cqrs.EventMessage{ev1, ev2})
		})
	})
}

func TestReleaseEvents(t *testing.T) {
	Convey("ReLeaseEvents should return all Events and make equal CurrentVersion and OriginalVersion", t, func() {
		agg := cqrs.NewAggregate(uuid.NewString())
		ev1 := NewTestEventMessage(agg.AggregateID())
		ev2 := NewTestEventMessage(agg.AggregateID())

		agg.RecordEvent(ev1)
		agg.RecordEvent(ev2)

		Convey("OriginalVersion Before ReleaseEvents should be equal `-1`", func() {
			So(agg.OriginalVersion(), ShouldEqual, -1)
		})
		Convey("CurrentVersion  Before ReleaseEvents should be equal `1`", func() {
			So(agg.CurrentVersion(), ShouldEqual, 1)
		})
		events := agg.ReleaseEvents()
		Convey("ReleaseEvents should return both Events", func() {
			So(events, ShouldResemble, []cqrs.EventMessage{ev1, ev2})
		})
		Convey("OriginalVersion After ReleaseEvents should be equal `1`", func() {
			So(agg.OriginalVersion(), ShouldEqual, 1)
		})
		Convey("CurrentVersion After ReleaseEvents should be equal `1`", func() {
			So(agg.CurrentVersion(), ShouldEqual, 1)
		})
		events = agg.ReleaseEvents()
		Convey("ReleaseEvents should return empty []", func() {
			So(events, ShouldBeEmpty)
		})
		Convey("OriginalVersion After second ReleaseEvents should be equal `1`", func() {
			So(agg.OriginalVersion(), ShouldEqual, 1)
		})
		Convey("CurrentVersion After second ReleaseEvents should be equal `1`", func() {
			So(agg.CurrentVersion(), ShouldEqual, 1)
		})
	})
}

type SomeAggregate struct {
	cqrs.AggregateRoot
}

func NewSomeAggregate(id string) cqrs.AggregateRoot {
	return &SomeAggregate{
		AggregateRoot: cqrs.NewAggregate(id),
	}
}

func (t *SomeAggregate) Handle(command cqrs.CommandMessage) error {
	_ = command
	return nil
}
