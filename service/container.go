package service

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/business"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/category"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Container struct {
	Web      *gin.Engine
	DB       *gorm.DB
	Business business.BusinessInterface
	Category category.CategoryInterface
}

func New(mainDB *gorm.DB) *Container {
	ginEngine := gin.Default()

	business := business.CreateBusiness(mainDB)
	category := category.CreateCategory(mainDB)

	return &Container{
		Web:      ginEngine,
		DB:       mainDB,
		Business: business,
		Category: category,
	}
}
