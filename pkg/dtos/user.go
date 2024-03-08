package dtos

type UserCreateRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Birthday   string `json:"birthday"`
	ProfilePic string `json:"profile_pic"`
	Club       string `json:"club"`
}
