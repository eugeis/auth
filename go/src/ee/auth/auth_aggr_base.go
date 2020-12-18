package auth

import (
	"github.com/go-ee/utils/eh"
	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/commandhandler/bus"
)

const AccountAggregateType eventhorizon.AggregateType = "Account"

type AccountAggrInitializer struct {
	*eh.AggregateInitializer
	*AccountCommandHandler
	*AccountEventHandler
	ProjectorHandler *AccountEventHandler
}

func (o *AccountAggrInitializer) RegisterForSentDisabledConfirmation(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountSentDisabledConfirmation())
}

func (o *AccountAggrInitializer) RegisterForSentEnabledConfirmation(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountSentEnabledConfirmation())
}

func (o *AccountAggrInitializer) RegisterForLogged(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountLogged())
}

func NewAccountAggrInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler,
	readRepos func(string, func() (ret eventhorizon.Entity)) (ret eventhorizon.ReadWriteRepo)) (ret *AccountAggrInitializer) {

	commandHandler := &AccountCommandHandler{}
	eventHandler := &AccountEventHandler{}
	entityFactory := func() eventhorizon.Entity { return NewAccountDefault() }
	ret = &AccountAggrInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
		func(id uuid.UUID) eventhorizon.Aggregate {
			return eh.NewAggregateBase(AccountAggregateType, id, commandHandler, eventHandler, entityFactory())
		}, entityFactory,
		AccountCommandTypes().Literals(), AccountEventTypes().Literals(), eventHandler,
		[]func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
		eventStore, eventBus, commandBus, readRepos), AccountCommandHandler: commandHandler, AccountEventHandler: eventHandler, ProjectorHandler: eventHandler,
	}

	return
}

type EsInitializer struct {
	eventStore             eventhorizon.EventStore
	eventBus               eventhorizon.EventBus
	commandBus             *bus.CommandHandler
	AccountAggrInitializer *AccountAggrInitializer
}

func NewEsInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler,
	readRepos func(string, func() (ret eventhorizon.Entity)) (ret eventhorizon.ReadWriteRepo)) (ret *EsInitializer) {
	accountAggrInitializer := NewAccountAggrInitializer(eventStore, eventBus, commandBus, readRepos)
	ret = &EsInitializer{
		eventStore:             eventStore,
		eventBus:               eventBus,
		commandBus:             commandBus,
		AccountAggrInitializer: accountAggrInitializer,
	}
	return
}

func (o *EsInitializer) Setup() (err error) {

	if err = o.AccountAggrInitializer.Setup(); err != nil {
		return
	}

	return
}
