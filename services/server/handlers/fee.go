package handlers

import (
	"github.com/button-tech/blockchain-fee-service/dto"
	"github.com/button-tech/blockchain-fee-service/dto/errors"
	"github.com/button-tech/blockchain-fee-service/dto/fee/responses"
	"github.com/button-tech/blockchain-fee-service/services/feeCalculator"
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
	params := paramsProcessing(body.ReceiversCount, body.FromAddress, body.Amount, body.Speed, "")

	res, apiErr, err := feeCalculator.GetBitcoinFee(params)

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
	params := paramsProcessing(body.ReceiversCount, body.FromAddress, body.Amount, body.Speed, "")

	res, apiErr, err := feeCalculator.GetLitecoinFee(params)
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
	params := paramsProcessing(body.ReceiversCount, body.FromAddress, body.Amount, body.Speed, "")

	res, apiErr, err := feeCalculator.GetBitcoinCashFee(params)
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
	params := paramsProcessing(body.ReceiversCount, body.FromAddress, body.Amount, body.Speed, "")

	res, apiErr, err := feeCalculator.GetEthereumFee(params)
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

	params := paramsProcessing(0, body.FromAddress, body.Amount, body.Speed, body.TokenAddress)

	res, apiErr, err := feeCalculator.GetTokenFee(params)

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
	params := paramsProcessing(body.ReceiversCount, body.FromAddress, body.Amount, body.Speed, "")

	res, apiErr, err := feeCalculator.GetEthereumClassicFee(params)

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
	params := paramsProcessing(body.ReceiversCount, body.FromAddress, body.Amount, body.Speed, "")

	res, apiErr, err := feeCalculator.GetWavesFee(params)
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
	params := paramsProcessing(body.ReceiversCount, body.FromAddress, body.Amount, body.Speed, "")

	res, apiErr, err := feeCalculator.GetStellarFee(params)

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

func paramsProcessing(i int, address, amount, speed, tokenAddress string) *feeCalculator.Params {
	var p feeCalculator.Params
	if speed != "" {
		p.Speed = speed
	} else {
		p.Speed = "1.5"
	}
	p.TokenAddress = tokenAddress
	p.Address = address
	p.ReceiversCount = i
	p.Amount = amount

	return &p
}
