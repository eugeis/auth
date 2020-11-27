package auth

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	Name      *PersonName `json:"name,omitempty" eh:"optional"`
	Username  string      `json:"username,omitempty" eh:"optional"`
	Password  string      `json:"password,omitempty" eh:"optional"`
	Email     string      `json:"email,omitempty" eh:"optional"`
	Roles     []string    `json:"roles,omitempty" eh:"optional"`
	Disabled  bool        `json:"disabled,omitempty" eh:"optional"`
	Id        uuid.UUID   `json:"id,omitempty" eh:"optional"`
	DeletedAt *time.Time  `json:"deletedAt,omitempty" eh:"optional"`
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
