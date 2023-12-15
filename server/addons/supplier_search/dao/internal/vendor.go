// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VendorDao is the data access object for table hg_fm_vendor.
type VendorDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns VendorColumns // columns contains all the column names of Table for convenient usage.
}

// VendorColumns defines and stores column names for table hg_fm_vendor.
type VendorColumns struct {
	Id             string // 自增ID
	VendorName     string // 供应商名称
	AllColumn      string // 完整字段，多个用英文逗号隔开
	RequiredColumn string // 隐藏字段，多个用英文逗号隔开
	IsDeleted      string // 是否删除，0：未删除，1：已删除
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	CreateBy       string // 创建人
	UpdateBy       string // 更新人
}

// vendorColumns holds the columns for table hg_fm_vendor.
var vendorColumns = VendorColumns{
	Id:             "id",
	VendorName:     "vendor_name",
	AllColumn:      "all_column",
	RequiredColumn: "required_column",
	IsDeleted:      "is_deleted",
	CreatedAt:      "create_time",
	UpdatedAt:      "update_time",
	CreateBy:       "create_by",
	UpdateBy:       "update_by",
}

// NewVendorDao creates and returns a new DAO object for table data access.
func NewVendorDao() *VendorDao {
	return &VendorDao{
		group:   "default",
		table:   "hg_fm_vendor",
		columns: vendorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VendorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VendorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VendorDao) Columns() VendorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VendorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VendorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VendorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
