package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
)

type AccountQueryRepository struct {
	repo eventhorizon.ReadRepo
	ctx  context.Context
}

func NewAccountQueryRepositoryFull(repo eventhorizon.ReadRepo, ctx context.Context) (ret *AccountQueryRepository) {
	ret = &AccountQueryRepository{
		repo: repo,
		ctx:  ctx,
	}
	return
}

func (o *AccountQueryRepository) FindAll() (ret []*Account, err error) {
	var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.ctx); err == nil {
		ret = make([]*Account, len(result))
		for i, e := range result {
			ret[i] = e.(*Account)
		}
	}
	return
}

func (o *AccountQueryRepository) FindById(id uuid.UUID) (ret *Account, err error) {
	var result eventhorizon.Entity
	if result, err = o.repo.Find(o.ctx, id); err == nil {
		ret = result.(*Account)
	}
	return
}

func (o *AccountQueryRepository) CountAll() (ret int, err error) {
	var result []*Account
	if result, err = o.FindAll(); err == nil {
		ret = len(result)
	}
	return
}

func (o *AccountQueryRepository) CountById(id uuid.UUID) (ret int, err error) {
	var result *Account
	if result, err = o.FindById(id); err == nil && result != nil {
		ret = 1
	}
	return
}

func (o *AccountQueryRepository) ExistAll() (ret bool, err error) {
	var result int
	if result, err = o.CountAll(); err == nil {
		ret = result > 0
	}
	return
}

func (o *AccountQueryRepository) ExistById(id uuid.UUID) (ret bool, err error) {
	var result int
	if result, err = o.CountById(id); err == nil {
		ret = result > 0
	}
	return
}
