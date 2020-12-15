package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	defaultCost = 10
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (user *User) BeforeSave(db *gorm.DB) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), defaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashPassword)
	return nil
}

func (user *User) ComparePassword(password string) error {
	hashPassword := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(hashPassword, []byte(password)); err != nil {
		return err
	}

	return nil
}
