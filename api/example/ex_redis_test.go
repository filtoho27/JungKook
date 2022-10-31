package example

import (
	Type "jungkook/commonType"
	Error "jungkook/error"
	"jungkook/test/mock_foundation"
	"jungkook/test/mock_redis"
	"reflect"
	"testing"
)

func TestExRedis(t *testing.T) {
	t.Run("TestExRedisSuccess", func(t *testing.T) {
		// 參數
		module := &Type.ModuleType{}
		name := "testUser"

		// mock time
		mockTime := mock_foundation.MockTime{}
		mockTime.Time = mock_foundation.GetTime("2022-12-16 16:35:00")
		module.Time = &mockTime

		// mock redis
		mockRedis := mock_redis.MockRedis{}
		mockRedis.SetExRedisData_Result = mock_redis.SetExRedisData_Result{
			Result: []error{nil},
		}
		mockRedis.GetExRedisData_Result = mock_redis.GetExRedisData_Result{
			Result1: []Type.ExRedis{
				{Name: name, DateTime: "2022-12-16 16:35:00"},
			},
			Result2: []error{nil},
		}
		module.Redis = &mockRedis

		// 期望值
		expectResult := Type.ExRedis{
			Name:     "testUser",
			DateTime: "2022-12-16 16:35:00",
		}

		// 比對
		result, err := exRedis(module, name)
		if err != nil || !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"API ExRedis Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got %+v",
				expectResult, result, nil, err,
			)
		}
	})

	t.Run("TestExRedisSetFail", func(t *testing.T) {
		// 參數
		module := &Type.ModuleType{}
		name := "testUser"

		// mock time
		mockTime := mock_foundation.MockTime{}
		mockTime.Time = mock_foundation.GetTime("2022-12-16 16:35:00")
		module.Time = &mockTime

		// mock redis
		mockRedis := mock_redis.MockRedis{}
		mockRedis.SetExRedisData_Result = mock_redis.SetExRedisData_Result{
			Result: []error{Error.CustomError{ErrMsg: "SET_EX_DATA_ERROR", ErrCode: 1002}},
		}
		module.Redis = &mockRedis

		// 期望值
		expectResult := Type.ExRedis{}
		expectErr := Error.CustomError{ErrMsg: "SET_EX_DATA_ERROR", ErrCode: 1002}

		// 比對
		result, err := exRedis(module, name)
		if err.Error() != expectErr.Error() || !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"API ExRedis Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got %+v",
				expectResult, result, expectErr, err,
			)
		}
	})

	t.Run("TestExRedisGetFail", func(t *testing.T) {
		// 參數
		module := &Type.ModuleType{}
		name := "testUser"

		// mock time
		mockTime := mock_foundation.MockTime{}
		mockTime.Time = mock_foundation.GetTime("2022-12-16 16:35:00")
		module.Time = &mockTime

		// mock redis
		mockRedis := mock_redis.MockRedis{}
		mockRedis.SetExRedisData_Result = mock_redis.SetExRedisData_Result{
			Result: []error{nil},
		}
		mockRedis.GetExRedisData_Result = mock_redis.GetExRedisData_Result{
			Result1: []Type.ExRedis{{}},
			Result2: []error{Error.CustomError{ErrMsg: "GET_EX_DATA_ERROR", ErrCode: 1003}},
		}
		module.Redis = &mockRedis

		// 期望值
		expectResult := Type.ExRedis{}
		expectErr := Error.CustomError{ErrMsg: "GET_EX_DATA_ERROR", ErrCode: 1003}

		// 比對
		result, err := exRedis(module, name)
		if err.Error() != expectErr.Error() || !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"API ExRedis Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got %+v",
				expectResult, result, expectErr, err,
			)
		}
	})
}
