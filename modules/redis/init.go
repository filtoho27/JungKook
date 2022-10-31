package redis

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	Type "jungkook/commonType"
	customLog "jungkook/log"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type P struct {
	Central MasterSlaveRB
	Member  MasterSlaveRB
}

type MasterSlaveRB struct {
	Master *redis.Pool
	Slave  *redis.Pool
}

type RedisLib struct{}

var redisPool *P
var config *viper.Viper
var redisLib Type.RedisInterface

// 初始化連線
func Init() {
	getConfig()
	wg := new(sync.WaitGroup)

	redisPool = &P{}
	rt := reflect.TypeOf(*redisPool)
	rv := reflect.ValueOf(redisPool)
	rbNum := rt.NumField()
	wg.Add(rbNum)

	for i := 0; i < rbNum; i++ {
		rbName := rt.Field(i).Name
		go func() {
			defer wg.Done()
			conn := newMasterSlaveRB(rbName)
			rv.Elem().FieldByName(rbName).Set(reflect.ValueOf(conn))
		}()
	}

	wg.Wait()
}

// 取得Redis物件
func GetRedis() Type.RedisInterface {
	if redisLib == nil {
		redisLib = &RedisLib{}
	}
	return redisLib
}

// 取得設定檔
func getConfig() {
	config = viper.New()
	config.SetConfigName("redis")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config/modules/")
	err := config.ReadInConfig()
	if err != nil {
		customLog.WriteRedisLog("", "GET_REDIS_CONFIG_FAILED", err)
	}
}

// 建立連線
func newMasterSlaveRB(rbName string) (masterSlaveRB MasterSlaveRB) {
	masterSlaveRB.Master = newPool(rbName, "master")
	masterSlaveRB.Slave = newPool(rbName, "slave")
	return
}

func newPool(rbName string, rbType string) (rb *redis.Pool) {
	rbKey := fmt.Sprintf("redis.%s.%s", rbName, rbType)
	rbConfig := config.GetStringMapString(rbKey)
	host := rbConfig["host"]
	port := rbConfig["port"]
	// auth := rbConfig["password"]
	index := rbConfig["index"]

	rb = &redis.Pool{
		Wait:        true,
		MaxIdle:     config.GetInt("redis.maxIdle"),   // 設置空閒連線的最大數量
		MaxActive:   config.GetInt("redis.maxActive"), // 設置打開連線的最大數量
		IdleTimeout: 3 * time.Second,                  // 設置空閒連接超時時間
		Dial: func() (redis.Conn, error) {
			server := host + ":" + port
			rb, err := redis.Dial("tcp",
				server,
				redis.DialConnectTimeout(time.Duration(1)*time.Minute),
				redis.DialReadTimeout(time.Duration(3)*time.Second),
				redis.DialWriteTimeout(time.Duration(3)*time.Second),
			)

			if err != nil {
				customLog.WriteRedisLog(rbName, "REDIS_CONNECTION_ERROR", err)
				return rb, err
			}

			// _, err = rb.Do("AUTH", auth)
			// if err != nil {
			// 	customLog.WriteRedisLog(rbName, "REDIS_AUTH_ERROR", err)
			// 	return rb, err
			// }

			_, err = rb.Do("SELECT", index)
			if err != nil {
				customLog.WriteRedisLog(rbName, "REDIS_SELECT_INDEX_ERROR", err)
				return rb, err
			}

			return rb, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return
}
