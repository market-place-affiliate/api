package mocks

import (
	"github.com/market-place-affiliate/commonlib/lazada"
	"github.com/stretchr/testify/mock"
)

type MockLazadaRepository struct {
	mock.Mock
}

func (m *MockLazadaRepository) GetProductFeed(cred lazada.LazadaCredentials, productId string, page, limit int) (lazada.LazadaResponse[[]lazada.ProductFeedResponse], error) {
	args := m.Called(cred, productId, page, limit)
	return args.Get(0).(lazada.LazadaResponse[[]lazada.ProductFeedResponse]), args.Error(1)
}

func (m *MockLazadaRepository) GetBatchPromoteLink(cred lazada.LazadaCredentials, inputType, inputValue string, sub [6]string) (lazada.LazadaResponse[lazada.BatchPromoteLinkResponse], error) {
	args := m.Called(cred, inputType, inputValue, sub)
	return args.Get(0).(lazada.LazadaResponse[lazada.BatchPromoteLinkResponse]), args.Error(1)
}
