package db_service

import (
	"fmt"
	"github.com/coldfight/ab-invoicer/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DriverName = "sqlite3"
	DbFilePath = "./storage/test.db"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DbFilePath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}
	return db
}

func CreateInitialDatabase() {
	db := GetConnection()

	// Migrate the schemas
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Owner{})
	db.AutoMigrate(&models.Expense{})
	db.AutoMigrate(&models.Labour{})
	db.AutoMigrate(&models.Invoice{})
}
