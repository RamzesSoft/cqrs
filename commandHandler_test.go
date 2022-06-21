package cqrs_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/RamzesSoft/cqrs"
)

type TestCommandHandler struct {
	command cqrs.CommandMessage
}

func (t *TestCommandHandler) Handle(command cqrs.CommandMessage) error {
	t.command = command
	return nil
}

type MockRepository struct {
	aggregates map[string]cqrs.AggregateRoot
}

func (m *MockRepository) Load(aggregateType string, id string) (cqrs.AggregateRoot, error) {
	_ = aggregateType
	return m.aggregates[id], nil
}

func (m *MockRepository) Save(aggregate cqrs.AggregateRoot) error {
	m.aggregates[aggregate.AggregateID()] = aggregate
	return nil
}

func TestCommandHandler2(t *testing.T) {
	Convey("CommandHandler test", t, func() {

	})
}
