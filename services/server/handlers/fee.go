package handlers

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/services/feeCalculator"
	"dev.azure.com/moon-pay/dto/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBitcoinFee(c *gin.Context) {
	var body dto.GetFeeRequest
	body.ReceiversCount = 1
	if ok, statusCode, message := handleError(errors.BadRequest{
		Error:   c.ShouldBindJSON(&body),
		Message: errors.BadRequestMessage,
	}); ok != false {
		c.JSON(statusCode, message)
		return
	}
	res, apiErr, err := feeCalculator.GetBitcoinFee(body.FromAddress, body.Amount, body.ReceiversCount)
	if ok, statusCode, message := handleError(err); ok != false {
		c.JSON(statusCode, message)
		return
	}
	if ok, statusCode, message := handleError(apiErr); ok != false {
		c.JSON(statusCode, message)
		return
	}
	c.JSON(http.StatusOK, res)
}
