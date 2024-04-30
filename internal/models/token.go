package models

import (
	"time"
)

const TokenLifetime = 3 * 24 * time.Hour

type Token struct {
	Plaintext string    // `json:"token"`
	Hash      []byte    // `json:"-"`
	UserID    int64     // `json:"-"`
	Expiry    time.Time // `json:"expiry"`
	Scope     string    // `json:"-"`
}
