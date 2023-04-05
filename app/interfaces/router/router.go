package router

import (
	"ddd-demo/app/application/service"
	"ddd-demo/app/interfaces/controller"
	"ddd-demo/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/v1")
	{
		userRepo := repository.NewUserRepository(db)
		userService := service.NewUserService(userRepo)
		userController := controller.NewUserController(userService)
		userController.RegisterRoutes(userGroup)

		orderRepo := repository.NewOrderRepository(db)
		orderService := service.NewOrderService(orderRepo)
		orderController := controller.NewOrderController(orderService)
		orderController.RegisterRoutes(userGroup)

	}

	return r
}
