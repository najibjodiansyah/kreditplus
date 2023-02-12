package domain

import "context"

type Transaction struct {
	ID           uint   `gorm:"primaryKey"`
	Nik          string `gorm:"unique" json:"nik"`
	OTR          int    `json:"otr"`
	AdminFee     int    `json:"admin_fee"`
	Instalment   int    `json:"instalment"`
	BankInterest int    `json:"bank_interest"`
	AssetName    string `json:"asset_name"`
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
