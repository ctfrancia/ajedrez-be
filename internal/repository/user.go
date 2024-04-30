package repository

import (
	"ctfrancia/ajedrez-be/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) Create(user *models.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return ErrDuplicateEmail

		default:
			return result.Error
		}
	}

	return result.Error
}
