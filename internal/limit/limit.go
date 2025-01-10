package limit

import (
	"github.com/najibjodiansyah/kreditplus/internal/user"

	"gorm.io/gorm"
)

type Limit struct {
	gorm.Model
	User_id uint      `gorm:"column:user_id` // foreign key
	Tenor   int       `gorm:"column:tenor"`
	Limit   int       `gorm:"column:limit"`
	User    user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
