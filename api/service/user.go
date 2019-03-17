package service

import (
	"api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(email, password, name string) (err error)
}

func NewUserService() UserService {
	return userService{
		repository: repository.NewUserRepository(),
	}
}

func NewUserServiceWithRepository(repository repository.UserRepository) UserService {
	return userService{
		repository: repository,
	}
}

type userService struct {
	repository repository.UserRepository
}

func (u userService) SignUp(email, password, name string) (err error) {
	// encrypt password
	digest, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	passwordDigest := string(digest)
	return u.repository.CreateUser(email, passwordDigest, name)
}
