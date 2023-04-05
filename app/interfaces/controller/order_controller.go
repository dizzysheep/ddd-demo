package controller

import (
	"ddd-demo/app/application/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderServiceImpl
}

func NewOrderController(orderService *service.OrderServiceImpl) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/order/:id", c.getOrderByID)
}

func (c *OrderController) getOrderByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDto, err := c.orderService.GetOrderById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userDto)
}
