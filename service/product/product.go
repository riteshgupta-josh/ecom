package service

import (
	"net/http"

	"ecom/models"
	"ecom/validator"

	"github.com/spf13/cast"
)

type Product struct {
	productsDataStore *models.Product
}

func NewProduct() ProductIntf {
	return &Product{
		productsDataStore: models.InitiateProductList(),
	}
}

func (p *Product) GetAllProductDetails() []models.ProductRequest {

	values := []models.ProductRequest{}
	for _, value := range p.productsDataStore.ProductList {
		values = append(values, value)
	}
	return values
}

func (p *Product) AddProductDetails(requestBody models.ProductRequest) models.AddProductResponse {

	if _, ok := p.productsDataStore.ProductList[cast.ToString(requestBody.Id)]; !ok {
		if p.productsDataStore.ProductList == nil {
			p.productsDataStore.ProductList = make(map[string]models.ProductRequest)
		}

		if validator.ProductCategoryValid(requestBody.Category) {
			p.productsDataStore.ProductList[requestBody.Id] = requestBody
			return models.AddProductResponse{
				Code:      http.StatusOK,
				ProductId: requestBody.Id,
				Msg:       "Product added Successfully",
			}
		} else {
			return models.AddProductResponse{
				Code: http.StatusBadRequest,
				Msg:  "Invalid Product category",
			}
		}
	} else {
		return models.AddProductResponse{
			Code: http.StatusBadRequest,
			Msg:  "Duplicate Product ID",
		}
	}
}
