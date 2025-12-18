package user

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	userSvc := service.NewUserService()
	userCtrl := controller.NewUserController(userSvc)

	// public routes
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.GET("/otp")
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/login")
		userRouterPublic.GET("/:id", middlewares.AuthMiddleware(), userCtrl.GetUserByID)
	}
	// private routes
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/profile")
		userRouterPrivate.POST("/logout")
	}
}
