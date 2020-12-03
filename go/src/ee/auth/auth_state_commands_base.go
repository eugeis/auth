package auth

type AccountConfirmationDisabledExecutor struct {
}

func NewAccountConfirmationDisabledExecutorDefault() (ret *AccountConfirmationDisabledExecutor) {
	ret = &AccountConfirmationDisabledExecutor{}
	return
}

type AccountConfirmationEnabledExecutor struct {
}

func NewAccountConfirmationEnabledExecutorDefault() (ret *AccountConfirmationEnabledExecutor) {
	ret = &AccountConfirmationEnabledExecutor{}
	return
}

type AccountConfirmationInitialExecutor struct {
}

func NewAccountConfirmationInitialExecutorDefault() (ret *AccountConfirmationInitialExecutor) {
	ret = &AccountConfirmationInitialExecutor{}
	return
}

type AccountDeletedExecutor struct {
}

func NewAccountDeletedExecutorDefault() (ret *AccountDeletedExecutor) {
	ret = &AccountDeletedExecutor{}
	return
}

type AccountDisabledExecutor struct {
}

func NewAccountDisabledExecutorDefault() (ret *AccountDisabledExecutor) {
	ret = &AccountDisabledExecutor{}
	return
}

type AccountEnabledExecutor struct {
}

func NewAccountEnabledExecutorDefault() (ret *AccountEnabledExecutor) {
	ret = &AccountEnabledExecutor{}
	return
}

type AccountExistExecutor struct {
}

func NewAccountExistExecutorDefault() (ret *AccountExistExecutor) {
	ret = &AccountExistExecutor{}
	return
}

type AccountInitialExecutor struct {
}

func NewAccountInitialExecutorDefault() (ret *AccountInitialExecutor) {
	ret = &AccountInitialExecutor{}
	return
}
