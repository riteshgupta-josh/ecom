package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/riteshgupta-josh/logger"
	service "github.com/riteshgupta-josh/service/order"
)

func OrderHandler(router *gin.Engine) {
	funcName := "handler.OrderHandler"
	logger.I(funcName)
	order := router.Group("/order")
	{
		order.GET("/customerdetails/:cid", service.GetAllOrderDetailsOfCustomer)
		order.GET("/details/:id", service.GetOrderDetailsById)
		order.POST("/add", service.AddOrderDetails)
		order.POST("/updateStatus", service.UpdateOrderStatus)
	}
}
