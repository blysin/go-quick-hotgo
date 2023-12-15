// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VendorUploadFileDao is the data access object for table hg_fm_vendor_upload_file.
type VendorUploadFileDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns VendorUploadFileColumns // columns contains all the column names of Table for convenient usage.
}

// VendorUploadFileColumns defines and stores column names for table hg_fm_vendor_upload_file.
type VendorUploadFileColumns struct {
	Id                  string // 自增ID
	VendorId            string // 供应商主表id
	FileName            string // 文件名称
	FileId              string // 文件id
	ExceptionDataFileId string // 异常数据文件id
	ValidNum            string // 正常数据条数
	ExceptionNum        string // 异常数据条数
	AllColumn           string // 完整字段，多个用英文逗号隔开
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
	CreateBy            string // 创建人
	UpdateBy            string // 更新人
}

// vendorUploadFileColumns holds the columns for table hg_fm_vendor_upload_file.
var vendorUploadFileColumns = VendorUploadFileColumns{
	Id:                  "id",
	VendorId:            "vendor_id",
	FileName:            "file_name",
	FileId:              "file_id",
	ExceptionDataFileId: "exception_data_file_id",
	ValidNum:            "valid_num",
	ExceptionNum:        "exception_num",
	AllColumn:           "all_column",
	CreatedAt:           "create_time",
	UpdatedAt:           "update_time",
	CreateBy:            "create_by",
	UpdateBy:            "update_by",
}

// NewVendorUploadFileDao creates and returns a new DAO object for table data access.
func NewVendorUploadFileDao() *VendorUploadFileDao {
	return &VendorUploadFileDao{
		group:   "default",
		table:   "hg_fm_vendor_upload_file",
		columns: vendorUploadFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VendorUploadFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VendorUploadFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VendorUploadFileDao) Columns() VendorUploadFileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VendorUploadFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VendorUploadFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VendorUploadFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
