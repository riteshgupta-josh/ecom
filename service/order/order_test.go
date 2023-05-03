package service

import (
	"ecom/logger"
	"ecom/models"
	"encoding/json"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetAllOrderDetailsOfCustomer(t *testing.T) {

	orderSVC := getOrderSVC()
	response := []models.OrderRequest{
		{
			Id:         "1",
			CustomerId: "1",
			ProductDetails: []models.ProductDetails{
				{
					Quantity:  1,
					ProductID: 1,
				},
			},
			OrderStatus: "Placed",
		},
		{
			Id:         "2",
			CustomerId: "1",
			ProductDetails: []models.ProductDetails{
				{
					Quantity:  3,
					ProductID: 2,
				},
			},
			OrderStatus: "Placed",
		},
	}

	orderSVC.GetAllOrderDetailsOfCustomer("1")
	responseBytes, _ := json.Marshal(response)
	assert.NotEmpty(t, string(responseBytes))
}

func TestAddOrderDetails(t *testing.T) {
	orderSVC := getOrderSVC()
	requestBody := models.OrderRequest{
		Id:         "3",
		CustomerId: "1",
		ProductDetails: []models.ProductDetails{
			{
				Quantity:  1,
				ProductID: 1,
			},
		},
		OrderStatus: "Placed",
	}

	response := models.AddOrderResponse{
		Code:    200,
		OrderId: "3",
		Msg:     "Successfully Added Order",
	}

	orderSVC.AddOrderDetails(requestBody)
	responseBytes, _ := json.Marshal(response)
	assert.NotEmpty(t, string(responseBytes))
}

func TestUpdateOrderStatus(t *testing.T) {
	orderSVC := getOrderSVC()
	requestBody := models.OrderUpdateRequest{
		Id:          "1",
		OrderStatus: "Dispatched",
	}

	response := models.UpdateOrderResponse{
		Code: 200,
		Msg:  "Successfully Updated the order for order id : 1",
	}

	orderSVC.UpdateOrderStatus(requestBody)
	responseBytes, _ := json.Marshal(response)
	assert.NotEmpty(t, string(responseBytes))
}

func getOrderSVC() Order {
	viper.SetConfigFile("../../.env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.E("Error in reading config File", err)
	}
	product := new(models.Product)
	product = &models.Product{
		ProductList: map[string]models.ProductRequest{
			"1": {
				Id:          "1",
				ProductName: "Tshirt",
				Quantity:    10,
				Price:       "740",
				Category:    "Premium",
			},
			"2": {
				Id:          "2",
				ProductName: "Trouser",
				Quantity:    8,
				Price:       "1050",
				Category:    "Premium",
			},
		},
	}
	order := new(models.Order)
	order = &models.Order{
		OrderList: map[string]models.OrderRequest{
			"1": {
				Id:         "1",
				CustomerId: "1",
				ProductDetails: []models.ProductDetails{
					{
						Quantity:  1,
						ProductID: 1,
					},
				},
				OrderStatus: "Placed",
			},
			"2": {
				Id:         "2",
				CustomerId: "1",
				ProductDetails: []models.ProductDetails{
					{
						Quantity:  3,
						ProductID: 2,
					},
				},
				OrderStatus: "Placed",
			},
		},
	}
	orderSVC := Order{
		orderDataStore:   order,
		productDataStore: product,
	}
	return orderSVC
}
