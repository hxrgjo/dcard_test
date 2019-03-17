package repository

import (
	"api/model"

	"github.com/go-xorm/xorm"
)

type UserRepository interface {
	CreateUser(email, passwordDigest, name string) (err error)
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: model.GetDB(),
	}
}

type userRepository struct {
	db *xorm.Engine
}

func (u *userRepository) CreateUser(email, passwordDigest, name string) (err error) {
	user := model.User{Email: email, PasswordDigest: passwordDigest, Name: name}
	if _, err = u.db.InsertOne(&user); err != nil {
		return
	}
	return
}
