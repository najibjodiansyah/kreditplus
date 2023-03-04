package mysql

import (
	"context"

	"github.com/najibjodiansyah/kreditplus/domain"
	"gorm.io/gorm"
)

type mysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransactionRepository(conn *gorm.DB) domain.TransactionRepository {
	return &mysqlTransactionRepository{Conn: conn}
}

func (m *mysqlTransactionRepository) GetByNik(ctx context.Context, Id int) (domain.Transaction, error) {
	var user domain.Transaction
	result := m.Conn.Where("nik = ?", Id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (m *mysqlTransactionRepository) Create(ctx context.Context, tr domain.Transaction) error {
	result := m.Conn.Create(&tr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *mysqlTransactionRepository) Update(ctx context.Context, tr domain.Transaction) error {
	result := m.Conn.Save(&tr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
