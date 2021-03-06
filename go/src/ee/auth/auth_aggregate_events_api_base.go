package auth

import (
	"fmt"
	"github.com/go-ee/utils/enum"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type AccountEventType struct {
	name    string
	ordinal int
}

func (o *AccountEventType) Name() string {
	return o.name
}

func (o *AccountEventType) Ordinal() int {
	return o.ordinal
}

func (o *AccountEventType) IsAccountCreated() bool {
	return o.name == _accountEventTypes.AccountCreated().name
}

func (o *AccountEventType) IsAccountDeleted() bool {
	return o.name == _accountEventTypes.AccountDeleted().name
}

func (o *AccountEventType) IsAccountDisabled() bool {
	return o.name == _accountEventTypes.AccountDisabled().name
}

func (o *AccountEventType) IsAccountEnabled() bool {
	return o.name == _accountEventTypes.AccountEnabled().name
}

func (o *AccountEventType) IsAccountLogged() bool {
	return o.name == _accountEventTypes.AccountLogged().name
}

func (o *AccountEventType) IsAccountSentDisabledConfirmation() bool {
	return o.name == _accountEventTypes.AccountSentDisabledConfirmation().name
}

func (o *AccountEventType) IsAccountSentEnabledConfirmation() bool {
	return o.name == _accountEventTypes.AccountSentEnabledConfirmation().name
}

func (o *AccountEventType) IsAccountUpdated() bool {
	return o.name == _accountEventTypes.AccountUpdated().name
}

func (o *AccountEventType) MarshalJSON() (ret []byte, err error) {
	ret = []byte(fmt.Sprintf("\"%v\"", o.name))
	return
}

func (o *AccountEventType) UnmarshalJSON(data []byte) (err error) {
	name := string(data)
	//remove quotes
	name = name[1 : len(name)-1]
	if v, ok := AccountEventTypes().ParseAccountEventType(name); ok {
		*o = *v
	} else {
		err = fmt.Errorf("invalid AccountEventType %q", name)
	}
	return
}

func (o *AccountEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AccountEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
	if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AccountEventTypes().ParseAccountEventType(lit); ok {
			*o = *v
		} else {
			err = fmt.Errorf("invalid AccountEventType %q", lit)
		}
	}
	return
}

type accountEventTypes struct {
	values           []*AccountEventType
	valuesAsLiterals []enum.Literal
}

var _accountEventTypes = &accountEventTypes{values: []*AccountEventType{
	{name: "AccountCreated", ordinal: 0},
	{name: "AccountDeleted", ordinal: 1},
	{name: "AccountDisabled", ordinal: 2},
	{name: "AccountEnabled", ordinal: 3},
	{name: "AccountLogged", ordinal: 4},
	{name: "AccountSentDisabledConfirmation", ordinal: 5},
	{name: "AccountSentEnabledConfirmation", ordinal: 6},
	{name: "AccountUpdated", ordinal: 7}},
}

func AccountEventTypes() *accountEventTypes {
	return _accountEventTypes
}

func (o *accountEventTypes) Values() []*AccountEventType {
	return o.values
}

func (o *accountEventTypes) AccountCreated() *AccountEventType {
	return o.values[0]
}

func (o *accountEventTypes) AccountDeleted() *AccountEventType {
	return o.values[1]
}

func (o *accountEventTypes) AccountDisabled() *AccountEventType {
	return o.values[2]
}

func (o *accountEventTypes) AccountEnabled() *AccountEventType {
	return o.values[3]
}

func (o *accountEventTypes) AccountLogged() *AccountEventType {
	return o.values[4]
}

func (o *accountEventTypes) AccountSentDisabledConfirmation() *AccountEventType {
	return o.values[5]
}

func (o *accountEventTypes) AccountSentEnabledConfirmation() *AccountEventType {
	return o.values[6]
}

func (o *accountEventTypes) AccountUpdated() *AccountEventType {
	return o.values[7]
}

func (o *accountEventTypes) ParseAccountEventType(name string) (ret *AccountEventType, ok bool) {
	for _, lit := range o.Values() {
		if strings.EqualFold(lit.Name(), name) {
			return lit, true
		}
	}
	return nil, false
}

// we have to convert the instances to Literal interface, because it is not a other way in Go
func (o *accountEventTypes) Literals() []enum.Literal {
	if o.valuesAsLiterals == nil {
		o.valuesAsLiterals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.valuesAsLiterals[i] = item
		}
	}
	return o.valuesAsLiterals
}
