package mysql

import (
	Error "jungkook/error"
)

func (pd *productMethod) GetProduct() (product []ProductType, err error) {
	db := mysqlConn.ProductDB.Slave
	result := db.
		Table("Product").
		Find(&product)

	err = result.Error

	if err != nil {
		err = Error.CustomError{ErrMsg: "GET_PRODUCT_FAIL", ErrCode: 1020003}
	}
	return
}
