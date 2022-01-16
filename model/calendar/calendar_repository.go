//go:generate mockgen -source=$GOFILE -package=mock -destination=mock/$GOFILE
package calendar

import "github.com/aokuyama/go-stock_jp/common"

type CalendarRepository interface {
	LoadByDateRange(date_range *common.DateRange) (*Calendar, error)
}
