package request

import "time"

type UpdateUserRequest struct {
	Id        int    `validate:"required"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isadmin"`
}

type UpdateBookRequest struct {
	Id      int        `json:"id" validate:"required"`
	Title   string     `json:"title" validate:"required"`
	Author  string     `json:"author" validate:"required"`
	Release *time.Time `json:"release"`
	Resume  string     `json:"resume"`
	Stock   bool       `json:"stock"`
	Price   float64    `json:"price" validate:"required,min=0"`
}
