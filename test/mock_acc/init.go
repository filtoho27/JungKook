package mock_acc

type MockAcc struct {
	GetDomainData_Mock
}

func checknow(nowType int, max int) (next int, now int) {
	now = nowType
	if next = nowType + 1; next >= max {
		next = nowType
	}
	return
}
