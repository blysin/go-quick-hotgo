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

// Delete 删除供应商检索
func (c *cVendor) Delete(ctx context.Context, req *ven.DeleteReq) (res *ven.DeleteRes, err error) {
	err = service.SysVendor().Delete(ctx, &req.VendorDeleteInp)
	return
}

// Save 更新供应商检索
func (c *cVendor) Save(ctx context.Context, req *ven.SaveReq) (res *ven.SaveRes, err error) {
	data, err := service.VenService.Save(ctx, &req.VenSaveInp)
	return &ven.SaveRes{Id: data.Id}, err
}
