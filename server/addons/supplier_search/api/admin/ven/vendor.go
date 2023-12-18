// Package vendor
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package ven

import (
	"hotgo/addons/supplier_search/model/input/sysin"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询供应商检索列表
type ListReq struct {
	g.Meta `path:"/vendor/list" method:"get" tags:"供应商检索" summary:"获取供应商检索列表"`
	sysin.VendorListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.VendorListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出供应商检索列表
type ExportReq struct {
	g.Meta `path:"/vendor/export" method:"get" tags:"供应商检索" summary:"导出供应商检索列表"`
	sysin.VendorListInp
}

type ExportRes struct{}

// ViewReq 获取供应商检索指定信息
type ViewReq struct {
	g.Meta `path:"/vendor/view" method:"get" tags:"供应商检索" summary:"获取供应商检索指定信息"`
	sysin.VendorViewInp
}

type ViewRes struct {
	*sysin.VendorViewModel
}

// EditReq 修改/新增供应商检索
type EditReq struct {
	g.Meta `path:"/vendor/edit" method:"post" tags:"供应商检索" summary:"修改/新增供应商检索"`
	sysin.VendorEditInp
}

type EditRes struct{}

// DeleteReq 删除供应商检索
type DeleteReq struct {
	g.Meta `path:"/vendor/delete" method:"post" tags:"供应商检索" summary:"删除供应商检索"`
	sysin.VendorDeleteInp
}

type DeleteRes struct{}

// EditReq 修改/新增供应商检索
type SaveReq struct {
	g.Meta `path:"/vendor/save" method:"post" tags:"供应商检索" summary:"修改/新增供应商检索"`
	venin.VenSaveInp
}

type SaveRes struct {
	Id int64 `json:"id"`
}
