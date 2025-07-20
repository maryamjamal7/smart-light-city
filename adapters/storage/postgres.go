// package storage

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // ConnectPostgres sets up the GORM PostgreSQL connection
// func ConnectPostgres() (*gorm.DB, error) {
// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		getEnv("DB_HOST", "localhost"),
// 		getEnv("DB_USER", "mariam"),
// 		getEnv("DB_PASSWORD", "123"),
// 		getEnv("DB_NAME", "city_light_db"),
// 		getEnv("DB_PORT", "5432"),
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ping the DB to ensure connection is valid
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetMaxOpenConns(100)
// 	sqlDB.SetConnMaxLifetime(time.Hour)

// 	if err := sqlDB.Ping(); err != nil {
// 		return nil, err
// 	}

// 	log.Println("✅ Connected to PostgreSQL successfully")
// 	return db, nil
// }

// // getEnv returns fallback if the env var is not set
//
//	func getEnv(key, fallback string) string {
//		if value := os.Getenv(key); value != "" {
//			return value
//		}
//		return fallback
//	}
package storage

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=mariam password=123 dbname=city_lights port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Failed to connect to DB: %v", err)
		return nil, err
	}
	return db, nil
}
