package initialize

import (
	"log"
	"os"

	"go-ecommerce-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config %s: %v", env, err)
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
}
