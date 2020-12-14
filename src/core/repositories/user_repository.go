package repositories

import (
	"ginapp/core/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (repo *UserRepository) FindByEmail(user *models.User) error {
	err := repo.DB.Where("email = ?", user.Email).First(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) All() []models.User {
	var result []models.User
	repo.DB.Find(&result)

	return result
}

func (repo *UserRepository) Create(user *models.User) error {
	err := repo.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}
