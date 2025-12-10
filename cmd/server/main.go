package main

import (
	"fmt"
	"log"

	"go-ecommerce-backend-api/configs"
	"go-ecommerce-backend-api/internal/routers"
	"go-ecommerce-backend-api/pkg/loggers"

	"go.uber.org/zap"
)

func main() {
	router := routers.SetupRouter()
	configs.InitConfig()

	config := configs.GetConfig()

	loggers.Info("config", zap.Any("config", config))

	var port string = fmt.Sprintf(":%d", config.Server.Port)

	if err := router.Run(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
