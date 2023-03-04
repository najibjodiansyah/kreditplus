package mysql

import (
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/domain"
	"gorm.io/gorm"
)

type mysqlLimitRepository struct {
	Conn *gorm.DB
}

func NewMysqlLimitRepository(conn *gorm.DB) domain.LimitRepository {
	return &mysqlLimitRepository{Conn: conn}
}

func (us *mysqlLimitRepository) Create(ctx echo.Context, lm domain.Limit) error {
	result := us.Conn.Create(&lm)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (us *mysqlLimitRepository) GetById(ctx echo.Context, id int) ([]domain.Limit, error) {
	var lms []domain.Limit
	result := us.Conn.Where("user_id = ?", id).Find(&lms)
	if result.Error != nil {
		return lms, result.Error
	}
	return lms, nil
}
