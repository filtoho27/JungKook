package commontype

type RedisInterface interface {
	SetEmailCode(userName string, code int) (err error)
	GetEmailCode(userName string) (result EmailCode, err error)
	DelEmailCode(userName string) (err error)
	SetUserToken(userID int, userToken string) (err error)
}

type EmailCode struct {
	Code      int   `json:"code"`
	Timestamp int64 `json:"timestamp"`
}
