package log

import (
	"fmt"
	Error "jungkook/error"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func WriteLog(folder string, file string, err error, customErr error, paramMsg string, params ...interface{}) {
	dirPath := fmt.Sprintf("./txt/%s", folder)
	_ = checkDir(dirPath)
	logPath := getPath(dirPath, file)
	fileOpen, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer fileOpen.Close()
	nowTime, usEastTime := getTime()
	logMsg := fmt.Sprintf(
		"type=%s|datetime=%s|datetime_gmt=%s|",
		file, nowTime, usEastTime,
	)
	if customErr != nil {
		customError := customErr.(Error.CustomError)
		logMsg += fmt.Sprintf("ErrorCode=%d|ErrorMsg=%s|", customError.ErrCode, customError.ErrMsg)
	}
	if err != nil {
		logMsg += fmt.Sprintf("OriginError=%+v|", err)
	}
	if paramMsg != "" {
		logMsg += fmt.Sprintf("%s|", paramMsg)
	}
	logger := log.New(fileOpen, "", 0)
	logger.Printf(logMsg, params...)
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

func getTime() (now string, usEast string) {
	now = time.Now().Format("2006-01-02 15:04:05")
	usEast = time.Now().UTC().Add(time.Hour * -4).Format("2006-01-02 15:04:05")
	return
}
