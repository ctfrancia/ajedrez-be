package data

import (
	"database/sql"
	// "strings"
)

type Club struct {
	ClubID       int          `json:"club_id,omitempty" db:"club_id"`
	IsActive     bool         `json:"is_active,omitempty" db:"is_active"`
	CreatedAt    sql.NullTime `json:"created_at,omitempty" db:"created_at"`
	DeletedAt    sql.NullTime `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
	IsVerified   bool         `json:"is_verified,omitempty" db:"is_verified"`
	UpdatedAt    sql.NullTime `json:"updated_at,omitempty" db:"updated_at"`
	Code         string       `json:"code" binding:"required" db:"code"`
	Club         string       `json:"club" binding:"required" db:"name"`
	Address      string       `json:"address" db:"address"`
	Observations string       `json:"observations" db:"observations"`
	City         string       `json:"city" db:"city"`
	Country      string       `json:"country" db:"country"`
}

// Define a ClubModel struct type which wraps a sql.DB connection pool.
type ClubModel struct {
	DB *sql.DB
}

func (m ClubModel) Insert(club *Club) error {
	query := `
        INSERT INTO clubs (
            code,
            name,
            address,
            observations,
            city
        )
        VALUES ($1, $2, $3, $4, $5)
        RETURNING club_id, code, created_at`
	args := []interface{}{
		club.Code,
		club.Club,
		club.Address,
		club.Observations,
		club.City,
	}
	return m.DB.QueryRow(query, args...).Scan(&club.ClubID, &club.Code, &club.CreatedAt)
}

func (m ClubModel) GetByName(name string, club *Club) error {
	query := `
        SELECT * FROM clubs
        WHERE name = $1`
	return m.DB.QueryRow(query, name).Scan(
		&club.ClubID,
		&club.IsActive,
		&club.IsVerified,
		&club.CreatedAt,
		&club.UpdatedAt,
		&club.DeletedAt,
		&club.Code,
		&club.Club,
		&club.Address,
		&club.Observations,
		&club.City,
		&club.Country,
	)
}
