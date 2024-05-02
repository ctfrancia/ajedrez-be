package repository

import (
	"context"
	"crypto/sha256"
	"ctfrancia/ajedrez-be/internal/models"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) Create(user *models.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return ErrDuplicateEmail

		default:
			return result.Error
		}
	}

	return result.Error
}

func (r UserRepository) Update(user map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result := r.DB.WithContext(ctx).Model(&models.User{}).Updates(user)
	if result.Error != nil {
		// if the error is not nil, we need to check if the error is a gorm.ErrRecordNotFound
		// with all the others to normalize the error
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return ErrRecordNotFound
		default:
			return result.Error
		}
	}

	return nil
}

func (r UserRepository) GetForToken(tokenScope, tokenPlainText string) (*models.User, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlainText))
	var user models.User

	/*
			query := `
		        SELECT users.id, users.created_at, users.last_name, users.email, users.password, users.activated, users.version
		        FROM users
		        INNER JOIN tokens
		        ON users.id = tokens.user_id
		        WHERE tokens.hash = $1
		        AND tokens.scope = $2
		        AND tokens.expiry > $3`
	*/

	// args := []any{tokenHash[:], tokenScope, time.Now()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res := r.DB.WithContext(ctx).Model(&models.User{}).Select("users.id, users.created_at, users.last_name, users.email, users.password, users.activated, users.version").Joins("INNER JOIN tokens ON users.id = tokens.user_id").Where("tokens.hash = $1 AND tokens.scope = $2 AND tokens.expiry = $3", tokenHash[:], tokenScope, time.Now()).Scan(&user)

	if errors.Is(r.DB.Error, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}

	return &user, res.Error
}
