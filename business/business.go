package business

import (
	"errors"
	"github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman/core"
	"gorm.io/gorm"
)

type Business struct {
	db *gorm.DB
}

type BusinessInterface interface {
	CreateBusiness(business *core.Business) error
	UpdateBusiness(business *core.Business) error
	DeleteBusiness(business *core.Business) error
	GetBusinessByParam(location, category, sortBy, price string, openNow bool, limit, offset int, longitude, latitude float64) ([]core.Business, error)
}

func CreateBusiness(db *gorm.DB) BusinessInterface {
	return &Business{
		db: db,
	}
}

func (b *Business) CreateBusiness(business *core.Business) error {
	return b.db.Save(business).Error
}

func (b *Business) UpdateBusiness(business *core.Business) error {
	return b.db.Updates(business).Error
}

func (b *Business) DeleteBusiness(business *core.Business) error {
	return b.db.Delete(business).Error
}

func (b *Business) GetBusinessByParam(location, category, sortBy, price string, openNow bool, limit, offset int, longitude, latitude float64) ([]core.Business, error) {
	var business []core.Business

	dbModel := b.db.Model(&core.Business{})
	if location != "" {
		dbModel = dbModel.Where("businesses.state = ?", location)
	}
	if latitude != 0 && longitude != 0 {
		dbModel = dbModel.Where("businesses.latitude = ?", latitude).Where("businesses.longitude = ?", longitude)
	}
	if category != "" {
		dbModel = dbModel.Joins("JOIN business_categories ON business_categories.business_id = businesses.id").Joins("JOIN categories ON categories.id = business_categories.category_id").Where("categories.alias = ?", category)
	}
	if price != "" {
		dbModel = dbModel.Where("businesses.price = ?", price)
	}
	if openNow == true {
		dbModel = dbModel.Where("businesses.is_closed = ?", !openNow)
	} else {
		dbModel = dbModel.Where("businesses.is_closed = ?", openNow)
	}
	if limit != 0 {
		dbModel = dbModel.Limit(limit)
	}
	if offset != 0 {
		dbModel = dbModel.Offset(offset)
	}
	if sortBy != "" {
		dbModel = dbModel.Order(sortBy)
	}
	err := dbModel.Find(&business).Error
	if err != nil || len(business) == 0 {
		return nil, errors.New("no business found")
	}
	return business, nil
}
