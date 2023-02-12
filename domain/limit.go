package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Limit struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Nik       string         `gorm:"unique" json:"nik"`
	Tenor     int            `json:"tenor"`
	Limit     int            `json:"limit"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LimitService interface {
	Transaction(ctx context.Context) ([]Limit, error)
	GetByNik(ctx context.Context, Id int) (Limit, error)
	Update(ctx context.Context, lm Limit) error
}

type LimitRepository interface {
	Transaction(ctx context.Context) ([]Limit, error)
	GetByNik(ctx context.Context, Id int) (Limit, error)
	Update(ctx context.Context, lm Limit) error
}
