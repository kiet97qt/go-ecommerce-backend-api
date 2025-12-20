package user

import (
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	userCtrl, err := wire.InitUserController()
	if err != nil {
		panic(err)
	}

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
