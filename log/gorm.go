package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type GormLog struct {
	DBName string
}

func (g *GormLog) Printf(text string, Vars ...interface{}) {
	file := "DBError"
	dirPath := "./txt/modules"
	_ = checkDir(dirPath)
	logPath := getPath(dirPath, file)
	nowTime, usEastTime := getTime()
	text = strings.Replace(text, "\n", "", -1)
	gormLog := fmt.Sprintf(text, Vars...)
	dbLog := fmt.Sprintf(
		"type=%s|datetime=%s|datetime_gmt=%s|db_connect_info=%s|gorm_err=%s",
		file, nowTime, usEastTime, g.DBName, gormLog,
	)
	fileOpen, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	defer fileOpen.Close()
	logger := log.New(fileOpen, "", 0)
	logger.Printf(dbLog)
}
