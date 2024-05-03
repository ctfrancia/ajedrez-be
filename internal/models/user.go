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
	ID                  int64     `json:"id" gorm:"primaryKey"`
	IsActive            bool      `json:"is_active"`
	Activated           bool      `json:"activated"`
	IsVerified          bool      `json:"is_verified"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	SoftDeleted         bool      `json:"soft_deleted"`
	UserCode            string    `json:"user_code"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	Username            string    `json:"username"`
	Password            []byte    `json:"password"`
	PasswordResetToken  string    `json:"password_reset_token"`
	Email               *string   `json:"email" gorm:"unique;not null"`
	Avatar              string    `json:"avatar"`
	DateOfBirth         time.Time `json:"date_of_birth"`
	AboutMe             string    `json:"about_me"`
	Language            string    `json:"language" gorm:"not null"`
	Sex                 string    `json:"sex"`
	ClubID              int       `json:"club_id"`
	ChessAgeCategory    string    `json:"chess_age_category"`
	FideTitle           string    `json:"fide_title"`
	ELOFideStandard     int       `json:"elo_fide_standard"`
	ELOFideRapid        int       `json:"elo_fide_rapid"`
	ELOFideBlitz        int       `json:"elo_fide_blitz"`
	ELOFideBullet       int       `json:"elo_fide_bullet"`
	NationalTitle       string    `json:"national_title"`
	ELONationalStandard int       `json:"elo_national_standard"`
	ELONationalRapid    int       `json:"elo_national_rapid"`
	ELONationalBlitz    int       `json:"elo_national_blitz"`
	ELONationalBullet   int       `json:"elo_national_bullet"`
	RegionalTitle       string    `json:"regional_title"`
	ELORegionalStandard int       `json:"elo_regional_standard"`
	ELORegionalRapid    int       `json:"elo_regional_rapid"`
	ELORegionalBlitz    int       `json:"elo_regional_blitz"`
	ELORegionalBullet   int       `json:"elo_regional_bullet"`
	IsCoach             bool      `json:"is_coach"`
	PricePerHour        float32   `json:"price_per_hour"`
	Currency            string    `json:"currency"`
	ChessComUsername    string    `json:"chess_com_username"`
	LichessUsername     string    `json:"lichess_username"`
	Chess24Username     string    `json:"chess24_username"`
	Country             string    `json:"country"`
	Province            string    `json:"province"`
	City                string    `json:"city"`
	Neighborhood        string    `json:"neighborhood"`
	Version             int16     `json:"version" gorm:"default:1"`
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

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}
