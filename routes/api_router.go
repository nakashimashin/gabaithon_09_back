package routes

import (
	"gabaithon-09-back/controllers"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetApiRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:  []string{"http://localhost:50241"},
	// 	AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:  []string{"Origin", "Content-Type", "Content-Length"},
	// 	ExposeHeaders: []string{"Content-Length"},
	// }))

	handler := &controllers.Handler{
		DB: db,
	}

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", handler.SignUpHandler)
			auth.POST("/login", handler.LoginHandler)
		}
	}

	return router
}
