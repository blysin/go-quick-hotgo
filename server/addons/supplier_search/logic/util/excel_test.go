package util

import (
	"fmt"
	"github.com/antlabs/strsim"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/xuri/excelize/v2"
	"hotgo/addons/supplier_search/model/input/venin"
	"testing"
)

func TestReadTitle(t *testing.T) {

	path := "D:\\Downloads\\vendor.xlsx"

	f, err := excelize.OpenFile(path)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = f.Close()
	}()

	//"Row"（行）和 "Cell"（单元格）

	// 读取第一行作为标题
	i := f.GetActiveSheetIndex()
	sheetName := f.GetSheetName(i)

	rows, err := f.Rows(sheetName)
	if err != nil {
		t.Error(err)
	}

	// 获取第一行
	firstRows, err := getFirstRows(rows)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(firstRows)

	//品牌名	条码	英文名	供货价	销售价	供应商
	presetColumn := venin.PresetColumn{}

	presetColumn.BrandName = matchKeyWord("品牌", firstRows)
	presetColumn.BarCode = matchKeyWord("条码", firstRows)
	presetColumn.EnName = matchKeyWord("英文名", firstRows)
	presetColumn.SupplyPrice = matchKeyWord("供货价", firstRows)
	presetColumn.SalePrice = matchKeyWord("销售价", firstRows)
	presetColumn.VendorName = matchKeyWord("供应商", firstRows)

	fmt.Println(gjson.MustEncodeString(presetColumn))
}

func matchKeyWord(key string, keywords []string) string {
	//匹配最佳的关键字
	for _, keyword := range keywords {
		if key == keyword {
			return keyword
		}
	}
	//计算相似度
	result := strsim.FindBestMatch(key, keywords)
	return result.Match.S
}

func TestStrSim(t *testing.T) {
	result := strsim.FindBestMatch("少坤", []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"})
	t.Log(result.Match.S)
}

func getFirstRows(rows *excelize.Rows) ([]string, error) {
	results, cur := make([][]string, 0, 64), 0
	for rows.Next() {
		cur++
		row, err := rows.Columns()
		if err != nil {
			break
		}
		results = append(results, row)
		break
	}
	return results[0], rows.Close()
}
