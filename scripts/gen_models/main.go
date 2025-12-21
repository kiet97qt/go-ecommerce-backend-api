package main

import (
	"log"

	"go-ecommerce-backend-api/internal/initialize"
)

// Script nhỏ để chạy GenerateModels độc lập với quá trình khởi động app.
//
// Cách dùng:
//
//	cd /Users/kietle/Projects/go-ecommerce-backend-api
//	go run ./scripts/gen_models
func main() {
	// 1. Load config
	initialize.LoadConfig()

	// 2. Init logger (set global.Logger)
	initialize.InitLogger()

	// 3. Init MySQL (dùng global.Logger để log)
	initialize.InitMySQL()

	// 4. Generate models
	log.Println("Generating GORM models from database...")
	initialize.GenerateModels()
	log.Println("Done generating models.")
}
