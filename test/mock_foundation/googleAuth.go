package mock_foundation

type MockGoogle struct {
	VerifyCode_Mock
	GetQrcodeUrl_Mock
	GetSecret_Mock
}

type VerifyCode_Mock struct {
	Result bool
	Error  error
}

func (mg MockGoogle) VerifyCode(secret string, code string) (result bool, err error) {
	result = mg.VerifyCode_Mock.Result
	err = mg.VerifyCode_Mock.Error
	return
}

type GetQrcodeUrl_Mock struct {
	Result string
}

func (mg MockGoogle) GetQrcodeUrl(user, secret string) (result string) {
	result = mg.GetQrcodeUrl_Mock.Result
	return
}

type GetSecret_Mock struct {
	Result string
}

func (mg MockGoogle) GetSecret() (result string) {
	result = mg.GetSecret_Mock.Result
	return
}
