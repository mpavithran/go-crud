package services

import (
	"github.com/mpavithran/go-crud/models"
	"github.com/mpavithran/go-crud/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
