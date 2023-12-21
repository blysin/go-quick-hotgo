package ven

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"hotgo/addons/supplier_search/api/admin/ven"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/addons/supplier_search/service"
	"hotgo/internal/model/input/sysin"
	intService "hotgo/internal/service"
)

var ControllerVen = new(cVen)

type cVen struct{}

func (c *cVen) List(ctx context.Context, req *ven.FileListReq) (res *ven.FileListRes, err error) {
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
	listRes := ven.FileListRes(&model)

	return &listRes, nil
}

func (c *cVen) Upload(ctx context.Context, req *ven.UploadReq) (res *ven.UploadRes, err error) {
	glog.Info(ctx, "收到请求，上传excel文件")

	if req.Id == nil || *req.Id <= 0 {
		// 新文件上传
		fileName, attr, err := saveFile(ctx)
		if err != nil {
			return nil, err
		}

		// 保存文件信息到数据库
		res, err = service.VenFileService.UploadFile(ctx, fileName, attr)
		if err != nil {
			glog.Error(ctx, "文件信息保存失败", err)
			return nil, gerror.New("文件信息保存失败")
		}
	} else {
		//更新文件
		// 获取主表数据
		vendor := service.VenService.Get(ctx, *req.Id)
		if vendor == nil {
			return nil, gerror.New("数据异常，没有找到供应商信息")
		}

		_, attr, err := saveFile(ctx)
		if err != nil {
			return nil, err
		}

		// 保存文件信息到数据库
		_, err = service.VenFileService.ReUploadFile(ctx, vendor, attr)
		if err != nil {
			glog.Error(ctx, "文件信息保存失败", err)
			return nil, err
		}
		res = &ven.UploadRes{}
	}

	return
}

func saveFile(ctx context.Context) (fileName string, attr *sysin.AttachmentListModel, err error) {
	r := g.RequestFromCtx(ctx)

	file := r.GetUploadFile("file")

	if file == nil {
		err = gerror.New("没有找到上传的文件")
		return
	}

	// 获取文件名
	fileName = file.Filename

	// 如果不是.xlsx文件，返回错误
	if fileName[len(fileName)-5:] != ".xlsx" {
		err = gerror.New("请上传.xlsx格式的文件")
		return
	}

	uploadType := "doc"
	attr, err = intService.CommonUpload().UploadFile(ctx, uploadType, file)
	if err != nil {
		glog.Error(ctx, "文件存储失败", err)
		return "", nil, gerror.New("文件存储失败")
	}
	return
}

func (c *cVen) ReUpload(ctx context.Context, req *ven.ReUploadReq) (res *ven.ReUploadRes, err error) {
	// 获取主表数据
	vendor := service.VenService.Get(ctx, req.Id)
	if vendor == nil {
		return nil, gerror.New("数据异常，没有找到供应商信息")
	}

	_, attr, err := saveFile(ctx)
	if err != nil {
		return nil, err
	}

	// 保存文件信息到数据库
	res, err = service.VenFileService.ReUploadFile(ctx, vendor, attr)
	if err != nil {
		glog.Error(ctx, "文件信息保存失败", err)
		return nil, err
	}
	return
}

func (c *cVen) Download(ctx context.Context, req *ven.DownloadReq) (res *ven.DownloadRes, err error) {
	attr, _, err := service.VenFileService.GetAttrByFileId(ctx, req.Id)
	if err != nil {
		return
	}

	return &ven.DownloadRes{
		Url: attr.FileUrl,
	}, nil
}
