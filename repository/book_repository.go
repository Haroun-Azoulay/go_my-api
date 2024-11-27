package repository

import (
	"my-api/model"
)

type BookRepository interface {
	Save(book model.Book)
	Update(book model.Book)
	Delete(book int)
	FindById(bookId int) (book model.Book, err error)
	FindAll() []model.Book
	FindByTitle(title string) ([]model.Book, error)
	FindByTitleGuess(title string) ([]model.Book, error)
	FindAllGuess() []model.Book
}
