package commontype

import (
	"time"
)

type TimeInterface interface {
	Now() time.Time
	Sleep(duration time.Duration)
	NowTaipei() time.Time
	NowUSEast() time.Time
}

type BcryptInterface interface {
	GenerateFromPassword(pwd string) (encodePwd string, err error)
	CompareHashAndPassword(pwd string, password string) (err error)
}

type GoogleInterface interface {
	VerifyCode(secret string, code string) (bool, error)
	GetQrcodeUrl(user, secret string) string
	GetSecret() string
}
