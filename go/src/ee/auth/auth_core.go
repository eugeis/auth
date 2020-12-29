package auth

import (
	"github.com/go-ee/utils/crypt"
	"github.com/go-ee/utils/eh"
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

func (o *AccountAggregateEngine) ImplementSendCommands() {
	o.AggregateExecutors.Enabled.SendEnabledConfirmationHandler =
		func(cmd *SendEnabledConfirmationAccount, account *Account, event eh.AggregateStoreEvent) (err error) {

			return
		}
}
