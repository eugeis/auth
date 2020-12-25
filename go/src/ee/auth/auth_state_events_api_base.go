package auth

import (
	"errors"
	"fmt"
	"github.com/looplab/eventhorizon"
)

type AccountAggregateInitialHandler struct {
	CreatedHandler func(eventhorizon.Event, *AccountCreated, *Account) (err error)
}

func NewAccountAggregateInitialHandlerDefault() (ret *AccountAggregateInitialHandler) {
	ret = &AccountAggregateInitialHandler{}
	return
}

func (o *AccountAggregateInitialHandler) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Initial()
	return
}

func (o *AccountAggregateInitialHandler) Apply(event eventhorizon.Event, account *Account) (ret *AccountAggregateStateType, err error) {

	switch event.EventType() {
	case AccountCreatedEvent:
		err = o.CreatedHandler(event, event.Data().(*AccountCreated), account)
		if account.Disabled {
			ret = AccountAggregateStateTypes().Disabled()
		} else if !(account.Disabled) {
			ret = AccountAggregateStateTypes().Enabled()
		}
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), account))
	}
	return
}

func (o *AccountAggregateInitialHandler) SetupEventHandler() (err error) {

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

type AccountAggregateDeletedHandler struct {
}

func NewAccountAggregateDeletedHandlerDefault() (ret *AccountAggregateDeletedHandler) {
	ret = &AccountAggregateDeletedHandler{}
	return
}

func (o *AccountAggregateDeletedHandler) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Deleted()
	return
}

func (o *AccountAggregateDeletedHandler) Apply(event eventhorizon.Event, account *Account) (ret *AccountAggregateStateType, err error) {

	return
}

func (o *AccountAggregateDeletedHandler) SetupEventHandler() (err error) {
	return
}

type AccountAggregateDisabledHandler struct {
	EnabledHandler func(eventhorizon.Event, *Account) (err error)
}

func NewAccountAggregateDisabledHandlerDefault() (ret *AccountAggregateDisabledHandler) {
	ret = &AccountAggregateDisabledHandler{}
	return
}

func (o *AccountAggregateDisabledHandler) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Disabled()
	return
}

func (o *AccountAggregateDisabledHandler) Apply(event eventhorizon.Event, account *Account) (ret *AccountAggregateStateType, err error) {

	switch event.EventType() {
	case AccountEnabledEvent:
		err = o.EnabledHandler(event, account)
		ret = AccountAggregateStateTypes().Enabled()
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), account))
	}
	return
}

func (o *AccountAggregateDisabledHandler) SetupEventHandler() (err error) {

	//default handler implementation
	o.EnabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {

		entity.Disabled = false
		return
	}
	return
}

type AccountAggregateEnabledHandler struct {
	DeletedHandler  func(eventhorizon.Event, *Account) (err error)
	DisabledHandler func(eventhorizon.Event, *Account) (err error)
}

func NewAccountAggregateEnabledHandlerDefault() (ret *AccountAggregateEnabledHandler) {
	ret = &AccountAggregateEnabledHandler{}
	return
}

func (o *AccountAggregateEnabledHandler) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Enabled()
	return
}

func (o *AccountAggregateEnabledHandler) Apply(event eventhorizon.Event, account *Account) (ret *AccountAggregateStateType, err error) {

	switch event.EventType() {
	case AccountDeletedEvent:
		err = o.DeletedHandler(event, account)
		ret = AccountAggregateStateTypes().Deleted()
	case AccountDisabledEvent:
		err = o.DisabledHandler(event, account)
		ret = AccountAggregateStateTypes().Disabled()
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), account))
	}
	return
}

func (o *AccountAggregateEnabledHandler) SetupEventHandler() (err error) {

	//default handler implementation
	o.DeletedHandler = func(event eventhorizon.Event, entity *Account) (err error) {

		*entity = *NewAccountDefault()
		return
	}

	//default handler implementation
	o.DisabledHandler = func(event eventhorizon.Event, entity *Account) (err error) {

		entity.Disabled = true
		return
	}
	return
}

type AccountAggregateExistHandler struct {
	DeletedHandler func(eventhorizon.Event, *Account) (err error)
	UpdatedHandler func(eventhorizon.Event, *AccountUpdated, *Account) (err error)
}

func NewAccountAggregateExistHandlerDefault() (ret *AccountAggregateExistHandler) {
	ret = &AccountAggregateExistHandler{}
	return
}

func (o *AccountAggregateExistHandler) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Exist()
	return
}

func (o *AccountAggregateExistHandler) Apply(event eventhorizon.Event, account *Account) (ret *AccountAggregateStateType, err error) {

	switch event.EventType() {
	case AccountDeletedEvent:
		err = o.DeletedHandler(event, account)
		ret = AccountAggregateStateTypes().Deleted()
	case AccountUpdatedEvent:
		err = o.UpdatedHandler(event, event.Data().(*AccountUpdated), account)
	default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), account))
	}
	return
}

func (o *AccountAggregateExistHandler) SetupEventHandler() (err error) {

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
