package data

import (
	"database/sql" // New import
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Username     string    `json:"username"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Avatar       string    `json:"avatar"`
	Club         string    `json:"club"`
	Sex          string    `json:"sex"`
	AboutMe      string    `json:"about_me"`
	ELOFide      int       `json:"elo_fide"`
	ELONational  int       `json:"elo_national"`
	ELORegional  int       `json:"elo_regional"`
	Country      string    `json:"country"`
	Province     string    `json:"province"`
	City         string    `json:"city"`
	Neighborhood string    `json:"neighborhood"`
}

// Define a UserModel struct type which wraps a sql.DB connection pool.
type UserModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the users table.
func (m UserModel) Insert(user *User) error {
	return nil
}

// Add a placeholder method for fetching a specific record from the users table.
func (m UserModel) Get(id int64) (*User, error) {
	return nil, nil
}

// Add a placeholder method for updating a specific record in the users table.
func (m UserModel) Update(movie *User) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the users table.
func (m UserModel) Delete(id int64) error {
	return nil
}
