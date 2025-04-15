package services

import (
	"errors"
	"pet_project_1/models"
	"pet_project_1/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.User) error
	Login(username, password string) (*models.User, error)
	Update(user *models.User) error
	Delete(userID string) error
	GetUserByID(userID string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	_, err := s.repo.FindByUsername(user.Username)
	if err == nil {
		return ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}

func (s *userService) Login(username, password string) (*models.User, error) {
	dbUser, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return dbUser, nil
}

func (s *userService) Update(user *models.User) error {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return s.repo.Update(user)
}

func (s *userService) Delete(userID string) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

var (
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

func (s *userService) GetUserByID(userID string) (*models.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
