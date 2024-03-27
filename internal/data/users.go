package data

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID                  int       `json:"user_id,omitempty" db:"user_id"`
	IsActive            bool      `json:"is_active,omitempty" db:"is_active"`
	IsVerified          bool      `json:"is_verified,omitempty" db:"is_verified"`
	CreatedAt           time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at,omitempty" db:"updated_at"`
	SoftDeleted         bool      `json:"-" db:"soft_deleted"`
	UserCode            string    `json:"user_code,omitempty" db:"user_code"`
	FirstName           string    `json:"first_name,omitempty" binding:"required" db:"first_name"`
	LastName            string    `json:"last_name,omitempty" binding:"required" db:"last_name"`
	Username            string    `json:"username,omitempty" db:"username"`
	Password            string    `json:"password,omitempty"   db:"password"`
	PasswordResetToken  string    `json:"password_reset_token,omitempty" db:"password_reset_token"`
	Email               string    `json:"email,omitempty" db:"email"`
	Avatar              string    `json:"avatar,omitempty" db:"avatar"`
	DateOfBirth         time.Time `json:"date_of_birth,omitempty" db:"dob"`
	AboutMe             string    `json:"about_me,omitempty" db:"about_me"`
	Sex                 string    `json:"sex,omitempty" db:"sex"`
	ClubID              int       `json:"club_id,-" db:"club_id"`
	ChessAgeCategory    string    `json:"chess_age_category,omitempty" db:"chess_age_category"`
	ELOFideStandard     int       `json:"elo_fide_standard,omitempty" db:"elo_fide_standard"`
	ELOFideRapid        int       `json:"elo_fide_rapid,omitempty" db:"elo_fide_rapid"`
	ELOFideBlitz        int       `json:"elo_fide_blitz,omitempty" db:"elo_fide_blitz"`
	ELOFideBullet       int       `json:"elo_fide_bullet,omitempty" db:"elo_fide_bullet"`
	ELONationalStandard int       `json:"elo_national_standard,omitempty" db:"elo_national_standard"`
	ELONationalRapid    int       `json:"elo_national_rapid,omitempty" db:"elo_national_rapid"`
	ELONationalBlitz    int       `json:"elo_national_blitz,omitempty" db:"elo_national_blitz"`
	ELONationalBullet   int       `json:"elo_national_bullet,omitempty" db:"elo_national_bullet"`
	ELORegionalStandard int       `json:"elo_regional_standard,omitempty" db:"elo_regional_standard"`
	ELORegionalRapid    int       `json:"elo_regional_rapid,omitempty" db:"elo_regional_rapid"`
	ELORegionalBlitz    int       `json:"elo_regional_blitz,omitempty" db:"elo_regional_blitz"`
	EloRegionalBullet   int       `json:"elo_regional_bullet,omitempty" db:"elo_regional_bullet"`
	IsArbiter           bool      `json:"is_arbiter,omitempty" db:"is_arbiter"`
	IsCoach             bool      `json:"is_coach,omitempty" db:"is_coach"`
	PricePerHour        float32   `json:"price_per_hour,omitempty" db:"price_per_hour"`
	Currency            string    `json:"currency,omitempty" db:"currency"`
	ChessComUsername    string    `json:"chess_com_username,omitempty" db:"chess_com_username"`
	LichessUsername     string    `json:"lichess_username,omitempty" db:"lichess_username"`
	Chess24Username     string    `json:"chess24_username,omitempty" db:"chess24_username"`
	Country             string    `json:"country,omitempty" db:"country"`
	Province            string    `json:"province,omitempty" db:"province"`
	City                string    `json:"city,omitempty" db:"city"`
	Neighborhood        string    `json:"neighborhood,omitempty" db:"neighborhood"`
	Version             int       `json:"version,omitempty" db:"version"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	fmt.Println("user", user)
	query := `
        INSERT INTO users (
            is_active,
            is_verified,
            created_at,
            updated_at,
            soft_deleted,
            user_code,
            first_name,
            last_name,
            username,
            password,
            password_reset_token,
            email,
            avatar,
            dob,
            about_me,
            sex,
            club_id,
            chess_age_category,
            elo_fide_standard,
            elo_fide_rapid,
            elo_fide_blitz,
            elo_fide_bullet,
            elo_national_standard,
            elo_national_rapid,
            elo_national_blitz,
            elo_national_bullet,
            elo_regional_standard,
            elo_regional_rapid,
            elo_regional_blitz,
            elo_regional_bullet,
            is_arbiter,
            is_coach,
            price_per_hour,
            currency,
            chess_com_username,
            lichess_username,
            chess24_username,
            country,
            province,
            city,
            neighborhood
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
        $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
        $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41)
        RETURNING user_id, created_at, user_code`

	args := []any{
		user.IsActive,
		user.IsVerified,
		user.CreatedAt,
		user.UpdatedAt,
		user.SoftDeleted,
		user.UserCode,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Password,
		user.PasswordResetToken,
		user.Email,
		user.Avatar,
		user.DateOfBirth,
		user.AboutMe,
		user.Sex,
		user.ClubID,
		user.ChessAgeCategory,
		user.ELOFideStandard,
		user.ELOFideRapid,
		user.ELOFideBlitz,
		user.ELOFideBullet,
		user.ELONationalStandard,
		user.ELONationalRapid,
		user.ELONationalBlitz,
		user.ELONationalBullet,
		user.ELORegionalStandard,
		user.ELORegionalRapid,
		user.ELORegionalBlitz,
		user.EloRegionalBullet,
		user.IsArbiter,
		user.IsCoach,
		user.PricePerHour,
		user.Currency,
		user.ChessComUsername,
		user.LichessUsername,
		user.Chess24Username,
		user.Country,
		user.Province,
		user.City,
		user.Neighborhood,
	}
	return m.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt, &user.UserCode)
}

func (m UserModel) GetByEmail(user *User) error {
	query := `
        SELECT * FROM users
        WHERE email = $1`
	return m.DB.QueryRow(query, user.Email).Scan(
		&user.Email,
		&user.Password,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Avatar,
		&user.ClubID,
		&user.Sex,
		&user.AboutMe,
		&user.Country,
		&user.Province,
		&user.City,
		&user.Neighborhood,
		&user.UserCode,
		&user.ChessAgeCategory,
		&user.ELOFideStandard,
		&user.ELOFideRapid,
		&user.ELONationalStandard,
		&user.ELONationalRapid,
		&user.ELORegionalStandard,
		&user.ELORegionalRapid,
		&user.Version,
	)
}

func (m UserModel) Update(nd *User) error {
	fmt.Println("user", nd)
	query := `
            UPDATE users
		        SET
		        is_active = CASE WHEN $1. IS NULL THEN is_active ELSE $1 END,
		        is_verified = CASE WHEN $2 IS NULL THEN is_verified ELSE $2 END,
		        is_admin_of_club = CASE WHEN $3 IS NULL THEN is_admin_of_club ELSE $3 END,
		        club_admin_of = CASE WHEN $4 IS NULL THEN club_admin_of ELSE $4 END,
		        deleted_at = CASE WHEN $5 IS NULL THEN deleted_at ELSE $5 END,
		        first_name = CASE WHEN $6 IS NULL THEN first_name ELSE $6 END,
		        last_name = CASE WHEN $7 IS NULL THEN last_name ELSE $7 END,
		        dob = CASE WHEN $8 IS NULL THEN dob ELSE $8 END,
		        sex = CASE WHEN $9 IS NULL THEN sex ELSE $9 END,
		        username = CASE WHEN $10 IS NULL THEN username ELSE $10 END,
		        email = CASE WHEN $11 IS NULL THEN email ELSE $11 END,
		        password = CASE WHEN $12 IS NULL THEN password ELSE $12 END,
		        password_reset_token = CASE WHEN $13 IS NULL THEN password_reset_token ELSE $13 END,
		        avatar = CASE WHEN $14 IS NULL THEN avatar ELSE $14 END,
		        club_id = CASE WHEN $15 IS NULL THEN club_id ELSE $15 END,
		        club_role_id = CASE WHEN $16 IS NULL THEN club_role_id ELSE $16 END,
		        about_me = CASE WHEN $17 IS NULL THEN about_me ELSE $17 END,
		        is_arbiter = CASE WHEN $18 IS NULL THEN is_arbiter ELSE $18 END,
		        is_coach = CASE WHEN $19 IS NULL THEN is_coach ELSE $19 END,
		        price_per_hour = CASE WHEN $20 IS NULL THEN price_per_hour ELSE $20 END,
		        chess_com_username = CASE WHEN $21 IS NULL THEN chess_com_username ELSE $21 END,
		        lichess_username = CASE WHEN $22 IS NULL THEN lichess_username ELSE $22 END,
		        chess24_username = CASE WHEN $23 IS NULL THEN chess24_username ELSE $23 END,
		        country = CASE WHEN $24 IS NULL THEN country ELSE $24 END,
		        province = CASE WHEN $25 IS NULL THEN province ELSE $25 END,
		        city = CASE WHEN $26 IS NULL THEN city ELSE $26 END,
		        neighborhood = CASE WHEN $27 IS NULL THEN neighborhood ELSE $27 END,
		        elo_fide_standard = CASE WHEN $28 IS NULL THEN elo_fide_standard ELSE $28 END,
		        elo_fide_rapid = CASE WHEN $29 IS NULL THEN elo_fide_rapid ELSE $29 END,
		        elo_national_standard = CASE WHEN $30 IS NULL THEN elo_national_standard ELSE $30 END,
		        elo_national_rapid = CASE WHEN $31 IS NULL THEN elo_national_rapid ELSE $31 END,
		        elo_regional_standard = CASE WHEN $32 IS NULL THEN elo_regional_standard ELSE $32 END,
		        club_user_code = CASE WHEN $33 IS NULL THEN club_user_code ELSE $33 END,
		        chess_age_category = CASE WHEN $34 IS NULL THEN chess_age_category ELSE $34 END,
		        elo_regional_rapid = CASE WHEN $35 IS NULL THEN elo_regional_rapid ELSE $35 END,
		        version = version + 1
            WHERE user_code = $36
            RETURNING version`
	/*
			query := `
		        UPDATE users

		        SET
		         is_active = $1,
		         is_verified = $2,
		         is_admin_of_club = $3,
		         club_admin_of = $4,
		         deleted_at = $5,
		         first_name = $6,
		         last_name = $7,
		         dob = $8,
		         sex = $9,
		         username = $10,
		         email = $11,
		         password = $12,
		         password_reset_token = $13,
		         avatar = $14,
		         club_id = $15,
		         club_role_id = $16,
		         about_me = $17,
		         is_arbiter = $18,
		         is_coach = $19,
		         price_per_hour = $20,
		         chess_com_username = $21,
		         lichess_username = $22,
		         chess24_username = $23,
		         country = $24,
		         province = $25,
		         city = $26,
		         neighborhood = $27,
		         elo_fide_standard = $28,
		         elo_fide_rapid = $29,
		         elo_national_standard = $30,
		         elo_national_rapid = $31,
		         elo_regional_standard = $32,
		         club_user_code = $33,
		         chess_age_category = $34,
		         elo_regional_rapid = $35,
		         version = version + 1
		        WHERE user_code = $36
		        RETURNING version`
	*/

	args := []any{
		nd.IsActive,
		nd.IsVerified,
		nd.FirstName,
		nd.LastName,
		nd.DateOfBirth,
		nd.Sex,
		nd.Username,
		nd.Email,
		nd.Password,
		nd.PasswordResetToken,
		nd.Avatar,
		nd.ClubID,
		nd.AboutMe,
		nd.IsArbiter,
		nd.IsCoach,
		nd.PricePerHour,
		nd.ChessComUsername,
		nd.LichessUsername,
		nd.Chess24Username,
		nd.Country,
		nd.Province,
		nd.City,
		nd.Neighborhood,
		nd.ELOFideStandard,
		nd.ELOFideRapid,
		nd.ELONationalStandard,
		nd.ELONationalRapid,
		nd.ELORegionalStandard,
		nd.ChessAgeCategory,
		nd.ELORegionalRapid,
		nd.UserCode,
	}
	a := fmt.Sprintf("%#v", nd)
	fmt.Println("asdasd", a)
	return m.DB.QueryRow(query, args...).Scan(&nd.Version)
	return nil
}

func (m UserModel) Delete(email string) error {
	query := `
        DELETE FROM users
        WHERE email = $1`
	_, err := m.DB.Exec(query, email)
	return err
}

func (m UserModel) GetByUserCode(code string) (*User, error) {
	var u User
	fmt.Printf("newData that's coming in: %#v", code)
	query := `
    SELECT
        is_active,
        is_verified,
        is_admin_of_club,
        club_admin_of,
        created_at,
        updated_at,
        deleted_at,
        first_name,
        last_name,
        dob,
        sex,
        username,
        email,
        avatar,
        club_id,
        club_role_id,
        about_me,
        is_arbiter,
        is_coach,
        price_per_hour,
        chess_com_username,
        lichess_username,
        chess24_username,
        country,
        province,
        city,
        neighborhood,
        elo_fide_standard,
        elo_fide_rapid,
        elo_national_standard,
        elo_national_rapid,
        elo_regional_standard,
        elo_regional_rapid,
        club_user_code,
        chess_age_category,
        version,
        user_code
    FROM users
    WHERE user_code = $1`

	err := m.DB.QueryRow(query, code).Scan(
		&u.IsActive,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.FirstName,
		&u.LastName,
		&u.DateOfBirth,
		&u.Sex,
		&u.Username,
		&u.Email,
		&u.Avatar,
		&u.ClubID,
		&u.AboutMe,
		&u.IsArbiter,
		&u.IsCoach,
		&u.PricePerHour,
		&u.ChessComUsername,
		&u.LichessUsername,
		&u.Chess24Username,
		&u.Country,
		&u.Province,
		&u.City,
		&u.Neighborhood,
		&u.ELOFideStandard,
		&u.ELOFideRapid,
		&u.ELONationalStandard,
		&u.ELONationalRapid,
		&u.ELORegionalStandard,
		&u.ELORegionalRapid,
		&u.ChessAgeCategory,
		&u.Version,
		&u.UserCode,
	)
	// fmt.Printf("\n user after being fetched: %#v\n", u)
	// fmt.Printf("error after being fetched: %#v", err)
	return &u, err
}
