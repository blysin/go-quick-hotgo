// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Vendor is the golang structure for table vendor.
type Vendor struct {
	Id             uint        `json:"id"             description:"自增ID"`
	VendorName     string      `json:"vendorName"     description:"供应商名称"`
	AllColumn      string      `json:"allColumn"      description:"完整字段，多个用英文逗号隔开"`
	RequiredColumn string      `json:"requiredColumn" description:"隐藏字段，多个用英文逗号隔开"`
	IsDeleted      int         `json:"isDeleted"      description:"是否删除，0：未删除，1：已删除"`
	CreateTime     *gtime.Time `json:"createTime"     description:"创建时间"`
	UpdateTime     *gtime.Time `json:"updateTime"     description:"更新时间"`
	CreateBy       string      `json:"createBy"       description:"创建人"`
	UpdateBy       string      `json:"updateBy"       description:"更新人"`
}
