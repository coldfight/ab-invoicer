package db_service

import (
	"database/sql"
	"fmt"
	"github.com/coldfight/ab-invoicer/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DriverName     = "sqlite3"
	DbFilePath     = "./storage/invoices.db"
	DbFilePathTest = "./storage/test.db"
)

func GetConnection() (*sql.DB, error) {
	return sql.Open(DriverName, DbFilePath)
}

func CreateInitialDatabase() {
	db, err := gorm.Open(sqlite.Open(DbFilePathTest), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}

	// Migrate the schemas
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Owner{})
	db.AutoMigrate(&models.Expense{})
	db.AutoMigrate(&models.Labour{})
	db.AutoMigrate(&models.Invoice{})
}
