package usecase

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/app/config/middleware"
	"github.com/najibjodiansyah/kreditplus/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) domain.UserUsecase {
	return &UserUseCase{
		UserRepo: ur,
	}
}

func (us *UserUseCase) Login(ctx echo.Context, nik, pass string) (string, error) {
	user, err := us.UserRepo.Login(ctx, nik)
	if err != nil {
		return "", err
	}

	if user == (domain.User{}) {
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

func (us *UserUseCase) Create(ctx echo.Context, user domain.User) error {

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

func (us *UserUseCase) Update(ctx echo.Context, user domain.User) error {
	current_user, err := us.UserRepo.Login(ctx, user.Nik)
	if err != nil {
		return err
	}

	user, err = UpdateUserValidation(user.Nik, user, current_user)
	if err != nil {
		return err
	}

	err = us.UserRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil

}

func UpdateUserValidation(nik string, user, current_user domain.User) (domain.User, error) {
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
