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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: body.ReceiversCount,
		Speed:          body.Speed,
		TokenAddress:   "",
	})

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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: body.ReceiversCount,
		Speed:          body.Speed,
		TokenAddress:   "",
	})

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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: body.ReceiversCount,
		Speed:          body.Speed,
		TokenAddress:   "",
	})

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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: body.ReceiversCount,
		Speed:          body.Speed,
		TokenAddress:   "",
	})

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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: 0,
		Speed:          body.Speed,
		TokenAddress:   body.TokenAddress,
	})

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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: body.ReceiversCount,
		Speed:          body.Speed,
		TokenAddress:   "",
	})

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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: body.ReceiversCount,
		Speed:          body.Speed,
		TokenAddress:   "",
	})

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
	params := paramsProcessing(feeCalculator.Params{
		Address:        body.FromAddress,
		Amount:         body.Amount,
		ReceiversCount: body.ReceiversCount,
		Speed:          body.Speed,
		TokenAddress:   "",
	})

	res, apiErr, err := feeCalculator.GetStellarFee(params)

	isErr, statusCode, error := ErrorCheck(err, apiErr)
	if isErr {
		c.JSON(statusCode, error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetZilliqaFee(c *gin.Context) {
	var body dto.GetFeeRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestMessage)
		return
	}

	res, apiErr, err := feeCalculator.GetZilliqaFee(body.FromAddress, body.Amount)
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

func paramsProcessing(params feeCalculator.Params) *feeCalculator.Params {
	if params.Speed != "low" && params.Speed != "average" && params.Speed != "fast" || params.Speed == "" {
		params.Speed = "average"
	}
	return &params
}
