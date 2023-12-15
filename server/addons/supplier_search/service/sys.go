// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/supplier_search/model"
	"hotgo/addons/supplier_search/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysConfig interface {
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetConfigByGroup 获取指定分组配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
	ISysVendor interface {
		// Model 供应商检索ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取供应商检索列表
		List(ctx context.Context, in *sysin.VendorListInp) (list []*sysin.VendorListModel, totalCount int, err error)
		// Export 导出供应商检索
		Export(ctx context.Context, in *sysin.VendorListInp) (err error)
		// Edit 修改/新增供应商检索
		Edit(ctx context.Context, in *sysin.VendorEditInp) (err error)
		// Delete 删除供应商检索
		Delete(ctx context.Context, in *sysin.VendorDeleteInp) (err error)
		// View 获取供应商检索指定信息
		View(ctx context.Context, in *sysin.VendorViewInp) (res *sysin.VendorViewModel, err error)
	}
)

var (
	localSysConfig ISysConfig
	localSysIndex  ISysIndex
	localSysVendor ISysVendor
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysIndex() ISysIndex {
	if localSysIndex == nil {
		panic("implement not found for interface ISysIndex, forgot register?")
	}
	return localSysIndex
}

func RegisterSysIndex(i ISysIndex) {
	localSysIndex = i
}

func SysVendor() ISysVendor {
	if localSysVendor == nil {
		panic("implement not found for interface ISysVendor, forgot register?")
	}
	return localSysVendor
}

func RegisterSysVendor(i ISysVendor) {
	localSysVendor = i
}
