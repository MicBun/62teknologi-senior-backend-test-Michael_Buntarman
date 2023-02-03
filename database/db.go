package database

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
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
