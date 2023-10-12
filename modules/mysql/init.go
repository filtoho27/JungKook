package mysql

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"jungkook/foundation"
	customLog "jungkook/log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type P struct {
	AccountDB  MasterSlaveDB
	ProductDB  MasterSlaveDB
}

type MasterSlaveDB struct {
	Master *gorm.DB
	Slave  *gorm.DB
}

var mysqlConn *P
var config *viper.Viper

// 初始化連線
func Init() {
	config = foundation.GetViperConfig("./config/modules/", "mysql", "yaml")
	wg := new(sync.WaitGroup)

	mysqlConn = &P{}
	rt := reflect.TypeOf(*mysqlConn)
	rv := reflect.ValueOf(mysqlConn)
	dbNum := rt.NumField()
	var connStatus = make(chan string, dbNum)
	wg.Add(dbNum)

	for i := 0; i < dbNum; i++ {
		dbKey := rt.Field(i).Name
		go func() {
			defer wg.Done()
			retryTime := 0
			for retryTime < 2 {
				retryTime++
				conn, mErr, sErr := newMasterSlaveDB(dbKey)
				rv.Elem().FieldByName(dbKey).Set(reflect.ValueOf(conn))
				if retryTime == 1 && (mErr != nil || sErr != nil) {
					customLog.WriteLog("modules", "DBRetry", nil, nil, "DBName=%s|MasterErr=%+v|SlaveErr=%+v", dbKey, mErr, sErr)
					time.Sleep(500 * time.Microsecond)
					continue
				}
				msg := ""
				if mErr != nil {
					msg = fmt.Sprintf("Master: %s %+v,", dbKey, mErr)
				}
				if sErr != nil {
					msg = fmt.Sprintf("%s Slave: %s %+v", msg, dbKey, sErr)
				}
				connStatus <- msg
				break
			}
		}()
	}
	wg.Wait()
	close(connStatus)
	shutdown := false
	shutdownMsg := "Service Init Shutdown:"
	for val := range connStatus {
		if val != "" {
			shutdown = true
			shutdownMsg = fmt.Sprintf("%s\n%s", shutdownMsg, val)
		}
	}
	if shutdown {
		log.Fatal(shutdownMsg)
	}
}

// 建立連線
func newMasterSlaveDB(dbKey string) (masterSlaveDB MasterSlaveDB, mErr error, sErr error) {
	masterSlaveDB.Master, mErr = newOrm(dbKey, "m")
	masterSlaveDB.Slave, sErr = newOrm(dbKey, "s")
	return
}

func newOrm(dbKey string, dbType string) (db *gorm.DB, err error) {
	dbConfig := config.GetStringMapString(dbKey)
	dbName := dbConfig["db"]
	if dbName == "" {
		dbName = dbKey
	}
	connectName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=30s",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host_"+dbType],
		dbConfig["port"],
		dbName,
	)

	newLogger := logger.New(
		&customLog.GormLog{DBName: connectName},
		logger.Config{
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,  // 過濾first find ...等撈不到資料的錯誤回傳
			Colorful:                  false, // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(connectName), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return
	}

	conn, err := db.DB()
	if err != nil {
		return
	}

	conn.SetMaxIdleConns(config.GetInt("idleConns")) // 設置空閒連線的最大數量
	conn.SetMaxOpenConns(config.GetInt("openConns")) // 設置打開連線的最大數量

	return db, nil
}
