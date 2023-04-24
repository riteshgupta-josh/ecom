package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/riteshgupta-josh/models"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var order = models.InitiateOrderList()
var product = models.InitiateProductList()

func GetAllOrderDetailsOfCustomer(c *gin.Context) {
	cid := c.Param("cid")

	if _, ok := order.OrderList[cid]; ok {
		c.JSON(http.StatusOK, order.OrderList[cid])
	} else {
		c.JSON(http.StatusNotFound, "Order not found")
	}

}
func GetOrderDetailsById(c *gin.Context) {
	id := c.Param("id")

	if _, ok := order.OrderList[id]; ok {
		c.JSON(http.StatusOK, order.OrderList[id])
	} else {
		c.JSON(http.StatusNotFound, "Order not found")
	}

}
func AddOrderDetails(c *gin.Context) {
	requestBody := models.OrderRequest{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if order.OrderList == nil {
		order.OrderList = make(map[string]models.OrderRequest)
	}
	var count int
	for _, requestValue := range requestBody.ProductDetails {

		if productValue, ok := product.ProductList[cast.ToString(requestValue.ProductID)]; ok {

			if requestValue.Quantity > viper.GetInt("PRODUCT_MAX_QUANTITY") {
				c.JSON(410, "Cannot order more than "+cast.ToString(viper.GetInt("PRODUCT_MAX_QUANTITY")))
				return

			} else if requestValue.Quantity > productValue.Quantity {
				c.JSON(410, "Out of Stock")
				return
			}
			if productValue.Category == "Premium" {
				count++
			}
			productValue.Quantity -= requestValue.Quantity
			product.ProductList[cast.ToString(requestValue.ProductID)] = productValue

		} else {
			c.JSON(http.StatusNotFound, "Product not found")
			return
		}
	}
	if count > 2 {
		requestBody.TotalAmount = cast.ToString(cast.ToFloat32(requestBody.TotalAmount) - cast.ToFloat32(requestBody.TotalAmount)*0.1)
	}
	requestBody.OrderDate = time.Now()
	order.OrderList[requestBody.Id] = requestBody
	c.JSON(http.StatusOK, "Successfully Added Order")
}
func UpdateOrderStatus(c *gin.Context) {
	requestStatusBody := models.OrderUpdateRequest{}
	if err := c.BindJSON(&requestStatusBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if orderValue, ok := order.OrderList[requestStatusBody.Id]; ok {
		if requestStatusBody.OrderStatus == "Dispatched" {
			orderValue.DispatchDate = time.Now()
		} else if requestStatusBody.OrderStatus == "Cancelled" {
			for _, productDetails := range orderValue.ProductDetails {
				if productValue, ok := product.ProductList[cast.ToString(productDetails.ProductID)]; ok {
					productValue.Quantity += productDetails.Quantity
					// fmt.Println(productValue.Quantity)
					product.ProductList[cast.ToString(productDetails.ProductID)] = productValue
				}
			}
		}
		orderValue.OrderStatus = requestStatusBody.OrderStatus
		order.OrderList[cast.ToString(requestStatusBody.Id)] = orderValue
		c.JSON(http.StatusOK, "Successfully Updated Order")
	} else {
		c.JSON(http.StatusNotFound, "Order not found")
		return
	}

}
