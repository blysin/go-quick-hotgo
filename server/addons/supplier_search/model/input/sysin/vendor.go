// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package sysin

import (
	"context"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VendorUpdateFields 修改供应商检索字段过滤
type VendorUpdateFields struct {
	VendorName     string `json:"vendorName"     dc:"供应商名称"`
	AllColumn      string `json:"allColumn"      dc:"完整字段，多个用英文逗号隔开"`
	RequiredColumn string `json:"requiredColumn" dc:"隐藏字段，多个用英文逗号隔开"`
}

// VendorInsertFields 新增供应商检索字段过滤
type VendorInsertFields struct {
	VendorName     string `json:"vendorName"     dc:"供应商名称"`
	AllColumn      string `json:"allColumn"      dc:"完整字段，多个用英文逗号隔开"`
	RequiredColumn string `json:"requiredColumn" dc:"隐藏字段，多个用英文逗号隔开"`
}

// VendorEditInp 修改/新增供应商检索
type VendorEditInp struct {
	entity.Vendor
}

func (in *VendorEditInp) Filter(ctx context.Context) (err error) {
	// 验证供应商名称
	if err := g.Validator().Rules("required").Data(in.VendorName).Messages("供应商名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证完整字段，多个用英文逗号隔开
	if err := g.Validator().Rules("required").Data(in.AllColumn).Messages("完整字段，多个用英文逗号隔开不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type VendorEditModel struct{}

// VendorDeleteInp 删除供应商检索
type VendorDeleteInp struct {
	Id int64 `json:"id" v:"required#自增ID不能为空" dc:"自增ID"`
}

func (in *VendorDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type VendorDeleteModel struct{}

// VendorViewInp 获取指定供应商检索信息
type VendorViewInp struct {
	Id int `json:"id" v:"required#自增ID不能为空" dc:"自增ID"`
}

func (in *VendorViewInp) Filter(ctx context.Context) (err error) {
	return
}

type VendorViewModel struct {
	entity.Vendor
}

// VendorListInp 获取供应商检索列表
type VendorListInp struct {
	form.PageReq
	Id         int64         `json:"id"        dc:"自增ID"`
	VendorName string        `json:"vendorName" dc:"供应商名称"`
	Status     *int          `json:"status"    dc:"状态：0-新增，-1-删除，2已发布"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *VendorListInp) Filter(ctx context.Context) (err error) {
	return
}

type VendorListModel struct {
	Id         int64       `json:"id"         dc:"自增ID"`
	VendorName string      `json:"vendorName" dc:"供应商名称"`
	Status     int         `json:"status"  dc:"状态：0-新增，-1-删除，2已发布"`
	Currency   string      `json:"currency"  dc:"币种"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
	CreateBy   int64       `json:"createBy"   dc:"创建人"`
	UpdateBy   int64       `json:"updateBy"   dc:"更新人"`
}

// VendorExportModel 导出供应商检索
type VendorExportModel struct {
	Id         int64       `json:"id"         dc:"自增ID"`
	VendorName string      `json:"vendorName" dc:"供应商名称"`
	IsDeleted  int         `json:"isDeleted"  dc:"是否删除，0：未删除，1：已删除"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
	CreateBy   int64       `json:"createBy"   dc:"创建人"`
	UpdateBy   int64       `json:"updateBy"   dc:"更新人"`
}

type VendorDetailListModel struct {
	Id           int    `json:"id"         dc:"自增ID"`
	Brand        string `json:"brand"            dc:"品牌"`
	Barcode      string `json:"barcode"          dc:"条码"`
	EnglishName  string `json:"englishName"      dc:"英文名称"`
	Cost         int64  `json:"cost"             dc:"成本、供货价"`
	SellingPrice int64  `json:"sellingPrice"     dc:"销售价格"`
	Vendor       string `json:"vendor"           dc:"供应商"`
	Status       int    `json:"status"  dc:"状态：0-新增，-1-删除，2已发布"`
}
