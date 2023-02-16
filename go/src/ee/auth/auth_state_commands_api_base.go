package auth

import (
	"errors"
	"fmt"
	"github.com/go-ee/utils/eh"
	"github.com/looplab/eventhorizon"
	"time"
)

type AccountAggregateInitialExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
	CreateHandler    func(*CreateAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountAggregateInitialExecutorDefault() (ret *AccountAggregateInitialExecutor) {
	ret = &AccountAggregateInitialExecutor{}
	return
}

func (o *AccountAggregateInitialExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
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

func (o *AccountAggregateInitialExecutor) AddCreatePreparer(preparer func(*CreateAccount, *Account) (err error)) {
	prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateInitialExecutor) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Initial()
	return
}

func (o *AccountAggregateInitialExecutor) Execute(cmd eventhorizon.Command, account *Account, store eh.AggregateStoreEvent) (err error) {
	if o.CommandsPreparer != nil {
		if err = o.CommandsPreparer(cmd, account); err != nil {
			return
		}
	}

	switch cmd.CommandType() {
	case CreateAccountCommand:
		err = o.CreateHandler(cmd.(*CreateAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' in state 'Initial' for entity '%v", cmd.CommandType(), account))
	}
	return
}

func (o *AccountAggregateInitialExecutor) SetupCommandHandler() (err error) {
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

type AccountAggregateDeletedExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
}

func NewAccountAggregateDeletedExecutorDefault() (ret *AccountAggregateDeletedExecutor) {
	ret = &AccountAggregateDeletedExecutor{}
	return
}

func (o *AccountAggregateDeletedExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
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

func (o *AccountAggregateDeletedExecutor) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Deleted()
	return
}

func (o *AccountAggregateDeletedExecutor) Execute(cmd eventhorizon.Command, account *Account, store eh.AggregateStoreEvent) (err error) {
	if o.CommandsPreparer != nil {
		if err = o.CommandsPreparer(cmd, account); err != nil {
			return
		}
	}
	err = errors.New(fmt.Sprintf("Not supported command type '%v' in state 'Deleted' for entity '%v", cmd.CommandType(), account))
	return
}

func (o *AccountAggregateDeletedExecutor) SetupCommandHandler() (err error) {
	return
}

type AccountAggregateDisabledExecutor struct {
	CommandsPreparer                func(eventhorizon.Command, *Account) (err error)
	EnableHandler                   func(*EnableAccount, *Account, eh.AggregateStoreEvent) (err error)
	SendDisabledConfirmationHandler func(*SendDisabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountAggregateDisabledExecutorDefault() (ret *AccountAggregateDisabledExecutor) {
	ret = &AccountAggregateDisabledExecutor{}
	return
}

func (o *AccountAggregateDisabledExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
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

func (o *AccountAggregateDisabledExecutor) AddEnablePreparer(preparer func(*EnableAccount, *Account) (err error)) {
	prevHandler := o.EnableHandler
	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateDisabledExecutor) AddSendDisabledConfirmationPreparer(preparer func(*SendDisabledConfirmationAccount, *Account) (err error)) {
	prevHandler := o.SendDisabledConfirmationHandler
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateDisabledExecutor) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Disabled()
	return
}

func (o *AccountAggregateDisabledExecutor) Execute(cmd eventhorizon.Command, account *Account, store eh.AggregateStoreEvent) (err error) {
	if o.CommandsPreparer != nil {
		if err = o.CommandsPreparer(cmd, account); err != nil {
			return
		}
	}

	switch cmd.CommandType() {
	case EnableAccountCommand:
		err = o.EnableHandler(cmd.(*EnableAccount), account, store)
	case SendDisabledConfirmationAccountCommand:
		err = o.SendDisabledConfirmationHandler(cmd.(*SendDisabledConfirmationAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' in state 'Disabled' for entity '%v", cmd.CommandType(), account))
	}
	return
}

func (o *AccountAggregateDisabledExecutor) SetupCommandHandler() (err error) {
	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountEnabledEvent, nil, time.Now())
		return
	}
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountSentDisabledConfirmationEvent, nil, time.Now())
		return
	}
	return
}

type AccountAggregateEnabledExecutor struct {
	CommandsPreparer               func(eventhorizon.Command, *Account) (err error)
	DeleteHandler                  func(*DeleteAccount, *Account, eh.AggregateStoreEvent) (err error)
	DisableHandler                 func(*DisableAccount, *Account, eh.AggregateStoreEvent) (err error)
	SendEnabledConfirmationHandler func(*SendEnabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountAggregateEnabledExecutorDefault() (ret *AccountAggregateEnabledExecutor) {
	ret = &AccountAggregateEnabledExecutor{}
	return
}

func (o *AccountAggregateEnabledExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
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

func (o *AccountAggregateEnabledExecutor) AddDeletePreparer(preparer func(*DeleteAccount, *Account) (err error)) {
	prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateEnabledExecutor) AddDisablePreparer(preparer func(*DisableAccount, *Account) (err error)) {
	prevHandler := o.DisableHandler
	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateEnabledExecutor) AddSendEnabledConfirmationPreparer(preparer func(*SendEnabledConfirmationAccount, *Account) (err error)) {
	prevHandler := o.SendEnabledConfirmationHandler
	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateEnabledExecutor) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Enabled()
	return
}

func (o *AccountAggregateEnabledExecutor) Execute(cmd eventhorizon.Command, account *Account, store eh.AggregateStoreEvent) (err error) {
	if o.CommandsPreparer != nil {
		if err = o.CommandsPreparer(cmd, account); err != nil {
			return
		}
	}

	switch cmd.CommandType() {
	case DeleteAccountCommand:
		err = o.DeleteHandler(cmd.(*DeleteAccount), account, store)
	case DisableAccountCommand:
		err = o.DisableHandler(cmd.(*DisableAccount), account, store)
	case SendEnabledConfirmationAccountCommand:
		err = o.SendEnabledConfirmationHandler(cmd.(*SendEnabledConfirmationAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' in state 'Enabled' for entity '%v", cmd.CommandType(), account))
	}
	return
}

func (o *AccountAggregateEnabledExecutor) SetupCommandHandler() (err error) {
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountDeletedEvent, nil, time.Now())
		return
	}
	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountDisabledEvent, nil, time.Now())
		return
	}
	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		store.AppendEvent(AccountSentEnabledConfirmationEvent, nil, time.Now())
		return
	}
	return
}

type AccountAggregateExistExecutor struct {
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
	DeleteHandler    func(*DeleteAccount, *Account, eh.AggregateStoreEvent) (err error)
	UpdateHandler    func(*UpdateAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func NewAccountAggregateExistExecutorDefault() (ret *AccountAggregateExistExecutor) {
	ret = &AccountAggregateExistExecutor{}
	return
}

func (o *AccountAggregateExistExecutor) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
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

func (o *AccountAggregateExistExecutor) AddDeletePreparer(preparer func(*DeleteAccount, *Account) (err error)) {
	prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateExistExecutor) AddUpdatePreparer(preparer func(*UpdateAccount, *Account) (err error)) {
	prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountAggregateExistExecutor) StateType() (ret *AccountAggregateStateType) {
	ret = AccountAggregateStateTypes().Exist()
	return
}

func (o *AccountAggregateExistExecutor) Execute(cmd eventhorizon.Command, account *Account, store eh.AggregateStoreEvent) (err error) {
	if o.CommandsPreparer != nil {
		if err = o.CommandsPreparer(cmd, account); err != nil {
			return
		}
	}

	switch cmd.CommandType() {
	case DeleteAccountCommand:
		err = o.DeleteHandler(cmd.(*DeleteAccount), account, store)
	case UpdateAccountCommand:
		err = o.UpdateHandler(cmd.(*UpdateAccount), account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' in state 'Exist' for entity '%v", cmd.CommandType(), account))
	}
	return
}

func (o *AccountAggregateExistExecutor) SetupCommandHandler() (err error) {
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
