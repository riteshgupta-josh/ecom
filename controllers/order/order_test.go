package order

import (
	"bytes"
	"ecom/models"
	serviceMock "ecom/service/mock"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllOrderDetailsOfCustomer(t *testing.T) {
	mockOrder := serviceMock.MockOrder{}
	response := []models.OrderRequest{
		{
			Id:         "1",
			CustomerId: "1",
			ProductDetails: []models.ProductDetails{{
				Quantity:  1,
				ProductID: 1,
			}},
			OrderStatus: "Placed",
		},
	}
	mockOrder.On("GetAllOrderDetailsOfCustomer", mock.Anything).Return(response)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key:   "cid",
			Value: "1",
		},
	}
	responseBytes, _ := json.Marshal(response)
	GetAllOrderDetailsOfCustomer(mockOrder)(c)
	assert.Equal(t, w.Body.String(), string(responseBytes))
}

func TestAddOrderDetails(t *testing.T) {
	mockOrder := serviceMock.MockOrder{}
	response := models.AddOrderResponse{
		Code:    200,
		OrderId: "1",
		Msg:     "Successfully Added Order",
	}
	mockOrder.On("AddOrderDetails", mock.Anything).Return(response)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	requestBody := models.OrderRequest{
		Id:         "1",
		CustomerId: "1",
		ProductDetails: []models.ProductDetails{
			{
				Quantity:  1,
				ProductID: 1,
			},
			{
				Quantity:  3,
				ProductID: 2,
			},
		},
		OrderStatus: "Placed",
	}
	jsonbytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	responseBytes, _ := json.Marshal(response)
	AddOrderDetails(mockOrder)(c)
	assert.NotEmpty(t, string(responseBytes))
}

func TestUpdateOrderStatus(t *testing.T) {
	mockOrder := serviceMock.MockOrder{}
	response := models.UpdateOrderResponse{
		Code: 200,
		Msg:  "Successfully Updated the order for order id : 1",
	}
	mockOrder.On("UpdateOrderStatus", mock.Anything).Return(response)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")

	requestBody := models.OrderUpdateRequest{
		Id:          "1",
		OrderStatus: "Dispatched",
	}
	jsonbytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	responseBytes, _ := json.Marshal(response)
	UpdateOrderStatus(mockOrder)(c)
	assert.NotEmpty(t, string(responseBytes))
}
