package data

import (
	"database/sql"
	"time"
)

type User struct {
	ID                  int64     `json:"id" db:"user_id"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	Email               string    `json:"email" db:"email"`
	Password            string    `json:"password" db:"password"`
	Username            string    `json:"username" db:"username"`
	FirstName           string    `json:"first_name" binding:"required" db:"first_name"`
	LastName            string    `json:"last_name" binding:"required" db:"last_name"`
	DateOfBirth         time.Time `json:"date_of_birth" db:"dob"`
	Avatar              string    `json:"avatar" db:"avatar"`
	ClubID              int       `json:"club_id" db:"club_id"`
	Sex                 string    `json:"sex" db:"sex"`
	AboutMe             string    `json:"about_me" db:"about_me"`
	Country             string    `json:"country" binding:"required" db:"country"`
	Province            string    `json:"province" db:"province"`
	City                string    `json:"city" db:"city"`
	Neighborhood        string    `json:"neighborhood" db:"neighborhood"`
	ClubUserID          string    `json:"club_user_code" db:"club_user_code"`
	ChessAgeCategory    string    `json:"chess_age_category" db:"chess_age_category"`
	ELOFideStandard     int       `json:"elo_fide_standard" db:"elo_fide_standard"`
	ELOFideRapid        int       `json:"elo_fide_rapid" db:"elo_fide_rapid"`
	ELONationalStandard int       `json:"elo_national_standard" db:"elo_national_standard"`
	ELONationalRapid    int       `json:"elo_national_rapid" db:"elo_national_rapid"`
	ELORegionalStandard int       `json:"elo_regional_standard" db:"elo_regional_standard"`
	ELORegionalRapid    int       `json:"elo_regional_rapid" db:"elo_regional_rapid"`
}
type UserModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the users table.
func (m UserModel) Insert(user *User) error {
	query := `
        INSERT INTO users (
            email,
            password,
            username,
            first_name,
            last_name,
            avatar,
            club_id,
            sex,
            about_me,
            country,
            province,
            city,
            neighborhood,
            club_user_code,
            chess_age_category,
            elo_fide_standard,
            elo_fide_rapid,
            elo_national_standard,
            elo_national_rapid,
            elo_regional_standard,
            elo_regional_rapid
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
        RETURNING user_id, created_at`

	args := []any{
		user.Email,
		user.Password,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Avatar,
		user.ClubID,
		user.Sex,
		user.AboutMe,
		user.Country,
		user.Province,
		user.City,
		user.Neighborhood,
		user.ClubUserID,
		user.ChessAgeCategory,
		user.ELOFideStandard,
		user.ELOFideRapid,
		user.ELONationalStandard,
		user.ELONationalRapid,
		user.ELORegionalStandard,
		user.ELORegionalRapid,
	}
	return m.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt)
}
