package account

import (
	Type "jungkook/commonType"
	Error "jungkook/error"
	"jungkook/modules/mysql"
	"regexp"
	"fmt"
	"net/smtp"
	"math/rand"
	"time"
)

func sendEmailCode(module *Type.ModuleType, userName string, passWord string, passWordRepeat string) (err error) {
	// 檢查帳號格式
	err = checkUserNameType(userName)
	if err != nil {
		return
	}
	// 檢查一分內是否申請過
	emailCode, _ := module.Redis.GetEmailCode(userName)
	if emailCode.Timestamp != 0 {
		if time.Now().Unix() - emailCode.Timestamp < 60 {
			err = Error.CustomError{ErrMsg: "GET_EMAIL_CODE_FREQUENTLY", ErrCode: 1010009}
			return
		}
	}
	// 檢查帳號是否已被使用
	accountDB := mysql.GetAccountDB()
	memberShip, err := accountDB.GetMemberShipByUserName(userName)
	if memberShip.UserID != 0 {
		err = Error.CustomError{ErrMsg: "USER_NAME_EXIST", ErrCode: 1010002}
		return
	}

	// 檢查密碼
	if passWord != passWordRepeat {
		err = Error.CustomError{ErrMsg: "PASSWORD_AND_PASSWORDREPEAT_IS_DIFFERENT", ErrCode: 1010003}
		return
	}
	passWordReg := regexp.MustCompile("^[a-zA-Z0-9]{8,16}$")
	if !passWordReg.MatchString(passWord) {
		err = Error.CustomError{ErrMsg: "PASSWORD_TYPE_ERROR", ErrCode: 1010004}
		return
	}

	// 產生驗證碼
	code := generateRandomSixDigits()

	// 寄信
	from := "wenflowtw@gmail.com"
	pass := "pwzmjpfpdysnucek"
	subject := "WenFlow Account Confirm!"
	context := `
	Thank you for creating an account!
	Please enter the code on your application:

	%d
	`
	msg := "From: " + from + "\r\n" +
		   "To: " + userName + "\r\n" +
		   "Subject: " + subject + "\r\n" +
		   "\r\n" +
		   fmt.Sprintf(context, code)

	err = smtp.SendMail("smtp.gmail.com:587",
		  smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		  from, []string{userName}, []byte(msg))

	if err != nil {
		err = Error.CustomError{ErrMsg: "SEND_MAIL_FAIL", ErrCode: 1010005}
		return
	}

	// 儲存驗證碼
	err = module.Redis.SetEmailCode(userName, code)
	return
}

func register(module *Type.ModuleType, userName string, passWord string, emailCode int) (err error) {
	userEmailCode, _ := module.Redis.GetEmailCode(userName)
	// 驗證碼是否存在
	if userEmailCode.Code == 0 {
		err = Error.CustomError{ErrMsg: "EMAIL_CODE_NOT_EXIST", ErrCode: 1010010}
		return
	}
	// 驗證碼是否正確
	if userEmailCode.Code != emailCode {
		err = Error.CustomError{ErrMsg: "EMAIL_CODE_IS_INCORRECT", ErrCode: 1010011}
		return
	}
	// 註冊
	accountDB := mysql.GetAccountDB()
	createData := mysql.MemberShip{
		UserName: userName,
		PassWord: passWord,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err = accountDB.CreateMemberShip(createData)
	if err != nil {
		return
	}
	// 刪除驗證碼
	_ = module.Redis.DelEmailCode(userName)
	return
}

func checkUserNameType(userName string) (err error) {
	userNameReg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !userNameReg.MatchString(userName) {
		err = Error.CustomError{ErrMsg: "USER_NAME_TYPE_ERROR", ErrCode: 1010001}
	}
	return
}

func generateRandomSixDigits() int {
	// 設置隨機種子，使用時間作為種子值，以確保每次運行產生不同的隨機數
	rand.Seed(time.Now().UnixNano())
	// 產生 6 位隨機數，範圍在 100000 到 999999 之間
	return rand.Intn(900000) + 100000
}
