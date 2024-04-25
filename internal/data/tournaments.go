package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"time"
)

type Tournament struct {
	ID                    *int           `json:"id,omitempty" db:"id"`
	IsActive              *bool          `json:"is_active,omitempty" db:"is_active"`
	IsVerified            *bool          `json:"is_verified,omitempty" db:"is_verified"`
	CreatedAt             *time.Time     `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt             *time.Time     `json:"-" db:"updated_at"`
	UpdatedBy             *sql.NullInt64 `json:"-" db:"updated_by"`
	DeletedAt             *sql.NullTime  `json:"-" db:"deleted_at"`
	IsDeleted             *bool          `json:"is_deleted,omitempty" db:"is_deleted"`
	IsOnline              *bool          `json:"is_online,omitempty" db:"is_online"`
	OnlineLink            *string        `json:"online_link,omitempty" db:"online_link"`
	IsOTB                 *bool          `json:"is_otb,omitempty" db:"is_otb"`
	IsHybrid              *bool          `json:"is_hybrid,omitempty" db:"is_hybrid"`
	IsTeam                *bool          `json:"is_team,omitempty" db:"is_team"`
	IsIndividual          *bool          `json:"is_individual,omitempty" db:"is_individual"`
	IsRated               *bool          `json:"is_rated,omitempty" db:"is_rated"`
	IsUnrated             *bool          `json:"is_unrated,omitempty" db:"is_unrated"`
	MatchMaking           *string        `json:"match_making,omitempty" db:"match_making"`
	IsPrivate             *bool          `json:"is_private,omitempty" db:"is_private"`
	IsPublic              *bool          `json:"is_public,omitempty" db:"is_public"`
	MemberCost            *int32         `json:"member_cost,omitempty" db:"member_cost"`
	PublicCost            *int32         `json:"public_cost,omitempty" db:"public_cost"`
	Currency              *string        `json:"currency,omitempty" db:"currency"`
	IsOpen                *bool          `json:"is_open,omitempty" db:"is_open"`
	IsClosed              *bool          `json:"is_closed,omitempty" db:"is_closed"`
	Code                  *string        `json:"code,omitempty" db:"code"`
	Name                  *string        `json:"name,omitempty" db:"name"`
	Poster                *string        `json:"poster,omitempty" db:"poster"`
	Dates                 []time.Time    `json:"dates,omitempty" db:"dates"`
	Location              *string        `json:"location,omitempty" db:"location"`
	RegistrationStartDate *time.Time     `json:"registration_start_date,omitempty" db:"registration_start_date"`
	RegistrationEndDate   *time.Time     `json:"registration_end_date,omitempty" db:"registration_end_date"`
	AgeCategory           *string        `json:"age_category,omitempty" db:"age_category"`
	TimeControl           *string        `json:"time_control,omitempty" db:"time_control"`
	Type                  *string        `json:"type,omitempty" db:"type"`
	Rounds                *int8          `json:"rounds,omitempty" db:"rounds"`
	Organizer             *string        `json:"organizer,omitempty" db:"organizer"`
	UserOrganizer         *int           `json:"user_organizer,omitempty" db:"user_organizer_id"`
	ContactEmail          *string        `json:"contact_email,omitempty" db:"contact_email"`
	ContactPhone          *string        `json:"contact_phone,omitempty" db:"contact_phone"`
	Country               *string        `json:"country,omitempty" db:"country"`
	Province              *string        `json:"province,omitempty" db:"province"`
	City                  *string        `json:"city,omitempty" db:"city"`
	Address               *string        `json:"address,omitempty" db:"address"`
	PostalCode            *string        `json:"postal_code,omitempty" db:"postal_code"`
	Observations          *string        `json:"observations,omitempty" db:"observations"`
	IsCancelled           *bool          `json:"is_cancelled,omitempty" db:"is_cancelled"`
	Players               []string       `json:"players,omitempty" db:"players"` // user.id
	Teams                 []string       `json:"teams,omitempty" db:"teams"`     // team.id
	MaxAttendees          *int           `json:"max_attendees,omitempty" db:"max_attendees"`
	MinAttendees          *int           `json:"min_attendees,omitempty" db:"min_attendees"`
	Completed             *bool          `json:"completed,omitempty" db:"completed"`
	IsDraft               *bool          `json:"is_draft,omitempty" db:"is_draft"`
	IsPublished           *bool          `json:"is_published,omitempty" db:"is_published"`
	Version               *int16         `json:"-" db:"version"`
}

type TournamentModel struct {
	DB *sql.DB
}

func (m TournamentModel) Insert(t *Tournament) error {
	now := time.Now()
	query := `
        INSERT INTO tournaments (
            name,
            code,
            created_at,
            updated_at
        ) VALUES ($1, $2, $3, $4)
        RETURNING code, created_at`
	args := []any{t.Name, t.Code, now, now}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&t.Code, &t.CreatedAt)
	if err != nil {
		return err
	}
	if t.Teams == nil {
		t.Teams = []string{}
	}

	if t.Players == nil {
		t.Players = []string{}
	}

	if t.Dates == nil {
		t.Dates = []time.Time{}
	}

	return nil
}

func (m TournamentModel) AddPlayersToTournament(t Tournament) error {
	players := &t.Players[0]
	tt := Tournament{}
	query := `
        UPDATE tournaments
        SET players = array_append(players, $1)
        WHERE code = $2
        RETURNING code, created_at, players`
	args := []any{players, t.Code}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("players: ", players)
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&tt.Code, &tt.CreatedAt, &tt.Players)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	return nil
}

// SavePlayers saves the players of a tournament
func (m TournamentModel) SavePlayers(players []string, code string) error {
	query := `
        UPDATE tournaments
        SET players = $1,
        VERSION = VERSION + 1
        WHERE code = $2
        `
	args := []any{players, code}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan()
	if err != nil {
		return err
	}

	return nil
}

// SaveTeams saves the teams of a tournament
func (m TournamentModel) SaveTeams(teams []string, code string) error {
	query := `
        UPDATE tournaments
        SET teams = $1,
        VERSION = VERSION + 1
        WHERE code = $2
        `
	args := []any{teams, code}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan()
	if err != nil {
		return err
	}
	return nil
}

func (m TournamentModel) GetPlayersAndTeams(code string) ([]int, []int, error) {
	players := []int{}
	teams := []int{}
	query := `
        SELECT players, teams
        FROM tournaments
        WHERE code = $1`
	args := []any{code}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&players, &teams)
	if err != nil {
		return []int{}, []int{}, err
	}

	return players, teams, nil
}

func (m TournamentModel) GetPlayers(code string) ([]int, error) {
	players := []int{}
	query := `
        SELECT players
        FROM tournaments
        WHERE code = $1`
	args := []any{code}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (m TournamentModel) GetTeams() ([]*Tournament, error) {
	return nil, nil
}

func (m TournamentModel) SaveDates() ([]*Tournament, error) {
	return nil, nil
}

func (m TournamentModel) AddTeamsToTournament(code string, teams []string) (*Tournament, error) {
	return nil, nil
}

func (m TournamentModel) RemoveTeamsToTournament(code string, teams []string) (*Tournament, error) {
	return nil, nil
}

// GetByCode returns a tournament by its uuid-code
func (m TournamentModel) GetByCode(code string) (*Tournament, error) {
	return nil, nil
}

func (m TournamentModel) GetByID(id int) (Tournament, error) {
	t := Tournament{}

	query := `
        SELECT
            name,
            is_active,
            is_verified,
            created_at,
            updated_at,
            updated_by,
            deleted_at,
            is_deleted,
            is_online,
            online_link,
            is_otb,
            is_hybrid,
            is_team,
            is_individual,
            is_rated,
            is_unrated,
            match_making,
            is_private,
            is_public,
            member_cost,
            public_cost,
            currency,
            is_open,
            is_closed,
            code,
            poster,
            dates,
            location,
            registration_start_date,
            registration_end_date,
            age_category,
            time_control,
            type,
            rounds,
            organizer,
            user_organizer,
            contact_email,
            contact_phone,
            country,
            province,
            city,
            address,
            postal_code,
            observations,
            is_cancelled,
            players,
            teams,
            max_attendees,
            min_attendees,
            completed,
            is_draft,
            is_published
        FROM tournaments
        WHERE id = $1`
	args := []any{id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&t.Name,
		&t.IsActive,
		&t.IsVerified,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.UpdatedBy,
		&t.DeletedAt,
		&t.IsDeleted,
		&t.IsOnline,
		&t.OnlineLink,
		&t.IsOTB,
		&t.IsHybrid,
		&t.IsTeam,
		&t.IsIndividual,
		&t.IsRated,
		&t.IsUnrated,
		&t.MatchMaking,
		&t.IsPrivate,
		&t.IsPublic,
		&t.MemberCost,
		&t.PublicCost,
		&t.Currency,
		&t.IsOpen,
		&t.IsClosed,
		&t.Code,
		&t.Poster,
		&t.Dates,
		&t.Location,
		&t.RegistrationStartDate,
		&t.RegistrationEndDate,
		&t.AgeCategory,
		&t.TimeControl,
		&t.Type,
		&t.Rounds,
		&t.Organizer,
		&t.UserOrganizer,
		&t.ContactEmail,
		&t.ContactPhone,
		&t.Country,
		&t.Province,
		&t.City,
		&t.Address,
		&t.PostalCode,
		&t.Observations,
		&t.IsCancelled,
		&t.Players,
		&t.Teams,
		&t.MaxAttendees,
		&t.MinAttendees,
		&t.Completed,
		&t.IsDraft,
		&t.IsPublished,
	)

	if err != nil {
		return t, err
	}
	return t, nil
}

// Update needs to be
func (m TournamentModel) Update(nt map[string]interface{}) (Tournament, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	qs := psql.Update("tournaments")
	t := Tournament{}
	now := time.Now()
	for key, value := range nt {
		if key == "players" {
			qs = qs.Set(key, pq.Array(value))
			continue
		}

		if key == "teams" {
			qs = qs.Set(key, pq.Array(value))
			continue
		}

		if key == "dates" {
			qs = qs.Set(key, pq.Array(value))
			continue
		}

		qs = qs.Set(key, value)
	}
	qs = qs.Set("updated_at", now)
	qs = qs.Set("version", sq.Expr("version + 1"))
	qs = qs.Where(sq.Eq{"code": nt["code"]})
	qs = qs.Suffix("RETURNING code, created_at")

	query, args, err := qs.ToSql()
	fmt.Println("query: ", query)
	if err != nil {
		return t, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = m.DB.QueryRowContext(ctx, query, args...).Scan(&t.Code, &t.CreatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return t, ErrRecordNotFound
		default:
			return t, err
		}
	}

	return t, nil
}
