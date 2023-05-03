package product

import (
	"ecom/models"
	productService "ecom/service/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProductDetails(productSVC productService.ProductIntf) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		serviceResponse := productSVC.GetAllProductDetails()
		if len(serviceResponse) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "No Product added yet"})
			return
		}
		ctx.JSON(http.StatusOK, serviceResponse)
	}
}

func AddProductDetails(productSVC productService.ProductIntf) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.ProductRequest{}
		if err := c.BindJSON(&requestBody); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		serviceResponse := productSVC.AddProductDetails(requestBody)
		if serviceResponse.Code != 200 {
			c.JSON(http.StatusBadRequest, gin.H{"message": serviceResponse.Msg})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Response": serviceResponse})
	}
}
