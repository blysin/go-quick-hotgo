// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VendorUploadFile is the golang structure for table vendor_upload_file.
type VendorUploadFile struct {
	Id                  int64       `json:"id"                  description:"自增ID"`
	VendorId            int64       `json:"vendorId"            description:"供应商主表id"`
	FileName            string      `json:"fileName"            description:"文件名称"`
	FileId              int64       `json:"fileId"              description:"文件id"`
	ExceptionDataFileId string      `json:"exceptionDataFileId" description:"异常数据文件id"`
	ValidNum            int         `json:"validNum"            description:"正常数据条数"`
	ExceptionNum        int         `json:"exceptionNum"        description:"异常数据条数"`
	AllColumn           string      `json:"allColumn"           description:"完整字段，多个用英文逗号隔开"`
	CreatedAt           *gtime.Time `json:"createdAt"          description:"创建时间"`
	UpdatedAt           *gtime.Time `json:"updatedAt"          description:"更新时间"`
	CreateBy            int64       `json:"createBy"            description:"创建人"`
	UpdateBy            int64       `json:"updateBy"            description:"更新人"`
}
