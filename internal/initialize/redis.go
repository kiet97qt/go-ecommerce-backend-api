package initialize

import (
	"context"
	"fmt"
	"time"

	"go-ecommerce-backend-api/global"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	if global.Config == nil {
		panic("config is nil, load config before initializing Redis")
	}

	cfg := global.Config.Redis

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolSize:     10,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("failed to connect to Redis at %s: %v", addr, err))
	}

	global.RedisClient = client
}
