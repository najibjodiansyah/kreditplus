package domain

import "context"

type Limit struct {
	Id    int
	Nik   int
	Tenor int
	Limit int
}

type LimitService interface {
	Transaction(ctx context.Context) ([]Limit, error)
	GetByNik(ctx context.Context, Id int) (Limit, error)
	Update(ctx context.Context, lm Limit) error
}

type LimitRepository interface {
	Transaction(ctx context.Context) ([]Limit, error)
	GetByNik(ctx context.Context, Id int) (Limit, error)
	Update(ctx context.Context, lm Limit) error
}
