//go:build wireinject
// +build wireinject

package wire

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/internal/utils/sendto"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func newRedisClient() *redis.Client {
	return global.RedisClient
}

func newSMTPConfig() sendto.SMTPConfig {
	return sendto.LoadSMTPConfigFromEnv()
}

var userSet = wire.NewSet(
	newRedisClient,
	newSMTPConfig,
	sendto.NewMandrillEmailSender,
	wire.Bind(new(sendto.EmailSender), new(*sendto.MandrillEmailSender)),
	service.NewUserService,
	controller.NewUserController,
)

func InitUserController() (*controller.UserController, error) {
	wire.Build(userSet)
	return &controller.UserController{}, nil
}
