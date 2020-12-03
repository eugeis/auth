package auth

import (
	"errors"
	"fmt"
	"github.com/go-ee/utils"
	"github.com/go-ee/utils/eh"
	"github.com/looplab/eventhorizon"
	"time"
)

type AccountEventHandler struct {
	CreatedHandler                  func(eventhorizon.Event, *AccountCreated, *Account) (err error)
	DeletedHandler                  func(eventhorizon.Event, *Account) (err error)
	DisabledHandler                 func(eventhorizon.Event, *Account) (err error)
	EnabledHandler                  func(eventhorizon.Event, *Account) (err error)
	LoggedHandler                   func(eventhorizon.Event, *AccountLogged, *Account) (err error)
	SentCreatedConfirmationHandler  func(eventhorizon.Event, *Account) (err error)
	SentDisabledConfirmationHandler func(eventhorizon.Event, *Account) (err error)
	SentEnabledConfirmationHandler  func(eventhorizon.Event, *Account) (err error)
	UpdatedHandler                  func(eventhorizon.Event, *AccountUpdated, *Account) (err error)
}

func (o *AccountEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountCreatedEvent:
		err = o.CreatedHandler(event, event.Data().(*AccountCreated), entity.(*Account))
	case AccountEnabledEvent:
		err = o.EnabledHandler(event, entity.(*Account))
	case AccountDisabledEvent:
		err = o.DisabledHandler(event, entity.(*Account))
	case AccountUpdatedEvent:
		err = o.UpdatedHandler(event, event.Data().(*AccountUpdated), entity.(*Account))
	case AccountDeletedEvent:
		err = o.DeletedHandler(event, entity.(*Account))
	case AccountSentEnabledConfirmationEvent:
		err = o.SentEnabledConfirmationHandler(event, entity.(*Account))
	case AccountSentDisabledConfirmationEvent:
		err = o.SentDisabledConfirmationHandler(event, entity.(*Account))
	case AccountLoggedEvent:
		err = o.LoggedHandler(event, event.Data().(*AccountLogged), entity.(*Account))
	case AccountSentCreatedConfirmationEvent:
		err = o.SentCreatedConfirmationHandler(event, entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountEventHandler) SetupEventHandler() (err error) {
	//register event object factory
	eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

	//default handler implementation
	o.CreatedHandler = func(event eventhorizon.Event, eventData *AccountCreated, entity *Account) (err error) {
		entity.Id = event.AggregateID()
		entity.Name = eventData.Name
		entity.Username = eventData.Username
		entity.Password = eventData.Password
		entity.Email = eventData.Email
		entity.Roles = eventData.Roles
		return
	}
	//default handler implementation
	o.EnabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		entity.Disabled = false
		return
	}
	//default handler implementation
	o.DisabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		entity.Disabled = true
		return
	}
	//register event object factory
	eventhorizon.RegisterEventData(AccountUpdatedEvent, func() eventhorizon.EventData {
		return &AccountUpdated{}
	})

	//default handler implementation
	o.UpdatedHandler = func(event eventhorizon.Event, eventData *AccountUpdated, entity *Account) (err error) {
		entity.Name = eventData.Name
		entity.Username = eventData.Username
		entity.Password = eventData.Password
		entity.Email = eventData.Email
		entity.Roles = eventData.Roles
		return
	}
	//default handler implementation
	o.DeletedHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		entity.DeletedAt = utils.PtrTime(time.Now())
		return
	}
	//default handler implementation
	o.SentEnabledConfirmationHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		err = eh.EventHandlerNotImplemented(AccountSentEnabledConfirmationEvent)
		return
	}
	//default handler implementation
	o.SentDisabledConfirmationHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		err = eh.EventHandlerNotImplemented(AccountSentDisabledConfirmationEvent)
		return
	}
	//register event object factory
	eventhorizon.RegisterEventData(AccountLoggedEvent, func() eventhorizon.EventData {
		return &AccountLogged{}
	})

	//default handler implementation
	o.LoggedHandler = func(event eventhorizon.Event, eventData *AccountLogged, entity *Account) (err error) {
		err = eh.EventHandlerNotImplemented(AccountLoggedEvent)
		return
	}
	//default handler implementation
	o.SentCreatedConfirmationHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		err = eh.EventHandlerNotImplemented(AccountSentCreatedConfirmationEvent)
		return
	}
	return
}
