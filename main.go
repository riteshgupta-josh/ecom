package main

import (
	"ecom/handler"
	"ecom/logger"
	"ecom/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	funcName := "main.main"
	logger.I(funcName)
	ConstPORT := utils.GetConfigValue("PORT")
	router := gin.New()
	handler.ProductHandler(router)
	handler.OrderHandler(router)
	router.Run(":" + ConstPORT)
}
