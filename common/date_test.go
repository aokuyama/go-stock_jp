package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnabledDate(t *testing.T) {
	var d *Date
	var err error
	d, err = NewDate("2000-01-01")
	assert.Equal(t, "2000-01-01", d.String(), "有効な日付")
	assert.NoError(t, err)
	d, err = NewDate("2021-12-25")
	assert.Equal(t, "2021-12-25", d.String(), "有効な日付")
	assert.NoError(t, err)
}

func TestDisabledDate(t *testing.T) {
	var err error
	_, err = NewDate("1")
	assert.Error(t, err)
	_, err = NewDate("50000")
	assert.Error(t, err)
}

func TestDisabledFormat(t *testing.T) {
	var err error
	_, err = NewDate("2000/01/01")
	assert.Error(t, err)
	_, err = NewDate("2000-1-1")
	assert.Error(t, err)
}
