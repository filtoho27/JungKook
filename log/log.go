package log

import (
	"fmt"
	Error "jungkook/error"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"time"
)

func WriteLog(folderName string, logName string, logData interface{}) {
	dirPath := fmt.Sprintf("./txt/%s", folderName)
	_ = checkDir(dirPath)
	logPath := getPath(dirPath, logName)
	logContent := getLogContent(logName, logData)
	fileOpen, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer fileOpen.Close()
	logger := log.New(fileOpen, "", 0)
	logger.Println(logContent)
}

func WriteAccLog(funcName string, customErr error, err error, accReturn string, param []interface{}) {
	file := "acc-error"
	dirPath := "./txt/modules"
	_ = checkDir(dirPath)
	logPath := getPath(dirPath, file)
	nowTime, usEastTime := getTime()
	customError := customErr.(Error.CustomError)
	paramText := ""
	for idx, val := range param {
		if idx > 0 {
			paramText += ", "
		}
		paramText = fmt.Sprintf("%s%v", paramText, val)
	}
	accLog := fmt.Sprintf(
		"type=%s|datetime=%s|datetime_gmt=%s|funcName=%s|param=(%+v)|accReturn=%s",
		file, nowTime, usEastTime, funcName, paramText, accReturn,
	)
	if customErr != nil {
		accLog += fmt.Sprintf("|error_code=%d|error_msg=%s", customError.ErrCode, customError.ErrMsg)
	}
	if err != nil {
		accLog += fmt.Sprintf("|origin_error=%+v", err)
	}
	fileOpen, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer fileOpen.Close()
	logger := log.New(fileOpen, "", 0)
	logger.Println(accLog)
}

func WriteRedisLog(rbName string, errMsg string, err error) {
	file := "redis-error"
	dirPath := "./txt/modules"
	_ = checkDir(dirPath)
	logPath := getPath(dirPath, file)
	nowTime, usEastTime := getTime()
	redisLog := fmt.Sprintf(
		"type=%s|datetime=%s|datetime_gmt=%s|rbName=%s|err_message=%s",
		file, nowTime, usEastTime, rbName, errMsg,
	)
	if err != nil {
		redisLog += fmt.Sprintf("|origin_error=%+v", err)
	}
	fileOpen, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer fileOpen.Close()
	logger := log.New(fileOpen, "", 0)
	logger.Println(redisLog)
}

func WriteApiLog(funcName string, customErr error, err error, apiReturn string, param []interface{}) {
	file := "api-error"
	dirPath := "./txt/modules"
	_ = checkDir(dirPath)
	logPath := getPath(dirPath, file)
	nowTime, usEastTime := getTime()
	customError := customErr.(Error.CustomError)
	paramText := ""
	for idx, val := range param {
		if idx > 0 {
			paramText += ", "
		}
		paramText = fmt.Sprintf("%s%v", paramText, val)
	}
	apiLog := fmt.Sprintf(
		"type=%s|datetime=%s|datetime_gmt=%s|funcName=%s|param=(%+v)|apiReturn=%s",
		file, nowTime, usEastTime, funcName, paramText, apiReturn,
	)
	if customErr != nil {
		apiLog += fmt.Sprintf("|error_code=%d|error_msg=%s", customError.ErrCode, customError.ErrMsg)
	}
	if err != nil {
		apiLog += fmt.Sprintf("|origin_error=%+v", err)
	}
	fileOpen, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer fileOpen.Close()
	logger := log.New(fileOpen, "", 0)
	logger.Println(apiLog)
}

func WritePanicLog(r *http.Request, errMsg string) {
	dirPath := "./txt/sys"
	file := "panic"
	_ = checkDir(dirPath)
	logPath := getPath(dirPath, file)
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("panic log open err : %+v", err)
	}
	defer logFile.Close()
	_ = r.ParseForm()
	nowTime, usEastTime := getTime()
	msg := fmt.Sprintf(
		"type=%s|datetime=%s|datetime_gmt=%s|method=%s|param=%+v|url=%s|error=%s",
		file, nowTime, usEastTime, r.Method, r.Form.Encode(), r.URL.Path, errMsg,
	)
	trace := ""
	for skip := 1; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		if file[len(file)-1] == 'c' {
			continue
		}
		f := runtime.FuncForPC(pc)
		if skip != 1 {
			trace += "<-"
		}
		trace += fmt.Sprintf("%s:%d %s()", file, line, f.Name())
	}
	msg = fmt.Sprintf("%s|trace=%s", msg, trace)
	logger := log.New(logFile, "", 0)
	logger.Println(msg)
	if os.Getenv("MODE") == "dev" {
		log.Printf("%s", msg)
	}
}

func getLogContent(logName string, info interface{}) (logContent string) {
	nowTime, usEastTime := getTime()
	logContent = fmt.Sprintf("type=%s|datetime=%s|datetime_gmt=%s|", logName, nowTime, usEastTime)
	infoType := reflect.TypeOf(info)
	infoValue := reflect.ValueOf(info)
	for i := 0; i < infoType.NumField(); i++ {
		field := infoType.Field(i)
		value := infoValue.Field(i).Interface()
		logContent += fmt.Sprintf("%s=%v|", field.Name, value)
	}
	return
}

func getTime() (now string, usEast string) {
	now = time.Now().Format("2006-01-02 15:04:05")
	usEast = time.Now().UTC().Add(time.Hour * -4).Format("2006-01-02 15:04:05")
	return
}
