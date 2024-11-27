package repository

import (
	"errors"
	"gorm.io/gorm"
	"my-api/helper"
	"my-api/model"
	"my-api/data/request"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepositoryImpl(Db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

func (t BookRepositoryImpl) Save(Book model.Book) {

	result := t.Db.Create(&Book)
	helper.ErrorPanic(result.Error)

}

func (t BookRepositoryImpl) Update(Book model.Book) {
    updateBook := request.UpdateBookRequest{
        Id:      Book.Id,
        Title:   Book.Title,
        Author:  Book.Author,
        Release: Book.Release,
        Resume:  Book.Resume,
        Stock:   Book.Stock,
        Price:   Book.Price,
    }
    result := t.Db.Model(&Book).Updates(updateBook)
    helper.ErrorPanic(result.Error)
}

func (t BookRepositoryImpl) Delete(BookId int) {
	var Book model.Book
	result := t.Db.Where("id = ?", BookId).Delete(&Book)
	helper.ErrorPanic(result.Error)
}

func (t BookRepositoryImpl) FindById(BookId int) (model.Book, error) {
	var tag model.Book
	result := t.Db.Find(&tag, BookId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("Book is not found")
	}
}

func (t BookRepositoryImpl) FindAll() []model.Book {
	var Book []model.Book
	results := t.Db.Find(&Book)
	helper.ErrorPanic(results.Error)
	return Book
}

func (t BookRepositoryImpl) FindByTitle(title string) ([]model.Book, error) {
    var books []model.Book
    result := t.Db.Where("title LIKE ?", "%"+title+"%").Find(&books)
    if result.Error != nil {
        return nil, result.Error
    }
    return books, nil
}

func (t BookRepositoryImpl) FindByTitleGuess(title string) ([]model.Book, error) {
    var books []model.Book
    result := t.Db.Where("title LIKE ?", "%"+title+"%").Find(&books)
    if result.Error != nil {
        return nil, result.Error
    }
    return books, nil
}

func (t BookRepositoryImpl) FindAllGuess() []model.Book {
	var Book []model.Book
	results := t.Db.Find(&Book)
	helper.ErrorPanic(results.Error)
	return Book
}

