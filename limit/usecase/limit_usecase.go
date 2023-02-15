package usecase

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/domain"
)

type LimitUseCase struct {
	TransactionRepo domain.TransactionRepository
	UserRepo        domain.UserRepository
	LimitRepo       domain.LimitRepository
}

func NewUserService(ur domain.UserRepository, lm domain.LimitRepository, tr domain.TransactionRepository) domain.LimitUsecase {
	return &LimitUseCase{
		UserRepo:        ur,
		LimitRepo:       lm,
		TransactionRepo: tr,
	}
}

func (us *LimitUseCase) Create(ctx context.Context, lm domain.Limit) error {
	panic("implement me")
}

func (us *LimitUseCase) GetByNik(ctx context.Context, Id int) (domain.Limit, error) {
	panic("implement me")
}

func (us *LimitUseCase) Update(ctx context.Context, lm domain.Limit) error {
	panic("implement me")
}
