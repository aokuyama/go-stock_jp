package value

import (
	"errors"
)

type Market string

func NewMarket(v string) (*Market, error) {
	codes := GetMarketCodes()
	for _, c := range codes {
		if v == c {
			m := Market(v)
			return &m, nil
		}
	}
	return nil, errors.New("無効な市場コード")
}

func GetMarketCodes() [6]string {
	return [...]string{
		"jpx",
		"ose",
		"jasdaq",
		"nse",
		"sse",
		"fse",
	}
}
