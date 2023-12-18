// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Vendor is the golang structure for table vendor.
type Vendor struct {
	Id             int64       `json:"id"             description:"自增ID"`
	VendorName     string      `json:"vendorName"     description:"供应商名称"`
	AllColumn      string      `json:"allColumn"      description:"完整字段，多个用英文逗号隔开"`
	RequiredColumn string      `json:"requiredColumn" description:"比填列，json格式"`
	Status         int         `json:"status"      description:"状态：0-新增，-1-删除，2已发布"`
	CreatedAt      *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"     description:"更新时间"`
	CreateBy       int64       `json:"createBy"       description:"创建人"`
	UpdateBy       int64       `json:"updateBy"       description:"更新人"`
}
