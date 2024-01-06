// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package api

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/addons/supplier_search/service"
)

var (
	VendorMobile = cVendorMobile{}
)

type cVendorMobile struct{}

type PageDetailReq struct {
	g.Meta `path:"/index/list" method:"get" tags:"供应商检索" summary:"获取供应商检索列表"`
	venin.VenPageApiInp
}

type PageDetailRes struct {
	venin.PageApiModel
}

// List 查看供应商检索列表
func (c *cVendorMobile) List(ctx context.Context, req *PageDetailReq) (res *PageDetailRes, err error) {
	s := service.VenIndexService
	list, err := s.Page(ctx, req.VenPageApiInp)
	if err != nil {
		return
	}

	res = new(PageDetailRes)
	res.PageApiModel = *list
	return
}
