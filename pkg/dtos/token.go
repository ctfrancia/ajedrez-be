package dtos

import "time"

type ActivateTokenDTO struct {
	Token string `json:"token"`
}

type AuthenticateUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenExpiryDTO struct {
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`
}
