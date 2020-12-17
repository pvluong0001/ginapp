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

func (service *AuthService) RegisterUserHandle(user *models.User) error {
	err := service.UserRepository.Create(user)

	if err != nil {
		return err
	}

	// send register mail
	/*go func() {
		email := emails2.RegisterMail{
			Recipient: user.Email,
			From:      "pvluong0001@gmail.com",
			Subject:   "Register Email",
			Message:   "Message body",
		}
		email.Handle()
	}()*/

	return nil
}

func (service AuthService) GetUserByEmail(user *models.User) error {
	err := service.UserRepository.FindByEmail(user)

	if err != nil {
		return err
	}

	return nil
}
