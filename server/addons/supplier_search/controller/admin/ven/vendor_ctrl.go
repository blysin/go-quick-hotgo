// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package ven

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/addons/supplier_search/api/admin/ven"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/addons/supplier_search/model/input/venin"
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
	err = service.VenService.ChangeStatus(ctx, req.Id, service.DELETE)
	return
}

// Save 更新供应商检索
func (c *cVendor) Save(ctx context.Context, req *ven.SaveReq) (res *ven.SaveRes, err error) {
	var data *entity.Vendor
	if req.VenSaveInp.Id != gconv.Int64(0) {
		data, err = service.VenService.Update(ctx, &req.VenSaveInp)
	} else {
		if req.FileId == 0 {
			err = gerror.New("请先上传文件")
			return
		}
		data, err = service.VenService.Save(ctx, &req.VenSaveInp)
	}

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

// ChangeStatus 修改状态
func (c *cVendor) ChangeStatus(ctx context.Context, req *ven.ChangeStatusReq) (res *ven.ChangeStatusRes, err error) {
	if req.DetailId != nil && *req.DetailId != gconv.Int64(0) {
		err = service.VenDetailService.ChangeStatus(ctx, *req.DetailId, req.Status)
	} else {
		err = service.VenService.ChangeStatus(ctx, req.VendorId, req.Status)
	}

	return
}

// GetVendor 获取明细表分页数据
func (c *cVendor) GetVendor(ctx context.Context, req *ven.VenViewReq) (res *ven.VenViewRes, err error) {
	vendor := service.VenService.Get(ctx, req.Id)
	if vendor == nil {
		err = gerror.New("数据异常，没有找到供应商信息")
		return
	}
	list, err := service.VenFileService.List(ctx, venin.VenFileListInp{
		VendorId: vendor.Id,
	})
	if err != nil {
		return
	}
	res = &ven.VenViewRes{
		Vendor: vendor,
		Files:  list,
	}

	return
}
