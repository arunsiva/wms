package data

import (
	"wms/model"
)

//InventoryDataProvider Inventory interface
type InventoryDataProvider interface {
	GetByID(id int) (l model.Inventory)
}

//InventoryDL Inventory dl
type InventoryDL struct {
	dbc DBContext
}

//GetByID get Inventory
func (idl InventoryDL) GetByID(id int) (i model.Inventory) {

	return i
}

//GetByProduct get Inventory
func (idl InventoryDL) GetByProduct(productNo string) (il model.InventoryList) {
	return il
}

//GetByUniqueKey get Inventory
func (idl InventoryDL) GetByUniqueKey(productNo string, location string) (i model.Inventory) {
	return i
}

//AddInventory add Inventory
func (idl InventoryDL) AddInventory(inv model.Inventory) (i model.Inventory) {
	return i
}

//UpdateInventory change Inventory
func (idl InventoryDL) UpdateInventory(inv model.Inventory) (i model.Inventory) {
	return i
}

//DeleteInventory delete Inventory
func (idl InventoryDL) DeleteInventory(inv model.Inventory) (i model.Inventory) {
	return i
}
