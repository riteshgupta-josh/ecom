package service

import (
	"ecom/models"
)

type ProductIntf interface {
	GetAllProductDetails() []models.ProductRequest
	AddProductDetails(requestBody models.ProductRequest) models.AddProductResponse
}
