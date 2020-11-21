package auth

import (
	"fmt"
	"github.com/google/uuid"
)

func NewAccountDefaultsByPropNames(count int) []*Account {
	items := make([]*Account, count)
	for i := 0; i < count; i++ {
		items[i] = NewAccountDefaultByPropNames(i)
	}
	return items
}

func NewAccountDefaultByPropNames(intSalt int) (ret *Account) {
	ret = NewAccountDefault()
	ret.Name = NewPersonNameDefault()
	ret.Username = fmt.Sprintf("Username %v", intSalt)
	ret.Password = fmt.Sprintf("Password %v", intSalt)
	ret.Email = fmt.Sprintf("Email %v", intSalt)
	ret.Roles = []string{}
	ret.Id = uuid.New()
	return
}

func NewUserCredentialsDefaultsByPropNames(count int) []*UserCredentials {
	items := make([]*UserCredentials, count)
	for i := 0; i < count; i++ {
		items[i] = NewUserCredentialsDefaultByPropNames(i)
	}
	return items
}

func NewUserCredentialsDefaultByPropNames(intSalt int) (ret *UserCredentials) {
	ret = NewUserCredentialsDefault()
	ret.Username = fmt.Sprintf("Username %v", intSalt)
	ret.Password = fmt.Sprintf("Password %v", intSalt)
	return
}

func NewPersonNameDefaultsByPropNames(count int) []*PersonName {
	items := make([]*PersonName, count)
	for i := 0; i < count; i++ {
		items[i] = NewPersonNameDefaultByPropNames(i)
	}
	return items
}

func NewPersonNameDefaultByPropNames(intSalt int) (ret *PersonName) {
	ret = NewPersonNameDefault()
	ret.First = fmt.Sprintf("First %v", intSalt)
	ret.Last = fmt.Sprintf("Last %v", intSalt)
	return
}
