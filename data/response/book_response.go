package response

import "time"

type BookResponse struct {
	Id        int    `json:"id"`
	Title   string    `json:"title" validate:"required"`
	Author  string    `json:"author" validate:"required"`
	Release *time.Time `json:"release" validate:"required"`
	Resume  string    `json:"resume"`
	Stock   bool       `json:"stock"`
	Price   float64   `json:"price" validate:"required,min=0"`
}

type GuessBookResponse struct {
	Id        int    `json:"id"`
	Title   string    `json:"title" validate:"required"`
	Author  string    `json:"author" validate:"required"`
	Resume  string    `json:"resume"`
	Price   float64   `json:"price" validate:"required,min=0"`
}

type GuessAllBookResponse struct {
	Id        int    `json:"id"`
	Title   string    `json:"title" validate:"required"`
	Author  string    `json:"author" validate:"required"`
}
