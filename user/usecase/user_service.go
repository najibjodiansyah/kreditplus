package usecase

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/app/config/middleware"
	"github.com/najibjodiansyah/kreditplus/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo  domain.UserRepository
	LimitRepo domain.LimitRepository
}

func NewUserService(ur domain.UserRepository, lr domain.LimitRepository) domain.UserUsecase {
	return &UserService{
		UserRepo:  ur,
		LimitRepo: lr,
	}
}

func (us *UserService) Login(ctx echo.Context, nik, pass string) (string, error) {
	user, err := us.UserRepo.Login(ctx, nik)
	if err != nil {
		return "", err
	}

	if user == (domain.User{}) {
		return "", errors.New("user Not Empty")
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

func (us *UserService) Create(ctx echo.Context, user domain.User) error {

	hashedPassword, errEncrypt := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errEncrypt != nil {
		return errEncrypt
	}

	user.Password = string(hashedPassword)

	err := us.UserRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	// if gaji := user.Wages; gaji < 10000000 {
	// 	return domain.ErrGajiKurang
	// }

	// create limit

	return nil
}

func (us *UserService) Update(ctx echo.Context, user domain.User) error {
	nik, err := middleware.GetNik(ctx)
	if err != nil {
		return err
	}

	current_user, err := us.UserRepo.Login(ctx, nik)
	if err != nil {
		return err
	}

	user, err = UpdateUserValidation(nik, user, current_user)
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

	if user.Nik == "" {
		user.Nik = current_user.Nik
	}

	if user.FullName == "" {
		user.FullName = current_user.FullName
	}
	if user.Password == "" {
		hashedPassword, errEncrypt := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if errEncrypt != nil {
			return domain.User{}, errEncrypt
		}
		user.Password = string(hashedPassword)
	}

	if user.LegalName == "" {
		user.LegalName = current_user.LegalName
	}
	if user.BirthPlace == "" {
		user.BirthPlace = current_user.BirthPlace
	}
	if user.BirthDate == "" {
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
