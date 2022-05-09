package service_test

import (
	"errors"
	"testing"

	"github.com/aokuyama/go-generic_subdomains/event"
	"github.com/aokuyama/go-stock_jp/model/order"
	"github.com/aokuyama/go-stock_jp/model/order/mock"
	. "github.com/aokuyama/go-stock_jp/service"
	"github.com/aokuyama/go-stock_jp/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func getNewTestOrdering(t *testing.T) *Ordering {
	gomock.NewController(t)
	mc := gomock.NewController(t)
	defer mc.Finish()
	repo := mock.NewMockOrderRepository(mc)
	return NewOrdering(repo)
}

func TestOrderingOrders(t *testing.T) {
	s := getNewTestOrdering(t)
	p := event.NewPublisher()
	sub := test.GetMockSubscriber(t)
	p.Register(sub)
	o1 := test.GetOrder()
	o2 := test.GetOrder()
	odrs, _ := order.NewCollection(o1, o2)

	sub.EXPECT().Type().Return("")
	sub.EXPECT().Type().Return("")
	s.Repository.(*mock.MockOrderRepository).EXPECT().Update(o1.Ordering(), o1).Return(nil)
	s.Repository.(*mock.MockOrderRepository).EXPECT().Update(o2.Ordering(), o2).Return(nil)

	odred, err := s.OrderingOrders(odrs, p)
	assert.NoError(t, err)
	assert.Equal(t, "ordering", (*odred)[0].Status())
	assert.Equal(t, "ordering", (*odred)[1].Status())
}
func TestOrdering(t *testing.T) {
	s := getNewTestOrdering(t)
	p := event.NewPublisher()
	sub := test.GetMockSubscriber(t)
	p.Register(sub)
	o := test.GetOrder()

	t.Run("success", func(t *testing.T) {
		sub.EXPECT().Type().Return("OrderingEvent")
		sub.EXPECT().Subscribe(gomock.Any()).Return(nil)
		s.Repository.(*mock.MockOrderRepository).EXPECT().Update(o.Ordering(), o).Return(nil)
		oo, err := s.Ordering(o, p)
		assert.NoError(t, err)
		assert.Equal(t, "ordering", oo.Status())
	})

	t.Run("publish error", func(t *testing.T) {
		sub.EXPECT().Type().Return("OrderingEvent")
		sub.EXPECT().Subscribe(gomock.Any()).Return(errors.New("err"))
		oo, err := s.Ordering(o, p)
		assert.Error(t, err)
		assert.Nil(t, oo)
	})

	t.Run("save error", func(t *testing.T) {
		sub.EXPECT().Type().Return("OrderingEvent")
		sub.EXPECT().Subscribe(gomock.Any()).Return(nil)
		s.Repository.(*mock.MockOrderRepository).EXPECT().Update(o.Ordering(), o).Return(errors.New("err"))
		oo, err := s.Ordering(o, p)
		assert.Error(t, err)
		assert.Nil(t, oo)
	})
}
