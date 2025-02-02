package connection

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DATABASE")
	if dsn == "" {
		log.Fatal("DATABASE environment variable not set")
		return nil
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil
	}
	return db

}

func DB() *sql.DB {
	sqlDB, err := Connect().DB()

	if err != nil {
		log.Fatal("Database connection is failed.")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return sqlDB
}
