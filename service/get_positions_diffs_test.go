package service_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/service"

	"github.com/aokuyama/go-stock_jp/model/position"
	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	t.Run("no diff", func(t *testing.T) {
		ps1 := position.NewPositions()
		p1, _ := position.NewPosition("spot_buy", "1234", 500)
		p2, _ := position.NewPosition("margin_buy", "1234", 400)
		p3, _ := position.NewPosition("margin_sell", "1234", 200)
		p4, _ := position.NewPosition("margin_sell", "1234", 100)
		p5, _ := position.NewPosition("margin_buy", "1235", 100)
		*ps1 = append(*ps1, p1, p2, p3, p4, p5)
		ps2 := position.NewPositions()
		*ps2 = append(*ps2, p5, p4, p3, p2, p1)
		diff1 := DiffPositions(ps1, ps2)
		assert.Equal(t, 0, len(*diff1), diff1.String())
		diff2 := DiffPositions(ps1, ps2)
		assert.Equal(t, 0, len(*diff2), diff2.String())
	})

	t.Run("case 1", func(t *testing.T) {
		ps1 := position.NewPositions()
		p1, _ := position.NewPosition("spot_buy", "1234", 500)
		p2, _ := position.NewPosition("margin_buy", "1234", 400)
		p3, _ := position.NewPosition("margin_sell", "1234", 200)
		p4, _ := position.NewPosition("margin_sell", "1234", 100)
		p5, _ := position.NewPosition("margin_buy", "1235", 100)
		ps1.AddPosition(p1)
		ps1.AddPosition(p2)
		ps1.AddPosition(p3)
		ps1.AddPosition(p4)
		ps1.AddPosition(p5)
		ps2 := position.NewPositions()
		ps2.AddPosition(p1)
		ps2.AddPosition(p2)
		ps2.AddPosition(p3)
		ps2.AddPosition(p5)
		diff1 := DiffPositions(ps1, ps2)
		assert.Equal(t, 1, len(*diff1))
		assert.Equal(t, 300, (*diff1)[0].Quantity.Int())
		diff2 := DiffPositions(ps2, ps1)
		assert.Equal(t, 1, len(*diff2))
		assert.Equal(t, 200, (*diff2)[0].Quantity.Int())
	})
	t.Run("case 2", func(t *testing.T) {
		ps1 := position.NewPositions()
		p1, _ := position.NewPosition("spot_buy", "1234", 500)
		p2, _ := position.NewPosition("margin_buy", "1234", 400)
		p3, _ := position.NewPosition("margin_sell", "1234", 300)
		p4, _ := position.NewPosition("margin_buy", "1235", 200)
		ps1.AddPosition(p1)
		ps1.AddPosition(p2)
		ps1.AddPosition(p3)
		ps2 := position.NewPositions()
		ps2.AddPosition(p2)
		ps2.AddPosition(p3)
		ps2.AddPosition(p4)
		assert.Equal(t, 1200, ps1.SumQuantity())
		assert.Equal(t, 900, ps2.SumQuantity())
		diff1 := DiffPositions(ps1, ps2)
		assert.Equal(t, 1, len(*diff1))
		assert.Equal(t, 500, (*diff1)[0].Quantity.Int())
		diff2 := DiffPositions(ps2, ps1)
		assert.Equal(t, 1, len(*diff2))
		assert.Equal(t, 200, (*diff2)[0].Quantity.Int())
		assert.Equal(t, 3, len(*ps1))
		assert.Equal(t, 3, len(*ps2))
		assert.Equal(t, 1200, ps1.SumQuantity())
		assert.Equal(t, 900, ps2.SumQuantity())
	})
}
