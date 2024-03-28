package data

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
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

func (m UserModel) Update(nd map[string]interface{}) (int, error) {
	var user *User
	user = &User{}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	u := psql.Update("users")
	for key, value := range nd {
		if key == "'user_code'" {
			continue
		}
		u = u.Set(key, value)
	}
	u = u.Set("updated_at", time.Now())
	u = u.Set("version", sq.Expr("version + 1"))
	u = u.Where(squirrel.Eq{"user_code": nd["user_code"]})
	u = u.Suffix("RETURNING \"version\"")

	query, args, err := u.ToSql()
	if err != nil {
		fmt.Println("error creating query: ", err)
		return 0, err
	}

	err = m.DB.QueryRow(query, args...).Scan(&user.Version)

	return user.Version, err
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
