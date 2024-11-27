package request

type CreateUserRequest struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Email     string `json:"email"`
    Password  string `json:"password"`
	IsAdmin bool `json:"isadmin"`
}
