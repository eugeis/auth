package auth

import (
	"errors"
	"fmt"
	"github.com/go-ee/utils/eh"
	"github.com/looplab/eventhorizon"
	"time"
)

type AccountCommandHandler struct {
	CommandsPreparer                func(eventhorizon.Command, *Account) (err error)
	SendEnabledConfirmationHandler  func(*SendEnabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
	SendDisabledConfirmationHandler func(*SendDisabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
	LoginHandler                    func(*LoginAccount, *Account, eh.AggregateStoreEvent) (err error)
	CreateHandler                   func(*CreateAccount, *Account, eh.AggregateStoreEvent) (err error)
	DeleteHandler                   func(*DeleteAccount, *Account, eh.AggregateStoreEvent) (err error)
	EnableHandler                   func(*EnableAccount, *Account, eh.AggregateStoreEvent) (err error)
	DisableHandler                  func(*DisableAccount, *Account, eh.AggregateStoreEvent) (err error)
	UpdateHandler                   func(*UpdateAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func (o *AccountCommandHandler) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
}

func (o *AccountCommandHandler) AddSendEnabledConfirmationPreparer(preparer func(*SendEnabledConfirmationAccount, *Account) (err error)) {
	prevHandler := o.SendEnabledConfirmationHandler
	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddSendDisabledConfirmationPreparer(preparer func(*SendDisabledConfirmationAccount, *Account) (err error)) {
	prevHandler := o.SendDisabledConfirmationHandler
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddLoginPreparer(preparer func(*LoginAccount, *Account) (err error)) {
	prevHandler := o.LoginHandler
	o.LoginHandler = func(command *LoginAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddCreatePreparer(preparer func(*CreateAccount, *Account) (err error)) {
	prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddDeletePreparer(preparer func(*DeleteAccount, *Account) (err error)) {
	prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddEnablePreparer(preparer func(*EnableAccount, *Account) (err error)) {
	prevHandler := o.EnableHandler
	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddDisablePreparer(preparer func(*DisableAccount, *Account) (err error)) {
	prevHandler := o.DisableHandler
	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddUpdatePreparer(preparer func(*UpdateAccount, *Account) (err error)) {
	prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}

	switch cmd.CommandType() {
	case SendEnabledConfirmationAccountCommand:
		err = o.SendEnabledConfirmationHandler(cmd.(*SendEnabledConfirmationAccount), account, store)
	case SendDisabledConfirmationAccountCommand:
		err = o.SendDisabledConfirmationHandler(cmd.(*SendDisabledConfirmationAccount), account, store)
	case LoginAccountCommand:
		err = o.LoginHandler(cmd.(*LoginAccount), account, store)
	case CreateAccountCommand:
		err = o.CreateHandler(cmd.(*CreateAccount), account, store)
	case DeleteAccountCommand:
		err = o.DeleteHandler(cmd.(*DeleteAccount), account, store)
	case EnableAccountCommand:
		err = o.EnableHandler(cmd.(*EnableAccount), account, store)
	case DisableAccountCommand:
		err = o.DisableHandler(cmd.(*DisableAccount), account, store)
	case UpdateAccountCommand:
		err = o.UpdateHandler(cmd.(*UpdateAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
	return
}

func (o *AccountCommandHandler) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
		return
	}

	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountSentEnabledConfirmationEvent, nil, time.Now())
		return
	}
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountSentDisabledConfirmationEvent, nil, time.Now())
		return
	}
	o.LoginHandler = func(command *LoginAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountLoggedEvent, &AccountLogged{
			Username: command.Username,
			Email:    command.Email,
			Password: command.Password}, time.Now())
		return
	}
	o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountCreatedEvent, &AccountCreated{
			Name:     command.Name,
			Username: command.Username,
			Password: command.Password,
			Email:    command.Email,
			Roles:    command.Roles}, time.Now())
		return
	}
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountDeletedEvent, nil, time.Now())
		return
	}
	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountEnabledEvent, nil, time.Now())
		return
	}
	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountDisabledEvent, nil, time.Now())
		return
	}
	o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountUpdatedEvent, &AccountUpdated{
			Name:     command.Name,
			Username: command.Username,
			Password: command.Password,
			Email:    command.Email,
			Roles:    command.Roles}, time.Now())
		return
	}
	return
}
