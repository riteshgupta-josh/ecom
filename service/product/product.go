package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riteshgupta-josh/logger"
	"github.com/riteshgupta-josh/models"
	"github.com/riteshgupta-josh/validator"
	"github.com/spf13/cast"
)

var product = models.InitiateProductList()

func GetAllProductDetails(c *gin.Context) {

	if product.ProductList != nil {
		values := []models.ProductRequest{}
		for _, value := range product.ProductList {
			values = append(values, value)
		}
		c.JSON(http.StatusOK, values)
	} else {
		c.JSON(http.StatusBadRequest, "No Product added yet")
	}
}

func GetProductDetailsById(c *gin.Context) {
	id := c.Param("id")
	logger.I(id)
	logger.I(product.ProductList)

	if _, ok := product.ProductList[id]; ok {
		c.JSON(http.StatusOK, product.ProductList[id])
	} else {
		c.JSON(http.StatusBadRequest, "Product not found")
	}

}

func AddProductDetails(c *gin.Context) {
	requestBody := models.ProductRequest{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	logger.I(requestBody)
	if _, ok := product.ProductList[cast.ToString(requestBody.Id)]; !ok {
		if product.ProductList == nil {
			product.ProductList = make(map[string]models.ProductRequest)
		}

		if validator.ProductCategoryValid(requestBody.Category) {
			product.ProductList[requestBody.Id] = requestBody
			c.JSON(http.StatusOK, gin.H{"message": "Product added Successfully"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Product category"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Duplicate Product ID"})
	}
}
