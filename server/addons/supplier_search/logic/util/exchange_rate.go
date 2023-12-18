package util

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type ExchangeRateResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *ExchangeRate
}

type ExchangeRate struct {
	Name       string      `json:"name"`
	From       string      `json:"from"`
	To         string      `json:"to"`
	UpdateTime *string     `json:"updateTime"`
	Time       *gtime.Time `json:"time"`
	Price      string      `json:"price"`
	PriceDig   *float64    `json:"priceDig"`
}

// FromExchangeRate 获取汇率，from 转成 CNY
func FromExchangeRate(ctx context.Context, from string) (*ExchangeRate, error) {
	key := "ex_rate_f:" + from
	ttl := 24 * time.Hour
	v, err := gcache.GetOrSetFunc(ctx, key, func(ctx context.Context) (interface{}, error) {
		// 如果是人民币，直接返回
		if from == "CNY" {
			dig := 1.0
			return &ExchangeRate{
				Name:     "人民币",
				From:     "CNY",
				To:       "CNY",
				Time:     gtime.Now(),
				PriceDig: &dig,
			}, nil
		}

		appId := "tloqllfhllnrcjlf"
		appSecret := "DjVf11uDBjwZstIKptq17JNaeJ2eoRe7"

		url := "https://www.mxnzp.com/api/exchange_rate/aim?from=%s&to=CNY&app_id=%s&app_secret=%s"
		url = fmt.Sprintf(url, from, appId, appSecret)

		resp := &ExchangeRateResp{}
		err := GetObject(ctx, url, resp)
		if err != nil {
			return nil, err
		}

		data := resp.Data

		dig := StringToFloat64(data.Price)
		data.PriceDig = &dig

		ut := data.UpdateTime
		t := formatTime(ut)
		data.UpdateTime = nil
		data.Time = t
		return data, nil
	}, ttl)

	if err != nil {
		return nil, err
	}

	rate := &ExchangeRate{}
	err = v.Struct(rate)

	return rate, err
}

func formatTime(ut *string) *gtime.Time {
	t := gtime.Now()
	if ut != nil {
		arr := gstr.SplitAndTrim(*ut, " ")

		y := "2023"
		mm := "01"
		dd := "01"
		h := "00"
		m := "00"
		s := "00"

		unitDate := gstr.SplitAndTrim(arr[0], "-")
		if len(unitDate) >= 1 {
			//年，如果是9:1，补0
			y = unitDate[0]
		}
		if len(unitDate) >= 2 {
			//月，如果是9:1，补0
			mm = unitDate[1]
			if len(mm) == 1 {
				mm = "0" + mm
			}
		}
		if len(unitDate) >= 3 {
			//日，如果是9:1，补0
			dd = unitDate[2]
			if len(dd) == 1 {
				dd = "0" + dd
			}
		}

		if len(arr) > 1 {
			//解析时间，支持：9:1,9:01,09:01,09:01:00
			unitTime := gstr.SplitAndTrim(arr[1], ":")
			if len(unitTime) >= 1 {
				//小时，如果是9:1，补0
				h = unitTime[0]
				if len(h) == 1 {
					h = "0" + h
				}
			}
			if len(unitTime) >= 2 {
				//分钟，如果是9:1，补0
				m = unitTime[1]
				if len(m) == 1 {
					m = "0" + m
				}
			}
			if len(unitTime) >= 3 {
				//秒，如果是9:1，补0
				s = unitTime[2]
				if len(s) == 1 {
					s = "0" + s
				}
			}
		}
		fullDt := y + "-" + mm + "-" + dd + " " + h + ":" + m + ":" + s
		t, _ = gtime.StrToTime(fullDt)
	}
	return t
}

func StringToFloat64(price string) float64 {
	if price == "" {
		return 0
	}
	return gconv.Float64(price)
}
