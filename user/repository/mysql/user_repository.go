package mysql

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) domain.UserService {
	return &mysqlUserRepository{Conn: conn}
}

func (m *mysqlUserRepository) GetByNik(ctx context.Context, nik int) (domain.User, error) {
	var user domain.User
	result := m.Conn.Where("nik = ?", nik).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (m *mysqlUserRepository) Create(ctx context.Context, user domain.User) error {
	result := m.Conn.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, user domain.User) error {
	result := m.Conn.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
