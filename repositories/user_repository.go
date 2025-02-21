package repositories

import (
	"go-app/config"
	"go-app/models"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}
