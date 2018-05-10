package data

import "wms/model"

//LocationDataProvider Location interface
type LocationDataProvider interface {
	GetByID(id int) (l model.Location)
}

//LocationDL Location dl
type LocationDL struct {
	dbc DBContext
}

//GetByID get location
func (ldl LocationDL) GetByID(id uint) (l model.Location) {
	l.ID = id
	l.Location = "adasd"
	return l
}

//GetByName get location
func (ldl LocationDL) GetByName(name string) (l model.Location) {
	l.ID = 1
	l.Location = name
	return l
}

//GetByType get location
func (ldl LocationDL) GetByType(t string) (l model.LocationList) {
	return l
}
