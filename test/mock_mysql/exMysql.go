package mock_mysql

import (
	"jungkook/modules/mysql"
)

type GetUserInfo_Result struct {
	now     int
	Result1 []mysql.UserInfo
	Result2 []error
}

func (m *MockCoreDB) GetUserInfo(acc string) (userInfo mysql.UserInfo, err error) {
	var now int
	m.GetUserInfo_Result.now, now = checknow(m.GetUserInfo_Result.now, len(m.GetUserInfo_Result.Result1))
	userInfo = m.GetUserInfo_Result.Result1[now]
	err = m.GetUserInfo_Result.Result2[now]
	return
}
