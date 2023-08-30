package mysql

import (
	"fmt"
	"log"
	"reflect"
	"sync"

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
var connStatus = make(chan error, 6)

// 初始化連線
func Init() {
	getConfig()
	wg := new(sync.WaitGroup)

	mysqlConn = &P{}
	rt := reflect.TypeOf(*mysqlConn)
	rv := reflect.ValueOf(mysqlConn)
	dbNum := rt.NumField()
	wg.Add(dbNum)

	for i := 0; i < dbNum; i++ {
		dbName := rt.Field(i).Name
		go func() {
			defer wg.Done()
			conn, Merr, Serr := newMasterSlaveDB(dbName)
			rv.Elem().FieldByName(dbName).Set(reflect.ValueOf(conn))
			if Merr != nil {
				connStatus <- Merr
			} else if Serr != nil {
				connStatus <- Serr
			}
		}()
	}
	wg.Wait()

	close(connStatus)
	for val := range connStatus {
		if val != nil {
			msg := fmt.Sprintf("DB_INIT_CONNECTION_FAILED, %+v", val)
			log.Fatalf(msg)
		}
	}
}

// 取得設定檔
func getConfig() {
	config = viper.New()
	config.SetConfigName("mysql")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config/modules/")
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("GET_MYSQL_CONFIG_FAILED, %+v", err)
	}
}

// 建立連線
func newMasterSlaveDB(dbName string) (masterSlaveDB MasterSlaveDB, Merr error, Serr error) {
	masterSlaveDB.Master, Merr = newOrm(dbName, "master")
	masterSlaveDB.Slave, Serr = newOrm(dbName, "slave")
	return
}

func newOrm(dbName string, dbType string) (db *gorm.DB, err error) {
	dbKey := fmt.Sprintf("mysql.%s.%s", dbName, dbType)
	dbConfig := config.GetStringMapString(dbKey)
	connectName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=30s",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
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

	conn.SetMaxIdleConns(config.GetInt("mysql.idleConns")) // 設置空閒連線的最大數量
	conn.SetMaxOpenConns(config.GetInt("mysql.openConns")) // 設置打開連線的最大數量
	conn.SetConnMaxLifetime(1)                             // 設置連線的生命週期

	return db, nil
}
