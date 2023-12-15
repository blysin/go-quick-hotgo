// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package ven

import (
	"context"
	"hotgo/addons/supplier_search/api/admin/ven"
	"hotgo/addons/supplier_search/service"
)

var (
	Vendor = cVendor{}
)

type cVendor struct{}

// List 查看供应商检索列表
func (c *cVendor) List(ctx context.Context, req *ven.ListReq) (res *ven.ListRes, err error) {
	list, totalCount, err := service.SysVendor().List(ctx, &req.VendorListInp)
	if err != nil {
		return
	}

	res = new(ven.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出供应商检索列表
func (c *cVendor) Export(ctx context.Context, req *ven.ExportReq) (res *ven.ExportRes, err error) {
	err = service.SysVendor().Export(ctx, &req.VendorListInp)
	return
}

// Edit 更新供应商检索
func (c *cVendor) Edit(ctx context.Context, req *ven.EditReq) (res *ven.EditRes, err error) {
	err = service.SysVendor().Edit(ctx, &req.VendorEditInp)
	return
}

// View 获取指定供应商检索信息
func (c *cVendor) View(ctx context.Context, req *ven.ViewReq) (res *ven.ViewRes, err error) {
	data, err := service.SysVendor().View(ctx, &req.VendorViewInp)
	if err != nil {
		return
	}

	res = new(ven.ViewRes)
	res.VendorViewModel = data
	return
}

// Delete 删除供应商检索
func (c *cVendor) Delete(ctx context.Context, req *ven.DeleteReq) (res *ven.DeleteRes, err error) {
	err = service.SysVendor().Delete(ctx, &req.VendorDeleteInp)
	return
}
