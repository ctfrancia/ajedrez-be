package repository

import (
	"gorm.io/gorm"
)

type TokensRepository struct {
	db *gorm.DB
}
