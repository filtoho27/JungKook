package example

import (
	"encoding/json"
	Type "jungkook/commonType"
	Error "jungkook/error"
	customLog "jungkook/log"
)

type HallInfo struct {
	Name      string `json:"name"`
	LoginCode string `json:"login_code"`
}

func exAcc1(module *Type.ModuleType, hallID int) (result HallInfo, err error) {
	acc := module.Acc.GetByVendor(module.Vendor)
	accJson, err := acc.GetDomainData(hallID)
	if err != nil {
		return
	}
	var accData struct {
		Ret HallInfo `json:"ret"`
		Type.AccResult
	}
	_ = json.Unmarshal([]byte(accJson), &accData)
	if accData.AccResult.Result != "ok" {
		err = Error.CustomError{ErrMsg: "GET_DOMAIN_ERROR", ErrCode: 1006}
		customLog.WriteAccLog("GetDomainData", err, nil, accJson, []interface{}{hallID})
		return
	}
	result = accData.Ret
	return
}
