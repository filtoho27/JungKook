package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

const maxSize int64 = 52428800 // 50M

// 生成資料夾
func checkDir(name string) error {
	_, err := os.Stat(name)
	if err == nil {
		return err
	}
	if !os.IsNotExist(err) {
		return nil
	}
	err = os.MkdirAll(name, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 刪除檔案
func deleteDir(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return nil
	}
	err = os.RemoveAll(name)
	if err != nil {
		return err
	}
	return nil
}

// 取得檔案路徑
func getPath(dir string, name string) (originPath string) {
	originPath = fmt.Sprintf("%s/%s.log", dir, name)
	size, _, _ := getFileInfo(originPath)
	if size > maxSize {
		bakPath := fmt.Sprintf("%s/%s.bak", dir, name)
		_ = deleteDir(bakPath)
		// 備份檔案
		originFile, _ := os.Open(originPath)
		defer originFile.Close()
		bakFile, _ := os.OpenFile(bakPath, os.O_WRONLY|os.O_CREATE, 0777)
		defer bakFile.Close()
		_, _ = io.Copy(bakFile, originFile)
		// 清空原本檔案
		_ = os.Truncate(originPath, 0)
	}
	return
}

// 取得檔案資訊
func getFileInfo(name string) (size int64, modTime time.Time, err error) {
	fi, err := os.Stat(name)
	if err == nil {
		size = fi.Size()
		modTime = fi.ModTime()
	}
	return
}
