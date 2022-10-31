package mysql

import (
	Error "jungkook/error"
)

type UserInfo struct {
	Name       string
	Password   string
	FaSecret   string
	Enable     int
	LoginError int
}

func (cm *coreMethod) GetUserInfo(acc string) (userInfo UserInfo, err error) {
	db := mysqlConn.CoreDB.Slave
	rd3admin := RD3Admin{}
	result := db.
		Table("RD3Admin").
		Select("`Name`, `Password`, `Google2faSecret`, `Enable`, `LoginError`").
		Where("`Name` = ?", acc).
		First(&rd3admin)
	if result.Error != nil {
		err = Error.CustomError{ErrMsg: "GET_USERINFO_FAIL", ErrCode: 1001}
		return
	}
	userInfo = UserInfo{
		Name:       rd3admin.Name,
		Password:   rd3admin.Password,
		FaSecret:   rd3admin.Google2faSecret,
		Enable:     rd3admin.Enable,
		LoginError: rd3admin.LoginError,
	}
	return
}
