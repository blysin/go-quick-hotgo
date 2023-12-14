// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VendorDetail is the golang structure for table vendor_detail.
type VendorDetail struct {
	Id               uint        `json:"id"               description:"自增ID"`
	VendorId         uint        `json:"vendorId"         description:"供应商主表id"`
	Brand            string      `json:"brand"            description:"品牌"`
	Barcode          string      `json:"barcode"          description:"条码"`
	EnglishName      string      `json:"englishName"      description:"英文名称"`
	Cost             int64       `json:"cost"             description:"成本、供货价"`
	CostCny          int64       `json:"costCny"          description:"成本、供货价-人民币"`
	SellingPrice     int64       `json:"sellingPrice"     description:"销售价格"`
	SellingPriceCny  int64       `json:"sellingPriceCny"  description:"销售价格-人民币"`
	Vendor           string      `json:"vendor"           description:"供应商"`
	Currency         string      `json:"currency"         description:"币种"`
	ExchangeRate     float64     `json:"exchangeRate"     description:"汇率"`
	ExchangeRateTime *gtime.Time `json:"exchangeRateTime" description:"汇率时间"`
	VendorData       string      `json:"vendorData"       description:"工资内容,json格式存储"`
	CreateTime       *gtime.Time `json:"createTime"       description:"创建时间"`
	UpdateTime       *gtime.Time `json:"updateTime"       description:"更新时间"`
	CreateBy         string      `json:"createBy"         description:"创建人"`
	UpdateBy         string      `json:"updateBy"         description:"更新人"`
}
