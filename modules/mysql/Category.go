package mysql

import (
	Error "jungkook/error"
)

func (pd *productMethod) GetCategory() (category []Category, err error) {
	db := mysqlConn.ProductDB.Slave
	result := db.
		Table("Category").
		Find(&category)

	err = result.Error

	if err != nil {
		err = Error.CustomError{ErrMsg: "GET_CATEGORY_FAIL", ErrCode: 1020001}
	}
	return
}

func (pd *productMethod) GetSubcategory() (subcategory []Subcategory, err error) {
	db := mysqlConn.ProductDB.Slave
	result := db.
		Table("Subcategory").
		Find(&subcategory)

	err = result.Error

	if err != nil {
		err = Error.CustomError{ErrMsg: "GET_SUBCATEGORY_FAIL", ErrCode: 1020002}
	}
	return
}
