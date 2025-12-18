package global

import (
	"go-ecommerce-backend-api/pkg/settings"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config      *settings.Config
	Logger      *zap.Logger
	DB          *gorm.DB
	RedisClient *redis.Client
)
