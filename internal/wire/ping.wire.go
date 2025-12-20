//go:build wireinject
// +build wireinject

package wire

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/service"

	"github.com/google/wire"
)

func InitPingController() (*controller.PingController, error) {
	wire.Build(service.NewPingService, controller.NewPingController)
	return &controller.PingController{}, nil
}
