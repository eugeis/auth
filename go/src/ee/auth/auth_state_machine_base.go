package auth

import (
    "errors"
    "fmt"
    "github.com/go-ee/utils/eh"
    "github.com/looplab/eventhorizon"
)
type AccountConfirmationDisabledHandler struct {
    EnabledHandler func (*AccountEnabled, *Account) (err error)
}

func NewAccountConfirmationDisabledHandlerDefault() (ret *AccountConfirmationDisabledHandler) {
    ret = &AccountConfirmationDisabledHandler{}
    return
}

func (o *AccountConfirmationDisabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountEnabledEvent:
        err = o.EnabledHandler(event.Data().(*AccountEnabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountConfirmationDisabledHandler) SetupEventHandler() (err error){

    //register event object factory
    eventhorizon.RegisterEventData(AccountEnabledEvent, func() eventhorizon.EventData {
		return &AccountEnabled{}
	})

    //default handler implementation
    o.EnabledHandler = func(event *AccountEnabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountConfirmationDisabledExecutor struct {
}

func NewAccountConfirmationDisabledExecutorDefault() (ret *AccountConfirmationDisabledExecutor) {
    ret = &AccountConfirmationDisabledExecutor{}
    return
}


type AccountConfirmationEnabledHandler struct {
    DisabledHandler func (*AccountDisabled, *Account) (err error)
}

func NewAccountConfirmationEnabledHandlerDefault() (ret *AccountConfirmationEnabledHandler) {
    ret = &AccountConfirmationEnabledHandler{}
    return
}

func (o *AccountConfirmationEnabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountDisabledEvent:
        err = o.DisabledHandler(event.Data().(*AccountDisabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountConfirmationEnabledHandler) SetupEventHandler() (err error){

    //register event object factory
    eventhorizon.RegisterEventData(AccountDisabledEvent, func() eventhorizon.EventData {
		return &AccountDisabled{}
	})

    //default handler implementation
    o.DisabledHandler = func(event *AccountDisabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountConfirmationEnabledExecutor struct {
}

func NewAccountConfirmationEnabledExecutorDefault() (ret *AccountConfirmationEnabledExecutor) {
    ret = &AccountConfirmationEnabledExecutor{}
    return
}


type AccountConfirmationInitialHandler struct {
    CreatedHandler func (*AccountCreated, *Account) (err error)
}

func NewAccountConfirmationInitialHandlerDefault() (ret *AccountConfirmationInitialHandler) {
    ret = &AccountConfirmationInitialHandler{}
    return
}

func (o *AccountConfirmationInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountConfirmationInitialHandler) SetupEventHandler() (err error){

    //register event object factory
    eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AccountCreated, entity *Account) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
            entity.Id = event.Id
        }
        return
    }
    return
}


type AccountConfirmationInitialExecutor struct {
}

func NewAccountConfirmationInitialExecutorDefault() (ret *AccountConfirmationInitialExecutor) {
    ret = &AccountConfirmationInitialExecutor{}
    return
}


type AccountConfirmationHandlers struct {
    Disabled *AccountConfirmationDisabledHandler
    Enabled *AccountConfirmationEnabledHandler
    Initial *AccountConfirmationInitialHandler
}

func NewAccountConfirmationHandlersFull() (ret *AccountConfirmationHandlers) {
    disabled := NewAccountConfirmationDisabledHandlerDefault()
    enabled := NewAccountConfirmationEnabledHandlerDefault()
    initial := NewAccountConfirmationInitialHandlerDefault()
    ret = &AccountConfirmationHandlers{
        Disabled: disabled,
        Enabled: enabled,
        Initial: initial,
    }
    return
}


type AccountConfirmationExecutors struct {
    Disabled *AccountConfirmationDisabledExecutor
    Enabled *AccountConfirmationEnabledExecutor
    Initial *AccountConfirmationInitialExecutor
}

func NewAccountConfirmationExecutorsFull() (ret *AccountConfirmationExecutors) {
    disabled := NewAccountConfirmationDisabledExecutorDefault()
    enabled := NewAccountConfirmationEnabledExecutorDefault()
    initial := NewAccountConfirmationInitialExecutorDefault()
    ret = &AccountConfirmationExecutors{
        Disabled: disabled,
        Enabled: enabled,
        Initial: initial,
    }
    return
}


type AccountDeletedHandler struct {
}

func NewAccountDeletedHandlerDefault() (ret *AccountDeletedHandler) {
    ret = &AccountDeletedHandler{}
    return
}

func (o *AccountDeletedHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    
    return
}

func (o *AccountDeletedHandler) SetupEventHandler() (err error){
    return
}


type AccountDeletedExecutor struct {
}

func NewAccountDeletedExecutorDefault() (ret *AccountDeletedExecutor) {
    ret = &AccountDeletedExecutor{}
    return
}


type AccountDisabledHandler struct {
    EnabledHandler func (*AccountEnabled, *Account) (err error)
}

func NewAccountDisabledHandlerDefault() (ret *AccountDisabledHandler) {
    ret = &AccountDisabledHandler{}
    return
}

func (o *AccountDisabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountEnabledEvent:
        err = o.EnabledHandler(event.Data().(*AccountEnabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountDisabledHandler) SetupEventHandler() (err error){

    //register event object factory
    eventhorizon.RegisterEventData(AccountEnabledEvent, func() eventhorizon.EventData {
		return &AccountEnabled{}
	})

    //default handler implementation
    o.EnabledHandler = func(event *AccountEnabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountDisabledExecutor struct {
}

func NewAccountDisabledExecutorDefault() (ret *AccountDisabledExecutor) {
    ret = &AccountDisabledExecutor{}
    return
}


type AccountEnabledHandler struct {
    DisabledHandler func (*AccountDisabled, *Account) (err error)
}

func NewAccountEnabledHandlerDefault() (ret *AccountEnabledHandler) {
    ret = &AccountEnabledHandler{}
    return
}

func (o *AccountEnabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountDisabledEvent:
        err = o.DisabledHandler(event.Data().(*AccountDisabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountEnabledHandler) SetupEventHandler() (err error){

    //register event object factory
    eventhorizon.RegisterEventData(AccountDisabledEvent, func() eventhorizon.EventData {
		return &AccountDisabled{}
	})

    //default handler implementation
    o.DisabledHandler = func(event *AccountDisabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountEnabledExecutor struct {
}

func NewAccountEnabledExecutorDefault() (ret *AccountEnabledExecutor) {
    ret = &AccountEnabledExecutor{}
    return
}


type AccountExistHandler struct {
    DeletedHandler func (*AccountDeleted, *Account) (err error)
    UpdatedHandler func (*AccountUpdated, *Account) (err error)
}

func NewAccountExistHandlerDefault() (ret *AccountExistHandler) {
    ret = &AccountExistHandler{}
    return
}

func (o *AccountExistHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountDeletedEvent:
        err = o.DeletedHandler(event.Data().(*AccountDeleted), entity.(*Account))
    case AccountUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*AccountUpdated), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountExistHandler) SetupEventHandler() (err error){

    //register event object factory
    eventhorizon.RegisterEventData(AccountDeletedEvent, func() eventhorizon.EventData {
		return &AccountDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *AccountDeleted, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            *entity = *NewAccountDefault()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountUpdatedEvent, func() eventhorizon.EventData {
		return &AccountUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *AccountUpdated, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
        }
        return
    }
    return
}


type AccountExistExecutor struct {
}

func NewAccountExistExecutorDefault() (ret *AccountExistExecutor) {
    ret = &AccountExistExecutor{}
    return
}


type AccountInitialHandler struct {
    CreatedHandler func (*AccountCreated, *Account) (err error)
}

func NewAccountInitialHandlerDefault() (ret *AccountInitialHandler) {
    ret = &AccountInitialHandler{}
    return
}

func (o *AccountInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountInitialHandler) SetupEventHandler() (err error){

    //register event object factory
    eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AccountCreated, entity *Account) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
            entity.Id = event.Id
        }
        return
    }
    return
}


type AccountInitialExecutor struct {
}

func NewAccountInitialExecutorDefault() (ret *AccountInitialExecutor) {
    ret = &AccountInitialExecutor{}
    return
}


type AccountHandlers struct {
    Deleted *AccountDeletedHandler
    Disabled *AccountDisabledHandler
    Enabled *AccountEnabledHandler
    Exist *AccountExistHandler
    Initial *AccountInitialHandler
}

func NewAccountHandlersFull() (ret *AccountHandlers) {
    deleted := NewAccountDeletedHandlerDefault()
    disabled := NewAccountDisabledHandlerDefault()
    enabled := NewAccountEnabledHandlerDefault()
    exist := NewAccountExistHandlerDefault()
    initial := NewAccountInitialHandlerDefault()
    ret = &AccountHandlers{
        Deleted: deleted,
        Disabled: disabled,
        Enabled: enabled,
        Exist: exist,
        Initial: initial,
    }
    return
}


type AccountExecutors struct {
    Deleted *AccountDeletedExecutor
    Disabled *AccountDisabledExecutor
    Enabled *AccountEnabledExecutor
    Exist *AccountExistExecutor
    Initial *AccountInitialExecutor
}

func NewAccountExecutorsFull() (ret *AccountExecutors) {
    deleted := NewAccountDeletedExecutorDefault()
    disabled := NewAccountDisabledExecutorDefault()
    enabled := NewAccountEnabledExecutorDefault()
    exist := NewAccountExistExecutorDefault()
    initial := NewAccountInitialExecutorDefault()
    ret = &AccountExecutors{
        Deleted: deleted,
        Disabled: disabled,
        Enabled: enabled,
        Exist: exist,
        Initial: initial,
    }
    return
}









