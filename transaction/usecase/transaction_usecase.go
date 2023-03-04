package usecase

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/domain"
)

type TransactionUseCase struct {
	TransactionRepo domain.TransactionRepository
	UserRepo        domain.UserRepository
	LimitRepo       domain.LimitRepository
}

func NewUserService(ur domain.UserRepository, lm domain.LimitRepository, tr domain.TransactionRepository) domain.TransactionUsecase {
	return &TransactionUseCase{
		UserRepo:        ur,
		LimitRepo:       lm,
		TransactionRepo: tr,
	}
}

func (us *TransactionUseCase) GetByNik(ctx context.Context, Id int) (domain.Transaction, error) {
	panic("implement me")
}

func (us *TransactionUseCase) Create(ctx context.Context, user domain.Transaction) error {
	panic("implement me")
}

func (us *TransactionUseCase) Update(ctx context.Context, user domain.Transaction) error {
	panic("implement me")
}
