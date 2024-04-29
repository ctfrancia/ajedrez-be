package repository

import (
	"ctfrancia/ajedrez-be/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}
