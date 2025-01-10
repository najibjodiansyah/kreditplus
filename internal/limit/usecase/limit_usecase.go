package usecase

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/internal/limit"
	"github.com/najibjodiansyah/kreditplus/internal/user"
)

type LimitRepository interface {
	Create(ctx context.Context, lm limit.Limit) error
	GetById(ctx context.Context, Id int) ([]limit.Limit, error)
}

type UserRepository interface {
	Login(ctx context.Context, nik string) (user.User, error)
	Create(ctx context.Context, us user.User) error
	Update(ctx context.Context, us user.User) error
}

type LimitUseCase struct {
	UserRepo  UserRepository // dont fo this, open public function on user module instead
	LimitRepo LimitRepository
}

func NewUserService(ur UserRepository, lm LimitRepository) *LimitUseCase {
	return &LimitUseCase{
		UserRepo:  ur,
		LimitRepo: lm,
	}
}

func (us *LimitUseCase) Create(ctx context.Context, lm limit.Limit, nik string) error {
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

func (us *LimitUseCase) GetById(ctx context.Context, Id int) ([]limit.Limit, error) {
	limits, err := us.LimitRepo.GetById(ctx, Id)
	if err != nil {
		return nil, err
	}
	return limits, nil
}
