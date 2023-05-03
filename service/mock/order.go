package mock

import (
	"ecom/models"

	"github.com/stretchr/testify/mock"
)

type MockOrder struct {
	mock.Mock
}

func (c MockOrder) AddOrderDetails(requestBody models.OrderRequest) (res models.AddOrderResponse) {
	args := c.Called(requestBody)
	if args != nil {
		res = args.Get(0).(models.AddOrderResponse)
	}
	return
}
func (c MockOrder) GetAllOrderDetailsOfCustomer(cid string) (res []models.OrderRequest) {
	args := c.Called(cid)
	if args != nil {
		res = args.Get(0).([]models.OrderRequest)
	}
	return
}
func (c MockOrder) UpdateOrderStatus(requestStatusBody models.OrderUpdateRequest) (res models.UpdateOrderResponse) {
	args := c.Called(requestStatusBody)
	if args != nil {
		res = args.Get(0).(models.UpdateOrderResponse)
	}
	return
}
