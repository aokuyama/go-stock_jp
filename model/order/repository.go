//go:generate mockgen -source=$GOFILE -package=mock -destination=mock/$GOFILE
package order

import "github.com/aokuyama/go-stock_jp/common"

type OrderRepository interface {
	LoadByDate(date *common.Date) ([]*Order, error)
}
