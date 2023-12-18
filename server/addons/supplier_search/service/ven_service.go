package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/logic/util"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	intService "hotgo/internal/service"
)

type SVen struct{}

func NewVenService() *SVen {
	return &SVen{}
}

var VenService = NewVenService()

// Model 充值订单ORM模型
func (s *SVen) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Vendor.Ctx(ctx), option...)
}

func (s *SVen) Save(ctx context.Context, req *venin.VenSaveInp) (ven *entity.Vendor, err error) {
	mod := s.Model(ctx)

	// 解析excel
	details, head, err := s.analysisExcel(ctx, req.FileId, req.PresetColumn)
	if err != nil {
		return
	}

	userId := contexts.GetUserId(ctx)

	// 保存主表
	ven = &entity.Vendor{
		Id:             0,
		VendorName:     req.VendorName,
		AllColumn:      gjson.MustEncodeString(head),
		RequiredColumn: gjson.MustEncodeString(req.PresetColumn),
		Status:         0,
		CreateBy:       userId,
		UpdateBy:       userId,
	}
	id, err := mod.InsertAndGetId(&ven)
	if err != nil {
		return
	}
	ven.Id = id

	rate, err := util.FromExchangeRate(ctx, req.Exchange)
	if err != nil {
		return nil, gerror.Wrap(err, "获取汇率失败")
	}

	for _, detail := range *details {
		detail.VendorId = id
		detail.Status = 0
		detail.CreateBy = userId
		detail.UpdateBy = userId

		//查询汇率
		detail.Exchange = req.Exchange

	}

	// 保存明细表
	err = VenDetailService.SaveBatch(ctx, details)
	if err != nil {
		return
	}

	return
}

func (s *SVen) analysisExcel(ctx context.Context, fileId int64, presetColumn *venin.PresetColumn) (details *[]entity.VendorDetail, head *[]string, err error) {
	// 获取上传的文件
	venFile, err := VenFileService.GetById(ctx, fileId)
	if err != nil {
		return
	}

	fid := venFile.FileId
	res, err := intService.CommonUpload().GetFile(ctx, fid)
	if err != nil {
		return
	}
	path, err := getFilePath(ctx, res.Path)
	if err != nil {
		return
	}

	details, head, err = analysisExcelDetail(path, presetColumn)

	return
}
