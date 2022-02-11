package main

import (
	"CurrencyConverter/controller"
	"CurrencyConverter/models"
	"github.com/gin-gonic/gin"
)

func CurrencyRouter(route *gin.Engine){
	exchangeTable := models.CreateExchangeTable()
	route.POST("/currencyConverter", controller.ConvertCurrencyExchange(exchangeTable.CurrencyTable))
}
