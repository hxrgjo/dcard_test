package repository

import (
	"api/model"
	"fmt"

	"github.com/go-xorm/xorm"
)

type UserRepository interface {
	CreateUser(email, passwordDigest, name string) (err error)
	FindByEmail(email string) (user model.User, err error)
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
	q := `INSERT INTO users(email, password_digest, name, created_at, updated_at)
	      SELECT ?, ?, ?, NOW(), NOW()
	      WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = ?)`
	result, err := u.db.Exec(q, email, passwordDigest, name, email)
	if err != nil {
		return
	}
	if affected, _ := result.RowsAffected(); affected == 0 {
		return fmt.Errorf("email already exists")
	}
	return
}

func (u *userRepository) FindByEmail(email string) (user model.User, err error) {
	if _, err = u.db.Where("email = ?", email).Get(&user); err != nil {
		return
	}
	return
}
