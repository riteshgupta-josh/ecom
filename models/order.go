package models

import "time"

type Order struct {
	OrderList map[string]OrderRequest
}
type ProductDetails struct {
	Quantity  int `json:"quantity" binding:"required,min=1"`
	ProductID int `json:"productId" binding:"required"`
}
type OrderRequest struct {
	Id             string           `json:"id"`
	CustomerId     string           `json:"customerId" binding:"required"`
	ProductDetails []ProductDetails `json:"productDetails" binding:"required"`
	OrderDate      time.Time        `json:"orderDate"`
	DispatchDate   time.Time        `json:"dispatchDate"`
	TotalAmount    string           `json:"totalAmount"`
	OrderStatus    string           `json:"orderStatus" binding:"required"`
}
type OrderUpdateRequest struct {
	Id          string `json:"id" binding:"required"`
	OrderStatus string `json:"orderStatus" binding:"required"`
}

type AddOrderResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	OrderId string `json:"orderId"`
}

type UpdateOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var Order_List *Order

func InitiateOrderList() *Order {
	if Order_List == nil {
		Order_List = new(Order)
	}
	return Order_List
}
