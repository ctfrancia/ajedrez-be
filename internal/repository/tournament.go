package repository

import (
	"ctfrancia/ajedrez-be/internal/models"
	"gorm.io/gorm"
)

type TournamentsRepository struct {
	DB *gorm.DB
}

func (r TournamentsRepository) Create(tournament models.Tournament) error {
	return r.DB.Create(tournament).Error
}
