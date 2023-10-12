package kernel

import (
	"encoding/json"
	Error "jungkook/error"
	"net/http"
)

type responseType struct {
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data"`
}

type responseError struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

func FormatResult(w http.ResponseWriter, data interface{}, err error) {
	// 處理回傳
	if err != nil {
		handleError(w, err)
		return
	}
	handleResult(w, data)
}

func handleResult(w http.ResponseWriter, data interface{}) {
	result := responseType{
		ErrCode: 0,
		ErrMsg:  "",
		Data:    data,
	}
	jsondata, _ := json.Marshal(result)
	_, _ = w.Write(jsondata)
}

func handleError(w http.ResponseWriter, err error) {
	customErr, _ := err.(Error.CustomError)
	result := responseError{
		ErrCode: customErr.ErrCode,
		ErrMsg:  customErr.ErrMsg,
	}
	jsondata, _ := json.Marshal(result)
	_, _ = w.Write(jsondata)
}
