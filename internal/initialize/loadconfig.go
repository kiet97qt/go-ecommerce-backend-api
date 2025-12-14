package initialize

import (
	"go-ecommerce-backend-api/global"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
}
