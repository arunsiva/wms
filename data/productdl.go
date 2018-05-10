package data

import "wms/model"

//ProductDataProvider Product interface
type ProductDataProvider interface {
	GetByID(id int) (p model.Product)
	GetByPartNumber(partNumber string) (p model.Product)
}

//ProductDL product dl
type ProductDL struct {
	dBCtx *DBContext
}

//NewProductDL to initialize ProductDL
func NewProductDL(dBContext *DBContext) (pdl *ProductDL) {
	var p ProductDL
	p.dBCtx = dBContext
	return &p
}

//GetByID get part
func (pdl ProductDL) GetByID(id int) (p model.Part) {
	pdl.dBCtx.Db.First(&p, id)
	return p
}

//GetByPartNumber get part
func (pdl ProductDL) GetByPartNumber(partNumber string) (p model.Part) {
	DbCtx.Db.Where("partnumber = ?", partNumber).First(&p)
	return p
}

//GetByType get part
func (pdl ProductDL) GetByType(t string) (p model.ProductList) {
	return p
}
