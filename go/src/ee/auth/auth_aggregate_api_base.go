package auth

import (
	"context"
	"github.com/go-ee/utils/eh"
	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/eventhandler/projector"
)

const AccountAggregateType eventhorizon.AggregateType = "Account"

type AccountAggregateEngine struct {
	*eh.AggregateEngine
	AggregateExecutors *AccountAggregateExecutors
	AggregateHandlers  *AccountAggregateHandlers
}

func (o *AccountAggregateEngine) RegisterForCreated(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountCreated())
}

func (o *AccountAggregateEngine) RegisterForDeleted(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountDeleted())
}

func (o *AccountAggregateEngine) RegisterForDisabled(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountDisabled())
}

func (o *AccountAggregateEngine) RegisterForEnabled(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountEnabled())
}

func (o *AccountAggregateEngine) RegisterForLogged(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountLogged())
}

func (o *AccountAggregateEngine) RegisterForSentDisabledConfirmation(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountSentDisabledConfirmation())
}

func (o *AccountAggregateEngine) RegisterForSentEnabledConfirmation(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountSentEnabledConfirmation())
}

func (o *AccountAggregateEngine) RegisterForUpdated(handler eventhorizon.EventHandler) error {
	return o.RegisterForEvent(handler, AccountEventTypes().AccountUpdated())
}

func (o *AccountAggregateEngine) RegisterAccountProjector(projectorType string, listener AccountAggregateHandler, events []eventhorizon.EventType) (ret *AccountProjector, err error) {
	repo := o.Repos(projectorType, o.EntityFactory)
	ret = NewAccountProjector(projectorType, listener, repo)
	proj := projector.NewEventHandler(ret, repo)
	proj.SetEntityFactory(o.EntityFactory)
	err = o.RegisterForEvents(proj, events)
	return
}

type AccountProjector struct {
	AccountAggregateHandler
	projectorType projector.Type
	Repo          eventhorizon.ReadRepo
}

func NewAccountProjector(projectorType string, eventHandler AccountAggregateHandler, repo eventhorizon.ReadRepo) (ret *AccountProjector) {
	ret = &AccountProjector{
		AccountAggregateHandler: eventHandler,
		projectorType:           projector.Type(projectorType),
		Repo:                    repo,
	}
	return
}

func (o *AccountProjector) ProjectorType() projector.Type {
	return o.projectorType
}

func (o *AccountProjector) Project(
	ctx context.Context, event eventhorizon.Event, entity eventhorizon.Entity) (ret eventhorizon.Entity, err error) {

	ret = entity
	err = o.Apply(event, entity.(*Account))
	return
}

func NewAccountAggregateEngine(middleware *eh.Middleware) (ret *AccountAggregateEngine) {

	accountAggregateExecutors := NewAccountAggregateExecutorsFull()
	accountAggregateHandlers := NewAccountAggregateHandlersFull()

	entityFactory := func() eventhorizon.Entity { return NewAccountDefault() }
	aggregateEngine := eh.NewAggregateEngine(middleware, AccountAggregateType,
		func(id uuid.UUID) eventhorizon.Aggregate {
			return &AccountAggregate{
				AggregateBase:      events.NewAggregateBase(AccountAggregateType, id),
				Account:            NewAccountDefault(),
				AggregateExecutors: accountAggregateExecutors,
				AggregateHandlers:  accountAggregateHandlers,
			}
		}, entityFactory,
		AccountCommandTypes().Literals(), AccountEventTypes().Literals())

	ret = &AccountAggregateEngine{
		AggregateEngine:    aggregateEngine,
		AggregateExecutors: accountAggregateExecutors,
		AggregateHandlers:  accountAggregateHandlers,
	}
	return
}

func (o *AccountAggregateEngine) Setup() (err error) {
	if err = o.AggregateEngine.Setup(); err != nil {
		return
	}

	if err = o.AggregateExecutors.SetupCommandHandler(); err != nil {
		return
	}

	if err = o.AggregateHandlers.SetupEventHandler(); err != nil {
		return
	}
	return
}

type EsEngine struct {
	*eh.Middleware
	Account *AccountAggregateEngine
}

func NewEsEngine(middleware *eh.Middleware) (ret *EsEngine) {
	account := NewAccountAggregateEngine(middleware)
	ret = &EsEngine{
		Middleware: middleware,
		Account:    account,
	}
	return
}

func (o *EsEngine) Setup() (err error) {

	if err = o.Account.Setup(); err != nil {
		return
	}

	return
}
