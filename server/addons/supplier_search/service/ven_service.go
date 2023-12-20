package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
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
	// 解析excel
	details, head, venFile, err := s.analysisExcel(ctx, req.FileId, req.PresetColumn)
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

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)

		id, e := mod.InsertAndGetId(&ven)
		if e != nil {
			return e
		}
		ven.Id = id

		rate, e := util.FromExchangeRate(ctx, req.Exchange)
		if e != nil {
			return gerror.Wrap(e, "获取汇率失败")
		}

		glog.Info(ctx, "当前汇率：", gjson.MustEncodeString(rate))

		for i := range *details {
			// go的遍历机制，如果直接用for中拿去detail对象，拿到的是一个变量副本，对其修改不会影响到原对象，只能已这种方式，用数据+索引的方式去修改原对象
			(*details)[i].VendorId = id
			(*details)[i].Status = 0
			(*details)[i].CreateBy = userId
			(*details)[i].UpdateBy = userId

			//查询汇率
			(*details)[i].ExchangeRate = *rate.PriceDig
			(*details)[i].ExchangeRateTime = rate.Time
			(*details)[i].Currency = rate.From

			// 计算CNY价格
			(*details)[i].CostCny = int64(float64((*details)[i].Cost) * *rate.PriceDig)
			(*details)[i].SellingPriceCny = int64(float64((*details)[i].SellingPrice) * *rate.PriceDig)
		}

		// 保存明细表
		e = VenDetailService.SaveBatch(ctx, details)
		if e != nil {
			return e
		}

		// 更新file表数据
		venFile.VendorId = id
		e = VenFileService.Update(ctx, venFile)

		return e
	})

	glog.Info(ctx, "保存成功")

	return
}

func (s *SVen) analysisExcel(ctx context.Context, fileId int64, presetColumn *venin.PresetColumn) (details *[]entity.VendorDetail, head *[]string, venFileEntity *entity.VendorUploadFile, err error) {
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

	venFileEntity = &venFile
	details, head, err = analysisExcelDetail(path, presetColumn)

	return
}
