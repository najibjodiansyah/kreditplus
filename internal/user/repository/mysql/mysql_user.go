package mysql

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/internal/user"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) *mysqlUserRepository {
	return &mysqlUserRepository{Conn: conn}
}

func (m *mysqlUserRepository) Login(ctx context.Context, nik string) (user.User, error) {
	var user user.User
	result := m.Conn.Where("nik = ?", nik).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (m *mysqlUserRepository) GetByNik(ctx context.Context, nik string) (user.User, error) {
	var user user.User
	result := m.Conn.Where("nik = ?", nik).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (m *mysqlUserRepository) Create(ctx context.Context, user user.User) error {
	result := m.Conn.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, user user.User) error {
	result := m.Conn.Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
