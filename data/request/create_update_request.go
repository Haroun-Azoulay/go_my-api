package request

type UpdateUserRequest struct {
	Id   int    `validate:"required"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsAdmin bool `json:"isadmin"`
}