package model

import "time"

//Product p
type Product struct {
	ID          int
	ProductNo   string
	Description string
}

//ProductList list
type ProductList struct {
	Products []Product
}

// Part - Part
type Part struct {
	ID          uint `gorm:"primary_key"`
	Partnumber  string
	Description string
	CreatedBy   string
	CreatedOn   time.Time
	ModifiedBy  string
	ModifiedOn  time.Time
}
