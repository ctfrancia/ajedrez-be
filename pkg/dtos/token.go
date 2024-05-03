package dtos

type ActivateTokenDTO struct {
	Token string `json:"token"`
}

type AuthenticateUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
