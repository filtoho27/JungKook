package foundation

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

var b *Bcrypt = &Bcrypt{}

func GetBcrypt() *Bcrypt {
	if b == nil {
		b = &Bcrypt{}
	}
	return b
}

func (b Bcrypt) GenerateFromPassword(pwd string) (encodePwd string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	encodePwd = string(hash)
	return
}

func (b Bcrypt) CompareHashAndPassword(pwd string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(string(password)), []byte(pwd))
	return
}
