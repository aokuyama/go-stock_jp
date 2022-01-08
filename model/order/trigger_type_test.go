package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledTriggerType(t *testing.T) {
	var tt *TriggerType
	var err error
	tt, err = NewTriggerType("more")
	assert.Equal(t, "more", tt.String(), "有効な逆指値")
	assert.NoError(t, err)
	tt, err = NewTriggerType("less")
	assert.Equal(t, "less", tt.String(), "有効な逆指値")
	assert.NoError(t, err)

}

func TestDisabledTriggerType(t *testing.T) {
	var err error
	_, err = NewTriggerType("a")
	assert.Error(t, err)
	_, err = NewTriggerType("1")
	assert.Error(t, err)
}
