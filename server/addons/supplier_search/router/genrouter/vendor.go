// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.11.5
package genrouter

import (
	"hotgo/addons/supplier_search/controller/admin/ven"
)

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, ven.Vendor) // 供应商检索
}
