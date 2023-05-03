package service

import (
	"ecom/models"
	"encoding/json"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetAllProductDetails(t *testing.T) {

	productSVC := getProductSVC()
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

	actualResponse := productSVC.GetAllProductDetails()
	actualResponseBytes, _ := json.Marshal(actualResponse)
	responseBytes, _ := json.Marshal(response)
	assert.Equal(t, string(actualResponseBytes), string(responseBytes))
}

func TestAddProductDetails(t *testing.T) {
	productSVC := getProductSVC()
	requestBody := models.ProductRequest{
		Id:          "3",
		ProductName: "Shoes",
		Quantity:    10,
		Price:       "740",
		Category:    "Premium",
	}

	response := models.AddProductResponse{
		Code:      200,
		ProductId: "3",
		Msg:       "Product added Successfully",
	}

	actualResponse := productSVC.AddProductDetails(requestBody)
	actualResponseBytes, _ := json.Marshal(actualResponse)
	responseBytes, _ := json.Marshal(response)
	assert.Equal(t, string(actualResponseBytes), string(responseBytes))
}

func TestAddProductDetailsDuplicateIDScenario(t *testing.T) {
	productSVC := getProductSVC()
	requestBody := models.ProductRequest{
		Id:          "1",
		ProductName: "Shoes",
		Quantity:    10,
		Price:       "740",
		Category:    "Premium",
	}

	response := models.AddProductResponse{
		Code: 400,
		Msg:  "Duplicate Product ID",
	}

	actualResponse := productSVC.AddProductDetails(requestBody)
	actualResponseBytes, _ := json.Marshal(actualResponse)
	responseBytes, _ := json.Marshal(response)
	assert.Equal(t, string(actualResponseBytes), string(responseBytes))
}

func TestAddProductDetailsInvalidProductCategoryScenario(t *testing.T) {
	productSVC := getProductSVC()
	requestBody := models.ProductRequest{
		Id:          "3",
		ProductName: "Shoes",
		Quantity:    10,
		Price:       "740",
		Category:    "abc",
	}

	response := models.AddProductResponse{
		Code: 400,
		Msg:  "Invalid Product category",
	}

	actualResponse := productSVC.AddProductDetails(requestBody)
	actualResponseBytes, _ := json.Marshal(actualResponse)
	responseBytes, _ := json.Marshal(response)
	assert.Equal(t, string(actualResponseBytes), string(responseBytes))
}

func getProductSVC() Product {
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
	productSVC := Product{
		productsDataStore: product,
	}
	return productSVC
}
