package order

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
}

func TestDisabledQuantity(t *testing.T) {
	var err error
	_, err = NewQuantity(101)
	assert.Error(t, err, "100の倍数でないためエラー")
	_, err = NewQuantity(2010)
	assert.Error(t, err, "100の倍数でないためエラー")
	_, err = NewQuantity(0)
	assert.Error(t, err, "0はエラー")
	_, err = NewQuantity(-100)
	assert.Error(t, err, "0未満はエラー")
}
func TestNewErrorQuantity(t *testing.T) {
	q, _ := NewErrorQuantity(100)
	assert.Equal(t, -100, q.Int())
	assert.True(t, q.IsError())
	q2, _ := NewErrorQuantity(1300)
	assert.Equal(t, -1300, q2.Int())
	assert.True(t, q.IsError())
}
