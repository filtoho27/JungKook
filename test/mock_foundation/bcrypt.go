package mock_foundation

type MockBcrypt struct {
	GeneratePassword_Mock
	ComparePassword_Mock
}

type GeneratePassword_Mock struct {
	Result string
	Error  error
}

func (mb MockBcrypt) GenerateFromPassword(pwd string) (result string, err error) {
	result = mb.GeneratePassword_Mock.Result
	err = mb.GeneratePassword_Mock.Error
	return
}

type ComparePassword_Mock struct {
	Error error
}

func (mb MockBcrypt) CompareHashAndPassword(pwd string, password string) (err error) {
	err = mb.ComparePassword_Mock.Error
	return
}
