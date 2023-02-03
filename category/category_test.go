package category

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestCategory_CreateCategory(t *testing.T) {
	database.RunTest(func(db *gorm.DB) {
		c := CreateCategory(db)
		category := &core.Category{
			Alias: "Test Category Alias",
			Title: "Test Category Title",
		}
		err := c.CreateCategory(category)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, category.ID)

		var count int64
		db.Model(&core.Category{}).Count(&count)
		assert.Equal(t, int64(1), count)

		var category2 core.Category
		db.First(&category2, category.ID)
		assert.Equal(t, category.ID, category2.ID)
		assert.Equal(t, category.Alias, category2.Alias)
	})
}
