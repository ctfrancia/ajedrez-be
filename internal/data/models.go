package data

import (
	"database/sql"
	"errors"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a user that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoResultSet    = errors.New("no rows in result set")
)

// Create a Models struct which wraps the UserModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Users       UserModel
	Clubs       ClubModel
	Tokens      TokenModel
	Tournaments TournamentModel
	// M2M         M2MRelModel
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized models.
func NewModels(db *sql.DB) Models {
	return Models{
		Users:       UserModel{DB: db},
		Clubs:       ClubModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Tournaments: TournamentModel{DB: db},
		// 		M2M:         M2MRelModel{DB: db},
	}
}
