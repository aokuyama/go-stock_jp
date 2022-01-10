package position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledQuantity(t *testing.T) {
	var q *Quantity
	q, _ = NewQuantity(100)
	assert.Equal(t, 100, int(*q), "100の倍数であれば有効")
	q, _ = NewQuantity(1100)
	assert.Equal(t, 1100, int(*q), "100の倍数であれば有効")
	q, _ = NewQuantity(0)
	assert.Equal(t, 0, int(*q), "0でも有効")
}

func TestDisabledQuantity(t *testing.T) {
	var err error
	_, err = NewQuantity(101)
	assert.Error(t, err, "100の倍数でないためエラー")
	_, err = NewQuantity(2010)
	assert.Error(t, err, "100の倍数でないためエラー")
	_, err = NewQuantity(-100)
	assert.Error(t, err, "0未満はエラー")
}
