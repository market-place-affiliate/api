package mocks

import (
	"github.com/market-place-affiliate/commonlib/shopee"
	"github.com/stretchr/testify/mock"
)

type MockShopeeRepository struct {
	mock.Mock
}

func (m *MockShopeeRepository) GetProductOfferListV2(cred shopee.ShopeeCredentials, shopId, itemId string) (shopee.ShopeeGetProductOfferList, error) {
	args := m.Called(cred, shopId, itemId)
	return args.Get(0).(shopee.ShopeeGetProductOfferList), args.Error(1)
}

func (m *MockShopeeRepository) GetShortLink(cred shopee.ShopeeCredentials, originalUrl string, sub [5]string) (shopee.ShopeeGetShortLink, error) {
	args := m.Called(cred, originalUrl, sub)
	return args.Get(0).(shopee.ShopeeGetShortLink), args.Error(1)
}
