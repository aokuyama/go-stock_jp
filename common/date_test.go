package common_test

import (
	"testing"
	"time"

	. "github.com/aokuyama/go-stock_jp/common"

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

func TestDateBefore(t *testing.T) {
	d1, _ := NewDate("2021-01-01")
	d2, _ := NewDate("2021-01-02")
	assert.True(t, d1.Before(d2))
	assert.False(t, d2.Before(d1))
	assert.False(t, d1.Before(d1))
}

func TestDateAfter(t *testing.T) {
	d1, _ := NewDate("2020-09-13")
	d2, _ := NewDate("2020-09-14")
	assert.False(t, d1.After(d2))
	assert.True(t, d2.After(d1))
	assert.False(t, d1.After(d1))
}

func TestNewDateAdded(t *testing.T) {
	d1, _ := NewDate("2021-01-16")
	assert.Equal(t, "2021-01-16", d1.String())
	assert.Equal(t, "2021-01-17", NewDateAdded(d1, time.Hour*24*1).String())
	assert.Equal(t, "2021-01-16", d1.String(), "元の日付は変更されない")
	assert.Equal(t, "2021-01-11", NewDateAdded(d1, -time.Hour*24*5).String())
	assert.Equal(t, "2021-01-16", d1.String(), "元の日付は変更されない")
}
