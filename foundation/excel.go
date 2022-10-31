package foundation

import (
	Type "jungkook/commonType"
	customLog "jungkook/log"
	"net/http"

	"github.com/tealeg/xlsx"
)

type excelLog struct {
	FileName  string
	SheetName string
	ErrMsg    string
}

func MakeExcel(w http.ResponseWriter, data Type.ExcelType) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(data.SheetName)
	if err != nil {
		handleLogRecord(data.FileName, data.SheetName, err.Error())
		return
	}

	// 處理header欄位
	style := xlsx.NewStyle()
	style.Fill.PatternType = "solid"
	style.Font.Color = "FFFFFF"
	row := sheet.AddRow()
	for _, v := range data.Header {
		cell := row.AddCell()
		cell.Value = v
		cell.SetStyle(style)
	}

	// 處理資料欄位
	for _, info := range data.Data {
		row := sheet.AddRow()
		for _, v := range info {
			cell := row.AddCell()
			cell.Value = v
		}
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename="+data.FileName+".xlsx")
	w.Header().Set("Cache-Control", "max-age=0")
	err = file.Write(w)
	if err != nil {
		handleLogRecord(data.FileName, data.SheetName, err.Error())
	}
}

func handleLogRecord(fileName string, sheetName string, msg string) {
	logData := excelLog{
		FileName:  fileName,
		SheetName: sheetName,
		ErrMsg:    msg,
	}
	customLog.WriteLog("tool", "excel-error", logData)
}
