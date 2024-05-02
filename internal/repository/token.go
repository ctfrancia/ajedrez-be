package repository

import (
	"ctfrancia/ajedrez-be/internal/models"
	"time"

	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"gorm.io/gorm"
)

type TokensRepository struct {
	db *gorm.DB
}

func (r TokensRepository) Insert(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r TokensRepository) New(userID int64, ttl time.Duration, scope string) (*models.Token, error) {

	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = r.Insert(token)
	return token, err
}

func generateToken(userID int64, lifetime time.Duration, scope string) (*models.Token, error) {
	token := models.Token{
		UserID: userID,
		Expiry: time.Now().Add(lifetime),
		Scope:  scope,
	}

	// Generate a random 16-byte authentication token.
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// sometimes it is necessary to remove the padding from the base32 encoding
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// Calculate the SHA-256 hash of the plaintext token.
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return &token, nil
}

func (r TokensRepository) DeleteAllForUser(scope string, userID int64) error {
	/* query := `
	   DELETE FROM tokens
	   WHERE scope = $1 AND user_id = $2`
	*/

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return r.db.WithContext(ctx).Delete(&models.Token{}, "scope = ? AND user_id = ?", scope, userID).Error
}
