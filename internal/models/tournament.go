package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Tournament struct {
	gorm.Model
	IsActive              *bool
	IsVerified            *bool
	IsDeleted             *bool
	IsOnline              *bool
	OnlineLink            *sql.NullString
	IsOTB                 *bool
	IsHybrid              *bool
	IsTeam                *bool
	IsIndividual          *bool
	IsRated               *bool
	IsUnrated             *bool
	MatchMaking           *string
	IsPrivate             *bool
	IsPublic              *bool
	MemberCost            *int32
	PublicCost            *int32
	Currency              *string
	IsOpen                *bool
	IsClosed              *bool
	Code                  *string
	Name                  *string
	Poster                *string
	Location              *string
	RegistrationStartDate *time.Time
	RegistrationEndDate   *time.Time
	AgeCategory           *string
	TimeControl           *string
	Type                  *string
	Rounds                *int8
	Organizer             *string
	UserOrganizer         *int
	ContactEmail          *string
	ContactPhone          *string
	Country               *string
	Province              *string
	City                  *string
	Address               *string
	PostalCode            *string
	Observations          *string
	IsCancelled           *bool
	MaxAttendees          *int
	MinAttendees          *int
	Completed             *bool
	IsDraft               *bool
	IsPublished           *bool
	Players               *[]User `gorm:"many2many:players_tournaments;"`
	Teams                 *[]Team `gorm:"many2many:teams_tournaments;"`
	Version               *int16
	/*
		ID                    uint          `json:"id,omitempty" db:"id"`
		IsActive              bool          `json:"is_active,omitempty" db:"is_active"`
		IsVerified            bool          `json:"is_verified,omitempty" db:"is_verified"`
		CreatedAt             time.Time     `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt             time.Time     `json:"-" db:"updated_at"`
		UpdatedBy             sql.NullInt64 `json:"-" db:"updated_by"`
		DeletedAt             sql.NullTime  `json:"-" db:"deleted_at"`
		IsDeleted             bool          `json:"is_deleted,omitempty" db:"is_deleted"`
		IsOnline              bool          `json:"is_online,omitempty" db:"is_online"`
		OnlineLink            string        `json:"online_link,omitempty" db:"online_link"`
		IsOTB                 bool          `json:"is_otb,omitempty" db:"is_otb"`
		IsHybrid              bool          `json:"is_hybrid,omitempty" db:"is_hybrid"`
		IsTeam                bool          `json:"is_team,omitempty" db:"is_team"`
		IsIndividual          bool          `json:"is_individual,omitempty" db:"is_individual"`
		IsRated               bool          `json:"is_rated,omitempty" db:"is_rated"`
		IsUnrated             bool          `json:"is_unrated,omitempty" db:"is_unrated"`
		MatchMaking           string        `json:"match_making,omitempty" db:"match_making"`
		IsPrivate             bool          `json:"is_private,omitempty" db:"is_private"`
		IsPublic              bool          `json:"is_public,omitempty" db:"is_public"`
		MemberCost            int32         `json:"member_cost,omitempty" db:"member_cost"`
		PublicCost            int32         `json:"public_cost,omitempty" db:"public_cost"`
		Currency              string        `json:"currency,omitempty" db:"currency"`
		IsOpen                bool          `json:"is_open,omitempty" db:"is_open"`
		IsClosed              bool          `json:"is_closed,omitempty" db:"is_closed"`
		Code                  string        `json:"code,omitempty" db:"code"`
		Name                  string        `json:"name,omitempty" db:"name"`
		Poster                string        `json:"poster,omitempty" db:"poster"`
		Dates                 []time.Time   `json:"dates,omitempty" db:"dates"`
		Location              string        `json:"location,omitempty" db:"location"`
		RegistrationStartDate time.Time     `json:"registration_start_date,omitempty" db:"registration_start_date"`
		RegistrationEndDate   time.Time     `json:"registration_end_date,omitempty" db:"registration_end_date"`
		AgeCategory           string        `json:"age_category,omitempty" db:"age_category"`
		TimeControl           string        `json:"time_control,omitempty" db:"time_control"`
		Type                  string        `json:"type,omitempty" db:"type"`
		Rounds                int8          `json:"rounds,omitempty" db:"rounds"`
		Organizer             string        `json:"organizer,omitempty" db:"organizer"`
		UserOrganizer         int           `json:"user_organizer,omitempty" db:"user_organizer_id"`
		ContactEmail          string        `json:"contact_email,omitempty" db:"contact_email"`
		ContactPhone          string        `json:"contact_phone,omitempty" db:"contact_phone"`
		Country               string        `json:"country,omitempty" db:"country"`
		Province              string        `json:"province,omitempty" db:"province"`
		City                  string        `json:"city,omitempty" db:"city"`
		Address               string        `json:"address,omitempty" db:"address"`
		PostalCode            string        `json:"postal_code,omitempty" db:"postal_code"`
		Observations          string        `json:"observations,omitempty" db:"observations"`
		IsCancelled           bool          `json:"is_cancelled,omitempty" db:"is_cancelled"`
		Players               []int         `json:"players,omitempty" db:"players"` // user.id
		Teams                 []int         `json:"teams,omitempty" db:"teams"`     // team.id
		MaxAttendees          int           `json:"max_attendees,omitempty" db:"max_attendees"`
		MinAttendees          int           `json:"min_attendees,omitempty" db:"min_attendees"`
		Completed             bool          `json:"completed,omitempty" db:"completed"`
		IsDraft               bool          `json:"is_draft,omitempty" db:"is_draft"`
		IsPublished           bool          `json:"is_published,omitempty" db:"is_published"`
		Version               int16         `json:"-" db:"version"`
	*/
}
