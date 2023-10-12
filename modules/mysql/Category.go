package mysql

import (
	Error "jungkook/error"
)

func (pd *productMethod) GetCategory() (category []CategoryType, err error) {
	db := mysqlConn.ProductDB.Slave
	result := db.
		Table("Category").
		Order("Sort").
		Find(&category)

	err = result.Error

	if err != nil {
		err = Error.CustomError{ErrMsg: "GET_CATEGORY_FAIL", ErrCode: 1020001}
	}
	return
}
