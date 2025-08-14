package database

import (
	"fmt"
	"go-fiber/app/model"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	// Get database configuration from environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "1433"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "sa"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "123Password"
	}

	// MSSQL connection string
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		user, password, host, port)

	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully!")
}

func Migrate() {
	// Auto migrate your models

	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed!")

}

// Drop all tables
func DropTables() {
	log.Println("Dropping all tables...")

	err := DB.Migrator().DropTable(
		&model.User{},
		// Add more models here in reverse order
	)

	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	log.Println("All tables dropped successfully!")
}

// Add to database package
func DropTable(model interface{}) {
	err := DB.Migrator().DropTable(model)
	if err != nil {
		log.Printf("Failed to drop table for model %T: %v", model, err)
	} else {
		log.Printf("Table for model %T dropped successfully", model)
	}
}

func HasTable(model interface{}) bool {
	return DB.Migrator().HasTable(model)
}

// Fresh migration (drop and recreate)
func FreshMigrate() {
	log.Println("Running fresh migration...")

	// First drop all tables
	DropTables()

	// Then run migration
	Migrate()

	log.Println("Fresh migration completed!")
}
