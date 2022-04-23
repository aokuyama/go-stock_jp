package common_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/common"

	"github.com/stretchr/testify/assert"
)

func TestEnabledLocale(t *testing.T) {
	var st *Locale
	var err error
	st, err = NewLocale("jp")
	assert.NoError(t, err)
	assert.Equal(t, "jp", st.String(), "有効な地域設定")
	st, err = NewLocale("ny")
	assert.NoError(t, err)
	assert.Equal(t, "ny", st.String(), "有効な地域設定")
}

func TestDisabledLocale(t *testing.T) {
	var st *Locale
	var err error
	_, err = NewLocale("aaa")
	assert.Nil(t, st)
	assert.Error(t, err)
	_, err = NewLocale("j")
	assert.Nil(t, st)
	assert.Error(t, err)
}
