//go:build wireinject
// +build wireinject

package wire

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/service"

	"github.com/google/wire"
)

func InitUserController() (*controller.UserController, error) {
	wire.Build(service.NewUserService, controller.NewUserController)
	return &controller.UserController{}, nil
}
