package handlers

import (
	"dev.azure.com/fee-service/services/fee"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBitcoinFee(c *gin.Context) {
	fee, err := fee.GetBitcoinFee()
	if ok, statusCode, message := handleError(err); ok != false {
		c.JSON(statusCode, message)
		return
	}
	c.JSON(http.StatusOK, fee)
}
