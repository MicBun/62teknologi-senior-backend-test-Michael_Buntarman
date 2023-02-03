package web

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/service"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/web/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterAPIRoutes(c *service.Container) {
	api := handlers.NewApiHandler(c)

	c.Web.GET("/hello", api.Hello)

	c.Web.POST("/business", api.CreateBusiness)
	c.Web.PUT("/business/:id", api.EditBusiness)
	c.Web.DELETE("/business/:id", api.DeleteBusiness)
	c.Web.GET("/business/search", api.GetBusinessesByParams)

	c.Web.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
