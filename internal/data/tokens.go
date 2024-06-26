package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"errors"
	"time"
)

const (
	ScopeActivation       = "activation"
	ScopeAuthentication   = "authentication"
	ErrTokenRequired      = "token must be provided"
	ErrTokenTooShort      = "token must be 26 bytes long"
	ErrTokenInvalid       = "invalid token"
	ErrInvalidCredentials = "invalid credentials"
)

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int64     `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

type TokenModel struct {
	DB *sql.DB
}

func generateToken(userID int64, lifetime time.Duration, scope string) (*Token, error) {
	// Initialize a new Token struct, setting the UserID, Expiry, and Scope fields.
	token := Token{
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

	// Encode the random bytes as a base32 string.
	// sometimes it is necessary to remove the padding from the base32 encoding
	// that's why we use the WithPadding(base32.NoPadding) method.
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// Calculate the SHA-256 hash of the plaintext token.
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return &token, nil
}

func ValidateTokenPlaintext(tokenPlaintext string) error {
	if tokenPlaintext == "" {
		return errors.New(ErrTokenRequired)
	}

	if len(tokenPlaintext) != 26 {
		return errors.New(ErrTokenTooShort)
	}

	return nil
}

func (m TokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = m.Insert(token)
	return token, err
}

func (m TokenModel) Insert(token *Token) error {
	query := `
        INSERT INTO tokens (hash, user_id, expiry, scope)
        VALUES ($1, $2, $3, $4)
    `
	args := []interface{}{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m TokenModel) DeleteAllForUser(scope string, userID int64) error {
	query := `
        DELETE FROM tokens 
        WHERE scope = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, scope, userID)
	return err
}
