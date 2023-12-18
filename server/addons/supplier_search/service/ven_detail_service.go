package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/model/entity"
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

func (s *SVenDetail) PageByVenId(ctx context.Context, venId int, page, pageSize int) (list []*entity.VendorDetail, total int, err error) {
	mod := s.Model(ctx)
	mod = mod.Where("vendor_id = ?", venId)
	err = mod.Page(page, pageSize).Scan(&list, &total)
	return
}

func (s *SVenDetail) SaveBatch(ctx context.Context, list *[]entity.VendorDetail) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)
		_, e := mod.Insert(list)
		return e
	})
}

func (s *SVenDetail) DeleteByVenId(ctx context.Context, venId int) (err error) {
	mod := s.Model(ctx)
	_, err = mod.Where("vendor_id = ?", venId).Delete()
	return
}

func (s *SVenDetail) GetById(ctx context.Context, id int) (detail *entity.VendorDetail, err error) {
	mod := s.Model(ctx)
	err = mod.Where("id = ?", id).Scan(&detail)
	return
}
