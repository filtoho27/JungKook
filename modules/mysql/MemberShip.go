package mysql

import (
	Error "jungkook/error"
)

func (cm *accountMethod) GetMemberShipByUserName(userName string) (memberShip MemberShip, err error) {
	db := mysqlConn.AccountDB.Slave
	result := db.
		Table("MemberShip").
		Where("UserName", userName).
		Find(&memberShip)

	err = result.Error

	if err != nil {
		err = Error.CustomError{ErrMsg: "GET_MEMBERSHOP_FAIL", ErrCode: 1010006}
	}
	return
}

func (cm *accountMethod) CreateMemberShip(memberShip MemberShip) (err error) {
	db := mysqlConn.AccountDB.Master
	result := db.
		Table("MemberShip").
		Create(&memberShip)

	if result.Error != nil {
		err = Error.CustomError{ErrMsg: "INSERT_MEMBERSHOP_FAIL", ErrCode: 1010012}
	}
	return
}
