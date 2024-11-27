package repository

import (
	"my-api/model"
)

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(user int)
	FindById(userId int) (user model.User, err error)
	FindAll() []model.User
	FindByCondition(condition string, args ...interface{}) *model.User
}
