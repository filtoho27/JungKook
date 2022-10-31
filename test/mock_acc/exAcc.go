package mock_acc

type GetDomainData_Mock struct {
	now    int
	Result []string
	Error  []error
}

func (m *MockAcc) GetDomainData(hallID int) (result string, err error) {
	var now int
	m.GetDomainData_Mock.now, now = checknow(m.GetDomainData_Mock.now, len(m.GetDomainData_Mock.Result))
	result = m.GetDomainData_Mock.Result[now]
	err = m.GetDomainData_Mock.Error[now]
	return
}
