package repository

import (
	"errors"
	"my-api/data/request"
	"my-api/helper"
	"my-api/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (t UserRepositoryImpl) Save(User model.User) {
	result := t.Db.Create(&User)
	helper.ErrorPanic(result.Error)

}

func (t UserRepositoryImpl) Update(User model.User) {
	var updateTag = request.UpdateUserRequest{
		Id:   User.Id,
		Firstname: User.Firstname,
		Lastname: User.Lastname,
		Email: User.Email,
		Password: User.Password,
		IsAdmin: User.IsAdmin,
	}
	result := t.Db.Model(&User).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t UserRepositoryImpl) Delete(UserId int) {
	var User model.User
	result := t.Db.Where("id = ?", UserId).Delete(&User)
	helper.ErrorPanic(result.Error)
}

func (t UserRepositoryImpl) FindById(UserId int) (model.User, error) {
	var tag model.User
	result := t.Db.Find(&tag, UserId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("user is not found")
	}
}

func (t UserRepositoryImpl) FindAll() []model.User {
	var User []model.User
	results := t.Db.Find(&User)
	helper.ErrorPanic(results.Error)
	return User
}