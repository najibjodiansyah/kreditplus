package mysql

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/internal/limit"
	"gorm.io/gorm"
)

type mysqlLimitRepository struct {
	Conn *gorm.DB
}

func NewMysqlLimitRepository(conn *gorm.DB) *mysqlLimitRepository {
	return &mysqlLimitRepository{Conn: conn}
}

func (us *mysqlLimitRepository) Create(ctx context.Context, lm limit.Limit) error {
	result := us.Conn.Create(&lm)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (us *mysqlLimitRepository) GetById(ctx context.Context, id int) ([]limit.Limit, error) {
	var lms []limit.Limit
	result := us.Conn.Where("user_id = ?", id).Find(&lms)
	if result.Error != nil {
		return lms, result.Error
	}
	return lms, nil
}
