package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-ee/utils/eh"
	"github.com/go-ee/utils/enum"
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type AccountAggregateHandlers struct {
	Initial        *AccountAggregateInitialHandler
	Deleted        *AccountAggregateDeletedHandler
	Disabled       *AccountAggregateDisabledHandler
	Enabled        *AccountAggregateEnabledHandler
	Exist          *AccountAggregateExistHandler
	EventsPreparer func(eventhorizon.Event, *Account) (err error)
}

func NewAccountAggregateHandlersFull() (ret *AccountAggregateHandlers) {
	initial := NewAccountAggregateInitialHandlerDefault()
	deleted := NewAccountAggregateDeletedHandlerDefault()
	disabled := NewAccountAggregateDisabledHandlerDefault()
	enabled := NewAccountAggregateEnabledHandlerDefault()
	exist := NewAccountAggregateExistHandlerDefault()
	ret = &AccountAggregateHandlers{
		Initial:  initial,
		Deleted:  deleted,
		Disabled: disabled,
		Enabled:  enabled,
		Exist:    exist,
	}
	return
}

func (o *AccountAggregateHandlers) AddEventsPreparer(preparer func(eventhorizon.Event, *Account) (err error)) {
	prevHandler := o.EventsPreparer
	o.EventsPreparer = func(event eventhorizon.Event, entity *Account) (err error) {
		if err = preparer(event, entity); err == nil {
			if prevHandler != nil {
				err = prevHandler(event, entity)
			}
		}
		return
	}
}

func (o *AccountAggregateHandlers) Apply(event eventhorizon.Event, account *Account) (err error) {

	currentAggregateState := account.AggregateState
	if currentAggregateState == nil {
		currentAggregateState = AccountAggregateStateTypes().Initial()
	}

	var newAggregateState *AccountAggregateStateType
	switch currentAggregateState {
	case AccountAggregateStateTypes().Initial():
		newAggregateState, err = o.Initial.Apply(event, account)
	case AccountAggregateStateTypes().Deleted():
		newAggregateState, err = o.Deleted.Apply(event, account)
	case AccountAggregateStateTypes().Disabled():
		newAggregateState, err = o.Disabled.Apply(event, account)
	case AccountAggregateStateTypes().Enabled():
		newAggregateState, err = o.Enabled.Apply(event, account)
	case AccountAggregateStateTypes().Exist():
		newAggregateState, err = o.Exist.Apply(event, account)
	default:
		err = errors.New(fmt.Sprintf("Not supported AggregateState '%v' for entity '%v", account.AggregateState, account))
	}

	if err == nil && newAggregateState != account.AggregateState {
		account.AggregateState = newAggregateState
	}
	return
}

func (o *AccountAggregateHandlers) SetupEventHandler() (err error) {
	if err = o.Initial.SetupEventHandler(); err != nil {
		return
	}
	if err = o.Deleted.SetupEventHandler(); err != nil {
		return
	}
	if err = o.Disabled.SetupEventHandler(); err != nil {
		return
	}
	if err = o.Enabled.SetupEventHandler(); err != nil {
		return
	}
	if err = o.Exist.SetupEventHandler(); err != nil {
		return
	}
	return
}

type AccountAggregateExecutors struct {
	Initial          *AccountAggregateInitialExecutor
	Deleted          *AccountAggregateDeletedExecutor
	Disabled         *AccountAggregateDisabledExecutor
	Enabled          *AccountAggregateEnabledExecutor
	Exist            *AccountAggregateExistExecutor
	CommandsPreparer func(eventhorizon.Command, *Account) (err error)
}

func NewAccountAggregateExecutorsFull() (ret *AccountAggregateExecutors) {
	initial := NewAccountAggregateInitialExecutorDefault()
	deleted := NewAccountAggregateDeletedExecutorDefault()
	disabled := NewAccountAggregateDisabledExecutorDefault()
	enabled := NewAccountAggregateEnabledExecutorDefault()
	exist := NewAccountAggregateExistExecutorDefault()
	ret = &AccountAggregateExecutors{
		Initial:  initial,
		Deleted:  deleted,
		Disabled: disabled,
		Enabled:  enabled,
		Exist:    exist,
	}
	return
}

func (o *AccountAggregateExecutors) AddCommandsPreparer(preparer func(eventhorizon.Command, *Account) (err error)) {
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

func (o *AccountAggregateExecutors) Execute(cmd eventhorizon.Command, account *Account, store eh.AggregateStoreEvent) (err error) {

	currentAggregateState := account.AggregateState
	if currentAggregateState == nil {
		currentAggregateState = AccountAggregateStateTypes().Initial()
	}

	switch currentAggregateState {
	case AccountAggregateStateTypes().Initial():
		err = o.Initial.Execute(cmd, account, store)
	case AccountAggregateStateTypes().Deleted():
		err = o.Deleted.Execute(cmd, account, store)
	case AccountAggregateStateTypes().Disabled():
		err = o.Disabled.Execute(cmd, account, store)
	case AccountAggregateStateTypes().Enabled():
		err = o.Enabled.Execute(cmd, account, store)
	case AccountAggregateStateTypes().Exist():
		err = o.Exist.Execute(cmd, account, store)
	default:
		err = errors.New(fmt.Sprintf("Not supported state '%v' for entity '%v", account.AggregateState, account))
	}
	return
}

func (o *AccountAggregateExecutors) SetupCommandHandler() (err error) {
	if err = o.Initial.SetupCommandHandler(); err != nil {
		return
	}
	if err = o.Deleted.SetupCommandHandler(); err != nil {
		return
	}
	if err = o.Disabled.SetupCommandHandler(); err != nil {
		return
	}
	if err = o.Enabled.SetupCommandHandler(); err != nil {
		return
	}
	if err = o.Exist.SetupCommandHandler(); err != nil {
		return
	}
	return
}

type AccountAggregate struct {
	*events.AggregateBase
	Account            *Account
	AggregateExecutors *AccountAggregateExecutors
	AggregateHandlers  *AccountAggregateHandlers
}

func NewAccountAggregateFull(aggregateBase *events.AggregateBase, account *Account, aggregateExecutors *AccountAggregateExecutors,
	aggregateHandlers *AccountAggregateHandlers) (ret *AccountAggregate) {
	ret = &AccountAggregate{
		AggregateBase:      aggregateBase,
		Account:            account,
		AggregateExecutors: aggregateExecutors,
		AggregateHandlers:  aggregateHandlers,
	}
	return
}

func (o *AccountAggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) (err error) {
	err = o.AggregateHandlers.Apply(event, o.Account)
	return
}

func (o *AccountAggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) (err error) {
	err = o.AggregateExecutors.Execute(cmd, o.Account, o.AggregateBase)
	return
}

type AccountAggregateStateType struct {
	name    string
	ordinal int
}

func (o *AccountAggregateStateType) Name() string {
	return o.name
}

func (o *AccountAggregateStateType) Ordinal() int {
	return o.ordinal
}

func (o *AccountAggregateStateType) IsInitial() bool {
	return o.name == _accountAggregateStateTypes.Initial().name
}

func (o *AccountAggregateStateType) IsDeleted() bool {
	return o.name == _accountAggregateStateTypes.Deleted().name
}

func (o *AccountAggregateStateType) IsDisabled() bool {
	return o.name == _accountAggregateStateTypes.Disabled().name
}

func (o *AccountAggregateStateType) IsEnabled() bool {
	return o.name == _accountAggregateStateTypes.Enabled().name
}

func (o *AccountAggregateStateType) IsExist() bool {
	return o.name == _accountAggregateStateTypes.Exist().name
}

func (o *AccountAggregateStateType) MarshalJSON() (ret []byte, err error) {
	ret = []byte(fmt.Sprintf("\"%v\"", o.name))
	return
}

func (o *AccountAggregateStateType) UnmarshalJSON(data []byte) (err error) {
	name := string(data)
	//remove quotes
	name = name[1 : len(name)-1]
	if v, ok := AccountAggregateStateTypes().ParseAccountAggregateStateType(name); ok {
		*o = *v
	} else {
		err = fmt.Errorf("invalid AccountAggregateStateType %q", name)
	}
	return
}

func (o *AccountAggregateStateType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AccountAggregateStateType) SetBSON(raw bson.Raw) (err error) {
	var lit string
	if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AccountAggregateStateTypes().ParseAccountAggregateStateType(lit); ok {
			*o = *v
		} else {
			err = fmt.Errorf("invalid AccountAggregateStateType %q", lit)
		}
	}
	return
}

type accountAggregateStateTypes struct {
	values           []*AccountAggregateStateType
	valuesAsLiterals []enum.Literal
}

var _accountAggregateStateTypes = &accountAggregateStateTypes{values: []*AccountAggregateStateType{
	{name: "Initial", ordinal: 0},
	{name: "Deleted", ordinal: 1},
	{name: "Disabled", ordinal: 2},
	{name: "Enabled", ordinal: 3},
	{name: "Exist", ordinal: 4}},
}

func AccountAggregateStateTypes() *accountAggregateStateTypes {
	return _accountAggregateStateTypes
}

func (o *accountAggregateStateTypes) Values() []*AccountAggregateStateType {
	return o.values
}

func (o *accountAggregateStateTypes) Initial() *AccountAggregateStateType {
	return o.values[0]
}

func (o *accountAggregateStateTypes) Deleted() *AccountAggregateStateType {
	return o.values[1]
}

func (o *accountAggregateStateTypes) Disabled() *AccountAggregateStateType {
	return o.values[2]
}

func (o *accountAggregateStateTypes) Enabled() *AccountAggregateStateType {
	return o.values[3]
}

func (o *accountAggregateStateTypes) Exist() *AccountAggregateStateType {
	return o.values[4]
}

func (o *accountAggregateStateTypes) ParseAccountAggregateStateType(name string) (ret *AccountAggregateStateType, ok bool) {
	for _, lit := range o.Values() {
		if strings.EqualFold(lit.Name(), name) {
			return lit, true
		}
	}
	return nil, false
}

// we have to convert the instances to Literal interface, because it is not a other way in Go
func (o *accountAggregateStateTypes) Literals() []enum.Literal {
	if o.valuesAsLiterals == nil {
		o.valuesAsLiterals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.valuesAsLiterals[i] = item
		}
	}
	return o.valuesAsLiterals
}
