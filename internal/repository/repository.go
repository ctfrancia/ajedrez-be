package repository

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoResultSet    = errors.New("no rows in result set")
)

type Repository struct {
	Users       UserRepository
	Clubs       ClubsRepository
	Tokens      TokensRepository
	Tournaments TournamentsRepository
	// Matches MatchesRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Tournaments: TournamentsRepository{db},
	}
}
