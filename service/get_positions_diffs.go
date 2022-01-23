package service

import (
	"github.com/aokuyama/go-stock_jp/model/position"
)

func GetPositionsDiffs(source *position.Positions, dest *position.Positions) ([2]*position.Positions, error) {
	var result [2]*position.Positions
	result[0] = DiffPositions(source, dest)
	result[1] = DiffPositions(dest, source)
	return result, nil
}

func DiffPositions(source *position.Positions, dest *position.Positions) *position.Positions {
	diffs := position.NewPositions()
	for _, psrc := range *source {
		if !isIn(psrc, dest) {
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
