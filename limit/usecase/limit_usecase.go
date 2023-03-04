package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/domain"
)

type LimitUseCase struct {
	UserRepo  domain.UserRepository
	LimitRepo domain.LimitRepository
}

func NewUserService(ur domain.UserRepository, lm domain.LimitRepository) domain.LimitUsecase {
	return &LimitUseCase{
		UserRepo:  ur,
		LimitRepo: lm,
	}
}

func (us *LimitUseCase) Create(ctx echo.Context, lm domain.Limit, nik string) error {
	user, err := us.UserRepo.Login(ctx, nik)
	if err != nil {
		return err
	}

	lm.User_id = user.ID

	err = us.LimitRepo.Create(ctx, lm)
	if err != nil {
		return err
	}

	return nil
}

func (us *LimitUseCase) GetById(ctx echo.Context, Id int) ([]domain.Limit, error) {
	limits, err := us.LimitRepo.GetById(ctx, Id)
	if err != nil {
		return nil, err
	}
	return limits, nil
}
