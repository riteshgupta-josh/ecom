package handler

import (
	productController "ecom/controllers/product"
	"ecom/logger"
	productService "ecom/service/product"

	"github.com/gin-gonic/gin"
)

func ProductHandler(router *gin.Engine) {
	funcName := "handler.ProductHandler"
	logger.I(funcName)
	productSVC := productService.NewProduct()
	product := router.Group("")
	{
		product.GET("/product", productController.GetAllProductDetails(productSVC))
		product.POST("/product", productController.AddProductDetails(productSVC))
	}
}
