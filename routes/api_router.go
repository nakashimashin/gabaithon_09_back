package routes

import (
	"gabaithon-09-back/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetApiRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	handler := &controllers.Handler{
		DB: db,
	}

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", handler.SignUpHandler)
		}
	}

	return router
}
