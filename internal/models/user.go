package models

import (
	"time"
	// "database/sql"
	// "errors"
)

type password struct {
	plaintext *string
	hashed    []byte
}

type User struct {
	ID                  int64     `json:"id,omitempty" db:"id"`
	IsActive            bool      `json:"is_active,omitempty" db:"is_active"`
	Activated           bool      `json:"is_activated,omitempty" db:"is_activated"`
	IsVerified          bool      `json:"is_verified,omitempty" db:"is_verified"`
	CreatedAt           time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at,omitempty" db:"updated_at"`
	SoftDeleted         bool      `json:"-" db:"soft_deleted"`
	UserCode            string    `json:"user_code,omitempty" db:"user_code"`
	FirstName           string    `json:"first_name,omitempty" binding:"required" db:"first_name"`
	LastName            string    `json:"last_name,omitempty" binding:"required" db:"last_name"`
	Username            string    `json:"username,omitempty" db:"username"`
	Password            password  `json:"-" db:"password"`
	PasswordResetToken  string    `json:"-" db:"password_reset_token"`
	Email               string    `json:"email,omitempty" db:"email"`
	Avatar              string    `json:"avatar,omitempty" db:"avatar"`
	DateOfBirth         time.Time `json:"date_of_birth,omitempty" db:"dob"`
	AboutMe             string    `json:"about_me,omitempty" db:"about_me"`
	Language            string    `json:"language,omitempty" db:"language"`
	Sex                 string    `json:"sex,omitempty" db:"sex"`
	ClubID              int       `json:"club_id,omitempty" db:"club_id"`
	ChessAgeCategory    string    `json:"chess_age_category,omitempty" db:"chess_age_category"`
	FideTitle           string    `json:"fide_title,omitempty" db:"fide_title"`
	ELOFideStandard     int       `json:"elo_fide_standard,omitempty" db:"elo_fide_standard"`
	ELOFideRapid        int       `json:"elo_fide_rapid,omitempty" db:"elo_fide_rapid"`
	ELOFideBlitz        int       `json:"elo_fide_blitz,omitempty" db:"elo_fide_blitz"`
	ELOFideBullet       int       `json:"elo_fide_bullet,omitempty" db:"elo_fide_bullet"`
	NationalTitle       string    `json:"national_title,omitempty" db:"national_title"`
	ELONationalStandard int       `json:"elo_national_standard,omitempty" db:"elo_national_standard"`
	ELONationalRapid    int       `json:"elo_national_rapid,omitempty" db:"elo_national_rapid"`
	ELONationalBlitz    int       `json:"elo_national_blitz,omitempty" db:"elo_national_blitz"`
	ELONationalBullet   int       `json:"elo_national_bullet,omitempty" db:"elo_national_bullet"`
	RegionalTitle       string    `json:"regional_title,omitempty" db:"regional_title"`
	ELORegionalStandard int       `json:"elo_regional_standard,omitempty" db:"elo_regional_standard"`
	ELORegionalRapid    int       `json:"elo_regional_rapid,omitempty" db:"elo_regional_rapid"`
	ELORegionalBlitz    int       `json:"elo_regional_blitz,omitempty" db:"elo_regional_blitz"`
	ELORegionalBullet   int       `json:"elo_regional_bullet,omitempty" db:"elo_regional_bullet"`
	IsArbiter           bool      `json:"is_arbiter,omitempty" db:"is_arbiter"`
	IsCoach             bool      `json:"is_coach,omitempty" db:"is_coach"`
	PricePerHour        float32   `json:"price_per_hour,omitempty" db:"price_per_hour"`
	Currency            string    `json:"currency,omitempty" db:"currency"`
	ChessComUsername    string    `json:"chess_com_username,omitempty" db:"chess_com_username"`
	LichessUsername     string    `json:"lichess_username,omitempty" db:"lichess_username"`
	Chess24Username     string    `json:"chess24_username,omitempty" db:"chess24_username"`
	Country             string    `json:"country,omitempty" db:"country"`
	Province            string    `json:"province,omitempty" db:"province"`
	City                string    `json:"city,omitempty" db:"city"`
	Neighborhood        string    `json:"neighborhood,omitempty" db:"neighborhood"`
	Version             int       `json:"-" db:"version"`
}
