package service

import (
	"github.com/aokuyama/go-stock_jp/model/position"
)

func GetPositionsDiffs(source *position.Positions, dest *position.Positions) ([2]*position.Positions, error) {
	var result [2]*position.Positions
	result[0] = diff(source, dest)
	result[1] = diff(dest, source)
	return result, nil
}

func diff(source *position.Positions, dest *position.Positions) *position.Positions {
	srcs := source.Compress()
	dsts := dest.Compress()
	diffs := position.NewPositions()
	for _, psrc := range *srcs {
		if !isIn(psrc, dsts) {
			*diffs = append(*diffs, psrc)
		}
	}
	return diffs
}

func isIn(pos *position.Position, positions *position.Positions) bool {
	for _, dpos := range *positions {
		if pos.IsEqual(dpos) {
			return true
		}
	}
	return false
}
