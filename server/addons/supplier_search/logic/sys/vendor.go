// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package sys

import (
	"context"
	"fmt"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/model/input/sysin"
	"hotgo/addons/supplier_search/service"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysVendor struct{}

func NewSysVendor() *sSysVendor {
	return &sSysVendor{}
}

func init() {
	service.RegisterSysVendor(NewSysVendor())
}

// Model 供应商检索ORM模型
func (s *sSysVendor) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Vendor.Ctx(ctx), option...)
}

// List 获取供应商检索列表
func (s *sSysVendor) List(ctx context.Context, in *sysin.VendorListInp) (list []*sysin.VendorListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自增ID
	if in.Id > 0 {
		mod = mod.Where(dao.Vendor.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Vendor.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 关联表fmVendorUploadFile
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.Vendor.Table(), dao.Vendor.Columns().Id, // 主表表名,关联字段
		dao.VendorUploadFile.Table(), "fmVendorUploadFile", dao.VendorUploadFile.Columns().VendorId, // 关联表表名,别名,关联字段
	)...)

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取供应商检索数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	// 关联表select
	fields, err := hgorm.GenJoinSelect(ctx, sysin.VendorListModel{}, &dao.Vendor, []*hgorm.Join{
		{Dao: &dao.VendorUploadFile, Alias: "fmVendorUploadFile"},
	})

	if err != nil {
		err = gerror.Wrap(err, "获取供应商检索关联字段失败，请稍后重试！")
		return
	}

	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.Vendor.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取供应商检索列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出供应商检索
func (s *sSysVendor) Export(ctx context.Context, in *sysin.VendorListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.VendorExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出供应商检索-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.VendorExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增供应商检索
func (s *sSysVendor) Edit(ctx context.Context, in *sysin.VendorEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.VendorUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改供应商检索失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.VendorInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增供应商检索失败，请稍后重试！")
	}
	return
}

// Delete 删除供应商检索
func (s *sSysVendor) Delete(ctx context.Context, in *sysin.VendorDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除供应商检索失败，请稍后重试！")
		return
	}
	return
}

// View 获取供应商检索指定信息
func (s *sSysVendor) View(ctx context.Context, in *sysin.VendorViewInp) (res *sysin.VendorViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取供应商检索信息，请稍后重试！")
		return
	}
	return
}
