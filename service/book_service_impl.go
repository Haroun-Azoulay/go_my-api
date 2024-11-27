package service

import (
	"github.com/go-playground/validator/v10"
	"my-api/data/request"
	"my-api/helper"
	"my-api/model"
	"my-api/repository"
	"my-api/data/response"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	Validate       *validator.Validate
}

func NewBookServiceImpl(bookRepository repository.BookRepository, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		Validate:       validate,
	}
}

func (t BookServiceImpl) Create(bookRequest request.CreateBookRequest) {
	err := t.Validate.Struct(bookRequest)
	helper.ErrorPanic(err)

	bookModel := model.Book{
		Title:   bookRequest.Title,
		Author:  bookRequest.Author,
		Release: bookRequest.Release,
		Resume:  bookRequest.Resume,
		Stock:   bookRequest.Stock,
		Price:   bookRequest.Price,
	}

	t.BookRepository.Save(bookModel)
}

func (t BookServiceImpl) Update(Book request.UpdateBookRequest) {
	BookData, err := t.BookRepository.FindById(Book.Id)
	helper.ErrorPanic(err)
	BookData.Title = Book.Title
	BookData.Author = Book.Author
	BookData.Release = Book.Release
	BookData.Stock = Book.Stock
	BookData.Price = Book.Price
	t.BookRepository.Update(BookData)
}

func (t BookServiceImpl) Delete(BookId int) {
	t.BookRepository.Delete(BookId)
}

func (t BookServiceImpl) FindById(BookId int) response.BookResponse {
	BookData, err := t.BookRepository.FindById(BookId)
	helper.ErrorPanic(err)

	BookResponse := response.BookResponse{
		Id:   BookData.Id,
		Title: BookData.Title,
		Author: BookData.Author,
		Release: BookData.Release,
		Resume: BookData.Resume,
		Stock: BookData.Stock,
		Price: BookData.Price,
	}
	return BookResponse
}

func (t BookServiceImpl) FindAll() []response.BookResponse {
	result := t.BookRepository.FindAll()

	var books []response.BookResponse
	for _, value := range result {
		book := response.BookResponse{
			Id:   value.Id,
			Title: value.Title,
			Author: value.Author,
			Release: value.Release,
			Resume: value.Resume,
			Stock: value.Stock,
			Price: value.Price,
		}
		books = append(books, book)
	}
	return books
}

func (t BookServiceImpl) FindByTitle(title string) []response.BookResponse {
    books, err := t.BookRepository.FindByTitle(title)
    helper.ErrorPanic(err)

    var bookResponses []response.BookResponse
    for _, book := range books {
        bookResponse := response.BookResponse{
            Id:      book.Id,
            Title:   book.Title,
            Author:  book.Author,
            Release: book.Release,
            Resume:  book.Resume,
            Stock:   book.Stock,
            Price:   book.Price,
        }
        bookResponses = append(bookResponses, bookResponse)
    }
    return bookResponses
}

func (t BookServiceImpl) FindByTitleGuess(title string) []response.GuessBookResponse {
    books, err := t.BookRepository.FindByTitleGuess(title)
    helper.ErrorPanic(err)

    var bookResponsesGuess []response.GuessBookResponse
    for _, book := range books {
        if book.Stock {
            bookResponseGuess := response.GuessBookResponse{
                Id:      book.Id,
                Title:   book.Title,
                Author:  book.Author,
				Resume:  book.Resume,
				Price:   book.Price,
            }
            bookResponsesGuess = append(bookResponsesGuess, bookResponseGuess)
        }
    }
    return bookResponsesGuess
}

func (t BookServiceImpl) FindAllGuess() []response.GuessAllBookResponse {
	result := t.BookRepository.FindAllGuess()

	var books []response.GuessAllBookResponse
	for _, value := range result {
		if value.Stock {
			book := response.GuessAllBookResponse{
				Id:     value.Id,
				Author: value.Author,
				Title:  value.Title,
			}
			books = append(books, book)
		}
	}
	return books
}
