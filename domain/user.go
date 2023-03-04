package domain

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nik          string    `gorm:"unique column:nik"`
	Password     string    `gorm:"unique column:password"`
	FullName     string    `gorm:"column:full_name"`
	LegalName    string    `gorm:"column:legal_name"`
	BirthPlace   string    `gorm:"column:birth_place"`
	BirthDate    time.Time `gorm:"column:birth_date"`
	Wages        int       `gorm:"column:wages"`
	Photo_ktp    string    `gorm:"column:photo_ktp"`
	Photo_selfie string    `gorm:"column:photo_selfie"`
}

//go:generate mockery --name UserUsecase
type UserUsecase interface {
	Login(ctx echo.Context, nik, pass string) (string, error)
	Create(ctx echo.Context, us User) error
	Update(ctx echo.Context, us User) error
}

//go:generate mockery --name UserRepository
type UserRepository interface {
	Login(ctx echo.Context, nik string) (User, error)
	Create(ctx echo.Context, us User) error
	Update(ctx echo.Context, us User) error
}
