package service

import (
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
)

func SeedData(c *Container) {
	c.Category.CreateCategory(&core.Category{
		Alias: "newamerican",
		Title: "New American",
	})

	c.Category.CreateCategory(&core.Category{
		Alias: "burgers",
		Title: "Burgers",
	})

	business := generateBusiness()
	c.Business.CreateBusiness(&business)
}

func generateBusiness() core.Business {
	return core.Business{
		Alias:       "fork-boise",
		Name:        "Fork",
		ImageURL:    "https://s3-media4.fl.yelpcdn.com/bphoto/P9mNoEUBfeSbgmJEla4jmQ/o.jpg",
		IsClosed:    false,
		URL:         "https://www.yelp.com/biz/fork-boise?adjust_creative=DSj6I8qbyHf-Zm2fGExuug&utm_campaign=yelp_api_v3&utm_medium=api_v3_business_search&utm_source=DSj6I8qbyHf-Zm2fGExuug",
		ReviewCount: 2069,
		//CategoriesID: []uint{
		//	1,
		//},
		Categories: core.Categories{
			{
				Alias: "newamerican",
				Title: "American (New)",
			},
			{
				Alias: "breakfast_brunch",
				Title: "Breakfast & Brunch",
			},
			{
				Alias: "burgers",
				Title: "Burgers",
			},
		},
		Rating: 4.0,
		Coordinates: core.Coordinates{
			Latitude:  43.616389,
			Longitude: -116.203056,
		},
		Transactions: []string{"delivery"},
		Price:        "$$",
		Location: core.Location{
			Address1:       "199 N 8th St",
			Address2:       "",
			Address3:       "",
			City:           "Boise",
			ZipCode:        "83702",
			Country:        "US",
			State:          "ID",
			DisplayAddress: []string{"199 N 8th St", "Boise, ID 83702"},
		},
		Phone:        "+12082871700",
		DisplayPhone: "(208) 287-1700",
		Distance:     314.400925836215,
	}
}
