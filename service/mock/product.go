package mock

import (
	"ecom/models"

	"github.com/stretchr/testify/mock"
)

type MockProduct struct {
	mock.Mock
}

func (c MockProduct) AddProductDetails(requestBody models.ProductRequest) (res models.AddProductResponse) {
	args := c.Called(requestBody)
	if args != nil {
		res = args.Get(0).(models.AddProductResponse)
	}
	return
}
func (c MockProduct) GetAllProductDetails() (res []models.ProductRequest) {
	args := c.Called()
	if args != nil {
		res = args.Get(0).([]models.ProductRequest)
	}
	return
}
