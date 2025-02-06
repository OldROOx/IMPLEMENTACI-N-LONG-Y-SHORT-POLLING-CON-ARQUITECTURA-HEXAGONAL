package application

import (
	"APIHEX_LPySP/domain"
	"APIHEX_LPySP/infrastructure/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUsers() ([]domain.User, error) {
	return s.userRepo.GetUsers()
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}
