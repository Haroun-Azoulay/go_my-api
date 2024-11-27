package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"my-api/data/request"
	"my-api/data/response"
	"my-api/database"
	"my-api/model"
	"my-api/helper"
	"my-api/service"
	"my-api/security"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

func (controller *UserController) Create(ctx *gin.Context) {
    createUserRequest := request.CreateUserRequest{}
    err := ctx.ShouldBindJSON(&createUserRequest)
	
    helper.ErrorPanic(err)

    isAdmin := createUserRequest.Email == "Saber.essakhori@inwi.ma"

	if createUserRequest.Email == "Saber.essakhori@inwi.ma" {
		createUserRequest.IsAdmin = true
	}

    if isAdmin {
        existingAdmin := controller.userService.FindAdmin()
        if existingAdmin != nil {
            webResponse := response.Response{
                Code:   http.StatusForbidden,
                Status: "An administrator already exists",
                Data:   nil,
            }

            ctx.JSON(http.StatusForbidden, webResponse)
            return
        }
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
    helper.ErrorPanic(err)
    createUserRequest.Password = string(hashedPassword)


    controller.userService.Create(createUserRequest)

    webResponse := response.Response{
        Code:   http.StatusCreated,
        Status: "The user has been created",
        Data:   nil,
    }

    ctx.JSON(http.StatusCreated, webResponse)
}


func Login(ctx *gin.Context) {

	var authInput model.User

	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound model.User
	database.DB.Where("firstname=?", authInput.Firstname).Find(&userFound)

	if userFound.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	tokenString, err := security.CreateToken(userFound.Firstname, userFound.IsAdmin)
    helper.ErrorPanic(err)


	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	webResponse := response.Response{
        Code:   200,
        Status: "The user has been login",
        Data:   nil,
		Token: tokenString,
    }


    ctx.JSON(http.StatusCreated, webResponse)
}



func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.userService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "The user are updated",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.userService.Delete(id)

	webResponse := response.Response{
		Code:   204,
		Status: "The user are removed or don't exist",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.userService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userResponse := controller.userService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}