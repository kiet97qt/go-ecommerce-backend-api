package global

import (
	"go-ecommerce-backend-api/pkg/settings"

	"go.uber.org/zap"
)

var (
	Config *settings.Config
	Logger *zap.Logger
)
