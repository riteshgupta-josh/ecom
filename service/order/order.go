package service

import (
	"net/http"
	"time"

	"ecom/logger"
	"ecom/models"
	"ecom/validator"

	"ecom/utils"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Order struct {
	orderDataStore   *models.Order
	productDataStore *models.Product
}

func NewOrder() OrderIntf {
	return &Order{
		orderDataStore:   models.InitiateOrderList(),
		productDataStore: models.InitiateProductList(),
	}
}

func (orderSVC *Order) GetAllOrderDetailsOfCustomer(cid string) []models.OrderRequest {
	ordersList := []models.OrderRequest{}
	for _, orderItem := range orderSVC.orderDataStore.OrderList {
		if orderItem.CustomerId == cid {
			ordersList = append(ordersList, orderItem)
		}
	}
	return ordersList
}

func (orderSVC *Order) AddOrderDetails(requestBody models.OrderRequest) models.AddOrderResponse {

	if orderSVC.orderDataStore.OrderList == nil {
		orderSVC.orderDataStore.OrderList = make(map[string]models.OrderRequest)
	}
	if _, ok := orderSVC.orderDataStore.OrderList[cast.ToString(requestBody.Id)]; !ok {
		var count int
		var order_total_amount int
		for _, requestValue := range requestBody.ProductDetails {

			if productValue, ok := orderSVC.productDataStore.ProductList[cast.ToString(requestValue.ProductID)]; ok {

				if requestValue.Quantity > viper.GetInt("PRODUCT_MAX_QUANTITY") {
					logger.E("Cannot order more than " + cast.ToString(viper.GetInt("PRODUCT_MAX_QUANTITY")))
					return models.AddOrderResponse{
						Code: http.StatusBadRequest,
						Msg:  "Cannot order more than " + cast.ToString(viper.GetInt("PRODUCT_MAX_QUANTITY")),
					}

				} else if requestValue.Quantity > productValue.Quantity {
					logger.E("Out of Stock")
					return models.AddOrderResponse{
						Code: http.StatusBadRequest,
						Msg:  "Out of Stock",
					}
				}
				if productValue.Category == "Premium" {
					count++
				}
				order_total_amount += cast.ToInt(productValue.Price) * requestValue.Quantity
				productValue.Quantity -= requestValue.Quantity
				orderSVC.productDataStore.ProductList[cast.ToString(requestValue.ProductID)] = productValue

			} else {
				logger.E("Product not found")
				return models.AddOrderResponse{
					Code: http.StatusBadRequest,
					Msg:  "Product not found",
				}
			}
		}
		if count > 2 {
			requestBody.TotalAmount = cast.ToString(cast.ToFloat32(requestBody.TotalAmount) - cast.ToFloat32(requestBody.TotalAmount)*0.1)
		}
		requestBody.Id = utils.GenerateUUID()
		requestBody.OrderDate = time.Now()
		requestBody.TotalAmount = cast.ToString(order_total_amount)
		orderSVC.orderDataStore.OrderList[requestBody.Id] = requestBody
		logger.I("Successfully Added Order")
		return models.AddOrderResponse{
			Code:    http.StatusOK,
			OrderId: requestBody.Id,
			Msg:     "Successfully Added Order",
		}
	} else {
		logger.E("Duplicate Order ID")
		return models.AddOrderResponse{
			Code: http.StatusBadRequest,
			Msg:  "Duplicate Order ID",
		}
	}
}
func (orderSVC *Order) UpdateOrderStatus(requestStatusBody models.OrderUpdateRequest) models.UpdateOrderResponse {
	logger.I(requestStatusBody.OrderStatus)
	if orderValue, ok := orderSVC.orderDataStore.OrderList[requestStatusBody.Id]; ok {
		if validator.OrderStatusValid(requestStatusBody.OrderStatus) {

			if requestStatusBody.OrderStatus != "Placed" {

				if requestStatusBody.OrderStatus == "Dispatched" {
					orderValue.DispatchDate = time.Now()
				} else if requestStatusBody.OrderStatus == "Cancelled" {
					for _, productDetails := range orderValue.ProductDetails {
						if productValue, ok := orderSVC.productDataStore.ProductList[cast.ToString(productDetails.ProductID)]; ok {
							productValue.Quantity += productDetails.Quantity
							// fmt.Println(productValue.Quantity)
							orderSVC.productDataStore.ProductList[cast.ToString(productDetails.ProductID)] = productValue
						}
					}
				}
				orderValue.OrderStatus = requestStatusBody.OrderStatus
				orderSVC.orderDataStore.OrderList[cast.ToString(requestStatusBody.Id)] = orderValue
				logger.I("Successfully Updated the order for order id : " + requestStatusBody.Id)
				return models.UpdateOrderResponse{
					Code: http.StatusOK,
					Msg:  "Successfully Updated the order for order id : " + requestStatusBody.Id,
				}
			} else {
				logger.E("Invalid request body. Please request to change the status for `Dispatched`, `Cancelled`, `Completed`")
				return models.UpdateOrderResponse{
					Code: http.StatusBadRequest,
					Msg:  "Invalid request body. Please request to change the status for `Dispatched`, `Cancelled`, `Completed`",
				}
			}
		} else {
			logger.E("Invalid request body. Please request to change the status for `Dispatched`, `Cancelled`, `Completed` and `Placed`")
			return models.UpdateOrderResponse{
				Code: http.StatusBadRequest,
				Msg:  "Invalid request body. Please request to change the status for `Dispatched`, `Cancelled`, `Completed` and `Placed`",
			}
		}
	} else {
		logger.E("Order not found")
		return models.UpdateOrderResponse{
			Code: http.StatusBadRequest,
			Msg:  "Order not found",
		}
	}
}
