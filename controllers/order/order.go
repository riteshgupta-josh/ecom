package order

import (
	"ecom/models"
	orderService "ecom/service/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrderDetailsOfCustomer(orderSVC orderService.OrderIntf) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		serviceResponse := orderSVC.GetAllOrderDetailsOfCustomer(cid)
		if len(serviceResponse) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Orders not found for customer"})
			return
		}
		ctx.JSON(http.StatusOK, serviceResponse)
	}
}
func AddOrderDetails(orderSVC orderService.OrderIntf) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.OrderRequest{}
		if err := c.BindJSON(&requestBody); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		serviceResponse := orderSVC.AddOrderDetails(requestBody)
		if serviceResponse.Code != 200 {
			c.JSON(http.StatusBadRequest, gin.H{"message": serviceResponse.Msg})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Response": serviceResponse})
	}
}

func UpdateOrderStatus(orderSVC orderService.OrderIntf) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.OrderUpdateRequest{}
		if err := c.BindJSON(&requestBody); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		serviceResponse := orderSVC.UpdateOrderStatus(requestBody)
		if serviceResponse.Code != 200 {
			c.JSON(http.StatusBadRequest, gin.H{"message": serviceResponse.Msg})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Response": serviceResponse})
	}
}
