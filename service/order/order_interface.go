package service

import (
	"ecom/models"
)

type OrderIntf interface {
	GetAllOrderDetailsOfCustomer(cid string) []models.OrderRequest
	AddOrderDetails(requestBody models.OrderRequest) models.AddOrderResponse
	UpdateOrderStatus(requestBody models.OrderUpdateRequest) models.UpdateOrderResponse
}
