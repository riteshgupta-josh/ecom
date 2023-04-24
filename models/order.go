package models

import "time"

type Order struct {
	OrderList map[string]OrderRequest
}
type ProductDetails struct {
	Quantity  int `json:"quantity"`
	ProductID int `json:"productId"`
}
type OrderRequest struct {
	Id             string           `json:"id"`
	ProductDetails []ProductDetails `json:"productDetails"`
	OrderDate      time.Time        `json:"orderDate"`
	DispatchDate   time.Time        `json:"dispatchDate"`
	TotalAmount    string           `json:"totalAmount"`
	OrderStatus    string           `json:"orderStatus"`
}
type OrderUpdateRequest struct {
	Id          string `json:"id"`
	OrderStatus string `json:"orderStatus"`
}

var Order_List *Order

func InitiateOrderList() *Order {
	if Order_List == nil {
		Order_List = new(Order)
	}
	return Order_List
}
