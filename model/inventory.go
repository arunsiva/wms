package model

import "time"

//Inventory I
/* type Inventory struct {
	ProductNo string
	Location  string
	Quantity  float32
	UOM       string

	ProductEntity  Product
	LocationEntity Location
} */

//InventoryList list
type InventoryList struct {
	Inventories []Inventory
}

// Inventory - inventory
type Inventory struct {
	ID         uint `gorm:"primary_key"`
	Partid     int
	Locationid int
	Quantity   int
	CreatedBy  string
	CreatedOn  time.Time
	ModifiedBy string
	ModifiedOn time.Time
	Location   func() *Location
	location   Location //`gorm:"ForeignKey:LocationID"`
	Part       Part     //`gorm:"ForeignKey:PartID"`
}

func (i *Inventory) GetLocation() *Location {

	return &i.location
}
