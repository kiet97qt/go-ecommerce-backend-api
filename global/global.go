package global

import (
	"go-ecommerce-backend-api/pkg/settings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *settings.Config
	Logger *zap.Logger
	DB     *gorm.DB
)
