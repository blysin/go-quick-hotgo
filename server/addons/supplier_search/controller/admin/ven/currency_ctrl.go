package ven

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	isc "hotgo/internal/service"
)

var (
	Currency = cCurrency{}
)

type cCurrency struct{}

type GetReq struct {
	g.Meta `path:"/currency/get" method:"get" tags:"币种" summary:"获取币种"`
}

type GetRes struct {
	List *[]CurrencyVO `json:"list"`
}

type CurrencyVO struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Get 更新供应商检索
func (c *cCurrency) Get(ctx context.Context, _ *GetReq) (res *GetRes, err error) {
	//调用这个接口可以查询到所有的币种
	//https://www.mxnzp.com/api/exchange_rate/configs?app_id=tloqllfhllnrcjlf&app_secret=DjVf11uDBjwZstIKptq17JNaeJ2eoRe7

	condition := c.newCondition()
	list, err := isc.SysAddonsConfig().GetConfigByCondition(ctx, condition)
	if err != nil {
		return
	}
	res = new(GetRes)

	jsonStr := (*list)[0].Value

	err = gjson.DecodeTo(jsonStr, &res.List)

	return
}

func (c *cCurrency) newCondition() *entity.SysAddonsConfig {
	condition := &entity.SysAddonsConfig{
		AddonName: "supplier_search",
		Group:     "basic",
		Key:       "currency",
	}
	return condition
}

type SaveReq struct {
	g.Meta `path:"/currency/save" method:"post" tags:"币种" summary:"获取币种"`
	List   *[]CurrencyVO `json:"list"`
}

type SaveRes struct {
}

// Save 更新供应商检索
func (c *cCurrency) Save(ctx context.Context, req *SaveReq) (res *SaveRes, err error) {
	if len(*req.List) == 0 {
		return nil, gerror.New("参数不能为空")
	}
	jsonStr := gjson.MustEncodeString(req.List)

	in := &sysin.UpdateAddonsConfigInp{
		AddonName: "supplier_search",
		Group:     "basic",
		List: g.Map{
			"currency": jsonStr,
		},
	}
	err = isc.SysAddonsConfig().UpdateConfigByGroup(ctx, in)
	if err != nil {
		return
	}

	return
}
