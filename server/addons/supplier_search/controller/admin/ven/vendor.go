// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package ven

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
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
	if err != nil {
		return
	}
	return &ven.SaveRes{
		Vendor: data,
	}, nil
}

// ListDetail 获取明细表分页数据
func (c *cVendor) ListDetail(ctx context.Context, req *ven.PageDetailReq) (res *ven.PageDetailRes, err error) {
	s := service.VenDetailService
	list, totalCount, err := s.List(ctx, req)
	if err != nil {
		return
	}

	res = new(ven.PageDetailRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// ChangeStatus 获取明细表分页数据
func (c *cVendor) ChangeStatus(ctx context.Context, req *ven.ChangeStatusReq) (res *ven.ChangeStatusRes, err error) {
	if req.DetailId != nil && *req.DetailId != gconv.Int64(0) {
		err = service.VenDetailService.ChangeStatus(ctx, *req.DetailId, req.Status)
	} else {
		err = service.VenService.ChangeStatus(ctx, req.VendorId, req.Status)
	}

	return
}
