package mock_redis

import (
	Type "jungkook/commonType"
)

type SetExRedisData_Result struct {
	now    int
	Result []error
}

func (m *MockRedis) SetExRedisData(name string, dtime string) (err error) {
	var now int
	m.SetExRedisData_Result.now, now = checknow(m.SetExRedisData_Result.now, len(m.SetExRedisData_Result.Result))
	err = m.SetExRedisData_Result.Result[now]
	return
}

type GetExRedisData_Result struct {
	now     int
	Result1 []Type.ExRedis
	Result2 []error
}

func (m *MockRedis) GetExRedisData() (result Type.ExRedis, err error) {
	var now int
	m.GetExRedisData_Result.now, now = checknow(m.GetExRedisData_Result.now, len(m.GetExRedisData_Result.Result1))
	result = m.GetExRedisData_Result.Result1[now]
	err = m.GetExRedisData_Result.Result2[now]
	return
}
