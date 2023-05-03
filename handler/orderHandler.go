package handler

import (
	orderController "ecom/controllers/order"
	"ecom/logger"
	orderService "ecom/service/order"

	"github.com/gin-gonic/gin"
)

func OrderHandler(router *gin.Engine) {
	funcName := "handler.OrderHandler"
	logger.I(funcName)
	orderSVC := orderService.NewOrder()
	order := router.Group("/order")
	{
		order.GET("/customer/:cid", orderController.GetAllOrderDetailsOfCustomer(orderSVC))
		order.POST("", orderController.AddOrderDetails(orderSVC))
		order.PUT("/status", orderController.UpdateOrderStatus(orderSVC))
	}
}
