package data

import (
	"database/sql"
)

type Club struct {
	ClubID       int    `json:"club_id,omitempty" db:"club_id"`
	IsActive     bool   `json:"is_active,omitempty" db:"is_active"`
	CreatedAt    string `json:"created_at,omitempty" db:"created_at"`
	IsVerified   bool   `json:"is_verified,omitempty" db:"is_verified"`
	UpdatedAt    string `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt    string `json:"deleted_at,omitempty" db:"deleted_at"`
	Code         string `json:"code" binding:"required" db:"code"`
	Club         string `json:"club" binding:"required" db:"name"`
	Address      string `json:"address" db:"address"`
	Observations string `json:"observations" db:"observations"`
	City         string `json:"city" binding:"required" db:"city"`
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
        RETURNING club_id, code, created_at, version`
	args := []interface{}{
		club.Code,
		club.Club,
		club.Address,
		club.Observations,
		club.City,
	}
	return m.DB.QueryRow(query, args...).Scan(&club.ClubID, &club.CreatedAt)
}
