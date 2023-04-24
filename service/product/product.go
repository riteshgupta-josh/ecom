package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riteshgupta-josh/logger"
	"github.com/riteshgupta-josh/models"
)

var product = models.InitiateProductList()

func GetAllProductDetails(c *gin.Context) {

	if product.ProductList != nil {
		c.JSON(http.StatusOK, product.ProductList)
	} else {
		c.JSON(http.StatusNotFound, "No Product added yet")
	}
}

func GetProductDetailsById(c *gin.Context) {
	id := c.Param("id")
	logger.I(id)
	logger.I(product.ProductList)

	if _, ok := product.ProductList[id]; ok {
		c.JSON(http.StatusOK, product.ProductList[id])
	} else {
		c.JSON(http.StatusNotFound, "Product not found")
	}

}

func AddProductDetails(c *gin.Context) {
	requestBody := models.ProductRequest{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	logger.I(requestBody)
	if product.ProductList == nil {
		product.ProductList = make(map[string]models.ProductRequest)
	}
	product.ProductList[requestBody.Id] = requestBody
	c.JSON(http.StatusOK, product.ProductList)
}
