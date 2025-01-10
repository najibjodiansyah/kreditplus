package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/najibjodiansyah/kreditplus/config/middleware"
	"github.com/najibjodiansyah/kreditplus/internal/user"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name UserRepository
type UserRepository interface {
	Login(ctx context.Context, nik string) (user.User, error)
	Create(ctx context.Context, us user.User) error
	Update(ctx context.Context, us user.User) error
}

type UserUseCase struct {
	UserRepo UserRepository
}

func NewUserUseCase(ur UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: ur,
	}
}

func (us *UserUseCase) Login(ctx context.Context, nik, pass string) (string, error) {
	user, err := us.UserRepo.Login(ctx, nik)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("user Empty")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return "", err
	}

	token, err := middleware.CreateToken(user.Nik)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserUseCase) Create(ctx context.Context, user user.User) error {

	hashedPassword, errEncrypt := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errEncrypt != nil {
		return errEncrypt
	}

	user.Password = string(hashedPassword)

	err := us.UserRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserUseCase) Update(ctx context.Context, user user.User) error {
	current_user, err := us.UserRepo.Login(ctx, user.Nik)
	if err != nil {
		return err
	}

	user, err = updateUserValidation(user, current_user)
	if err != nil {
		return err
	}

	err = us.UserRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil

}

func updateUserValidation(user, current_user user.User) (user.User, error) {
	user.ID = current_user.ID
	var zeroTime time.Time

	if user.Nik == "" {
		user.Nik = current_user.Nik
	}

	if user.FullName == "" {
		user.FullName = current_user.FullName
	}

	if user.LegalName == "" {
		user.LegalName = current_user.LegalName
	}
	if user.BirthPlace == "" {
		user.BirthPlace = current_user.BirthPlace
	}
	if user.BirthDate == zeroTime {
		user.BirthDate = current_user.BirthDate
	}
	if user.Wages == 0 {
		user.Wages = current_user.Wages
	}

	if user.Photo_ktp == "" {
		user.Photo_ktp = current_user.Photo_ktp
	}

	if user.Photo_selfie == "" {
		user.Photo_selfie = current_user.Photo_selfie
	}

	return user, nil
}
