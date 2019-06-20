package handlers

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/dto/errors"
	"dev.azure.com/fee-service/dto/fee/responses"
	"dev.azure.com/fee-service/services/feeCalculator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBitcoinFee(c *gin.Context) {
	var body dto.GetFeeRequest

	body.ReceiversCount = 1

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetBitcoinFee(body.FromAddress, body.Amount, body.ReceiversCount)

	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetLitecoinFee(c *gin.Context) {
	var body dto.GetFeeRequest

	body.ReceiversCount = 1

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetLitecoinFee(body.FromAddress, body.Amount, body.ReceiversCount)
	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetBitcoinCashFee(c *gin.Context) {
	var body dto.GetFeeRequest

	body.ReceiversCount = 1

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetBitcoinCashFee(body.FromAddress, body.Amount, body.ReceiversCount)
	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetEthereumFee(c *gin.Context) {
	var body dto.GetFeeRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetEthereumFee(body.FromAddress, body.Amount)
	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetTokenFee(c *gin.Context) {
	var body dto.GetTokenFeeRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetTokenFee(body.FromAddress, body.TokenAddress, body.Amount)

	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetEthereumClassicFee(c *gin.Context) {
	var body dto.GetFeeRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetEthereumClassicFee(body.FromAddress, body.Amount)

	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetWavesFee(c *gin.Context) {
	var body dto.GetFeeRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetWavesFee(body.FromAddress, body.Amount)
	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetStellarFee(c *gin.Context) {
	var body dto.GetFeeRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetStellarFee(body.FromAddress, body.Amount)

	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func ErrorCheck(err error, apiError responses.ResponseError) (bool, int, errors.ApiError) {
	if ok, statusCode, message := handleError(err); ok {
		return true, statusCode, message
	} else if ok, statusCode, message := handleError(apiError); ok {
		return true, statusCode, message
	}
	return false, 0, errors.ApiError{}
}
