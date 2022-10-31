package account

import (
	"jungkook/api"
	"jungkook/kernel"
	"net/http"
)

type AccountSt struct{}

func (ex *AccountSt) SendEmailCodeHandler(w http.ResponseWriter, r *http.Request) {
	module := api.GetModule()
	_ = r.ParseForm()
	userName := kernel.ParamString(r.Form, "userName")
	passWord := kernel.ParamString(r.Form, "passWord")
	passWordRepeat := kernel.ParamString(r.Form, "passWordRepeat")
	err := sendEmailCode(module, userName, passWord, passWordRepeat)
	kernel.FormatResult(w, nil, err)
}

func (ex *AccountSt) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	module := api.GetModule()
	_ = r.ParseForm()
	userName := kernel.ParamString(r.Form, "userName")
	passWord := kernel.ParamString(r.Form, "passWord")
	emailCode := kernel.ParamInt(r.Form, "emailCode")
	err := register(module, userName, passWord, emailCode)
	kernel.FormatResult(w, nil, err)
}

func (ex *AccountSt) LoginHandler(w http.ResponseWriter, r *http.Request) {
	module := api.GetModule()
	_ = r.ParseForm()
	userName := kernel.ParamString(r.Form, "userName")
	passWord := kernel.ParamString(r.Form, "passWord")
	loginType := kernel.ParamInt(r.Form, "loginType")
	userInfo, err := login(module, userName, passWord, loginType)
	kernel.FormatResult(w, userInfo, err)
}
