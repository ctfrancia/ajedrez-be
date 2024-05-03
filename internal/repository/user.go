package repository

import (
	"context"
	"crypto/sha256"
	"ctfrancia/ajedrez-be/internal/models"
	// "database/sql"
	"errors"
	"fmt"
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
	fmt.Println("user before save", user)

	result := r.DB.WithContext(ctx).Model(&models.User{}).Where("id", user["id"]).Updates(user)
	if result.Error != nil {
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
	sQ := "users.id, users.created_at, users.last_name, users.email, users.user_code, users.password, users.activated, users.version"
	jQ := "INNER JOIN tokens ON users.id = tokens.user_id"
	wQ := "tokens.hash = $1 AND tokens.scope = $2 AND tokens.expiry > $3"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res := r.DB.WithContext(ctx).Model(&models.User{}).Select(sQ).Joins(jQ).Where(wQ, tokenHash[:], tokenScope, time.Now()).Find(&user)

	if errors.Is(r.DB.Error, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}

	return &user, res.Error
}

func (r UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res := r.DB.WithContext(ctx).Where("email = ?", email).First(&user)

	if errors.Is(r.DB.Error, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}

	return &user, res.Error
}
