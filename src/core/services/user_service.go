package services

import (
	"ginapp/core/models"
	"ginapp/core/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func ProvideUserService(repo repositories.UserRepository) UserService {
	return UserService{UserRepository: repo}
}

func (service *UserService) GetAllUserHandle() []models.User {
	return service.UserRepository.All()
}
