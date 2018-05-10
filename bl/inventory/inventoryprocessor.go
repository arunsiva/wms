package inventory

import (
	"wms/data"
	"wms/model"
)

//Processor implements inventory adjustments
type Processor interface {
	Increment(inv model.Inventory) (i model.Inventory)
}

//InvBL Inventory bl
type InvBL struct {
	idl data.InventoryDL
}

//Increment increment Inventory
func (ibl InvBL) Increment(inv model.Inventory) (i model.Inventory, err error) {
	i = ibl.idl.GetByUniqueKey(inv.Part.Partnumber, inv.Part.Partnumber)
	if &i == nil {
		ibl.idl.AddInventory(inv)
	} else {
		i.Quantity += inv.Quantity
		ibl.idl.UpdateInventory(i)
	}
	return i, nil
}
