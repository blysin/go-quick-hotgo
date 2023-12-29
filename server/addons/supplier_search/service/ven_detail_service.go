package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/supplier_search/api/admin/ven"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/addons/supplier_search/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"
)

type SVenDetail struct{}

func NewVenDetailService() *SVenDetail {
	return &SVenDetail{}
}

var VenDetailService = NewVenDetailService()

// Model 充值订单ORM模型
func (s *SVenDetail) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.VendorDetail.Ctx(ctx), option...)
}

func (s *SVenDetail) PageByVenId(ctx context.Context, venId int64, page, pageSize int) (list []*entity.VendorDetail, total int, err error) {
	mod := s.Model(ctx)
	mod = mod.Where("vendor_id = ?", venId)
	err = mod.Page(page, pageSize).Scan(&list, &total)
	return
}

func (s *SVenDetail) ListByVenId(ctx context.Context, venId int64) (list []*entity.VendorDetail, err error) {
	mod := s.Model(ctx)
	err = mod.Where("vendor_id = ?", venId).Scan(&list)
	return
}

func (s *SVenDetail) SaveBatch(ctx context.Context, list []*entity.VendorDetail) (err error) {
	if len(list) == 0 {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)
		_, e := mod.Insert(list)
		return e
	})
}

func (s *SVenDetail) UpdateBatch(ctx context.Context, list []*entity.VendorDetail) error {
	if len(list) == 0 {
		return nil
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)

		for _, v := range list {
			_, e := mod.Data(v).Where("id = ?", v.Id).Update()
			if e != nil {
				return e
			}
		}
		return nil
	})
}

func (s *SVenDetail) DeleteByVenId(ctx context.Context, venId int64) (err error) {
	mod := s.Model(ctx)
	_, err = mod.Where("vendor_id = ?", venId).Delete()
	return
}

func (s *SVenDetail) GetById(ctx context.Context, id int64) (detail *entity.VendorDetail, err error) {
	mod := s.Model(ctx)
	err = mod.Where("id = ?", id).Scan(&detail)
	return
}

func (s *SVenDetail) ListByBrandNames(ctx context.Context, id int64, names *[]string) (list []*entity.VendorDetail, err error) {
	//去掉names中的空字符串
	for i := 0; i < len(*names); i++ {
		if (*names)[i] == "" {
			*names = append((*names)[:i], (*names)[i+1:]...)
			i--
		}
	}

	if len(*names) == 0 {
		return
	}

	mod := s.Model(ctx)
	err = mod.Where("vendor_id = ?", id).Where("brand in (?)", names).Scan(&list)
	return
}

func (s *SVenDetail) List(ctx context.Context, req *ven.PageDetailReq) (list []*sysin.VendorDetailListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	if req.VendorId != 0 {
		mod = mod.Where("vendor_id = ?", req.VendorId)
	}

	if req.Brand != "" {
		mod = mod.Where("brand like ?", "%"+req.Brand+"%")
	}

	if req.Vendor != "" {
		mod = mod.Where("vendor like ?", "%"+req.Vendor+"%")
	}

	if req.Barcode != "" {
		mod = mod.Where("barcode like ?", "%"+req.Barcode+"%")
	}

	if req.Status != -99 {
		mod = mod.Where("status = ?", req.Status)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取供应商检索数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(req.Page, req.PerPage).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取供应商检索列表失败，请稍后重试！")
		return
	}

	return
}

// ChangeStatus 修改供应商检索状态
func (s *SVenDetail) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	if (NORMAL != status) && (DELETE != status) && (PUBLISHED != status) {
		return gerror.New("状态值不正确")
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id = ?", id).Update()
		if err != nil {
			return err
		}

		// 获取detail表数据
		detail, err := s.GetById(ctx, id)
		if err != nil {
			return err
		}
		// 更新索引
		details := []*entity.VendorDetail{detail}

		err = VenIndexService.UpdateIndex(ctx, details)

		return err
	})

	return
}

// ChangeStatusByVenId 修改供应商检索状态
func (s *SVenDetail) ChangeStatusByVenId(ctx context.Context, venId int64, status int) (err error) {
	if (NORMAL != status) && (DELETE != status) && (PUBLISHED != status) {
		return gerror.New("状态值不正确")
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("vendor_id = ?", venId).Update()

		if err != nil {
			return err
		}

		var details []*entity.VendorDetail
		// 获取detail表数据
		err = s.Model(ctx).Where("vendor_id = ?", venId).Fields("barcode", "id").Scan(&details)
		if err != nil {
			return err
		}

		err = VenIndexService.UpdateIndex(ctx, details)

		return err
	})

	return
}
