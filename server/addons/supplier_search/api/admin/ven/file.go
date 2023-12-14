package ven

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/supplier_search/model/input/venin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/ven-file/list" method:"get" tags:"供应商" summary:"获取上传的文件列表"`
	venin.VenFileListInp
}

type ListRes *venin.VenFileListModel
