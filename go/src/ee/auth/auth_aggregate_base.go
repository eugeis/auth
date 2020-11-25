package auth

import (
    "errors"
    "fmt"
    "github.com/go-ee/utils/eh"
    "github.com/google/uuid"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type AccountCommandHandler struct {
    SendEnabledConfirmationHandler func (*SendEnabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
    SendDisabledConfirmationHandler func (*SendDisabledConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
    LoginHandler func (*LoginAccount, *Account, eh.AggregateStoreEvent) (err error)
    SendCreatedConfirmationHandler func (*SendCreatedConfirmationAccount, *Account, eh.AggregateStoreEvent) (err error)
    CreateHandler func (*CreateAccount, *Account, eh.AggregateStoreEvent) (err error)
    DeleteHandler func (*DeleteAccount, *Account, eh.AggregateStoreEvent) (err error)
    EnableHandler func (*EnableAccount, *Account, eh.AggregateStoreEvent) (err error)
    DisableHandler func (*DisableAccount, *Account, eh.AggregateStoreEvent) (err error)
    UpdateHandler func (*UpdateAccount, *Account, eh.AggregateStoreEvent) (err error)
}

func (o *AccountCommandHandler) AddSendEnabledConfirmationPreparer(preparer func (*SendEnabledConfirmationAccount, *Account) (err error)) {
    prevHandler := o.SendEnabledConfirmationHandler
	o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddSendDisabledConfirmationPreparer(preparer func (*SendDisabledConfirmationAccount, *Account) (err error)) {
    prevHandler := o.SendDisabledConfirmationHandler
	o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddLoginPreparer(preparer func (*LoginAccount, *Account) (err error)) {
    prevHandler := o.LoginHandler
	o.LoginHandler = func(command *LoginAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddSendCreatedConfirmationPreparer(preparer func (*SendCreatedConfirmationAccount, *Account) (err error)) {
    prevHandler := o.SendCreatedConfirmationHandler
	o.SendCreatedConfirmationHandler = func(command *SendCreatedConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddCreatePreparer(preparer func (*CreateAccount, *Account) (err error)) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddDeletePreparer(preparer func (*DeleteAccount, *Account) (err error)) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddEnablePreparer(preparer func (*EnableAccount, *Account) (err error)) {
    prevHandler := o.EnableHandler
	o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddDisablePreparer(preparer func (*DisableAccount, *Account) (err error)) {
    prevHandler := o.DisableHandler
	o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) AddUpdatePreparer(preparer func (*UpdateAccount, *Account) (err error)) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *AccountCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error){
    switch cmd.CommandType() {
    case SendEnabledConfirmationAccountCommand:
        err = o.SendEnabledConfirmationHandler(cmd.(*SendEnabledConfirmationAccount), entity.(*Account), store)
    case SendDisabledConfirmationAccountCommand:
        err = o.SendDisabledConfirmationHandler(cmd.(*SendDisabledConfirmationAccount), entity.(*Account), store)
    case LoginAccountCommand:
        err = o.LoginHandler(cmd.(*LoginAccount), entity.(*Account), store)
    case SendCreatedConfirmationAccountCommand:
        err = o.SendCreatedConfirmationHandler(cmd.(*SendCreatedConfirmationAccount), entity.(*Account), store)
    case CreateAccountCommand:
        err = o.CreateHandler(cmd.(*CreateAccount), entity.(*Account), store)
    case DeleteAccountCommand:
        err = o.DeleteHandler(cmd.(*DeleteAccount), entity.(*Account), store)
    case EnableAccountCommand:
        err = o.EnableHandler(cmd.(*EnableAccount), entity.(*Account), store)
    case DisableAccountCommand:
        err = o.DisableHandler(cmd.(*DisableAccount), entity.(*Account), store)
    case UpdateAccountCommand:
        err = o.UpdateHandler(cmd.(*UpdateAccount), entity.(*Account), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *AccountCommandHandler) SetupCommandHandler() (err error){
    o.SendEnabledConfirmationHandler = func(command *SendEnabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountSentEnabledConfirmationEvent, &AccountSentEnabledConfirmation{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendDisabledConfirmationHandler = func(command *SendDisabledConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountSentDisabledConfirmationEvent, &AccountSentDisabledConfirmation{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.LoginHandler = func(command *LoginAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountLoggedEvent, &AccountLogged{
                Username: command.Username,
                Email: command.Email,
                Password: command.Password,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.SendCreatedConfirmationHandler = func(command *SendCreatedConfirmationAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountSentCreatedConfirmationEvent, &AccountSentCreatedConfirmation{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.CreateHandler = func(command *CreateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountCreatedEvent, &AccountCreated{
                Name: command.Name,
                Username: command.Username,
                Password: command.Password,
                Email: command.Email,
                Roles: command.Roles,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountDeletedEvent, &AccountDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.EnableHandler = func(command *EnableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountEnabledEvent, &AccountEnabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DisableHandler = func(command *DisableAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountDisabledEvent, &AccountDisabled{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateAccount, entity *Account, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, AccountAggregateType); err == nil {
            store.AppendEvent(AccountUpdatedEvent, &AccountUpdated{
                Name: command.Name,
                Username: command.Username,
                Password: command.Password,
                Email: command.Email,
                Roles: command.Roles,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type AccountEventHandler struct {
    CreatedHandler func (*AccountCreated, *Account) (err error)
    DeletedHandler func (*AccountDeleted, *Account) (err error)
    DisabledHandler func (*AccountDisabled, *Account) (err error)
    EnabledHandler func (*AccountEnabled, *Account) (err error)
    LoggedHandler func (*AccountLogged, *Account) (err error)
    SentCreatedConfirmationHandler func (*AccountSentCreatedConfirmation, *Account) (err error)
    SentDisabledConfirmationHandler func (*AccountSentDisabledConfirmation, *Account) (err error)
    SentEnabledConfirmationHandler func (*AccountSentEnabledConfirmation, *Account) (err error)
    UpdatedHandler func (*AccountUpdated, *Account) (err error)
}

func (o *AccountEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error){
    switch event.EventType() {
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    case AccountEnabledEvent:
        err = o.EnabledHandler(event.Data().(*AccountEnabled), entity.(*Account))
    case AccountDisabledEvent:
        err = o.DisabledHandler(event.Data().(*AccountDisabled), entity.(*Account))
    case AccountUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*AccountUpdated), entity.(*Account))
    case AccountDeletedEvent:
        err = o.DeletedHandler(event.Data().(*AccountDeleted), entity.(*Account))
    case AccountSentEnabledConfirmationEvent:
        err = o.SentEnabledConfirmationHandler(event.Data().(*AccountSentEnabledConfirmation), entity.(*Account))
    case AccountSentDisabledConfirmationEvent:
        err = o.SentDisabledConfirmationHandler(event.Data().(*AccountSentDisabledConfirmation), entity.(*Account))
    case AccountLoggedEvent:
        err = o.LoggedHandler(event.Data().(*AccountLogged), entity.(*Account))
    case AccountSentCreatedConfirmationEvent:
        err = o.SentCreatedConfirmationHandler(event.Data().(*AccountSentCreatedConfirmation), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountEventHandler) SetupEventHandler() (err error){

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
    eventhorizon.RegisterEventData(AccountSentEnabledConfirmationEvent, func() eventhorizon.EventData {
		return &AccountSentEnabledConfirmation{}
	})

    //default handler implementation
    o.SentEnabledConfirmationHandler = func(event *AccountSentEnabledConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(AccountSentEnabledConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountSentDisabledConfirmationEvent, func() eventhorizon.EventData {
		return &AccountSentDisabledConfirmation{}
	})

    //default handler implementation
    o.SentDisabledConfirmationHandler = func(event *AccountSentDisabledConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(AccountSentDisabledConfirmationEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountLoggedEvent, func() eventhorizon.EventData {
		return &AccountLogged{}
	})

    //default handler implementation
    o.LoggedHandler = func(event *AccountLogged, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(AccountLoggedEvent)
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountSentCreatedConfirmationEvent, func() eventhorizon.EventData {
		return &AccountSentCreatedConfirmation{}
	})

    //default handler implementation
    o.SentCreatedConfirmationHandler = func(event *AccountSentCreatedConfirmation, entity *Account) (err error) {
        //err = eh.EventHandlerNotImplemented(AccountSentCreatedConfirmationEvent)
        return
    }
    return
}


const AccountAggregateType eventhorizon.AggregateType = "Account"

type AccountAggregateInitializer struct {
    *eh.AggregateInitializer
    *AccountCommandHandler
    *AccountEventHandler
    ProjectorHandler *AccountEventHandler
}


func (o *AccountAggregateInitializer) RegisterForSentEnabledConfirmation(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountSentEnabledConfirmation())
}

func (o *AccountAggregateInitializer) RegisterForSentDisabledConfirmation(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountSentDisabledConfirmation())
}

func (o *AccountAggregateInitializer) RegisterForLogged(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountLogged())
}

func (o *AccountAggregateInitializer) RegisterForSentCreatedConfirmation(handler eventhorizon.EventHandler){
    o.RegisterForEvent(handler, AccountEventTypes().AccountSentCreatedConfirmation())
}


func NewAccountAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity)) (ret eventhorizon.ReadWriteRepo)) (ret *AccountAggregateInitializer) {
    
    commandHandler := &AccountCommandHandler{}
    eventHandler := &AccountEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewAccountDefault() }
    ret = &AccountAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(AccountAggregateType,
        func(id uuid.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(AccountAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        AccountCommandTypes().Literals(), AccountEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, commandBus, readRepos), AccountCommandHandler: commandHandler, AccountEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type AuthEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore
    eventBus eventhorizon.EventBus
    commandBus *bus.CommandHandler
    AccountAggregateInitializer *AccountAggregateInitializer
}

func NewAuthEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity)) (ret eventhorizon.ReadWriteRepo)) (ret *AuthEventhorizonInitializer) {
    accountAggregateInitializer := NewAccountAggregateInitializer(eventStore, eventBus, commandBus, readRepos)
    ret = &AuthEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        commandBus: commandBus,
        AccountAggregateInitializer: accountAggregateInitializer,
    }
    return
}

func (o *AuthEventhorizonInitializer) Setup() (err error){
    
    if err = o.AccountAggregateInitializer.Setup(); err != nil {
        return
    }

    return
}









