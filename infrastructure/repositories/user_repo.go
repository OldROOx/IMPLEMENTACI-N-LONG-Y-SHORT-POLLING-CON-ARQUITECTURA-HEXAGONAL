package repositories

import (
	"APIHEX_LPySP/config"
	"APIHEX_LPySP/domain"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	result := config.DB.Preload("Products").First(&user, id)
	return &user, result.Error
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	result := config.DB.Create(user)
	return result.Error
}

func (r *UserRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	result := config.DB.Save(user)
	return result.Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	result := config.DB.Delete(&domain.User{}, id)
	return result.Error
}
