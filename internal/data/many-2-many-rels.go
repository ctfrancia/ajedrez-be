package data

import (
	"context"
	"database/sql"
	"time"
)

type M2MRelModel struct {
	DB *sql.DB
}

type DateTournament struct {
	DateID       string `json:"date_id" db:"date_id"`
	TournamentID int    `json:"tournament_id" db:"tournament_id"`
}

type PlayerTournament struct {
	PlayerID     int `json:"player_id" db:"player_id"`
	TournamentID int `json:"tournament_id" db:"tournament_id"`
}

func (m *M2MRelModel) InsertDateTournament(dt *DateTournament) error {
	query := `
        INSERT INTO date_tournament (date_id, tournament_id)
        VALUES ($1, $2)
    `
	args := []any{dt.DateID, dt.TournamentID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan()
	if err != nil {
		return err
	}

	return err
}
