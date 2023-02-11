package domain

import "context"

type Transaction struct {
	Id           int
	Nik          int
	OTR          int
	AdminFee     int
	Instalment   int
	BankInterest int
	AssetName    string
}

type TransactionService interface {
	Fetch(ctx context.Context) ([]Transaction, error)
	GetByNik(ctx context.Context, Id int) (Transaction, error)
	Create(ctx context.Context, tr Transaction) error
	Update(ctx context.Context, tr Transaction) error
}

type TransactionRepository interface {
	Fetch(ctx context.Context) ([]Transaction, error)
	GetByNik(ctx context.Context, Id int) (Transaction, error)
	Create(ctx context.Context, tr Transaction) error
	Update(ctx context.Context, tr Transaction) error
}
