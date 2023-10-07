package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	// Define the database connection string
	// Format: [username]:[password]@tcp([hostname]:[port])/[database_name]?charset=utf8&parseTime=True&loc=Local
	dbConnectionString := "root:password@tcp(localhost:3306)/lora?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"

	// Create a new migrate instance
	m, err := migrate.New(
		"file:///home/baoquoc/project/lora_wan/db/migrations", // Updated absolute path
		"mysql://"+dbConnectionString)
	if err != nil {
		log.Fatalf("migration failed to initialize: %v", err)
	}

	// Apply all up migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while migrating: %v", err)
	}
}
