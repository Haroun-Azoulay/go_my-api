package service

import (
	"my-api/data/request"
	"my-api/data/response"
)

type BookService interface {
	Create(book request.CreateBookRequest)
	Update(book request.UpdateBookRequest)
	Delete(bookId int)
	FindById(bookId int) response.BookResponse
	FindAll() []response.BookResponse
	FindByTitle(title string) []response.BookResponse
	FindByTitleGuess(title string) []response.GuessBookResponse
	FindAllGuess() []response.GuessAllBookResponse
}
