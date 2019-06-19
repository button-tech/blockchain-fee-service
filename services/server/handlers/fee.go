package handlers

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/dto/errors"
	"dev.azure.com/fee-service/services/feeCalculator"
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

func GetLitecoinFee(c *gin.Context) {
	var body dto.GetFeeRequest
	body.ReceiversCount = 1
	if ok, statusCode, message := handleError(errors.BadRequest{
		Error:   c.ShouldBindJSON(&body),
		Message: errors.BadRequestMessage,
	}); ok != false {
		c.JSON(statusCode, message)
		return
	}
	res, apiErr, err := feeCalculator.GetLitecoinFee(body.FromAddress, body.Amount, body.ReceiversCount)
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

func GetBitcoinCashFee(c *gin.Context) {
	var body dto.GetFeeRequest
	body.ReceiversCount = 1
	if ok, statusCode, message := handleError(errors.BadRequest{
		Error:   c.ShouldBindJSON(&body),
		Message: errors.BadRequestMessage,
	}); ok != false {
		c.JSON(statusCode, message)
		return
	}
	res, apiErr, err := feeCalculator.GetBitcoinCashFee(body.FromAddress, body.Amount, body.ReceiversCount)
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

func GetEthereumFee(c *gin.Context) {
	var body dto.GetFeeRequest
	if ok, statusCode, message := handleError(errors.BadRequest{
		Error:   c.ShouldBindJSON(&body),
		Message: errors.BadRequestMessage,
	}); ok != false {
		c.JSON(statusCode, message)
		return
	}
	res, apiErr, err := feeCalculator.GetEthereumFee(body.FromAddress, body.Amount)
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

func GetEthereumClassicFee(c *gin.Context) {
	var body dto.GetFeeRequest
	if ok, statusCode, message := handleError(errors.BadRequest{
		Error:   c.ShouldBindJSON(&body),
		Message: errors.BadRequestMessage,
	}); ok != false {
		c.JSON(statusCode, message)
		return
	}
	res, apiErr, err := feeCalculator.GetEthereumClassicFee(body.FromAddress, body.Amount)
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

func GetWavesFee(c *gin.Context) {
	var body dto.GetFeeRequest
	if ok, statusCode, message := handleError(errors.BadRequest{
		Error:   c.ShouldBindJSON(&body),
		Message: errors.BadRequestMessage,
	}); ok != false {
		c.JSON(statusCode, message)
		return
	}
	res, apiErr, err := feeCalculator.GetWavesFee(body.FromAddress, body.Amount)
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

func GetStellarFee(c *gin.Context) {
	var body dto.GetFeeRequest
	if ok, statusCode, message := handleError(errors.BadRequest{
		Error:   c.ShouldBindJSON(&body),
		Message: errors.BadRequestMessage,
	}); ok != false {
		c.JSON(statusCode, message)
		return
	}
	res, apiErr, err := feeCalculator.GetStellarFee(body.FromAddress, body.Amount)
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
