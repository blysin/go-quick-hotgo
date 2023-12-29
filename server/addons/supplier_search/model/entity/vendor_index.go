// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VendorIndex is the golang structure for table vendor_index.
type VendorIndex struct {
	Id          int64       `json:"id"             description:"自增ID"`
	UnitKey     string      `json:"unitKey"        description:"唯一标识，当前已barcode作为标识"`
	VenDetailId int64       `json:"venDetailId"    description:"关联的hg_fm_vendor_detail.id"`
	Status      int         `json:"status"         description:"状态：0-不可用，1-可用"`
	CreatedAt   *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"      description:"更新时间"`
}
