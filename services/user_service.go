package services

import (
	"go-app/models"
	"go-app/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: &repositories.UserRepository{},
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}
