// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/addons/supplier_search/dao/internal"
)

// internalVendorDetailDao is internal type for wrapping internal DAO implements.
type internalVendorDetailDao = *internal.VendorDetailDao

// vendorDetailDao is the data access object for table hg_fm_vendor_detail.
// You can define custom methods on it to extend its functionality as you wish.
type vendorDetailDao struct {
	internalVendorDetailDao
}

var (
	// VendorDetail is globally public accessible object for table hg_fm_vendor_detail operations.
	VendorDetail = vendorDetailDao{
		internal.NewVendorDetailDao(),
	}
)

// Fill with you ideas below.
