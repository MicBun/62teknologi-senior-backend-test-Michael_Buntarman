package business

import (
	"fmt"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestBusiness_CreateBusiness(t *testing.T) {
	database.RunTest(func(db *gorm.DB) {
		b := CreateBusiness(db)
		business := &core.Business{
			Name: "Test Business",
		}
		err := b.CreateBusiness(business)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, business.ID)

		var count int64
		db.Model(&core.Business{}).Count(&count)
		assert.Equal(t, int64(1), count)

		var business2 core.Business
		db.First(&business2, business.ID)
		assert.Equal(t, business.ID, business2.ID)
		assert.Equal(t, business.Name, business2.Name)
	})
}

func TestBusiness_UpdateBusiness(t *testing.T) {
	database.RunTest(func(db *gorm.DB) {
		b := CreateBusiness(db)
		business := &core.Business{
			Name: "Test Business",
		}
		err := b.CreateBusiness(business)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, business.ID)

		business.Name = "Updated Business"
		err = b.UpdateBusiness(business)
		assert.NoError(t, err)

		var business2 core.Business
		db.First(&business2, business.ID)
		assert.Equal(t, business.ID, business2.ID)
		assert.Equal(t, business.Name, business2.Name)
	})
}

func TestBusiness_DeleteBusiness(t *testing.T) {
	database.RunTest(func(db *gorm.DB) {
		b := CreateBusiness(db)
		business := &core.Business{
			Name: "Test Business",
		}
		err := b.CreateBusiness(business)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, business.ID)

		err = b.DeleteBusiness(business)
		assert.NoError(t, err)

		var count int64
		db.Model(&core.Business{}).Count(&count)
		assert.Equal(t, int64(0), count)
	})
}

func TestBusiness_GetBusinessByParam(t *testing.T) {
	database.RunTest(func(db *gorm.DB) {
		b := CreateBusiness(db)
		business := &core.Business{
			Name:     "Test Business",
			IsClosed: false,
			Location: core.Location{
				Address1: "Test Address 1",
				Address2: "Test Address 2",
				Address3: "Test Address 3",
				City:     "Test City",
				State:    "ID",
				ZipCode:  "Test Zip Code",
				Country:  "Test Country",
				DisplayAddress: []string{
					"Test Address 1",
					"Test Address 2",
				},
			},
			Coordinates: core.Coordinates{
				Latitude:  22.5,
				Longitude: 114.1,
			},
			Categories: []core.Category{
				{
					Alias: "Test Category Alias",
					Title: "Test Category Title",
				},
			},
			Price: "$$",
		}
		err := b.CreateBusiness(business)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, business.ID)

		business2, err := b.GetBusinessByParam("", "", "", "", true, 0, 0, 1, 0)
		fmt.Println(business2)
		assert.NoError(t, err)
		assert.Equal(t, business.ID, business2[0].ID)
		assert.Equal(t, business.Name, business2[0].Name)
		assert.Equal(t, business.IsClosed, business2[0].IsClosed)

		// filter by longitude and latitude
		business2, err = b.GetBusinessByParam("", "", "", "", true, 0, 0, 114.1, 22.5)
		fmt.Println(business2)
		assert.NoError(t, err)
		assert.Equal(t, business.ID, business2[0].ID)
		assert.Equal(t, business.Name, business2[0].Name)
		assert.Equal(t, business.IsClosed, business2[0].IsClosed)

		// filter by location
		business2, err = b.GetBusinessByParam("ID", "", "", "", true, 0, 0, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, business.ID, business2[0].ID)

		// filter by category
		business2, err = b.GetBusinessByParam("ID", "Test Category Alias", "", "", true, 0, 0, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, business.ID, business2[0].ID)

		// filter by price
		business2, err = b.GetBusinessByParam("ID", "Test Category Alias", "", "$$", true, 0, 0, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, business.ID, business2[0].ID)

		// add another business
		business3 := &core.Business{
			Name:     "Test Business 2",
			IsClosed: false,
			Location: core.Location{
				Address1: "Test Address 1",
			},
		}
		err = b.CreateBusiness(business3)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, business3.ID)

		// filter by limit and offset
		business2, err = b.GetBusinessByParam("ID", "Test Category Alias", "", "$$", true, 1, 0, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, len(business2), 1)

		business2, err = b.GetBusinessByParam("", "", "", "", true, 2, 0, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, len(business2), 2)

		business2, err = b.GetBusinessByParam("", "", "", "", true, 1, 1, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, len(business2), 1)

		// filter by sort by
		business2, err = b.GetBusinessByParam("", "", "name", "", true, 2, 0, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, business2[0].Name, "Test Business")
		assert.Equal(t, business2[1].Name, "Test Business 2")

		// filter not found
		business2, err = b.GetBusinessByParam("", "", "", "##", true, 2, 0, 0, 0)
		assert.Error(t, err)
		assert.Equal(t, 0, len(business2))

		// filter by open now
		business2, err = b.GetBusinessByParam("", "", "", "", false, 2, 0, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(business2))
	})
}
