package ven

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"hotgo/addons/supplier_search/api/admin/ven"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/addons/supplier_search/service"
	intService "hotgo/internal/service"
)

var ControllerVen = new(cVen)

type cVen struct{}

func (c *cVen) List(ctx context.Context, req *ven.ListReq) (res *ven.ListRes, err error) {
	glog.Info(ctx, "收到请求，获取上传的文件列表，参数：", req)

	list, err := service.VenFileService.List(ctx, req.VenFileListInp)
	if err != nil {
		return nil, err
	}

	// 创建一个 VenFileListModel 对象
	model := venin.VenFileListModel{
		List: list,
	}

	// 获取 VenFileListModel 对象的指针，并将其赋值给一个 ListRes 类型的变量
	listRes := ven.ListRes(&model)

	return &listRes, nil
}

func (c *cVen) Upload(ctx context.Context, _ *ven.UploadReq) (res *ven.UploadRes, err error) {
	glog.Info(ctx, "收到请求，上传excel文件")

	r := g.RequestFromCtx(ctx)

	file := r.GetUploadFile("file")

	if file == nil {
		err = gerror.New("没有找到上传的文件")
		return
	}

	// 获取文件名
	fileName := file.Filename

	// 如果不是.xlsx文件，返回错误
	if fileName[len(fileName)-5:] != ".xlsx" {
		err = gerror.New("支持称.xlsx格式的文件")
		return
	}

	uploadType := "doc"
	attr, err := intService.CommonUpload().UploadFile(ctx, uploadType, file)
	if err != nil {
		glog.Error(ctx, "文件存储失败", err)
		return nil, gerror.New("文件存储失败")
	}

	// 保存文件信息到数据库
	venFile, err := service.VenFileService.UploadFile(ctx, fileName, attr)
	if err != nil {
		glog.Error(ctx, "文件信息保存失败", err)
		return nil, gerror.New("文件信息保存失败")
	}
	return &ven.UploadRes{
		Id: venFile.Id,
	}, nil
}
