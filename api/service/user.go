package service

import (
	"api/auth"
	"api/repository"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(email, password, name string) (err error)
	SignIn(email, password string) (token string, err error)
}

func NewUserService() UserService {
	return &userService{
		repository: repository.NewUserRepository(),
	}
}

func NewUserServiceWithRepository(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

type userService struct {
	repository repository.UserRepository
}

func (u *userService) SignUp(email, password, name string) (err error) {
	// encrypt password
	digest, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	passwordDigest := string(digest)
	return u.repository.CreateUser(email, passwordDigest, name)
}

func (u *userService) SignIn(email, password string) (token string, err error) {
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return
	}
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password)) != nil {
		return "", fmt.Errorf("incorrect email or password")
	}

	token, err = auth.Sign(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return
	}

	return token, nil
}
