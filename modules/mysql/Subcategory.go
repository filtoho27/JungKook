package mysql

import (
	Error "jungkook/error"
)

func (pd *productMethod) GetSubcategory() (subcategory []SubcategoryType, err error) {
	db := mysqlConn.ProductDB.Slave
	result := db.
		Table("Subcategory").
		Order("CategoryID, Sort").
		Find(&subcategory)

	err = result.Error

	if err != nil {
		err = Error.CustomError{ErrMsg: "GET_SUBCATEGORY_FAIL", ErrCode: 1020002}
	}
	return
}
