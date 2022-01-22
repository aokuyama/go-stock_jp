package ordertype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledMarginType(t *testing.T) {
	c, _ := NewMarginType("system")
	assert.Equal(t, "system", c.String(), "有効な信用取引区分")
	c2, _ := NewMarginType("general")
	assert.Equal(t, "general", c2.String(), "有効な信用取引区分")
}

func TestDisabledMarginType(t *testing.T) {
	var err error
	_, err = NewMarginType("systems")
	assert.Error(t, err)
	_, err = NewMarginType("1")
	assert.Error(t, err)
}
