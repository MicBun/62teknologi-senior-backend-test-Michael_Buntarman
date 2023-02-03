package category

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
	"gorm.io/gorm"
)

type Category struct {
	db *gorm.DB
}

type CategoryInterface interface {
	CreateCategory(category *core.Category) error
}

func CreateCategory(db *gorm.DB) CategoryInterface {
	return &Category{
		db: db,
	}
}

func (c *Category) CreateCategory(category *core.Category) error {
	return c.db.Save(category).Error
}
