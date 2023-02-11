package domain

import (
	"context"
)

type User struct {
	Id         int
	Nik        int
	Password   string
	FullName   string
	LegalName  string
	BirthPlace string
	BirthDate  string
	Wages      int
}

type UserService interface {
	GetByNik(ctx context.Context, Id int) (User, error)
	Create(ctx context.Context, us User) error
	Update(ctx context.Context, us User) error
}

type UserRepository interface {
	GetByNik(ctx context.Context, Id int) (User, error)
	Create(ctx context.Context, us User) error
	Update(ctx context.Context, us User) error
}
