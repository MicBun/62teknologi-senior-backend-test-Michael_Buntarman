package database

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Connect() (*gorm.DB, error) {
	dbConnection := getEnv("DB_CONNECTION", "sqlite")
	postgresDSN := getEnv("DB_DSN", "")

	if dbConnection == "postgres" && postgresDSN != "" {
		return gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	}
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(core.Business{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(core.Categories{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(core.Category{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(core.Location{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(core.Coordinates{}); err != nil {
		return err
	}
	return nil
}
