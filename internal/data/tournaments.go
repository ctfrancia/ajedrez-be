package data

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
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
	Players               []string       `json:"players,omitempty" db:"players"` // user codes
	Teams                 []string       `json:"teams,omitempty" db:"teams"`     // team codes
	MaxAttendees          *int32         `json:"max_attendees,omitempty" db:"max_attendees"`
	MinAttendees          *int32         `json:"min_attendees,omitempty" db:"min_attendees"`
	Completed             *bool          `json:"completed,omitempty" db:"completed"`
	IsDraft               *bool          `json:"is_draft,omitempty" db:"is_draft"`
	IsPublished           *bool          `json:"is_published,omitempty" db:"is_published"`
	Version               *int32         `json:"-" db:"version"`
}

// validFields is a slice containing the valid fields that can be used to sort
// filter or modify the results of a query. This is used to prevent SQL injection
// attacks by ensuring that only valid fields are used in the ORDER BY clause
var validFields = []string{
	"is_active",
	"is_verified",
	"is_deleted",
	"is_online",
	"online_link",
	"is_otb",
	"is_hybrid",
	"is_team",
	"is_individual",
	"is_rated",
	"is_unrated",
	"match_making",
	"is_private",
	"is_public",
	"member_cost",
	"public_cost",
	"currency",
	"is_open",
	"is_closed",
	"name",
	"poster",
	"dates",
	"location",
	"registration_start_date",
	"registration_end_date",
	"age_category",
	"time_control",
	"type",
	"rounds",
	"organizer",
	"user_organizer_id",
	"contact_email",
	"contact_phone",
	"country",
	"province",
	"city",
	"address",
	"postal_code",
	"observations",
	"is_cancelled",
	"players",
	"teams",
	"max_attendees",
	"min_attendees",
	"completed",
	"is_draft",
	"is_published",
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
