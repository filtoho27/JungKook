package error

import "fmt"

type CustomError struct {
	ErrMsg  string `json:"errMsg"`
	ErrCode int    `json:"errCode"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("ErrMsg: %s, ErrCode: %d", e.ErrMsg, e.ErrCode)
}
