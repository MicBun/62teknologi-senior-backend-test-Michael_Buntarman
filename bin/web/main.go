package main

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/database"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/docs"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/service"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/web"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default env")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to db %v", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Unable to migrate db %v", err)
	}

	description := "This is a simple CRUD API Server with GIN and GORM. \n\n" +
		"Checkout my Github: https://github.com/MicBun\n\n" +
		"Checkout my Linkedin: https://www.linkedin.com/in/MicBun\n\n"

	docs.SwaggerInfo.Title = "Go CRUD API Server with GIN and GORM"
	docs.SwaggerInfo.Description = description

	c := service.New(db)
	service.SeedData(c)
	web.RegisterAPIRoutes(c)
	c.Web.Run()
}
