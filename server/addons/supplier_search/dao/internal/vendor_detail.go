// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VendorDetailDao is the data access object for table hg_fm_vendor_detail.
type VendorDetailDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns VendorDetailColumns // columns contains all the column names of Table for convenient usage.
}

// VendorDetailColumns defines and stores column names for table hg_fm_vendor_detail.
type VendorDetailColumns struct {
	Id               string // 自增ID
	VendorId         string // 供应商主表id
	Brand            string // 品牌
	Barcode          string // 条码
	EnglishName      string // 英文名称
	Cost             string // 成本、供货价
	CostCny          string // 成本、供货价-人民币
	SellingPrice     string // 销售价格
	SellingPriceCny  string // 销售价格-人民币
	Vendor           string // 供应商
	Currency         string // 币种
	ExchangeRate     string // 汇率
	ExchangeRateTime string // 汇率时间
	VendorData       string // 工资内容,json格式存储
	CreateTime       string // 创建时间
	UpdateTime       string // 更新时间
	CreateBy         string // 创建人
	UpdateBy         string // 更新人
}

// vendorDetailColumns holds the columns for table hg_fm_vendor_detail.
var vendorDetailColumns = VendorDetailColumns{
	Id:               "id",
	VendorId:         "vendor_id",
	Brand:            "brand",
	Barcode:          "barcode",
	EnglishName:      "english_name",
	Cost:             "cost",
	CostCny:          "cost_cny",
	SellingPrice:     "selling_price",
	SellingPriceCny:  "selling_price_cny",
	Vendor:           "vendor",
	Currency:         "currency",
	ExchangeRate:     "exchange_rate",
	ExchangeRateTime: "exchange_rate_time",
	VendorData:       "vendor_data",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
}

// NewVendorDetailDao creates and returns a new DAO object for table data access.
func NewVendorDetailDao() *VendorDetailDao {
	return &VendorDetailDao{
		group:   "default",
		table:   "hg_fm_vendor_detail",
		columns: vendorDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VendorDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VendorDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VendorDetailDao) Columns() VendorDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VendorDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VendorDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VendorDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
