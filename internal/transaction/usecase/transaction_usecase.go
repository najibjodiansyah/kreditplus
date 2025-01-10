package usecase

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/internal/transaction"
)

type TransactionRepository interface {
	GetByNik(ctx context.Context, Id int) (transaction.Transaction, error)
	Create(ctx context.Context, tr transaction.Transaction) error
	Update(ctx context.Context, tr transaction.Transaction) error
}

type UserUseCase interface{}

type LimitUseCase interface{}

type TransactionUseCase struct {
	TransactionRepo TransactionRepository
	UserRepo        UserUseCase  // must be user service interface defining in line 15
	LimitRepo       LimitUseCase // must be limit service interface defining in line 15
}

func NewUserService(ur UserUseCase, lm LimitUseCase, tr TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{
		UserRepo:        ur,
		LimitRepo:       lm,
		TransactionRepo: tr,
	}
}

func (us *TransactionUseCase) GetByNik(ctx context.Context, Id int) (transaction.Transaction, error) {
	panic("implement me")
}

func (us *TransactionUseCase) Create(ctx context.Context, user transaction.Transaction) error {
	panic("implement me")
}

func (us *TransactionUseCase) Update(ctx context.Context, user transaction.Transaction) error {
	panic("implement me")
}
