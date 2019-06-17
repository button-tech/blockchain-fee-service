package handlers

import (
	"dev.azure.com/fee-service/dto/fee/responses"
	"dev.azure.com/fee-service/services/fee"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBitcoinFee(c *gin.Context)  {
	var bitcoinFee responses.BitcoinFeeResponse
	response := fee.GetBitcoinFee()
	if ok, statusCode, message := handleError(response, &bitcoinFee); ok != false {
		c.JSON(statusCode, message)
		return
	}
	c.JSON(http.StatusOK, bitcoinFee)
}
