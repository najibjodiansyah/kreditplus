package domain

import (
	"context"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	User_id      uint   `gorm:"column:user_id"`
	OTR          int    `json:"otr"`
	AdminFee     int    `json:"admin_fee"`
	Instalment   int    `json:"instalment"`
	BankInterest int    `json:"bank_interest"`
	AssetName    string `json:"asset_name"`
}

type TransactionUsecase interface {
	GetByNik(ctx context.Context, Id int) (Transaction, error)
	Create(ctx context.Context, tr Transaction) error
	Update(ctx context.Context, tr Transaction) error
}

type TransactionRepository interface {
	GetByNik(ctx context.Context, Id int) (Transaction, error)
	Create(ctx context.Context, tr Transaction) error
	Update(ctx context.Context, tr Transaction) error
}
