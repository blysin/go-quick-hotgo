// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VendorIndexDao is the data access object for table hg_fm_vendor_index.
type VendorIndexDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns VendorIndexColumns // columns contains all the column names of Table for convenient usage.
}

// VendorIndexColumns defines and stores column names for table hg_fm_vendor_index.
type VendorIndexColumns struct {
	Id             string // 自增ID
	UnitKey        string // 唯一标识，当前已barcode作为标识
	VenDetailId    string // 关联的hg_fm_vendor_detail.id
	Status         string // 状态：0-正常，1-作废
	DataUploadData string // 数据上传日期
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	Currency       string // 币种
}

// vendorIndexColumns holds the columns for table hg_fm_vendor_index.
var vendorIndexColumns = VendorIndexColumns{
	Id:             "id",
	UnitKey:        "unit_key",
	VenDetailId:    "ven_detail_id",
	Status:         "status",
	DataUploadData: "data_upload_data",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Currency:       "currency",
}

// NewVendorIndexDao creates and returns a new DAO object for table data access.
func NewVendorIndexDao() *VendorIndexDao {
	return &VendorIndexDao{
		group:   "default",
		table:   "hg_fm_vendor_index",
		columns: vendorIndexColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VendorIndexDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VendorIndexDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VendorIndexDao) Columns() VendorIndexColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VendorIndexDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VendorIndexDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VendorIndexDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
