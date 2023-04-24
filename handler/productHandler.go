package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/riteshgupta-josh/logger"
	service "github.com/riteshgupta-josh/service/product"
	// productServices "github.com/riteshgupta-josh/service"
)

func ProductHandler(router *gin.Engine) {
	funcName := "handler.ProductHandler"
	logger.I(funcName)
	product := router.Group("/product")
	{
		product.GET("/alldetails", service.GetAllProductDetails)
		product.GET("/details/:id", service.GetProductDetailsById)
		product.POST("/add", service.AddProductDetails)
	}
}
