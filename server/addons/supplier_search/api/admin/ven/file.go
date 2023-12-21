package ven

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/supplier_search/model/input/venin"
)

// FileListReq 查询列表
type FileListReq struct {
	g.Meta `path:"/ven-file/list" method:"get" tags:"供应商" summary:"获取上传的文件列表"`
	// 嵌入 VenFileListInp 结构体，类似Java中的继承
	venin.VenFileListInp
}

// ListRes 查询列表返回，这里使用了指针，相当于给 VenFileListModel 结构体起了一个别名，
// 要实例化这个结构体，需要使用 ListRes(&venin.VenFileListModel{}) 的方式
type FileListRes *venin.VenFileListModel

type UploadReq struct {
	g.Meta `path:"/ven-file/upload" method:"post" tags:"供应商" summary:"上传excel文件"`
	Id     *int64 `json:"id"`
}

type UploadRes struct {
	venin.PresetColumn
	Id         int64    `json:"id"`
	FileName   string   `json:"file_name"`
	AllColumns []string `json:"all_columns"`
}

type ReUploadReq struct {
	g.Meta `path:"/ven-file/re-upload" method:"post" tags:"供应商" summary:"更新数据，重新上传excel文件"`
	Id     int64 `json:"id"`
}

type ReUploadRes struct {
}

type DownloadReq struct {
	g.Meta `path:"/ven-file/download" method:"GET" tags:"供应商" summary:"下载文件"`
	Id     int64 `json:"id"`
}

type DownloadRes struct {
	Url string `json:"url"`
}
