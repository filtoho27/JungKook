package example

import (
	Error "jungkook/error"
	"jungkook/modules/mysql"
)

func exDB(acc string, pwd string) (result bool, err error) {
	coreDB := mysql.GetCoreDB()
	userInfo, err := coreDB.GetUserInfo(acc)
	if err != nil {
		return
	}
	if userInfo.Name == "" {
		err = Error.CustomError{ErrMsg: "ACCOUNT_NOT_EXIST", ErrCode: 1004}
	} else if userInfo.Enable == 0 {
		err = Error.CustomError{ErrMsg: "ACCOUNT_IS_DISABLED", ErrCode: 1005}
	} else {
		result = true
	}
	return
}
