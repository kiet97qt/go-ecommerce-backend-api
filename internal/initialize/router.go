package initialize

import (
	"go-ecommerce-backend-api/internal/routers"
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	pingCtrl, err := wire.InitPingController()
	if err != nil {
		panic(err)
	}

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

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
