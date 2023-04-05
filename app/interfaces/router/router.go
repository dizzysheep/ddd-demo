package router

import (
	"ddd-demo/app/application/inject"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/v1")
	{
		userController := inject.InitializeUserController(db)
		userController.RegisterRoutes(userGroup)

		orderController := inject.InitializeOrderController(db)
		orderController.RegisterRoutes(userGroup)
	}

	return r
}
