package data

import (
	"database/sql" // New import
	"time"
)

type User struct {
	ID           int64     `json:"id" db:"user_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	Email        string    `json:"email" binding:"required" db:"email"`
	Password     string    `json:"password" binding:"required" db:"password"`
	Username     string    `json:"username" binding:"required" db:"username"`
	FirstName    string    `json:"first_name" binding:"required" db:"first_name"`
	LastName     string    `json:"last_name" binding:"required" db:"last_name"`
	DateOfBirth  time.Time `json:"date_of_birth" binding:"required" db:"dob"`
	Avatar       string    `json:"avatar" binding:"required" db:"avatar"`
	Club         string    `json:"club" binding:"required" db:"club"`
	Sex          string    `json:"sex" binding:"required" db:"sex"`
	AboutMe      string    `json:"about_me" binding:"required" db:"about_me"`
	ELOFide      int       `json:"elo_fide" db:"elo_fide"`
	ELONational  int       `json:"elo_national" db:"elo_national"`
	ELORegional  int       `json:"elo_regional" db:"elo_regional"`
	Country      string    `json:"country" binding:"required" db:"country"`
	Province     string    `json:"province" binding:"required" db:"province"`
	City         string    `json:"city" db:"city"`
	Neighborhood string    `json:"neighborhood" db:"neighborhood"`
}

// Define a UserModel struct type which wraps a sql.DB connection pool.
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
            dob,
            avatar,
            club,
            sex,
            about_me,
            elo_fide,
            elo_national,
            elo_regional,
            country,
            province,
            city,
            neighborhood
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
        RETURNING id, created_at`

	args := []any{
		user.Email,
		user.Password,
		user.Username,
		user.FirstName,
		user.LastName,
		user.DateOfBirth,
		user.Avatar,
		user.Club,
		user.Sex,
		user.AboutMe,
		user.ELOFide,
		user.ELONational,
		user.ELORegional,
		user.Country,
		user.Province,
		user.City,
		user.Neighborhood,
	}
	return m.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt)
}

// Add a placeholder method for fetching a specific record from the users table.
func (m UserModel) Get(id int64) (*User, error) {
	return nil, nil
}

// Add a placeholder method for updating a specific record in the users table.
func (m UserModel) Update(user *User) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the users table.
func (m UserModel) Delete(id int64) error {
	return nil
}
