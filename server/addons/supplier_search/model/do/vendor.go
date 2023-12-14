// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Vendor is the golang structure of table hg_fm_vendor for DAO operations like Where/Data.
type Vendor struct {
	g.Meta         `orm:"table:hg_fm_vendor, do:true"`
	Id             interface{} // 自增ID
	VendorName     interface{} // 供应商名称
	AllColumn      interface{} // 完整字段，多个用英文逗号隔开
	RequiredColumn interface{} // 隐藏字段，多个用英文逗号隔开
	IsDeleted      interface{} // 是否删除，0：未删除，1：已删除
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
	CreateBy       interface{} // 创建人
	UpdateBy       interface{} // 更新人
}
