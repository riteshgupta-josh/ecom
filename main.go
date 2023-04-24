package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riteshgupta-josh/handler"
	"github.com/riteshgupta-josh/logger"
	"github.com/riteshgupta-josh/utils"
)

// func init() {
// 	funcName := "main.init"
// 	var result = make(map[string]string)
// result["id"] = "1"
// result["product_name"] = "tshirt"
// result["availability"] = "yes"
// result["price"] = "500"
// result["category"] = "budget"
// 	logger.I(funcName)
// }

// func initializeResult() map[string]string {
// 	var result = make(map[string]string)
// 	result["id"] = "1"
// 	result["product_name"] = "tshirt"
// 	result["availability"] = "yes"
// 	result["price"] = "500"
// 	result["category"] = "budget"
// 	return result
// }

func main() {
	funcName := "main.main"
	logger.I(funcName)
	ConstPORT := utils.GetConfigValue("PORT")
	router := gin.New()
	handler.ProductHandler(router)
	handler.OrderHandler(router)
	// User Handler
	router.Run(":" + ConstPORT)

}
