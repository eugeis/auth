package auth

import (
	"github.com/looplab/eventhorizon"
)

const (
	AccountEnabledEvent                  eventhorizon.EventType = "AccountEnabled"
	AccountSentDisabledConfirmationEvent eventhorizon.EventType = "AccountSentDisabledConfirmation"
	AccountDisabledEvent                 eventhorizon.EventType = "AccountDisabled"
	AccountSentEnabledConfirmationEvent  eventhorizon.EventType = "AccountSentEnabledConfirmation"
	AccountDeletedEvent                  eventhorizon.EventType = "AccountDeleted"
	AccountUpdatedEvent                  eventhorizon.EventType = "AccountUpdated"
	AccountCreatedEvent                  eventhorizon.EventType = "AccountCreated"
	AccountLoggedEvent                   eventhorizon.EventType = "AccountLogged"
)

type AccountLogged struct {
	Username string `json:"username,omitempty" eh:"optional"`
	Email    string `json:"email,omitempty" eh:"optional"`
	Password string `json:"password,omitempty" eh:"optional"`
}

type AccountCreated struct {
	Name     *PersonName `json:"name,omitempty" eh:"optional"`
	Username string      `json:"username,omitempty" eh:"optional"`
	Password string      `json:"password,omitempty" eh:"optional"`
	Email    string      `json:"email,omitempty" eh:"optional"`
	Roles    []string    `json:"roles,omitempty" eh:"optional"`
}

func (o *AccountCreated) AddToRoles(item string) string {
	o.Roles = append(o.Roles, item)
	return item
}

type AccountEnabled struct {
}

type AccountDisabled struct {
}

type AccountUpdated struct {
	Name     *PersonName `json:"name,omitempty" eh:"optional"`
	Username string      `json:"username,omitempty" eh:"optional"`
	Password string      `json:"password,omitempty" eh:"optional"`
	Email    string      `json:"email,omitempty" eh:"optional"`
	Roles    []string    `json:"roles,omitempty" eh:"optional"`
}

func (o *AccountUpdated) AddToRoles(item string) string {
	o.Roles = append(o.Roles, item)
	return item
}
