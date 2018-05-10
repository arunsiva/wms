package model

import "time"

//Location p
/* type Location struct {
	ID          int
	Name        string
	Description string
} */

//LocationList list
type LocationList struct {
	Locations []Location
}

// Location - Location
type Location struct {
	ID         uint      `gorm:"primary_key"`
	Location   string    `gorm:"column:Location"`
	CreatedBy  string    `gorm:"column:CreatedBy"`
	CreatedOn  time.Time `gorm:"column:CreatedOn"`
	ModifiedBy string    `gorm:"column:ModifiedBy"`
	ModifiedOn time.Time `gorm:"column:ModifiedOn"`
}
