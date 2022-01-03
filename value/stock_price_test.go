package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledStockPrice(t *testing.T) {
	var q *StockPrice
	var err error
	q, err = NewStockPrice(100)
	assert.Equal(t, 100.0, q.Float(), "有効な株価")
	assert.NoError(t, err)
	q, err = NewStockPrice(100.1)
	assert.Equal(t, 100.1, q.Float(), "有効な株価")
	assert.NoError(t, err)
	q, err = NewStockPrice(9999999.9)
	assert.Equal(t, 9999999.9, q.Float(), "有効な株価")
	assert.NoError(t, err)
}
func TestDisabledStockPrice(t *testing.T) {
	var err error
	_, err = NewStockPrice(0)
	assert.Error(t, err, "0はエラー")
	_, err = NewStockPrice(-1)
	assert.Error(t, err, "負の数はエラー")
	_, err = NewStockPrice(100.12)
	assert.Error(t, err, "小数点は1桁まで")
}
