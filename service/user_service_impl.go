package service

import (
	"github.com/go-playground/validator/v10"
	"my-api/data/request"
	"my-api/data/response"
	"my-api/helper"
	"my-api/model"
	"my-api/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	Validate      *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		Validate:      validate,
	}
}

func (t UserServiceImpl) Create(user request.CreateUserRequest) {
	err := t.Validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.User{
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Email: user.Email,
		Password: user.Password,
	}
	t.userRepository.Save(userModel)
}

func (t UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := t.userRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Firstname = user.Firstname
	userData.Lastname = user.Lastname
	userData.Email = user.Email
	userData.Password = user.Password
	t.userRepository.Update(userData)
}

func (t UserServiceImpl) Delete(userId int) {
	t.userRepository.Delete(userId)
}

func (t UserServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := t.userRepository.FindById(userId)
	helper.ErrorPanic(err)

	userResponse := response.UserResponse{
		Id:   userData.Id,
		Firstname: userData.Firstname,
		Lastname: userData.Lastname,
		Email: userData.Email,
		Password: userData.Password,
	}
	return userResponse
}

func (t UserServiceImpl) FindAll() []response.UserResponse {
	result := t.userRepository.FindAll()

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:   value.Id,
			Firstname: value.Firstname,
			Lastname: value.Lastname,
			Email: value.Email,
			Password: value.Password,
		}
		users = append(users, user)
	}
	return users
}