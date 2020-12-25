package auth

import (
	"github.com/go-ee/utils/eh"
	"github.com/looplab/eventhorizon"
)

type AccountAggregateExecutor interface {
	Execute(cmd eventhorizon.Command, account *Account, store eh.AggregateStoreEvent) (err error)
}

type AccountAggregateHandler interface {
	Apply(event eventhorizon.Event, account *Account) (err error)
}
