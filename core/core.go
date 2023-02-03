package core

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Category
type Category struct {
	gorm.Model
	Alias string
	Title string
}

// Categories
type Categories []Category

// Coordinates
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Location
type Location struct {
	Address1       string
	Address2       string
	Address3       string
	City           string
	ZipCode        string
	Country        string
	State          string
	DisplayAddress pq.StringArray `gorm:"type:text[]"`
}

// Business
type Business struct {
	gorm.Model
	Alias        string
	Name         string `gorm:"unique;not null;type:varchar(100);default:null"`
	ImageURL     string
	IsClosed     bool
	URL          string
	ReviewCount  int
	Categories   []Category `gorm:"many2many:business_categories;"`
	Rating       float32
	Coordinates  Coordinates    `gorm:"embedded"`
	Transactions pq.StringArray `gorm:"type:text[]"`
	Price        string
	Location     Location `gorm:"embedded"`
	Phone        string
	DisplayPhone string
	Distance     float32
}
