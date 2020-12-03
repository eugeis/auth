package auth

type AccountConfirmationHandlers struct {
	Disabled *AccountConfirmationDisabledHandler
	Enabled  *AccountConfirmationEnabledHandler
	Initial  *AccountConfirmationInitialHandler
}

func NewAccountConfirmationHandlersFull() (ret *AccountConfirmationHandlers) {
	disabled := NewAccountConfirmationDisabledHandlerDefault()
	enabled := NewAccountConfirmationEnabledHandlerDefault()
	initial := NewAccountConfirmationInitialHandlerDefault()
	ret = &AccountConfirmationHandlers{
		Disabled: disabled,
		Enabled:  enabled,
		Initial:  initial,
	}
	return
}

type AccountConfirmationExecutors struct {
	Disabled *AccountConfirmationDisabledExecutor
	Enabled  *AccountConfirmationEnabledExecutor
	Initial  *AccountConfirmationInitialExecutor
}

func NewAccountConfirmationExecutorsFull() (ret *AccountConfirmationExecutors) {
	disabled := NewAccountConfirmationDisabledExecutorDefault()
	enabled := NewAccountConfirmationEnabledExecutorDefault()
	initial := NewAccountConfirmationInitialExecutorDefault()
	ret = &AccountConfirmationExecutors{
		Disabled: disabled,
		Enabled:  enabled,
		Initial:  initial,
	}
	return
}

type AccountHandlers struct {
	Deleted  *AccountDeletedHandler
	Disabled *AccountDisabledHandler
	Enabled  *AccountEnabledHandler
	Exist    *AccountExistHandler
	Initial  *AccountInitialHandler
}

func NewAccountHandlersFull() (ret *AccountHandlers) {
	deleted := NewAccountDeletedHandlerDefault()
	disabled := NewAccountDisabledHandlerDefault()
	enabled := NewAccountEnabledHandlerDefault()
	exist := NewAccountExistHandlerDefault()
	initial := NewAccountInitialHandlerDefault()
	ret = &AccountHandlers{
		Deleted:  deleted,
		Disabled: disabled,
		Enabled:  enabled,
		Exist:    exist,
		Initial:  initial,
	}
	return
}

type AccountExecutors struct {
	Deleted  *AccountDeletedExecutor
	Disabled *AccountDisabledExecutor
	Enabled  *AccountEnabledExecutor
	Exist    *AccountExistExecutor
	Initial  *AccountInitialExecutor
}

func NewAccountExecutorsFull() (ret *AccountExecutors) {
	deleted := NewAccountDeletedExecutorDefault()
	disabled := NewAccountDisabledExecutorDefault()
	enabled := NewAccountEnabledExecutorDefault()
	exist := NewAccountExistExecutorDefault()
	initial := NewAccountInitialExecutorDefault()
	ret = &AccountExecutors{
		Deleted:  deleted,
		Disabled: disabled,
		Enabled:  enabled,
		Exist:    exist,
		Initial:  initial,
	}
	return
}
