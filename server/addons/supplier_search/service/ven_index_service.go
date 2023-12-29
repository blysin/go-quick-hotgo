package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
)

type SVenIndex struct{}

func NewVenIndexService() *SVenIndex {
	return &SVenIndex{}
}

var VenIndexService = NewVenIndexService()

const (
	IndexStatusFail = iota
	IndexStatusOk
)

// Model ORM模型
func (s *SVenIndex) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.VendorIndex.Ctx(ctx), option...)
}

func (s *SVenIndex) AsyncUpdateIndex(details []*entity.VendorDetail) {
	go func() {
		err := s.UpdateIndex(gctx.New(), details)
		if err != nil {
			g.Log().Error(gctx.New(), "异步更新索引失败", details, err)
		}
	}()
}

func (s *SVenIndex) UpdateIndex(ctx context.Context, details []*entity.VendorDetail) (err error) {
	mod := s.Model(ctx)

	// 获取details中的Barcode，存到一个数组中，并且要去重
	set := hashset.New()
	for _, v := range details {
		set.Add(v.Barcode)
	}
	// 从数据库中搜索出来，所有的barcode，并且状态为已发布的，只查询barcode和id字段
	var currents = ([]entity.VendorDetail)(nil)
	where := VenDetailService.Model(ctx).Where("barcode in (?)", set.Values()).Where("status = ?", PUBLISHED)
	err = where.Fields("barcode", "id").Scan(&currents)
	if err != nil {
		g.Log().Error(ctx, "查询明细表失败", err)
		return
	}

	// currents转成map，key为barcode，value为VendorDetail
	currentMap := make(map[string]*entity.VendorDetail)
	for _, v := range currents {
		vCopy := v
		currentMap[v.Barcode] = &vCopy
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 遍历set，如果currentMap中没有，就将index状态改成不可用，如果有就修改关联字段

		var insertList []*entity.VendorIndex
		var updateList []*entity.VendorIndex
		var failStatusUnitKey []string

		for _, v := range set.Values() {
			if _, ok := currentMap[v.(string)]; !ok {
				// 不存在
				failStatusUnitKey = append(failStatusUnitKey, v.(string))
			} else {
				// 存在，先获取到当前的数据
				current := currentMap[v.(string)]
				// 根据barcode查询
				idx := new(entity.VendorIndex)
				err := mod.Where("unit_key = ?", v).Scan(&idx)
				if err != nil && !errors.Is(err, sql.ErrNoRows) {
					g.Log().Error(ctx, "查询索引失败", err)
					continue
				}
				if idx == nil || idx.Id == 0 {
					// 不存在，新增
					insertList = append(insertList, &entity.VendorIndex{
						UnitKey:     current.Barcode,
						VenDetailId: current.Id,
						Status:      IndexStatusOk,
					})
				} else {
					// 存在，修改
					idx.VenDetailId = current.Id
					idx.Status = IndexStatusOk
					updateList = append(updateList, idx)
				}
			}
		}
		if failStatusUnitKey != nil && len(failStatusUnitKey) > 0 {
			_, err = mod.Where("unit_key in (?)", failStatusUnitKey).Update(g.Map{"status": IndexStatusFail})
			if err != nil {
				g.Log().Error(ctx, "更新索引失败", err)
			}
		}
		if insertList != nil && len(insertList) > 0 {
			_, err = mod.Insert(insertList)
			if err != nil {
				g.Log().Error(ctx, "新增索引失败", err)
			}
		}
		if updateList != nil && len(updateList) > 0 {
			err = s.UpdateById(ctx, updateList)
		}
		return err
	})

	return
}

func (s *SVenIndex) UpdateById(ctx context.Context, list []*entity.VendorIndex) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		mod := s.Model(ctx)
		for _, v := range list {
			_, err = mod.Data(v).Where("id = ?", v.Id).Update()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func (s *SVenIndex) Page(ctx context.Context, req venin.VenPageApiInp) (res *venin.PageApiModel, err error) {
	//	TODO 调试
	var (
		list  []*entity.VendorDetail
		total int
	)
	mod := VenDetailService.Model(ctx)
	mod.InnerJoin("vendor_index b", "b.ven_detail_id = a.id")
	if req.VendorName != "" {
		mod = mod.Where("vendor_name like ?", "%"+req.VendorName+"%")
	}
	if req.CostStart > 0 {
		mod = mod.Where("cost >= ?", req.CostStart)
	}
	if req.CostEnd > 0 {
		mod = mod.Where("cost <= ?", req.CostEnd)
	}
	if req.CostCnyStart > 0 {
		mod = mod.Where("cost_cny >= ?", req.CostCnyStart)
	}
	if req.CostCnyEnd > 0 {
		mod = mod.Where("cost_cny <= ?", req.CostCnyEnd)
	}
	if req.OrderByPriceAsc == 1 {
		mod = mod.Order("cost_cny asc")
	} else {
		if req.OrderBy == 1 {
			mod = mod.Order("cost desc")
		}
	}
	err = mod.Page(req.Page, req.PerPage).Scan(&list)
	if err != nil {
		return
	}
	total, err = mod.Count()
	if err != nil {
		return
	}
	res = &venin.PageApiModel{
		PageRes: form.PageRes{
			PageReq:    req.PageReq,
			TotalCount: total,
		},
		List: list,
	}
	return
}
