package usecase

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(ur domain.UserRepository) domain.UserService {
	return &UserService{
		userRepo: ur,
	}
}

func (us *UserService) GetByNik(ctx context.Context, Id int) (domain.User, error) {
	panic("implement me")
}

func (us *UserService) Create(ctx context.Context, user domain.User) error {

	hashedPassword, errEncrypt := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errEncrypt != nil {
		return errEncrypt
	}

	user.Password = string(hashedPassword)

	err := us.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Update(ctx context.Context, user domain.User) error {
	panic("implement me")
}

func (us *UserService) SomeLogic(ctx context.Context, user domain.User) error {
	panic("implement me")
}
