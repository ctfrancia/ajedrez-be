package data

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type Tournament struct {
	ID                    int           `json:"id,omitempty" db:"id"`
	IsActive              bool          `json:"is_active" db:"is_active"`
	IsVerified            bool          `json:"is_verified" db:"is_verified"`
	CreatedAt             time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time     `json:"-" db:"updated_at"`
	UpdatedBy             sql.NullInt64 `json:"-" db:"updated_by"`
	DeletedAt             sql.NullTime  `json:"-" db:"deleted_at"`
	IsDeleted             bool          `json:"is_deleted,omitempty" db:"is_deleted"`
	IsOnline              bool          `json:"is_online" db:"is_online"`
	OnlineLink            string        `json:"online_link" db:"online_link"`
	IsOTB                 bool          `json:"is_otb" db:"is_otb"`
	IsHybrid              bool          `json:"is_hybrid" db:"is_hybrid"`
	IsTeam                bool          `json:"is_team" db:"is_team"`
	IsIndividual          bool          `json:"is_individual" db:"is_individual"`
	IsRated               bool          `json:"is_rated" db:"is_rated"`
	IsUnrated             bool          `json:"is_unrated" db:"is_unrated"`
	MatchMaking           string        `json:"match_making" db:"match_making"`
	IsPrivate             bool          `json:"is_private" db:"is_private"`
	IsPublic              bool          `json:"is_public" db:"is_public"`
	MemberCost            int32         `json:"member_cost" db:"member_cost"`
	PublicCost            int32         `json:"public_cost" db:"public_cost"`
	Currency              string        `json:"currency" db:"currency"`
	IsOpen                bool          `json:"is_open" db:"is_open"`
	IsClosed              bool          `json:"is_closed" db:"is_closed"`
	Code                  string        `json:"code" db:"code"`
	Name                  string        `json:"name" db:"name"`
	Poster                string        `json:"poster" db:"poster"`
	Dates                 []time.Time   `json:"dates" db:"dates"`
	Location              string        `json:"location" db:"location"`
	RegistrationStartDate time.Time     `json:"registration_start_date" db:"registration_start_date"`
	RegistrationEndDate   time.Time     `json:"registration_end_date" db:"registration_end_date"`
	AgeCategory           string        `json:"age_category" db:"age_category"`
	TimeControl           string        `json:"time_control" db:"time_control"`
	Type                  string        `json:"type" db:"type"`
	Rounds                int8          `json:"rounds" db:"rounds"`
	Organizer             string        `json:"organizer" db:"organizer"`
	UserOrganizer         int           `json:"user_organizer" db:"user_organizer_id"`
	ContactEmail          string        `json:"contact_email" db:"contact_email"`
	ContactPhone          string        `json:"contact_phone" db:"contact_phone"`
	Country               string        `json:"country" db:"country"`
	Province              string        `json:"province" db:"province"`
	City                  string        `json:"city" db:"city"`
	Address               string        `json:"address" db:"address"`
	PostalCode            string        `json:"postal_code" db:"postal_code"`
	Observations          string        `json:"observations" db:"observations"`
	IsCancelled           bool          `json:"is_cancelled" db:"is_cancelled"`
	Players               []string      `json:"players" db:"players"` // user codes
	Teams                 []string      `json:"teams" db:"teams"`     // team codes
	MaxAttendees          int32         `json:"max_attendees" db:"max_attendees"`
	MinAttendees          int32         `json:"min_attendees" db:"min_attendees"`
	Completed             bool          `json:"completed" db:"completed"`
	IsDraft               bool          `json:"is_draft" db:"is_draft"`
	IsPublished           bool          `json:"is_published" db:"is_published"`
	Version               int32         `json:"-" db:"version"`
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

func (m TournamentModel) Update(nt map[string]interface{}) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	t := &Tournament{}
	qs := psql.Update("tournaments")
	now := time.Now()
	for key, value := range nt {
		if key == "code" {
			continue
		}
		qs = qs.Set(key, value)
	}
	qs = qs.Set("updated_at", now)
	qs = qs.Set("version", sq.Expr("version + 1"))
	qs = qs.Where(sq.Eq{"code": nt["code"]})
	qs = qs.Suffix("RETURNING updated_at")

	query, args, err := qs.ToSql()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = m.DB.QueryRowContext(ctx, query, args...).Scan(&t.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
