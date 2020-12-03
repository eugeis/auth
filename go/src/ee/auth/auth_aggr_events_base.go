package auth

import (
	"errors"
	"fmt"
	"github.com/go-ee/utils"
	"github.com/looplab/eventhorizon"
	"time"
)

type AccountEventHandler struct {
	CreatedHandler                  func(eventhorizon.Event, *AccountCreated, *Account) (err error)
	DeletedHandler                  func(eventhorizon.Event, *AccountDeleted, *Account) (err error)
	DisabledHandler                 func(eventhorizon.Event, *AccountDisabled, *Account) (err error)
	EnabledHandler                  func(eventhorizon.Event, *AccountEnabled, *Account) (err error)
	LoggedHandler                   func(eventhorizon.Event, *AccountLogged, *Account) (err error)
	SentCreatedConfirmationHandler  func(eventhorizon.Event, *AccountSentCreatedConfirmation, *Account) (err error)
	SentDisabledConfirmationHandler func(eventhorizon.Event, *AccountSentDisabledConfirmation, *Account) (err error)
	SentEnabledConfirmationHandler  func(eventhorizon.Event, *AccountSentEnabledConfirmation, *Account) (err error)
	UpdatedHandler                  func(eventhorizon.Event, *AccountUpdated, *Account) (err error)
}

func (o *AccountEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountCreatedEvent:
		err = o.CreatedHandler(event, event.Data().(*AccountCreated), entity.(*Account))
	case AccountEnabledEvent:
		err = o.EnabledHandler(event, event.Data().(*AccountEnabled), entity.(*Account))
	case AccountDisabledEvent:
		err = o.DisabledHandler(event, event.Data().(*AccountDisabled), entity.(*Account))
	case AccountUpdatedEvent:
		err = o.UpdatedHandler(event, event.Data().(*AccountUpdated), entity.(*Account))
	case AccountDeletedEvent:
		err = o.DeletedHandler(event, event.Data().(*AccountDeleted), entity.(*Account))
	case AccountSentEnabledConfirmationEvent:
		err = o.SentEnabledConfirmationHandler(event, event.Data().(*AccountSentEnabledConfirmation), entity.(*Account))
	case AccountSentDisabledConfirmationEvent:
		err = o.SentDisabledConfirmationHandler(event, event.Data().(*AccountSentDisabledConfirmation), entity.(*Account))
	case AccountLoggedEvent:
		err = o.LoggedHandler(event, event.Data().(*AccountLogged), entity.(*Account))
	case AccountSentCreatedConfirmationEvent:
		err = o.SentCreatedConfirmationHandler(event, event.Data().(*AccountSentCreatedConfirmation), entity.(*Account))
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

	//register event object factory
	eventhorizon.RegisterEventData(AccountEnabledEvent, func() eventhorizon.EventData {
		return &AccountEnabled{}
	})

	//default handler implementation
	o.EnabledHandler = func(event eventhorizon.Event, eventData *AccountEnabled, entity *Account) (err error) {
		return
	}

	//register event object factory
	eventhorizon.RegisterEventData(AccountDisabledEvent, func() eventhorizon.EventData {
		return &AccountDisabled{}
	})

	//default handler implementation
	o.DisabledHandler = func(event eventhorizon.Event, eventData *AccountDisabled, entity *Account) (err error) {
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

	//register event object factory
	eventhorizon.RegisterEventData(AccountDeletedEvent, func() eventhorizon.EventData {
		return &AccountDeleted{}
	})

	//default handler implementation
	o.DeletedHandler = func(event eventhorizon.Event, eventData *AccountDeleted, entity *Account) (err error) {
		entity.DeletedAt = utils.PtrTime(time.Now())
		return
	}

	//register event object factory
	eventhorizon.RegisterEventData(AccountSentEnabledConfirmationEvent, func() eventhorizon.EventData {
		return &AccountSentEnabledConfirmation{}
	})

	//default handler implementation
	o.SentEnabledConfirmationHandler = func(event eventhorizon.Event, eventData *AccountSentEnabledConfirmation, entity *Account) (err error) {
		//err = eh.EventHandlerNotImplemented(AccountSentEnabledConfirmationEvent)
		return
	}

	//register event object factory
	eventhorizon.RegisterEventData(AccountSentDisabledConfirmationEvent, func() eventhorizon.EventData {
		return &AccountSentDisabledConfirmation{}
	})

	//default handler implementation
	o.SentDisabledConfirmationHandler = func(event eventhorizon.Event, eventData *AccountSentDisabledConfirmation, entity *Account) (err error) {
		//err = eh.EventHandlerNotImplemented(AccountSentDisabledConfirmationEvent)
		return
	}

	//register event object factory
	eventhorizon.RegisterEventData(AccountLoggedEvent, func() eventhorizon.EventData {
		return &AccountLogged{}
	})

	//default handler implementation
	o.LoggedHandler = func(event eventhorizon.Event, eventData *AccountLogged, entity *Account) (err error) {
		//err = eh.EventHandlerNotImplemented(AccountLoggedEvent)
		return
	}

	//register event object factory
	eventhorizon.RegisterEventData(AccountSentCreatedConfirmationEvent, func() eventhorizon.EventData {
		return &AccountSentCreatedConfirmation{}
	})

	//default handler implementation
	o.SentCreatedConfirmationHandler = func(event eventhorizon.Event, eventData *AccountSentCreatedConfirmation, entity *Account) (err error) {
		//err = eh.EventHandlerNotImplemented(AccountSentCreatedConfirmationEvent)
		return
	}
	return
}
