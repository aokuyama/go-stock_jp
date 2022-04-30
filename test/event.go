package test

import (
	"testing"

	"github.com/aokuyama/go-generic_subdomains/event/mock"
	"github.com/golang/mock/gomock"
)

func GetMockEvent(t *testing.T) *mock.MockEvent {
	gomock.NewController(t)
	mc := gomock.NewController(t)
	defer mc.Finish()
	return mock.NewMockEvent(mc)
}

func GetMockSubscriber(t *testing.T) *mock.MockSubscriber {
	gomock.NewController(t)
	mc := gomock.NewController(t)
	defer mc.Finish()
	return mock.NewMockSubscriber(mc)
}
