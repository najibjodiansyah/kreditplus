package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nik        string `gorm:"unique column:nik"`
	Password   string `gorm:"unique column:password"`
	FullName   string `gorm:"column:full_name"`
	LegalName  string `gorm:"column:legal_name"`
	BirthPlace string `gorm:"column:birth_place"`
	BirthDate  string `gorm:"column:birth_date"`
	Wages      int    `gorm:"column:wages"`
	PhotoID    uint   `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo      Photo
}

type Photo struct {
	gorm.Model
	Selfie string `db:"selfie"`
	Ktp    string `db:"ktp"`
}

type UserService interface {
	GetByNik(ctx context.Context, Id int) (User, error)
	Create(ctx context.Context, us User) error
	Update(ctx context.Context, us User) error
}

type UserRepository interface {
	GetByNik(ctx context.Context, Id int) (User, error)
	Create(ctx context.Context, us User) error
	CreatePhoto(ctx context.Context, us User) error
	Update(ctx context.Context, us User) error
}
