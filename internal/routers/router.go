package routers

import (
	"github.com/gin-gonic/gin"

	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/service"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	pingSvc := service.NewPingService()
	pingCtrl := controller.NewPingController(pingSvc)

	userSvc := service.NewUserService()
	userCtrl := controller.NewUserController(userSvc)

	router.GET("/ping", pingCtrl.Ping)
	router.GET("/users/:id", middlewares.AuthMiddleware(), userCtrl.GetUserByID)

	return router
}
