package auth

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	Name                     *PersonName                `json:"name,omitempty" eh:"optional"`
	Username                 string                     `json:"username,omitempty" eh:"optional"`
	Password                 string                     `json:"password,omitempty" eh:"optional"`
	Email                    string                     `json:"email,omitempty" eh:"optional"`
	Roles                    []string                   `json:"roles,omitempty" eh:"optional"`
	SentDisabledConfirmation bool                       `json:"sentDisabledConfirmation,omitempty" eh:"optional"`
	SentEnabledConfirmation  bool                       `json:"sentEnabledConfirmation,omitempty" eh:"optional"`
	Disabled                 bool                       `json:"disabled,omitempty" eh:"optional"`
	Id                       uuid.UUID                  `json:"id,omitempty" eh:"optional"`
	AggregateState           *AccountAggregateStateType `json:"aggregateState,omitempty" eh:"optional"`
	DeletedAt                *time.Time                 `json:"deletedAt,omitempty" eh:"optional"`
}

func NewAccountDefault() (ret *Account) {
	ret = &Account{}
	return
}

func (o *Account) AddToRoles(item string) string {
	o.Roles = append(o.Roles, item)
	return item
}
func (o *Account) EntityID() uuid.UUID { return o.Id }
func (o *Account) Deleted() *time.Time { return o.DeletedAt }

type Deleted struct {
}

func NewDeletedDefault() (ret *Deleted) {
	ret = &Deleted{}
	return
}

type Disabled struct {
	*Exist
}

func NewDisabledDefault() (ret *Disabled) {
	exist := NewExistDefault()
	ret = &Disabled{
		Exist: exist,
	}
	return
}

type Enabled struct {
	*Exist
}

func NewEnabledDefault() (ret *Enabled) {
	exist := NewExistDefault()
	ret = &Enabled{
		Exist: exist,
	}
	return
}

type Exist struct {
}

func NewExistFull() (ret *Exist) {
	ret = &Exist{}
	return
}

func NewExistDefault() (ret *Exist) {
	ret = &Exist{}
	return
}

type AccountHandler struct {
}

func NewAccountHandlerDefault() (ret *AccountHandler) {
	ret = &AccountHandler{}
	return
}

type Initial struct {
}

func NewInitialDefault() (ret *Initial) {
	ret = &Initial{}
	return
}

type UserCredentials struct {
	Username string `json:"username,omitempty" eh:"optional"`
	Password string `json:"password,omitempty" eh:"optional"`
}

func NewUserCredentialsDefault() (ret *UserCredentials) {
	ret = &UserCredentials{}
	return
}

type PersonName struct {
	First string `json:"first,omitempty" eh:"optional"`
	Last  string `json:"last,omitempty" eh:"optional"`
}

func NewPersonNameDefault() (ret *PersonName) {
	ret = &PersonName{}
	return
}
