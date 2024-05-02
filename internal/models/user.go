package models

import (
	"time"
	// "database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
)

var AnonymousUser = &User{}

type password struct {
	plaintext *string
	hashed    []byte
}

type User struct {
	ID                  int64 `gorm:"primaryKey"`
	IsActive            bool
	Activated           bool
	IsVerified          bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	SoftDeleted         bool
	UserCode            string
	FirstName           string
	LastName            string
	Username            string
	Password            []byte
	PasswordResetToken  string
	Email               *string `gorm:"unique;not null"`
	Avatar              string
	DateOfBirth         time.Time
	AboutMe             string
	Language            string
	Sex                 string
	ClubID              int
	ChessAgeCategory    string
	FideTitle           string
	ELOFideStandard     int
	ELOFideRapid        int
	ELOFideBlitz        int
	ELOFideBullet       int
	NationalTitle       string
	ELONationalStandard int
	ELONationalRapid    int
	ELONationalBlitz    int
	ELONationalBullet   int
	RegionalTitle       string
	ELORegionalStandard int
	ELORegionalRapid    int
	ELORegionalBlitz    int
	ELORegionalBullet   int
	IsArbiter           bool
	IsCoach             bool
	PricePerHour        float32
	Currency            string
	ChessComUsername    string
	LichessUsername     string
	Chess24Username     string
	Country             string
	Province            string
	City                string
	Neighborhood        string
	Version             int
}

func PasswordSet(plain string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), 12)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func PasswordMatches(hashed []byte, plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
