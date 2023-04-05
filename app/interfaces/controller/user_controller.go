package controller

import (
	"ddd-demo/app/application/dto"
	"ddd-demo/app/application/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserServiceImpl
}

func NewUserController(userService *service.UserServiceImpl) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/users", c.createUser)
	router.GET("/users/:id", c.getUserByID)
	router.GET("/users", c.getUserByEmail)
}

func (c *UserController) createUser(ctx *gin.Context) {
	var userDto dto.UserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.userService.CreateUser(&userDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c *UserController) getUserByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDto, err := c.userService.GetUserByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userDto)
}

func (c *UserController) getUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	userDto, err := c.userService.GetUserByEmail(email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userDto)
}
