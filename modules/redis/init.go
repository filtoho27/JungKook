package redis

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	Type "jungkook/commonType"
	"jungkook/foundation"
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
	config = foundation.GetViperConfig("./config/modules/", "redis", "yaml")
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

// 建立連線
func newMasterSlaveRB(rbName string) (masterSlaveRB MasterSlaveRB) {
	masterSlaveRB.Master = newPool(rbName, "m")
	masterSlaveRB.Slave = newPool(rbName, "s")
	return
}

func newPool(rbName string, rbType string) (rb *redis.Pool) {
	rbConfig := config.GetStringMapString(rbName)
	server := fmt.Sprintf("%s:%s", rbConfig["host_"+rbType], rbConfig["port"])

	rb = &redis.Pool{
		Wait:        true,
		MaxIdle:     config.GetInt("maxIdle"),   // 設置空閒連線的最大數量
		MaxActive:   config.GetInt("maxActive"), // 設置打開連線的最大數量
		IdleTimeout: 3 * time.Second,            // 設置空閒連接超時時間
		Dial: func() (redis.Conn, error) {
			var conn redis.Conn
			var err error
			retryTime := 0
			for retryTime < 2 {
				if retryTime == 1 {
					time.Sleep(500 * time.Microsecond)
				}
				retryTime++

				conn, err = redis.Dial("tcp",
					server,
					redis.DialConnectTimeout(30*time.Second),
					redis.DialReadTimeout(3*time.Second),
					redis.DialWriteTimeout(3*time.Second),
				)

				if err != nil {
					connLog := getRedisConnLog(rb, rbName, retryTime, "REDIS_CONNECTION_ERROR")
					customLog.WriteLog("modules", "RedisError", err, nil, connLog)
					continue
				}

				// _, err = conn.Do("AUTH", rbConfig["password"])
				// if err != nil {
				// 	connLog := getRedisConnLog(rb, rbName, retryTime, "REDIS_AUTH_ERROR")
				// 	customLog.WriteLog("modules", "RedisError", err, nil, connLog)
				// 	continue
				// }

				_, err = conn.Do("SELECT", rbConfig["index"])
				if err != nil {
					connLog := getRedisConnLog(rb, rbName, retryTime, "REDIS_SELECT_INDEX_ERROR")
					customLog.WriteLog("modules", "RedisError", err, nil, connLog)
					continue
				}
				break
			}
			return conn, err
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

func getRedisConnLog(p *redis.Pool, name string, time int, err string) string {
	status := p.Stats()
	msg := fmt.Sprintf(
		"RBName=%s|ErrorMsg=%s|ReConnTime:%d|Active=%d|Idle=%d|Wait=%d|WaitDuration=%d",
		name, err, time, status.ActiveCount, status.IdleCount, status.WaitCount, status.WaitDuration,
	)
	return msg
}
