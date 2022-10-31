package mysql

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetUserInfo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
		return
	}
	defer db.Close()

	mockGorm, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	mockP := &P{}
	mockP.CoreDB = MasterSlaveDB{
		Master: mockGorm,
		Slave:  mockGorm,
	}
	mysqlConn = mockP

	t.Run("GetUserInfoSuccess", func(t *testing.T) {
		// 參數
		name := "testUser"
		// Mock
		rows := sqlmock.NewRows([]string{"Name", "Password", "Google2faSecret", "Enable", "LoginError"}).
			AddRow("testUser", "testPwd", "testFaSecret", 1, 0)
		mock.ExpectQuery(
			regexp.QuoteMeta(
				"SELECT `Name`, `Password`, `Google2faSecret`, `Enable`, `LoginError` FROM `RD3Admin` WHERE `Name` = ? ORDER BY `RD3Admin`.`ID` LIMIT 1",
			),
		).WithArgs(
			name,
		).WillReturnRows(rows)
		// 比對
		expectResult := UserInfo{
			Name:       "testUser",
			Password:   "testPwd",
			FaSecret:   "testFaSecret",
			Enable:     1,
			LoginError: 0,
		}
		result, err := GetCoreDB().GetUserInfo(name)
		if err != nil {
			t.Errorf("Sql Test Fail: expect err: %+v, got: %s", nil, err.Error())
		} else if !reflect.DeepEqual(result, expectResult) {
			t.Errorf("Sql Test Fail: expect result: %+v, got: %+v", expectResult, result)
		}
	})

	t.Run("GetUserInfoError", func(t *testing.T) {
		// 參數
		name := "testUser"
		// Mock
		mock.ExpectQuery(
			regexp.QuoteMeta(
				"SELECT `Name`, `Password`, `Google2faSecret`, `Enable`, `LoginError` FROM `RD3Admin` WHERE `Name` = ? ORDER BY `RD3Admin`.`ID` LIMIT 1",
			),
		).WithArgs(
			name,
		).WillReturnError(errors.New("timeout"))
		// 比對
		expectErr := fmt.Sprintf("ErrMsg: %s, ErrCode: %d", "GET_USERINFO_FAIL", 1001)
		_, err := GetCoreDB().GetUserInfo(name)
		if err.Error() != expectErr {
			t.Errorf("Sql Test Fail: expect err: %s, got: %s", expectErr, err.Error())
		}
		// 判斷真的是否為預期的錯誤
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
