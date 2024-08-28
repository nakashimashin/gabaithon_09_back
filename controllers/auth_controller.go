package controllers

import (
	"gabaithon-09-back/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (handler *Handler) SignUpHandler(context *gin.Context) {
	var signUpInput models.SignUpInput
	err := context.ShouldBind(&signUpInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	newUser := &models.User{
		Name:     signUpInput.Name,
		Email:    signUpInput.Email,
		Password: signUpInput.Password,
	}

	err = newUser.Validate()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := newUser.Create(handler.DB)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"message": "Successfully created user",
	})
}

func (handler *Handler) LoginHandler(context *gin.Context) {
	var loginInput models.LoginInput
	if err := context.ShouldBind(&loginInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	user, err := models.FindUserByEmail(handler.DB, loginInput.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "User not found",
		})
		return
	}

	if !user.VerifyPassword(loginInput.Password) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
	})
}
