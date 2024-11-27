package controller

import (
	"github.com/gin-gonic/gin"
	"my-api/data/request"
	"my-api/data/response"
	"my-api/helper"
	"my-api/service"
	"net/http"
	"strconv"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{bookService: service}
}

func (controller *BookController) Create(ctx *gin.Context) {
	createBookRequest := request.CreateBookRequest{}
	err := ctx.ShouldBindJSON(&createBookRequest)
	helper.ErrorPanic(err)

	controller.bookService.Create(createBookRequest)

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "The book has been created",
		Data:   nil,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *BookController) Update(ctx *gin.Context) {
	updateBookRequest := request.UpdateBookRequest{}
	err := ctx.ShouldBindJSON(&updateBookRequest)
	helper.ErrorPanic(err)

	BookId := ctx.Param("BookId")
	id, err := strconv.Atoi(BookId)
	helper.ErrorPanic(err)

	updateBookRequest.Id = id

	controller.bookService.Update(updateBookRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "The Book are updated",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookController) Delete(ctx *gin.Context) {
	BookId := ctx.Param("BookId")
	id, err := strconv.Atoi(BookId)
	helper.ErrorPanic(err)
	controller.bookService.Delete(id)

	webResponse := response.Response{
		Code:   204,
		Status: "The Book are removed or don't exist",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *BookController) FindById(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	helper.ErrorPanic(err)

	bookResponse := controller.bookService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookController) FindAll(ctx *gin.Context) {
	bookResponse := controller.bookService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *BookController) FindByTitle(ctx *gin.Context) {
    title := ctx.Query("title")
    books := controller.bookService.FindByTitle(title)

    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Books retrieved successfully",
        Data:   books,
    }

    ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookController) FindByTitleGuess(ctx *gin.Context) {
    title := ctx.Query("title")
    books := controller.bookService.FindByTitleGuess(title)


    webResponse := response.Response{
        Code:   http.StatusOK,
        Status: "Books retrieved successfully",
        Data:   books,
    }

    ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookController) FindAllGuess(ctx *gin.Context) {
	bookResponse := controller.bookService.FindAllGuess()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
