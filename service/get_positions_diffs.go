package service

import (
	"github.com/aokuyama/go-stock_jp/model/position"
)

func GetPositionsDiffs(source *position.Positions, dest *position.Positions) (*position.Positions, error) {
	ps := diff(source, dest)
	return ps, nil
}

func diff(source *position.Positions, dest *position.Positions) *position.Positions {
	copy_src := source.Copy()
	copy_dest := dest.Copy()
	for _, dest := range *copy_dest {
		for _, src := range *copy_src {
			src.OffsetIfEqualPosition(dest)
		}
	}
	for _, dest := range *copy_dest {
		if !dest.IsCompleted() {
			dest.Quantity *= -1
			*copy_src = append(*copy_src, dest)
		}
	}
	return copy_src.Compress()
}
