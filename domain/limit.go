package domain

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Limit struct {
	gorm.Model
	User_id uint `gorm:"column:user_id` // foreign key
	Tenor   int  `gorm:"column:tenor"`
	Limit   int  `gorm:"column:limit"`
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type LimitUsecase interface {
	Create(ctx echo.Context, lm Limit, nik string) error
	GetById(ctx echo.Context, Id int) ([]Limit, error)
}

type LimitRepository interface {
	Create(ctx echo.Context, lm Limit) error
	GetById(ctx echo.Context, Id int) ([]Limit, error)
}
