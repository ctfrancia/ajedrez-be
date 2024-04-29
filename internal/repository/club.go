package repository

import (
	"gorm.io/gorm"
)

type ClubsRepository struct {
	db *gorm.DB
}
