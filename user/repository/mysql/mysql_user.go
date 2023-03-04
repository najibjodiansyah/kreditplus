package mysql

import (
	"github.com/labstack/echo/v4"
	"github.com/najibjodiansyah/kreditplus/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn: conn}
}

func (m *mysqlUserRepository) Login(ctx echo.Context, nik string) (domain.User, error) {
	var user domain.User
	result := m.Conn.Where("nik = ?", nik).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (m *mysqlUserRepository) GetByNik(ctx echo.Context, nik string) (domain.User, error) {
	var user domain.User
	result := m.Conn.Where("nik = ?", nik).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (m *mysqlUserRepository) Create(ctx echo.Context, user domain.User) error {
	result := m.Conn.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *mysqlUserRepository) Update(ctx echo.Context, user domain.User) error {
	result := m.Conn.Model(domain.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
