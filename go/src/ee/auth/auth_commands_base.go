package auth

import (
	"fmt"
	"github.com/go-ee/utils/enum"
	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

const (
	SendAccountEnabledConfirmationCommand  eventhorizon.CommandType = "SendAccountEnabledConfirmation"
	SendAccountDisabledConfirmationCommand eventhorizon.CommandType = "SendAccountDisabledConfirmation"
	LoginAccountCommand                    eventhorizon.CommandType = "LoginAccount"
	SendAccountCreatedConfirmationCommand  eventhorizon.CommandType = "SendAccountCreatedConfirmation"
	CreateAccountCommand                   eventhorizon.CommandType = "CreateAccount"
	DeleteAccountCommand                   eventhorizon.CommandType = "DeleteAccount"
	EnableAccountCommand                   eventhorizon.CommandType = "EnableAccount"
	DisableAccountCommand                  eventhorizon.CommandType = "DisableAccount"
	UpdateAccountCommand                   eventhorizon.CommandType = "UpdateAccount"
)

type SendAccountEnabledConfirmation struct {
	Id uuid.UUID
}

func (o *SendAccountEnabledConfirmation) AggregateID() uuid.UUID { return o.Id }
func (o *SendAccountEnabledConfirmation) AggregateType() eventhorizon.AggregateType {
	return AccountAggregateType
}
func (o *SendAccountEnabledConfirmation) CommandType() eventhorizon.CommandType {
	return SendAccountEnabledConfirmationCommand
}

type SendAccountDisabledConfirmation struct {
	Id uuid.UUID
}

func (o *SendAccountDisabledConfirmation) AggregateID() uuid.UUID { return o.Id }
func (o *SendAccountDisabledConfirmation) AggregateType() eventhorizon.AggregateType {
	return AccountAggregateType
}
func (o *SendAccountDisabledConfirmation) CommandType() eventhorizon.CommandType {
	return SendAccountDisabledConfirmationCommand
}

type LoginAccount struct {
	Username string
	Email    string
	Password string
	Id       uuid.UUID
}

func (o *LoginAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *LoginAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *LoginAccount) CommandType() eventhorizon.CommandType     { return LoginAccountCommand }

type SendAccountCreatedConfirmation struct {
	Id uuid.UUID
}

func (o *SendAccountCreatedConfirmation) AggregateID() uuid.UUID { return o.Id }
func (o *SendAccountCreatedConfirmation) AggregateType() eventhorizon.AggregateType {
	return AccountAggregateType
}
func (o *SendAccountCreatedConfirmation) CommandType() eventhorizon.CommandType {
	return SendAccountCreatedConfirmationCommand
}

type CreateAccount struct {
	Name     *PersonName
	Username string
	Password string
	Email    string
	Roles    []string
	Id       uuid.UUID
}

func (o *CreateAccount) AddToRoles(item string) string {
	o.Roles = append(o.Roles, item)
	return item
}
func (o *CreateAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *CreateAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *CreateAccount) CommandType() eventhorizon.CommandType     { return CreateAccountCommand }

type DeleteAccount struct {
	Id uuid.UUID
}

func (o *DeleteAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *DeleteAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *DeleteAccount) CommandType() eventhorizon.CommandType     { return DeleteAccountCommand }

type EnableAccount struct {
	Id uuid.UUID
}

func (o *EnableAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *EnableAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *EnableAccount) CommandType() eventhorizon.CommandType     { return EnableAccountCommand }

type DisableAccount struct {
	Id uuid.UUID
}

func (o *DisableAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *DisableAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *DisableAccount) CommandType() eventhorizon.CommandType     { return DisableAccountCommand }

type UpdateAccount struct {
	Name     *PersonName
	Username string
	Password string
	Email    string
	Roles    []string
	Id       uuid.UUID
}

func (o *UpdateAccount) AddToRoles(item string) string {
	o.Roles = append(o.Roles, item)
	return item
}
func (o *UpdateAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *UpdateAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *UpdateAccount) CommandType() eventhorizon.CommandType     { return UpdateAccountCommand }

type AccountCommandType struct {
	name    string
	ordinal int
}

func (o *AccountCommandType) Name() string {
	return o.name
}

func (o *AccountCommandType) Ordinal() int {
	return o.ordinal
}

func (o *AccountCommandType) IsSendAccountEnabledConfirmation() bool {
	return o.name == _accountCommandTypes.SendAccountEnabledConfirmation().name
}

func (o *AccountCommandType) IsSendAccountDisabledConfirmation() bool {
	return o.name == _accountCommandTypes.SendAccountDisabledConfirmation().name
}

func (o *AccountCommandType) IsLoginAccount() bool {
	return o.name == _accountCommandTypes.LoginAccount().name
}

func (o *AccountCommandType) IsSendAccountCreatedConfirmation() bool {
	return o.name == _accountCommandTypes.SendAccountCreatedConfirmation().name
}

func (o *AccountCommandType) IsCreateAccount() bool {
	return o.name == _accountCommandTypes.CreateAccount().name
}

func (o *AccountCommandType) IsDeleteAccount() bool {
	return o.name == _accountCommandTypes.DeleteAccount().name
}

func (o *AccountCommandType) IsEnableAccount() bool {
	return o.name == _accountCommandTypes.EnableAccount().name
}

func (o *AccountCommandType) IsDisableAccount() bool {
	return o.name == _accountCommandTypes.DisableAccount().name
}

func (o *AccountCommandType) IsUpdateAccount() bool {
	return o.name == _accountCommandTypes.UpdateAccount().name
}

func (o *AccountCommandType) MarshalJSON() (ret []byte, err error) {
	ret = []byte(fmt.Sprintf("\"%v\"", o.name))
	return
}

func (o *AccountCommandType) UnmarshalJSON(data []byte) (err error) {
	name := string(data)
	//remove quotes
	name = name[1 : len(name)-1]
	if v, ok := AccountCommandTypes().ParseAccountCommandType(name); ok {
		*o = *v
	} else {
		err = fmt.Errorf("invalid AccountCommandType %q", name)
	}
	return
}

func (o *AccountCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AccountCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
	if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AccountCommandTypes().ParseAccountCommandType(lit); ok {
			*o = *v
		} else {
			err = fmt.Errorf("invalid AccountCommandType %q", lit)
		}
	}
	return
}

type accountCommandTypes struct {
	values           []*AccountCommandType
	valuesAsLiterals []enum.Literal
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
	{name: "SendAccountEnabledConfirmation", ordinal: 0},
	{name: "SendAccountDisabledConfirmation", ordinal: 1},
	{name: "LoginAccount", ordinal: 2},
	{name: "SendAccountCreatedConfirmation", ordinal: 3},
	{name: "CreateAccount", ordinal: 4},
	{name: "DeleteAccount", ordinal: 5},
	{name: "EnableAccount", ordinal: 6},
	{name: "DisableAccount", ordinal: 7},
	{name: "UpdateAccount", ordinal: 8}},
}

func AccountCommandTypes() *accountCommandTypes {
	return _accountCommandTypes
}

func (o *accountCommandTypes) Values() []*AccountCommandType {
	return o.values
}

func (o *accountCommandTypes) SendAccountEnabledConfirmation() *AccountCommandType {
	return o.values[0]
}

func (o *accountCommandTypes) SendAccountDisabledConfirmation() *AccountCommandType {
	return o.values[1]
}

func (o *accountCommandTypes) LoginAccount() *AccountCommandType {
	return o.values[2]
}

func (o *accountCommandTypes) SendAccountCreatedConfirmation() *AccountCommandType {
	return o.values[3]
}

func (o *accountCommandTypes) CreateAccount() *AccountCommandType {
	return o.values[4]
}

func (o *accountCommandTypes) DeleteAccount() *AccountCommandType {
	return o.values[5]
}

func (o *accountCommandTypes) EnableAccount() *AccountCommandType {
	return o.values[6]
}

func (o *accountCommandTypes) DisableAccount() *AccountCommandType {
	return o.values[7]
}

func (o *accountCommandTypes) UpdateAccount() *AccountCommandType {
	return o.values[8]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
	for _, lit := range o.Values() {
		if strings.EqualFold(lit.Name(), name) {
			return lit, true
		}
	}
	return nil, false
}

// we have to convert the instances to Literal interface, because it is not a other way in Go
func (o *accountCommandTypes) Literals() []enum.Literal {
	if o.valuesAsLiterals == nil {
		o.valuesAsLiterals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.valuesAsLiterals[i] = item
		}
	}
	return o.valuesAsLiterals
}
