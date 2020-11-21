package auth

type AccountCli struct {
}

func NewAccountCliDefault() (ret *AccountCli) {
	ret = &AccountCli{}
	return
}

type AuthCli struct {
	AccountCli *AccountCli
}

func NewAuthCli() (ret *AuthCli) {

	accountCli := NewAccountCliDefault()
	ret = &AuthCli{
		AccountCli: accountCli,
	}
	return
}
