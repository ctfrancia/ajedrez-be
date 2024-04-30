package main

import (
	"ctfrancia/ajedrez-be/internal/data"
	"ctfrancia/ajedrez-be/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func apiResponse(c *gin.Context, httpStatus int, status, message string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func normalizeUser(u *models.User) {
	u.FirstName = strings.Trim(u.FirstName, " ")
	u.LastName = strings.Trim(u.LastName, " ")
	u.Username = strings.Trim(u.Username, " ")
	u.Country = strings.Trim(u.Country, " ")
}

func prepareUserUpdate(oldData *data.User, newData *data.User) *data.User {

	return &data.User{}
}

func (app *application) background(fn func()) {
	app.wg.Add(1)
	// Launch a background goroutine.
	go func() {
		app.wg.Done()
		// Recover any panic.
		defer func() {
			if err := recover(); err != nil {
				// TODO: Log the error using the app's logger.
				// app.logger.Error(fmt.Sprintf("%v", err))
			}
		}()

		// Execute the arbitrary function that we passed as the parameter.
		fn()
	}()
}

func prepareTournamentUpdate(t data.Tournament) map[string]interface{} {
	ut := make(map[string]interface{})

	if t.IsActive != nil {
		ut["is_active"] = *t.IsActive
	}

	if t.IsVerified != nil {
		ut["is_verified"] = *t.IsVerified
	}

	if t.IsDeleted != nil {
		ut["is_deleted"] = *t.IsDeleted
	}

	if t.IsOnline != nil {
		ut["is_online"] = *t.IsOnline
	}

	if t.OnlineLink != nil {
		ut["online_link"] = *t.OnlineLink
	}

	if t.IsOTB != nil {
		ut["is_otb"] = *t.IsOTB
	}

	if t.IsHybrid != nil {
		ut["is_hybrid"] = *t.IsHybrid
	}

	if t.IsTeam != nil {
		ut["is_team"] = *t.IsTeam
	}

	if t.IsIndividual != nil {
		ut["is_individual"] = *t.IsIndividual
	}

	if t.IsRated != nil {
		ut["is_rated"] = *t.IsRated
	}

	if t.IsUnrated != nil {
		ut["is_unrated"] = *t.IsUnrated
	}

	if t.MatchMaking != nil {
		ut["match_making"] = *t.MatchMaking
	}

	if t.IsPrivate != nil {
		ut["is_private"] = *t.IsPrivate
	}

	if t.IsPublic != nil {
		ut["is_public"] = *t.IsPublic
	}

	if t.MemberCost != nil {
		ut["member_cost"] = *t.MemberCost
	}

	if t.PublicCost != nil {
		ut["public_cost"] = *t.PublicCost
	}

	if t.Currency != nil {
		ut["currency"] = *t.Currency
	}

	if t.IsOpen != nil {
		ut["is_open"] = *t.IsOpen
	}

	if t.IsClosed != nil {
		ut["is_closed"] = *t.IsClosed
	}

	if t.Code != nil {
		ut["code"] = *t.Code
	}

	if t.Name != nil {
		ut["name"] = *t.Name
	}

	if t.Poster != nil {
		ut["poster"] = *t.Poster
	}

	if t.Dates != nil {
		ut["dates"] = t.Dates
	}

	if t.Location != nil {
		ut["location"] = *t.Location
	}

	if t.RegistrationStartDate != nil {
		ut["registration_start_date"] = *t.RegistrationStartDate
	}

	if t.RegistrationEndDate != nil {
		ut["registration_end_date"] = *t.RegistrationEndDate
	}

	if t.AgeCategory != nil {
		ut["age_category"] = *t.AgeCategory
	}

	if t.TimeControl != nil {
		ut["time_control"] = *t.TimeControl
	}

	if t.Type != nil {
		ut["type"] = *t.Type
	}

	if t.Rounds != nil {
		ut["rounds"] = *t.Rounds
	}

	if t.Organizer != nil {
		ut["organizer"] = *t.Organizer
	}

	if t.UserOrganizer != nil {
		ut["user_organizer"] = *t.UserOrganizer
	}

	if t.ContactEmail != nil {
		ut["contact_email"] = *t.ContactEmail
	}

	if t.ContactPhone != nil {
		ut["contact_phone"] = *t.ContactPhone
	}

	if t.Country != nil {
		ut["country"] = *t.Country
	}

	if t.Province != nil {
		ut["province"] = *t.Province
	}

	if t.City != nil {
		ut["city"] = *t.City
	}

	if t.Address != nil {
		ut["address"] = *t.Address
	}

	if t.PostalCode != nil {
		ut["postal_code"] = *t.PostalCode
	}

	if t.Observations != nil {
		ut["observations"] = *t.Observations
	}

	if t.IsCancelled != nil {
		ut["is_cancelled"] = *t.IsCancelled
	}

	if t.Players != nil {
		ut["players"] = t.Players
	}

	if t.Teams != nil {
		ut["teams"] = t.Teams
	}

	if t.MaxAttendees != nil {
		ut["max_attendees"] = *t.MaxAttendees
	}

	if t.MinAttendees != nil {
		ut["min_attendees"] = *t.MinAttendees
	}

	if t.Completed != nil {
		ut["completed"] = *t.Completed
	}

	if t.IsDraft != nil {
		ut["is_draft"] = *t.IsDraft
	}

	if t.IsPublished != nil {
		ut["is_published"] = *t.IsPublished
	}

	return ut
}
