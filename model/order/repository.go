//go:generate mockgen -source=$GOFILE -package=mock -destination=mock/$GOFILE
package order

import "github.com/aokuyama/go-stock_jp/model/common"

type OrderRepository interface {
	LoadByDate(date *common.Date) (*Collection, error)
	Insert(o *Order) error
	Update(after *Order, before *Order) error
}
