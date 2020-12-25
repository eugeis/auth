package auth

import (
	"github.com/go-ee/utils/crypt"
)

func (o *EsEngine) ActivatePasswordEncryption() {
	o.Account.ActivatePasswordEncryption()
}

func (o *AccountAggregateEngine) ActivatePasswordEncryption() {
	o.AggregateExecutors.Initial.AddCreatePreparer(
		func(cmd *CreateAccount, entity *Account) (err error) {
			cmd.Password, err = crypt.Hash(cmd.Password)
			return
		})

	o.AggregateExecutors.Exist.AddUpdatePreparer(
		func(cmd *UpdateAccount, entity *Account) (err error) {
			if len(cmd.Password) > 0 {
				cmd.Password, err = crypt.Hash(cmd.Password)
			}
			return
		})
}
