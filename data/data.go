package data

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//DBContext DB context
type DBContext struct {
	dbConnection string
	Db           *gorm.DB
}

//DbCtx default db
var DbCtx *DBContext

//InitDB initialize DB connection
func InitDB(dbContext *DBContext) {
	var err error
	dbContext.Db, err = gorm.Open("mysql", "root:phoenix@/phoenix?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	dbContext.Db.LogMode(true)
	dbContext.Db.BlockGlobalUpdate(true)

	dbContext.Db.SingularTable(true)

	dbContext.Db.Callback().Create().Replace("gorm:update_time_stamp", myBeforeCreateCallback)
	dbContext.Db.Callback().Update().Replace("gorm:update_time_stamp", myBeforeUpdateCallback)
	DbCtx = dbContext

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
