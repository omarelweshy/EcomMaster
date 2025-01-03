package service

import (
	"errors"

	"user-service/internal/model"
	"user-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Repo repository.UserRepository
}

var (
	ErrUsernameTaken      = errors.New("username already taken")
	ErrEmailRegistered    = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *UserService) RegisterUser(username, email, password, firstName, lastName string) error {
	_, err := s.Repo.GetUserByUsername(username)
	if err == nil {
		return ErrUsernameTaken
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	_, err = s.Repo.GetUserByEmail(email)
	if err == nil {
		return ErrEmailRegistered
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		FirstName:    firstName,
		LastName:     lastName,
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}
	return s.Repo.CreateUser(&user)
}

func (s *UserService) LoginUser(username, password string) (*model.User, error) {
	user, err := s.Repo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}
