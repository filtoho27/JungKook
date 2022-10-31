package commontype

type RedisInterface interface {
	GetExRedisData() (result ExRedis, err error)
	SetExRedisData(name string, dtime string) (err error)
	SetEmailCode(userName string, code int) (err error)
	GetEmailCode(userName string) (result EmailCode, err error)
	DelEmailCode(userName string) (err error)
	SetUserToken(userID int, userToken string) (err error)
}

type ExRedis struct {
	Name     string `json:"name"`
	DateTime string `json:"datetime"`
}

type EmailCode struct {
	Code      int   `json:"code"`
	Timestamp int64 `json:"timestamp"`
}
