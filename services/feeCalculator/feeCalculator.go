package feeCalculator

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/dto/fee/responses"
	"dev.azure.com/fee-service/services/api"
)

type feeCalculator struct {
	CalcFee       func(int, int, int) int
	MinFeePerByte int
	FeePerByte    int
}

func GetBitcoinFee(address string, amount float64, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {
	utxos, apiErr := api.GetBitcoinUtxo(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	feePerByte, apiErr := api.GetBitcoinFee()
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	fee, apiErr := calcUtxoFee(utxos.Utxo, amount, receiversCount, feeCalculator{
		CalcFee:       calcBitcoinFee,
		MinFeePerByte: 10,
		FeePerByte:    feePerByte.HalfHourFee,
	})
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	return fee, responses.ResponseError{}, nil
}

func GetLitecoinFee(address string, amount float64, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {
	utxos, apiErr := api.GetLitecoinUtxo(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	feePerByte, apiErr := api.GetLitecoinFee()
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	fee, apiErr := calcUtxoFee(utxos.Utxo, amount, receiversCount, feeCalculator{
		CalcFee:       calcLitecoinFee,
		MinFeePerByte: 8,
		FeePerByte:    feePerByte.MediumFeePerKb / 1024,
	})
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	return fee, responses.ResponseError{}, nil
}

func GetBitcoinCashFee(address string, amount float64, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {
	utxos, apiErr := api.GetBitcoinCashUtxo(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	fee, apiErr := calcUtxoFee(utxos.Utxo, amount, receiversCount, feeCalculator{
		CalcFee:       calcBitcoinCashFee,
		MinFeePerByte: 1,
		FeePerByte:    3,
	})
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	return fee, responses.ResponseError{}, nil
}

func calcBitcoinFee(inputCount, outputCount, feePerByte int) int {
	return (inputCount*148 + outputCount*34 + 10) * feePerByte
}

func calcLitecoinFee(inputCount, outputCount, feePerByte int) int {
	return (inputCount*148 + outputCount*34 + 10) * feePerByte
}

func calcBitcoinCashFee(inputCount, outputCount, feePerByte int) int {
	return (inputCount*148 + outputCount*34 + 10) * feePerByte
}
