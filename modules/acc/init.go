package acc

import (
	"fmt"
	"time"

	Type "jungkook/commonType"
	Error "jungkook/error"
	customLog "jungkook/log"

	resty "github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var config *viper.Viper

type accLib struct {
	client *resty.Client
	host   string
}

type AccType struct {
	BBIN Type.AccInterface
	BBGP Type.AccInterface
}

var acc *AccType = &AccType{}

// 初始化連線
func Init() {
	getConfig()
	bbinClient, bbinHost := newRestyClient("BBIN")
	bbgpClient, bbgpHost := newRestyClient("BBGP")
	acc.BBIN = &accLib{client: bbinClient, host: bbinHost}
	acc.BBGP = &accLib{client: bbgpClient, host: bbgpHost}
}

func GetAcc() *AccType {
	return acc
}

func (at *AccType) GetByVendor(vendor int) (acc Type.AccInterface) {
	if vendor == 1 {
		acc = at.BBIN
	} else if vendor == 2 {
		acc = at.BBGP
	}
	return
}

// 建立連線
func newRestyClient(pType string) (client *resty.Client, accHost string) {
	accKey := fmt.Sprintf("acc.%s", pType)
	accConfig := config.GetStringMapString(accKey)
	accHost = fmt.Sprintf("%s://%s:%s", accConfig["protocol"], accConfig["ip"], accConfig["port"])
	client = resty.New()
	client.SetTimeout(5 * time.Second)
	client.SetHeader("Sensitive-Data", "entrance=3&operator_id=0&operator=nobody&client_ip=&vendor=rdc")
	client.SetHeader("Host", accConfig["hostname"])
	client.SetHeader("Connection", "keep-alive")
	client.SetHeader("Keep-Alive", "300")
	client.SetHeader("Accept-Language", "zh-tw")
	client.SetHeader("Content-Type", "application/json")
	return
}

// 取得設定檔
func getConfig() {
	config = viper.New()
	config.SetConfigName("acc")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config/modules/")
	err := config.ReadInConfig()
	if err != nil {
		msg := fmt.Sprintf("GET_ACC_CONFIG_FAILED, %+v", err)
		customLog.WriteAccLog("getConfig", Error.CustomError{ErrMsg: msg}, err, "", nil)
	}
}

// 處理連線錯誤
func accResponse(res *resty.Response, funcName string, err error, response string) error {
	httpCode := res.StatusCode()
	if err != nil || response == "" || httpCode != 200 {
		customErr := Error.CustomError{ErrMsg: "Acc連線失敗", ErrCode: 1338000001}
		logMsg := fmt.Sprintf("%s|httpCode=%d", customErr.ErrMsg, httpCode)
		logErr := Error.CustomError{ErrMsg: logMsg, ErrCode: customErr.ErrCode}
		customLog.WriteAccLog(funcName, logErr, err, response, nil)
		return customErr
	}
	return nil
}
