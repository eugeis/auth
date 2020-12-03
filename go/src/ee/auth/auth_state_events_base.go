package auth

import (
	"errors"
	"fmt"
	"github.com/looplab/eventhorizon"
)

type AccountConfirmationDisabledHandler struct {
	EnabledHandler func(eventhorizon.Event, *Account) (err error)
}

func NewAccountConfirmationDisabledHandlerDefault() (ret *AccountConfirmationDisabledHandler) {
	ret = &AccountConfirmationDisabledHandler{}
	return
}

func (o *AccountConfirmationDisabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountEnabledEvent:
		err = o.EnabledHandler(event, entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountConfirmationDisabledHandler) SetupEventHandler() (err error) {

	//default handler implementation
	o.EnabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		entity.Disabled = false
		return
	}
	return
}

type AccountConfirmationEnabledHandler struct {
	DisabledHandler func(eventhorizon.Event, *Account) (err error)
}

func NewAccountConfirmationEnabledHandlerDefault() (ret *AccountConfirmationEnabledHandler) {
	ret = &AccountConfirmationEnabledHandler{}
	return
}

func (o *AccountConfirmationEnabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountDisabledEvent:
		err = o.DisabledHandler(event, entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountConfirmationEnabledHandler) SetupEventHandler() (err error) {

	//default handler implementation
	o.DisabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		entity.Disabled = true
		return
	}
	return
}

type AccountConfirmationInitialHandler struct {
	CreatedHandler func(eventhorizon.Event, *AccountCreated, *Account) (err error)
}

func NewAccountConfirmationInitialHandlerDefault() (ret *AccountConfirmationInitialHandler) {
	ret = &AccountConfirmationInitialHandler{}
	return
}

func (o *AccountConfirmationInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountCreatedEvent:
		err = o.CreatedHandler(event, event.Data().(*AccountCreated), entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountConfirmationInitialHandler) SetupEventHandler() (err error) {

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
	return
}

type AccountDeletedHandler struct {
}

func NewAccountDeletedHandlerDefault() (ret *AccountDeletedHandler) {
	ret = &AccountDeletedHandler{}
	return
}

func (o *AccountDeletedHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {

	return
}

func (o *AccountDeletedHandler) SetupEventHandler() (err error) {
	return
}

type AccountDisabledHandler struct {
	EnabledHandler func(eventhorizon.Event, *Account) (err error)
}

func NewAccountDisabledHandlerDefault() (ret *AccountDisabledHandler) {
	ret = &AccountDisabledHandler{}
	return
}

func (o *AccountDisabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountEnabledEvent:
		err = o.EnabledHandler(event, entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountDisabledHandler) SetupEventHandler() (err error) {

	//default handler implementation
	o.EnabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		entity.Disabled = false
		return
	}
	return
}

type AccountEnabledHandler struct {
	DisabledHandler func(eventhorizon.Event, *Account) (err error)
}

func NewAccountEnabledHandlerDefault() (ret *AccountEnabledHandler) {
	ret = &AccountEnabledHandler{}
	return
}

func (o *AccountEnabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountDisabledEvent:
		err = o.DisabledHandler(event, entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountEnabledHandler) SetupEventHandler() (err error) {

	//default handler implementation
	o.DisabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		entity.Disabled = true
		return
	}
	return
}

type AccountExistHandler struct {
	DeletedHandler func(eventhorizon.Event, *Account) (err error)
	UpdatedHandler func(eventhorizon.Event, *AccountUpdated, *Account) (err error)
}

func NewAccountExistHandlerDefault() (ret *AccountExistHandler) {
	ret = &AccountExistHandler{}
	return
}

func (o *AccountExistHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountDeletedEvent:
		err = o.DeletedHandler(event, entity.(*Account))
	case AccountUpdatedEvent:
		err = o.UpdatedHandler(event, event.Data().(*AccountUpdated), entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountExistHandler) SetupEventHandler() (err error) {

	//default handler implementation
	o.DeletedHandler = func(event eventhorizon.Event, entity *Account) (err error) {
		*entity = *NewAccountDefault()
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
	return
}

type AccountInitialHandler struct {
	CreatedHandler func(eventhorizon.Event, *AccountCreated, *Account) (err error)
}

func NewAccountInitialHandlerDefault() (ret *AccountInitialHandler) {
	ret = &AccountInitialHandler{}
	return
}

func (o *AccountInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
	switch event.EventType() {
	case AccountCreatedEvent:
		err = o.CreatedHandler(event, event.Data().(*AccountCreated), entity.(*Account))
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
	return
}

func (o *AccountInitialHandler) SetupEventHandler() (err error) {

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
	return
}
