package example

import (
	"encoding/json"
	Type "jungkook/commonType"
	Error "jungkook/error"
	"jungkook/modules/acc"
	"jungkook/test/mock_acc"
	"reflect"
	"testing"
)

func TestExAcc1(t *testing.T) {
	t.Run("TestExAccSuccess", func(t *testing.T) {
		// 參數
		module := &Type.ModuleType{Vendor: 1}
		hallID := 6

		// mock acc
		accData := struct {
			Ret HallInfo `json:"ret"`
			Type.AccResult
		}{
			AccResult: Type.AccResult{Result: "ok"},
			Ret:       HallInfo{Name: "Esball", LoginCode: "esb"},
		}
		accJson, _ := json.Marshal(accData)
		mockAcc := mock_acc.MockAcc{}
		mockAcc.GetDomainData_Mock = mock_acc.GetDomainData_Mock{
			Result: []string{string(accJson)},
			Error:  []error{nil},
		}
		module.Acc = &acc.AccType{
			BBIN: &mockAcc,
		}

		// 期望值
		expectResult := HallInfo{
			Name:      "Esball",
			LoginCode: "esb",
		}

		// 比對
		result, err := exAcc1(module, hallID)
		if err != nil || !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"API ExAcc Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got: %+v",
				expectResult, result, nil, err,
			)
		}
	})

	t.Run("TestExAccConnectFail", func(t *testing.T) {
		// 參數
		module := &Type.ModuleType{Vendor: 1}
		hallID := 6

		// mock acc
		mockAcc := mock_acc.MockAcc{}
		mockAcc.GetDomainData_Mock = mock_acc.GetDomainData_Mock{
			Result: []string{""},
			Error:  []error{Error.CustomError{ErrMsg: "Acc連線失敗", ErrCode: 1338000001}},
		}
		module.Acc = &acc.AccType{
			BBIN: &mockAcc,
		}

		// 期望值
		expectResult := HallInfo{
			Name:      "",
			LoginCode: "",
		}
		expectErr := Error.CustomError{ErrMsg: "Acc連線失敗", ErrCode: 1338000001}

		// 比對
		result, err := exAcc1(module, hallID)
		if err.Error() != expectErr.Error() || !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"API ExAcc Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got: %+v",
				expectResult, result, expectErr, err,
			)
		}
	})

	t.Run("TestExAccResultFail", func(t *testing.T) {
		// 參數
		module := &Type.ModuleType{Vendor: 1}
		hallID := 3

		// mock acc
		accData := struct {
			Type.AccResult
		}{
			AccResult: Type.AccResult{
				Result: "error",
				Code:   150360004,
				Msg:    "廳的設定不存在",
			},
		}
		accJson, _ := json.Marshal(accData)
		mockAcc := mock_acc.MockAcc{}
		mockAcc.GetDomainData_Mock = mock_acc.GetDomainData_Mock{
			Result: []string{string(accJson)},
			Error:  []error{nil},
		}
		module.Acc = &acc.AccType{
			BBIN: &mockAcc,
		}

		// 期望值
		expectResult := HallInfo{
			Name:      "",
			LoginCode: "",
		}
		expectErr := Error.CustomError{ErrMsg: "GET_DOMAIN_ERROR", ErrCode: 1006}

		// 比對
		result, err := exAcc1(module, hallID)
		if err.Error() != expectErr.Error() || !reflect.DeepEqual(result, expectResult) {
			t.Errorf(
				"API ExAcc Test Fail: ExpectResult: %+v, got: %+v, ExpectError: %+v got: %+v",
				expectResult, result, expectErr, err,
			)
		}
	})
}
