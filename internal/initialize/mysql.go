package initialize

import (
	"fmt"
	"time"

	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySQL() {
	if global.Config == nil {
		panic("config is nil, load config before initializing MySQL")
	}

	mysqlCfg := global.Config.MySQL

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to MySQL: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get sql.DB from gorm: %v", err))
	}

	sqlDB.SetMaxIdleConns(mysqlCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlCfg.ConnMaxLifetime) * time.Second)
	global.Logger.Info("MySQL initialized", zap.Any("db", db))
	global.DB = db
}
