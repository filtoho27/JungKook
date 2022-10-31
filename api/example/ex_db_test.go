package example

import (
	Error "jungkook/error"
	"jungkook/modules/mysql"
	"jungkook/test/mock_mysql"
	"testing"
)

func TestExDB(t *testing.T) {
	t.Run("TestExDBSuccess", func(t *testing.T) {
		// 參數
		acc := "testUser"
		pwd := "testPwd"

		// mock db
		mockMysql := mock_mysql.MockCoreDB{}
		mockMysql.GetUserInfo_Result = mock_mysql.GetUserInfo_Result{
			Result1: []mysql.UserInfo{
				{
					Name:       "testUser",
					Password:   "testPwd",
					FaSecret:   "testFaSecret",
					Enable:     1,
					LoginError: 0,
				},
			},
			Result2: []error{nil},
		}
		mysql.Core = &mockMysql

		// 期望值
		expectResult := true

		// 比對
		result, err := exDB(acc, pwd)
		if err != nil || result != expectResult {
			t.Errorf(
				"API ExDB Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got %+v",
				expectResult, result, nil, err,
			)
		}
	})

	t.Run("TestExDBWithDBFail", func(t *testing.T) {
		// 參數
		acc := "testUser"
		pwd := "testPwd"

		// mock db
		mockMysql := mock_mysql.MockCoreDB{}
		mockMysql.GetUserInfo_Result = mock_mysql.GetUserInfo_Result{
			Result1: []mysql.UserInfo{{}},
			Result2: []error{Error.CustomError{ErrMsg: "GET_USERINFO_FAIL", ErrCode: 1001}},
		}
		mysql.Core = &mockMysql

		// 期望值
		expectResult := false
		expectErr := Error.CustomError{ErrMsg: "GET_USERINFO_FAIL", ErrCode: 1001}

		// 比對
		result, err := exDB(acc, pwd)
		if err.Error() != expectErr.Error() || result != expectResult {
			t.Errorf(
				"API ExDB Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got %+v",
				expectResult, result, expectErr, err,
			)
		}
	})

	t.Run("TestExDBWithAccNotExist", func(t *testing.T) {
		// 參數
		acc := "testUser"
		pwd := "testPwd"

		// mock db
		mockMysql := mock_mysql.MockCoreDB{}
		mockMysql.GetUserInfo_Result = mock_mysql.GetUserInfo_Result{
			Result1: []mysql.UserInfo{
				{
					Name:       "",
					Password:   "",
					FaSecret:   "",
					Enable:     0,
					LoginError: 0,
				},
			},
			Result2: []error{nil},
		}
		mysql.Core = &mockMysql

		// 期望值
		expectResult := false
		expectErr := Error.CustomError{ErrMsg: "ACCOUNT_NOT_EXIST", ErrCode: 1004}

		// 比對
		result, err := exDB(acc, pwd)
		if err.Error() != expectErr.Error() || result != expectResult {
			t.Errorf(
				"API ExDB Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got %+v",
				expectResult, result, expectErr, err,
			)
		}
	})

	t.Run("TestExDBWithNotEnable", func(t *testing.T) {
		// 參數
		acc := "testUser"
		pwd := "testPwd"

		// mock db
		mockMysql := mock_mysql.MockCoreDB{}
		mockMysql.GetUserInfo_Result = mock_mysql.GetUserInfo_Result{
			Result1: []mysql.UserInfo{
				{
					Name:       "testUser",
					Password:   "testPwd",
					FaSecret:   "testFaSecret",
					Enable:     0,
					LoginError: 0,
				},
			},
			Result2: []error{nil},
		}
		mysql.Core = &mockMysql

		// 期望值
		expectResult := false
		expectErr := Error.CustomError{ErrMsg: "ACCOUNT_IS_DISABLED", ErrCode: 1005}

		// 比對
		result, err := exDB(acc, pwd)
		if err.Error() != expectErr.Error() || result != expectResult {
			t.Errorf(
				"API ExDB Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got %+v",
				expectResult, result, expectErr, err,
			)
		}
	})
}
