package service

import (
	"my-api/data/request"
	"my-api/data/response"
	"my-api/model"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
	FindAdmin() *model.User
}