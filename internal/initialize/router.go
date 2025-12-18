package initialize

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/routers"
	"go-ecommerce-backend-api/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	pingSvc := service.NewPingService()
	pingCtrl := controller.NewPingController(pingSvc)

	MainRouter := router.Group("/v1/api")
	{
		MainRouter.GET("/ping", pingCtrl.Ping)
	}
	{
		manageRouter.InitAdminRouter(MainRouter)
		manageRouter.InitUserRouter(MainRouter)
	}
	{
		userRouter.InitUserRouter(MainRouter)
		userRouter.InitProductRouter(MainRouter)
	}

	return router
}
