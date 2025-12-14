package initialize

import (
	"strconv"

	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	var config = global.Config

	InitLogger()
	global.Logger.Info("config loaded", zap.Any("config", config))
	InitMySQL()
	InitRouter()
	router := InitRouter()

	router.Run(":" + strconv.Itoa(config.Server.Port))
}
