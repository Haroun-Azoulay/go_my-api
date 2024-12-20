package repository

import (
	"errors"
	"gorm.io/gorm"
	"my-api/data/request"
	"my-api/helper"
	"my-api/model"
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
	var updateUser = request.UpdateUserRequest{
		Id:        User.Id,
		Firstname: User.Firstname,
		Lastname:  User.Lastname,
		Email:     User.Email,
		Password:  User.Password,
		IsAdmin:   User.IsAdmin,
	}
	result := t.Db.Model(&User).Updates(updateUser)
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

func (r *UserRepositoryImpl) FindByCondition(condition string, args ...interface{}) *model.User {
	var user model.User
	result := r.Db.Where(condition, args...).First(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}
