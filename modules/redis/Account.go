package redis

import (
	"time"
	"encoding/json"
	Type "jungkook/commonType"
	Error "jungkook/error"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *RedisLib) SetEmailCode(userName string, code int) (err error) {
	rb := redisPool.Member.Master.Get()
	defer rb.Close()

	redisKey := "EmailCode:" + userName
	redisData := Type.EmailCode{code, time.Now().Unix()}
	redisJson, _ := json.Marshal(redisData)

	_, err = redis.String(rb.Do("SET", redisKey, redisJson))
	if err != nil {
		err = Error.CustomError{ErrMsg: "SET_EMAIL_CODE_FAIL", ErrCode: 1010007}
		return
	}

	_, err = redis.Bool(rb.Do("EXPIRE", redisKey, 86400))
	if err != nil {
		err = Error.CustomError{ErrMsg: "SET_EMAIL_CODE_EXPIRE_FAIL", ErrCode: 1010008}
	}
	return
}

func (r *RedisLib) GetEmailCode(userName string) (result Type.EmailCode, err error) {
	rb := redisPool.Member.Slave.Get()
	defer rb.Close()

	redisKey := "EmailCode:" + userName
	redisJson, err := redis.String(rb.Do("GET", redisKey))
	_ = json.Unmarshal([]byte(redisJson), &result)
	return
}

func (r *RedisLib) DelEmailCode(userName string) (err error) {
	rb := redisPool.Member.Master.Get()
	defer rb.Close()

	redisKey := "EmailCode:" + userName
	_, err = redis.String(rb.Do("DEL", redisKey))
	if err != nil {
		err = Error.CustomError{ErrMsg: "DEL_EMAIL_CODE_FAIL", ErrCode: 1010013}
	}
	return
}

func (r *RedisLib) SetUserToken(userID int, userToken string) (err error) {
	rb := redisPool.Member.Master.Get()
	defer rb.Close()

	redisKey := "UserToken:" + strconv.Itoa(userID)

	_, err = redis.String(rb.Do("SET", redisKey, userToken))
	if err != nil {
		err = Error.CustomError{ErrMsg: "SET_USER_TOKEN_FAIL", ErrCode: 1010016}
		return
	}

	_, err = redis.Bool(rb.Do("EXPIRE", redisKey, 5184000))
	if err != nil {
		err = Error.CustomError{ErrMsg: "SET_USER_TOKEN_EXPIRE_FAIL", ErrCode: 1010017}
	}
	return
}