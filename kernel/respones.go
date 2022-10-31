package kernel

import (
	"encoding/json"
	Error "jungkook/error"
	"net/http"
)

type responseType struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
}

type responseError struct {
	Result bool              `json:"result"`
	Data   Error.CustomError `json:"data"`
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
		Result: true,
		Data:   data,
	}
	jsondata, _ := json.Marshal(result)
	_, _ = w.Write(jsondata)
}

func handleError(w http.ResponseWriter, err error) {
	customErr, _ := err.(Error.CustomError)
	result := responseError{
		Result: false,
		Data:   customErr,
	}
	jsondata, _ := json.Marshal(result)
	_, _ = w.Write(jsondata)
}
