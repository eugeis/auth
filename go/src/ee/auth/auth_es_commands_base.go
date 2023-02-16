package auth

import (
	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
)

const (
	SendEnabledConfirmationAccountCommand  eventhorizon.CommandType = "SendEnabledConfirmationAccount"
	SendDisabledConfirmationAccountCommand eventhorizon.CommandType = "SendDisabledConfirmationAccount"
	LoginAccountCommand                    eventhorizon.CommandType = "LoginAccount"
	CreateAccountCommand                   eventhorizon.CommandType = "CreateAccount"
	DeleteAccountCommand                   eventhorizon.CommandType = "DeleteAccount"
	EnableAccountCommand                   eventhorizon.CommandType = "EnableAccount"
	DisableAccountCommand                  eventhorizon.CommandType = "DisableAccount"
	UpdateAccountCommand                   eventhorizon.CommandType = "UpdateAccount"
)

type SendEnabledConfirmationAccount struct {
	Id uuid.UUID `json:"id,omitempty" eh:"optional"`
}

func (o *SendEnabledConfirmationAccount) AggregateID() uuid.UUID { return o.Id }
func (o *SendEnabledConfirmationAccount) AggregateType() eventhorizon.AggregateType {
	return AccountAggregateType
}
func (o *SendEnabledConfirmationAccount) CommandType() eventhorizon.CommandType {
	return SendEnabledConfirmationAccountCommand
}

type SendDisabledConfirmationAccount struct {
	Id uuid.UUID `json:"id,omitempty" eh:"optional"`
}

func (o *SendDisabledConfirmationAccount) AggregateID() uuid.UUID { return o.Id }
func (o *SendDisabledConfirmationAccount) AggregateType() eventhorizon.AggregateType {
	return AccountAggregateType
}
func (o *SendDisabledConfirmationAccount) CommandType() eventhorizon.CommandType {
	return SendDisabledConfirmationAccountCommand
}

type LoginAccount struct {
	Username string    `json:"username,omitempty" eh:"optional"`
	Email    string    `json:"email,omitempty" eh:"optional"`
	Password string    `json:"password,omitempty" eh:"optional"`
	Id       uuid.UUID `json:"id,omitempty" eh:"optional"`
}

func (o *LoginAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *LoginAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *LoginAccount) CommandType() eventhorizon.CommandType     { return LoginAccountCommand }

type CreateAccount struct {
	Name     *PersonName `json:"name,omitempty" eh:"optional"`
	Username string      `json:"username,omitempty" eh:"optional"`
	Password string      `json:"password,omitempty" eh:"optional"`
	Email    string      `json:"email,omitempty" eh:"optional"`
	Roles    []string    `json:"roles,omitempty" eh:"optional"`
	Id       uuid.UUID   `json:"id,omitempty" eh:"optional"`
}

func (o *CreateAccount) AddToRoles(item string) string {
	o.Roles = append(o.Roles, item)
	return item
}
func (o *CreateAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *CreateAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *CreateAccount) CommandType() eventhorizon.CommandType     { return CreateAccountCommand }

type DeleteAccount struct {
	Id uuid.UUID `json:"id,omitempty" eh:"optional"`
}

func (o *DeleteAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *DeleteAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *DeleteAccount) CommandType() eventhorizon.CommandType     { return DeleteAccountCommand }

type EnableAccount struct {
	Id uuid.UUID `json:"id,omitempty" eh:"optional"`
}

func (o *EnableAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *EnableAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *EnableAccount) CommandType() eventhorizon.CommandType     { return EnableAccountCommand }

type DisableAccount struct {
	Id uuid.UUID `json:"id,omitempty" eh:"optional"`
}

func (o *DisableAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *DisableAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *DisableAccount) CommandType() eventhorizon.CommandType     { return DisableAccountCommand }

type UpdateAccount struct {
	Name     *PersonName `json:"name,omitempty" eh:"optional"`
	Username string      `json:"username,omitempty" eh:"optional"`
	Password string      `json:"password,omitempty" eh:"optional"`
	Email    string      `json:"email,omitempty" eh:"optional"`
	Roles    []string    `json:"roles,omitempty" eh:"optional"`
	Id       uuid.UUID   `json:"id,omitempty" eh:"optional"`
}

func (o *UpdateAccount) AddToRoles(item string) string {
	o.Roles = append(o.Roles, item)
	return item
}
func (o *UpdateAccount) AggregateID() uuid.UUID                    { return o.Id }
func (o *UpdateAccount) AggregateType() eventhorizon.AggregateType { return AccountAggregateType }
func (o *UpdateAccount) CommandType() eventhorizon.CommandType     { return UpdateAccountCommand }
