package inventory

import "wms/model"

//Manager implements inventory adjustments
type Manager interface {
	Increment(inv model.Inventory) (i model.Inventory)
}

//InvManager Inventory bl
type InvManager struct {
	ibl InvBL
}

//Increment increment Inventory
func (im InvManager) Increment(inv model.Inventory) {
	im.ibl.Increment(inv)
}
