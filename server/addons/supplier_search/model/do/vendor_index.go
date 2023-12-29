// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VendorIndex is the golang structure of table hg_fm_vendor_index for DAO operations like Where/Data.
type VendorIndex struct {
	g.Meta      `orm:"table:hg_fm_vendor_index, do:true"`
	Id          interface{} // 自增ID
	UnitKey     interface{} // 唯一标识，当前已barcode作为标识
	VenDetailId interface{} // 关联的hg_fm_vendor_detail.id
	Status      interface{} // 状态：0-正常，1-作废
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
