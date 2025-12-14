package initialize

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/pkg/loggers"
)

func InitLogger() {
	if global.Config != nil {
		loggers.Setup(global.Config.Logging)
	}
	global.Logger = loggers.GetLogger()
}
