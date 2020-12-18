package auth

import (
	"errors"
	"fmt"
	"github.com/go-ee/utils/eh"
	"github.com/looplab/eventhorizon"
	"time"
)

type AccountConfirmationDisabledExecutor struct {
	CommandsPreparer                func(eventhorizon.Command, *Account) (err error)
	SendDisabledConfirmationHandler func(*SendDisabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountConfirmationDisabledExecutorDefault() (ret *AccountConfirmationDisabledExecutor) {
	ret = &AccountConfirmationDisabledExecutor{}
	return
}

func (o *AccountConfirmationDisabledExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountConfirmationDisabledExecutor) AddSendDisabledConfirmationPreparer(preparer func(*SendDisabledConfirmationAccount, *Account) (err error)) {
	prevHandler := o.SendDisabledConfirmationHandler
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountConfirmationDisabledExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	switch cmd.CommandType() {
	case SendDisabledConfirmationAccountCommand:
		err = o.SendDisabledConfirmationHandler(cmd.(*SendDisabledConfirmationAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
	return
}

func (o *AccountConfirmationDisabledExecutor) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
		return
	}

	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountSentDisabledConfirmationEvent, nil, time.Now())
		return
	}
	return
}

type AccountConfirmationEnabledExecutor struct {
	CommandsPreparer               func(eventhorizon.Command, *Account) (err error)
	SendEnabledConfirmationHandler func(*SendEnabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountConfirmationEnabledExecutorDefault() (ret *AccountConfirmationEnabledExecutor) {
	ret = &AccountConfirmationEnabledExecutor{}
	return
}

func (o *AccountConfirmationEnabledExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountConfirmationEnabledExecutor) AddSendEnabledConfirmationPreparer(preparer func(*SendEnabledConfirmationAccount, *Account) (err error)) {
	prevHandler := o.SendEnabledConfirmationHandler
	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountConfirmationEnabledExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	switch cmd.CommandType() {
	case SendEnabledConfirmationAccountCommand:
		err = o.SendEnabledConfirmationHandler(cmd.(*SendEnabledConfirmationAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
	return
}

func (o *AccountConfirmationEnabledExecutor) SetupCommandHandler() (err error) {
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
	return
}

type AccountConfirmationInitialExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
}

func NewAccountConfirmationInitialExecutorDefault() (ret *AccountConfirmationInitialExecutor) {
	ret = &AccountConfirmationInitialExecutor{}
	return
}

func (o *AccountConfirmationInitialExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountConfirmationInitialExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	return
}

func (o *AccountConfirmationInitialExecutor) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
		return
	}

	return
}

type AccountDeletedExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
}

func NewAccountDeletedExecutorDefault() (ret *AccountDeletedExecutor) {
	ret = &AccountDeletedExecutor{}
	return
}

func (o *AccountDeletedExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountDeletedExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	return
}

func (o *AccountDeletedExecutor) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
		return
	}

	return
}

type AccountDisabledExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
	EnableHandler    func(*EnableAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountDisabledExecutorDefault() (ret *AccountDisabledExecutor) {
	ret = &AccountDisabledExecutor{}
	return
}

func (o *AccountDisabledExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountDisabledExecutor) AddEnablePreparer(preparer func(*EnableAccount, *Account) (err error)) {
	prevHandler := o.EnableHandler
	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountDisabledExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	switch cmd.CommandType() {
	case EnableAccountCommand:
		err = o.EnableHandler(cmd.(*EnableAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
	return
}

func (o *AccountDisabledExecutor) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
		return
	}

	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountEnabledEvent, nil, time.Now())
		return
	}
	return
}

type AccountEnabledExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
	DisableHandler   func(*DisableAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountEnabledExecutorDefault() (ret *AccountEnabledExecutor) {
	ret = &AccountEnabledExecutor{}
	return
}

func (o *AccountEnabledExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountEnabledExecutor) AddDisablePreparer(preparer func(*DisableAccount, *Account) (err error)) {
	prevHandler := o.DisableHandler
	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountEnabledExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	switch cmd.CommandType() {
	case DisableAccountCommand:
		err = o.DisableHandler(cmd.(*DisableAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
	return
}

func (o *AccountEnabledExecutor) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
		return
	}

	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountDisabledEvent, nil, time.Now())
		return
	}
	return
}

type AccountExistExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
	DeleteHandler    func(*DeleteAccount, *Account, eh.AggregateStoreEvent) (err error)
	UpdateHandler    func(*UpdateAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountExistExecutorDefault() (ret *AccountExistExecutor) {
	ret = &AccountExistExecutor{}
	return
}

func (o *AccountExistExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountExistExecutor) AddDeletePreparer(preparer func(*DeleteAccount, *Account) (err error)) {
	prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountExistExecutor) AddUpdatePreparer(preparer func(*UpdateAccount, *Account) (err error)) {
	prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountExistExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	switch cmd.CommandType() {
	case DeleteAccountCommand:
		err = o.DeleteHandler(cmd.(*DeleteAccount), account, store)
	case UpdateAccountCommand:
		err = o.UpdateHandler(cmd.(*UpdateAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
	return
}

func (o *AccountExistExecutor) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
		return
	}

	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountDeletedEvent, nil, time.Now())
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

type AccountInitialExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
	CreateHandler    func(*CreateAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountInitialExecutorDefault() (ret *AccountInitialExecutor) {
	ret = &AccountInitialExecutor{}
	return
}

func (o *AccountInitialExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
	prevHandler := o.CommandsPreparer
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if err = preparer(cmd, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(cmd, entity)
			}
		}
		return
	}
}

func (o *AccountInitialExecutor) AddCreatePreparer(preparer func(*CreateAccount, *Account) (err error)) {
	prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountInitialExecutor) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
	account := entity.(*Account)
	if err = o.CommandsPreparer(cmd, account); err != nil {
		return
	}
	switch cmd.CommandType() {
	case CreateAccountCommand:
		err = o.CreateHandler(cmd.(*CreateAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
	return
}

func (o *AccountInitialExecutor) SetupCommandHandler() (err error) {
	o.CommandsPreparer = func(cmd eventhorizon.Command, entity *Account) (err error) {
		if entity.DeletedAt != nil {
			err = eh.CommandError{Err: eh.ErrAggregateDeleted, Cmd: cmd, Entity: entity}
		}
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
	return
}
