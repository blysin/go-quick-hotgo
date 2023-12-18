// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VendorDetail is the golang structure of table hg_fm_vendor_detail for DAO operations like Where/Data.
type VendorDetail struct {
	g.Meta           `orm:"table:hg_fm_vendor_detail, do:true"`
	Id               interface{} // 自增ID
	VendorId         interface{} // 供应商主表id
	Brand            interface{} // 品牌
	Barcode          interface{} // 条码
	EnglishName      interface{} // 英文名称
	Cost             interface{} // 成本、供货价
	CostCny          interface{} // 成本、供货价-人民币
	SellingPrice     interface{} // 销售价格
	SellingPriceCny  interface{} // 销售价格-人民币
	Vendor           interface{} // 供应商
	Currency         interface{} // 币种
	ExchangeRate     interface{} // 汇率
	ExchangeRateTime *gtime.Time // 汇率时间
	VendorData       interface{} // 工资内容,json格式存储
	Status           interface{} // 状态：0-新增，-1-删除，2已发布
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	CreateBy         interface{} // 创建人
	UpdateBy         interface{} // 更新人
}
