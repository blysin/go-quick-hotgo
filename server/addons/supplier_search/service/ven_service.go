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
)

type SVen struct{}

func NewVenService() *SVen {
	return &SVen{}
}

var VenService = NewVenService()

const (
	NORMAL    = 0
	DELETE    = -1
	PUBLISHED = 2
)

// Model 充值订单ORM模型
func (s *SVen) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Vendor.Ctx(ctx), option...)
}

func (s *SVen) Update(ctx context.Context, v *venin.VenSaveInp) (ven *entity.Vendor, err error) {
	//更新主表，
	ven = s.Get(ctx, v.Id)
	if ven == nil {
		return nil, gerror.New("数据异常，没有查到相关数据")
	}

	ven.VendorName = v.VendorName
	ven.Currency = v.Exchange
	ven.RequiredColumn = gjson.MustEncodeString(v.PresetColumn)

	userId := contexts.GetUserId(ctx)

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)

		_, err = mod.Data(g.Map{"vendor_name": ven.VendorName, "currency": ven.Currency, "required_column": ven.RequiredColumn}).Where("id = ?", v.Id).Update()
		if err != nil {
			return err
		}

		// 更新明细表
		details, err := VenDetailService.ListByVenId(ctx, v.Id)
		if err != nil {
			return err
		}

		if len(details) == 0 {
			return gerror.New("没有查到明细数据")
		}

		//获取汇率
		rate, err := util.FromExchangeRate(ctx, v.Exchange)
		if err != nil {
			return gerror.Wrap(err, "获取汇率失败")
		}
		_, err = setDetailValue(details, v.PresetColumn)
		if err != nil {
			return err
		}

		presetDetails(details, v.Id, userId, rate)

		return VenDetailService.UpdateBatch(ctx, details)
	})
	return
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
		VendorName:     req.VendorName,
		AllColumn:      gjson.MustEncodeString(head),
		RequiredColumn: gjson.MustEncodeString(req.PresetColumn),
		Currency:       req.Exchange,
		Status:         NORMAL,
		CreateBy:       userId,
		UpdateBy:       userId,
	}

	rate, err := util.FromExchangeRate(ctx, req.Exchange)
	if err != nil {
		return nil, gerror.Wrap(err, "获取汇率失败")
	}
	glog.Info(ctx, "当前汇率：", gjson.MustEncodeString(rate))

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)

		id, e := mod.InsertAndGetId(&ven)
		if e != nil {
			return e
		}
		ven.Id = id

		presetDetails(details, id, userId, rate)

		// 保存明细表
		e = VenDetailService.SaveBatch(ctx, details)
		if e != nil {
			return e
		}

		// 更新file表数据
		venFile.VendorId = id
		venFile.ValidNum = len(details)
		e = VenFileService.Update(ctx, venFile)

		return e
	})

	glog.Info(ctx, "保存成功")

	return
}

func presetDetails(details []*entity.VendorDetail, id int64, userId int64, rate *util.ExchangeRate) {
	for i := range details {
		// go的遍历机制，如果直接用for中拿去detail对象，拿到的是一个变量副本，对其修改不会影响到原对象，只能已这种方式，用数据+索引的方式去修改原对象

		if (details)[i].Id == 0 {
			//新增
			(details)[i].VendorId = id
			(details)[i].Status = NORMAL
			(details)[i].CreateBy = userId
		}

		(details)[i].UpdateBy = userId
		//查询汇率
		(details)[i].ExchangeRate = *rate.PriceDig
		(details)[i].ExchangeRateTime = rate.Time
		(details)[i].Currency = rate.From

		// 计算CNY价格
		(details)[i].CostCny = int64(float64((details)[i].Cost) * *rate.PriceDig)
		(details)[i].SellingPriceCny = int64(float64((details)[i].SellingPrice) * *rate.PriceDig)
	}
}

func (s *SVen) analysisExcel(ctx context.Context, fileId int64, presetColumn *venin.PresetColumn) (details []*entity.VendorDetail, head *[]string, venFileEntity *entity.VendorUploadFile, err error) {
	// 获取上传的文件

	res, venFileEntity, err := VenFileService.GetAttrByFileId(ctx, fileId)

	if err != nil {
		return
	}
	path, err := getFilePath(ctx, res.Path)
	if err != nil {
		return
	}

	details, head, _, err = analysisExcelDetail(path, presetColumn)

	return
}

func (s *SVen) Get(ctx context.Context, id int64) *entity.Vendor {
	mod := s.Model(ctx)

	ven := (*entity.Vendor)(nil)
	err := mod.Where("id", id).Scan(&ven)
	if err != nil {
		return nil
	}
	return ven
}

func (s *SVen) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	if (NORMAL != status) && (DELETE != status) && (PUBLISHED != status) {
		return gerror.New("状态值不正确")
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err := VenDetailService.ChangeStatusByVenId(ctx, id, status)
		if err != nil {
			return err
		}

		_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id = ?", id).Update()

		return err
	})
	return
}