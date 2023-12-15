// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VendorUploadFile is the golang structure of table hg_fm_vendor_upload_file for DAO operations like Where/Data.
type VendorUploadFile struct {
	g.Meta              `orm:"table:hg_fm_vendor_upload_file, do:true"`
	Id                  interface{} // 自增ID
	VendorId            interface{} // 供应商主表id
	FileName            interface{} // 文件名称
	FileId              interface{} // 文件id
	ExceptionDataFileId interface{} // 异常数据文件id
	ValidNum            interface{} // 正常数据条数
	ExceptionNum        interface{} // 异常数据条数
	AllColumn           interface{} // 完整字段，多个用英文逗号隔开
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
	CreateBy            interface{} // 创建人
	UpdateBy            interface{} // 更新人
}
