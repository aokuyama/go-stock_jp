package stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPosition(t *testing.T) {
	var s *Stock
	var err error
	s, err = NewStock("3662", "jpx")
	assert.Equal(t, `{"SequrityCode":"3662","Market":"jpx"}`, string(s.String()), "有効な株")
	assert.NoError(t, err)
	s, err = NewStock("7974", "ose")
	assert.Equal(t, `{"SequrityCode":"7974","Market":"ose"}`, string(s.String()), "有効な株")
	assert.NoError(t, err)
}
func TestErrorNewPosition(t *testing.T) {
	var s *Stock
	var err error
	s, err = NewStock("13012", "jpx")
	assert.Nil(t, s)
	assert.Error(t, err)
	_, err = NewStock("7974", "aaa")
	assert.Nil(t, s)
	assert.Error(t, err)
}
