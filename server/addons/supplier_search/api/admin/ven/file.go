package ven

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/supplier_search/model/input/venin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/ven-file/list" method:"get" tags:"供应商" summary:"获取上传的文件列表"`
	// 嵌入 VenFileListInp 结构体，类似Java中的继承
	venin.VenFileListInp
}

// ListRes 查询列表返回，这里使用了指针，相当于给 VenFileListModel 结构体起了一个别名，
// 要实例化这个结构体，需要使用 ListRes(&venin.VenFileListModel{}) 的方式
type ListRes *venin.VenFileListModel

type UploadReq struct {
	g.Meta `path:"/ven-file/upload" method:"post" tags:"供应商" summary:"上传excel文件"`
}

type UploadRes struct {
	Id int64 `json:"id"`
}
