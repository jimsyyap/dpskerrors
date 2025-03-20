package internal

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbPool *gorm.DB

func InitDBPool() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	dbPool, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func GetDB() *gorm.DB {
	return dbPool
}

func CloseDBPool() {
	if dbPool != nil {
		sqlDB, _ := dbPool.DB()
		sqlDB.Close()
	}
}
