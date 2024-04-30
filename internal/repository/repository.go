package repository

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoResultSet    = errors.New("no rows in result set")
	ErrDuplicateEmail = errors.New("duplicate email")
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
		Users:       UserRepository{DB: db},
		Clubs:       ClubsRepository{db},
		Tokens:      TokensRepository{db},
	}
}
