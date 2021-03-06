package order_test

import (
	"testing"

	. "github.com/aokuyama/go-stock_jp/model/order"

	"github.com/stretchr/testify/assert"
)

func TestEnabledSession(t *testing.T) {
	var sess *Session
	var err error
	sess, err = NewSession("morning")
	assert.NoError(t, err)
	assert.Equal(t, "morning", sess.String(), "前場")
	sess, err = NewSession("afternoon")
	assert.NoError(t, err)
	assert.Equal(t, "afternoon", sess.String(), "後場")
	sess, err = NewSession("anytime")
	assert.NoError(t, err)
	assert.Equal(t, "anytime", sess.String(), "いつでも")
}

func TestDisabledSession(t *testing.T) {
	var sess *Session
	var err error
	_, err = NewSession("noon")
	assert.Nil(t, sess)
	assert.Error(t, err)
	_, err = NewSession("evening")
	assert.Nil(t, sess)
	assert.Error(t, err)
}
