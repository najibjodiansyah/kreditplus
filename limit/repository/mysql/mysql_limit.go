package mysql

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/domain"
	"gorm.io/gorm"
)

type mysqlLimitRepository struct {
	Conn *gorm.DB
}

func NewMysqlLimitRepository(conn *gorm.DB) domain.LimitRepository {
	return &mysqlLimitRepository{Conn: conn}
}

func (us *mysqlLimitRepository) Create(ctx context.Context, lm domain.Limit) error {
	panic("implement me")
}

func (us *mysqlLimitRepository) GetByNik(ctx context.Context, Id int) (domain.Limit, error) {
	panic("implement me")
}

func (us *mysqlLimitRepository) Update(ctx context.Context, lm domain.Limit) error {
	panic("implement me")
}
