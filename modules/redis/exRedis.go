package redis

import (
	"encoding/json"
	Type "jungkook/commonType"
	Error "jungkook/error"

	"github.com/gomodule/redigo/redis"
)

func (r *RedisLib) SetExRedisData(name string, dtime string) (err error) {
	rb := redisPool.Central.Master.Get()
	defer rb.Close()

	redisKey := "ExRedis"
	redisData := Type.ExRedis{
		Name:     name,
		DateTime: dtime,
	}
	redisJson, _ := json.Marshal(redisData)
	_, err = redis.String(rb.Do("SET", redisKey, redisJson))
	if err != nil {
		err = Error.CustomError{ErrMsg: "SET_EX_DATA_ERROR", ErrCode: 1002}
	}
	return
}

func (r *RedisLib) GetExRedisData() (result Type.ExRedis, err error) {
	rb := redisPool.Central.Slave.Get()
	defer rb.Close()

	redisKey := "ExRedis"
	redisJson, err := redis.String(rb.Do("GET", redisKey))
	if err != nil {
		err = Error.CustomError{ErrMsg: "GET_EX_DATA_ERROR", ErrCode: 1003}
	}
	_ = json.Unmarshal([]byte(redisJson), &result)
	return
}
