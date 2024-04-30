package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name    string
	Members []User `gorm:"many2many:team_members;"`
}
