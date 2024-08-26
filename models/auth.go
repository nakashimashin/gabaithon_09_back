package models

import (
	"crypto/sha256"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Encrypt(char string) string {
	encryptText := fmt.Sprintf("%x", sha256.Sum256([]byte(char)))
	return encryptText
}

func (user *User) Create(db *gorm.DB) (User, error) {
	newUser := User{
		Name:     user.Name,
		Email:    user.Email,
		Password: Encrypt(user.Password),
	}
	result := db.Create(&newUser)

	return newUser, result.Error
}

func (user *User) Validate() error {
	err := validation.ValidateStruct(user,
		validation.Field(&user.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 255).Error("Name is too long"),
		),
		validation.Field(&user.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Email is invalid format"),
		),
		validation.Field(&user.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 255).Error("Password is less than 7 chars or more than 256 chars"),
		),
	)
	return err
}
