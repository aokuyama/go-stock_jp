package test

import "github.com/aokuyama/go-stock_jp/model/order"

func GetOrder() *order.Order {
	o, err := order.New(0, "3662", "jpx", "margin_buy", "system", "normal", 0, "more", 10000, 200, "2022-01-24", "afternoon", "not_ordered", false)
	if err != nil {
		panic(err)
	}
	return o
}
