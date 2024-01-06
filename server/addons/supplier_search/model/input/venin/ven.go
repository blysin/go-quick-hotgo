package venin

import (
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/internal/model/input/form"
)

type VenSaveInp struct {
	Id           int64         `json:"id"            description:"主表id"`
	VendorName   string        `json:"vendorName"            description:"供应商名称"`
	FileId       int64         `json:"fileId"            description:"文件id"`
	Exchange     string        `json:"exchange"            description:"货币类型"`
	PresetColumn *PresetColumn `json:"presetColumn"            description:"预设列"`
}

type VenPageApiInp struct {
	form.PageReq
	VendorName      string `json:"vendorName" dc:"供应商名称"`
	Brand           string `json:"brand" dc:"品牌名称"`
	OrderBy         int    `json:"orderBy" dc:"排序"`
	OrderByPriceAsc int    `json:"orderByPriceAsc" dc:"价格升序"`
	CostStart       int    `json:"costStart" dc:"价格区间-开始"`
	CostEnd         int    `json:"costEnd" dc:"价格区间-结束"`
	CostCnyStart    int    `json:"costCnyStart" dc:"价格区间-开始"`
	CostCnyEnd      int    `json:"costCnyEnd" dc:"价格区间-结束"`
}

type PageApiModel struct {
	form.PageRes
	List []*entity.VendorDetail `json:"list" dc:"数据列表"`
}
