package dtos

type UserCreateRequest struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Age        int    `json:"age" binding:"required"`
	Birthday   string `json:"birthday" binding:"required"`
	ProfilePic string `json:"profile_pic"`
	Club       string `json:"club" binding:"required"`
}
