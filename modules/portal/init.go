package portal

import (
	"fmt"
	Error "jungkook/error"
	customLog "jungkook/log"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var config *viper.Viper
var portal *portalLib = &portalLib{}

type portalLib struct {
	client *resty.Client
	host   string
}

func GetPortal() *portalLib {
	return portal
}

// 初始化連線
func Init() {
	getConfig()
	client, host := newRestyClient()
	portal = &portalLib{client: client, host: host}
}

// 建立連線
func newRestyClient() (client *resty.Client, host string) {
	ip := config.GetString("ip")
	hostname := config.GetString("hostname")
	host = fmt.Sprintf("http://%s", ip)
	client = resty.New()
	client.SetTimeout(3 * time.Second)
	client.SetHeader("Host", hostname)
	client.SetHeader("Connection", "keep-alive")
	client.SetHeader("Keep-Alive", "300")
	client.SetHeader("Accept-Language", "zh-tw")
	client.SetHeader("Content-Type", "application/json")
	return
}

// 取得設定檔
func getConfig() {
	config = viper.New()
	config.SetConfigName("portal")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config/modules/")
	err := config.ReadInConfig()
	if err != nil {
		msg := fmt.Sprintf("GET_PORTAL_CONFIG_FAILED, %+v", err)
		customLog.WriteApiLog("getConfig", Error.CustomError{ErrMsg: msg}, err, "", nil)
	}
}

// 處理連線錯誤
func portalResponse(res *resty.Response, funcName string, err error, response string) error {
	httpCode := res.StatusCode()
	if err != nil || response == "" || httpCode != 200 {
		customErr := Error.CustomError{ErrMsg: "Portal連線失敗", ErrCode: 1338000053}
		customLog.WriteApiLog(funcName, customErr, err, response, nil)
		return customErr
	}
	return nil
}
