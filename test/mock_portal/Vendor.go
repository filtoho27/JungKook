package mock_portal

type SetHallVendorRedis_Mock struct {
	Result string
	Error  error
}

func (m *MockPortal) SetHallVendorRedis() (result string, err error) {
	result = m.SetHallVendorRedis_Mock.Result
	err = m.SetHallVendorRedis_Mock.Error
	return
}
