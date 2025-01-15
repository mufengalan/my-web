package service

import (
	"encoding/json"
	"github.com/mufengalan/my-web/internal/convert"
	"github.com/xuri/excelize/v2"
)

type ExportService struct{}

func NewExportService() *ExportService {
	return &ExportService{}
}

func (s *ExportService) Export(data []byte) error {
	var exportRes convert.ExportRes
	err := json.Unmarshal(data, &exportRes)
	if err != nil {
		return err
	}
	showNameUrlMap := make(map[string]string)
	for _, v := range exportRes.Data {
		showNameUrlMap[v.ShowName] = v.URL
	}
	f, err := excelize.OpenFile("file/template_upload.xlsx")
	if err != nil {
		return err
	}
	//1: 获取excel行数据
	sheets := f.GetSheetList()
	sheetName := sheets[0]
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}
	//2: 修改excel数据
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for j, cell := range row {
			if url, ok := showNameUrlMap[cell]; ok {
				cellAddress, _ := excelize.CoordinatesToCellName(j+1, i+1)
				err = f.SetCellValue(sheetName, cellAddress, url)
				if err != nil {
					continue
				}
			}
		}
	}
	if err = f.Save(); err != nil {
		return err
	}

	return nil
}
