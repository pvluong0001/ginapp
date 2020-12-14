package services

import (
	"ginapp/core/models"
	"ginapp/core/repositories"
)

type AuthService struct {
	UserRepository repositories.UserRepository
}

func ProvideAuthService(repository repositories.UserRepository) AuthService {
	return AuthService{UserRepository: repository}
}

func (service AuthService) RegisterUserHandle(user *models.User) {
	err := service.UserRepository.Create(user)

	if err != nil {
		panic(err)
	}
}
