package main

import (
	"fmt"
	"time"
	"wms/bl/inventory"
	"wms/data"
	"wms/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func maina() {

	var dbCtx data.DBContext
	data.InitDB(&dbCtx)
	defer dbCtx.Db.Close()

	var pdl data.ProductDL

	var p = pdl.GetByID(1)
	println(p.Description)

	p = pdl.GetByPartNumber("1009830-00-B")
	println(p.Description)

}

func main() {

	var dbCtx data.DBContext
	data.InitDB(&dbCtx)

	var pdl = data.NewProductDL(&dbCtx)

	var p = pdl.GetByID(1)
	println(p.Description)

	var inv model.Inventory
	//inv.Location.Location = "Loc1"
	inv.Part.Partnumber = "Product1"
	inv.Quantity = 1.0

	var im inventory.InvManager
	im.Increment(inv)
	//var loc = inv.GetLocation()
	//println(loc)
	var err error
	db, err = gorm.Open("mysql", "root:phoenix@/phoenix?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)
	db.BlockGlobalUpdate(true)

	db.SingularTable(true)

	//db.Callback().Create().Register("my:create", myBeforeDeleteCallback)
	//db.Callback().Create().Register("gorm:create", myBeforeDeleteCallback)
	//gorm.DefaultCallback.Create().Register("a:before_create", aBeforeDeleteCallback)

	//gorm.DefaultCallback.Create().Before("a:create").Register("a:create", aBeforeDeleteCallback)
	//	db.Callback().Create().Replace("gorm:update_time_stamp", myBeforeCreateCallback)
	//	db.Callback().Update().Replace("gorm:update_time_stamp", myBeforeUpdateCallback)

	var part model.Part
	part.Partnumber = "1009830-00-A"
	part.Description = "Model S"
	part.CreatedBy = "asiva"
	part.ModifiedBy = "asiva"

	db.Create(&part)
	part.Partnumber = "1009830-00-B"

	println(&part)

	db.Save(&part)

	var loc model.Location

	loc.Location = "BBB"
	db.Create(&loc)

	var inv2 model.Inventory

	db.Debug().Where("ID =?", "5").Find(&inv2)
	fmt.Println(inv2)

	var l = GetLocation()
	println(l.Location)
	// var loc model.Location
	// db.Delete(&loc)

}

//GetLocation s
func GetLocation() model.Location {
	var loc model.Location
	db.First(&loc, 2)
	println(loc.Location)
	return loc
}

func myBeforeCreateCallback(scope *gorm.Scope) {
	fmt.Println(scope.TableName())

	if scope.HasColumn("ModifiedOn") {
		scope.SetColumn("ModifiedOn", time.Now())
	}
	if scope.HasColumn("CreatedOn") {
		scope.SetColumn("CreatedOn", time.Now())
	}
	fmt.Println(scope.Value)

}

func myBeforeUpdateCallback(scope *gorm.Scope) {
	fmt.Println(scope.TableName())

	if scope.HasColumn("ModifiedOn") {
		scope.SetColumn("ModifiedOn", time.Now())
	}
	fmt.Println(scope.Value)

}
func myBeforeDeleteCallback(scope *gorm.Scope) {
	fmt.Println(scope.TableName())

	if scope.HasColumn("ModifiedOn") {
		scope.SetColumn("ModifiedOn", time.Now())
	}
	fmt.Println(scope.Value)

	if !scope.HasError() {
		scope.CallMethod("BeforeCreate")
	}
	if !scope.HasError() {
		scope.CallMethod("Create")
	}

}
func aBeforeDeleteCallback(scope *gorm.Scope) {
	fmt.Println(scope.TableName())

	if scope.HasColumn("ModifiedOn") {
		scope.SetColumn("ModifiedOn", time.Now())
	}
	fmt.Println(scope.Value)

}
