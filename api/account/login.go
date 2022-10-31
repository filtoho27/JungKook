package account

import (
	Type "jungkook/commonType"
	Error "jungkook/error"
	"jungkook/modules/mysql"
	"time"
	"encoding/base64"
	"strconv"
)

type userInfoType struct {
	UserID     int    `json:"userID"`
	UserName   string `json:"userName"`
	CreateTime string `json:"createTime"`
	LoginType  int    `json:"loginType"`
	Token      string `json:"token"`
}

func login(module *Type.ModuleType, userName string, passWord string, loginType int) (userInfo userInfoType, err error) {
	// 檢查帳號格式
	err = checkUserNameType(userName)
	if err != nil {
		return
	}
	// 檢查帳號是否存在
	accountDB := mysql.GetAccountDB()
	memberShip, err := accountDB.GetMemberShipByUserName(userName)
	if memberShip.UserID == 0 {
		err = Error.CustomError{ErrMsg: "USER_NAME_NOT_EXIST", ErrCode: 1010014}
		return
	}

	// 檢查密碼是否正確
	if memberShip.PassWord != passWord{
		err = Error.CustomError{ErrMsg: "PASSWORD_IS_INCORRECT", ErrCode: 1010015}
	}

	// 產生Token
	token := createToken(memberShip.UserID)
	err = module.Redis.SetUserToken(memberShip.UserID, token)
	if err != nil {
		return
	}

	userInfo = userInfoType{memberShip.UserID, memberShip.UserName, memberShip.CreateTime, loginType, token}
	return
}

func createToken(userID int) string {
	now := time.Now().Format("2006-01-02 15:04:05")
	token := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(userID) + "JUNGKOOK" + now))
	return token
}
