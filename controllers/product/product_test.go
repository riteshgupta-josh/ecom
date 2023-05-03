package product

import (
	"bytes"
	"ecom/models"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	serviceMock "ecom/service/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllProductDetails(t *testing.T) {
	mockProduct := serviceMock.MockProduct{}
	response := []models.ProductRequest{
		{
			Id:          "1",
			ProductName: "Tshirt",
			Quantity:    10,
			Price:       "740",
			Category:    "Premium",
		},
		{
			Id:          "2",
			ProductName: "Trouser",
			Quantity:    8,
			Price:       "1050",
			Category:    "Premium",
		},
	}
	mockProduct.On("GetAllProductDetails").Return(response)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	responseBytes, _ := json.Marshal(response)
	GetAllProductDetails(mockProduct)(c)
	assert.Equal(t, w.Body.String(), string(responseBytes))
}

func TestAddProductDetails(t *testing.T) {
	mockProduct := serviceMock.MockProduct{}
	response := models.AddProductResponse{
		Code:      200,
		ProductId: "1",
		Msg:       "Product added Successfully",
	}
	mockProduct.On("AddProductDetails", mock.Anything).Return(response)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	requestBody := models.ProductRequest{
		Id:          "1",
		ProductName: "Tshirt",
		Quantity:    10,
		Price:       "740",
		Category:    "Premium",
	}
	jsonbytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	responseBytes, _ := json.Marshal(response)
	AddProductDetails(mockProduct)(c)
	assert.NotEmpty(t, string(responseBytes))
}
