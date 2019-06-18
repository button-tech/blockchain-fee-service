package feeCalculator

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/dto/fee/responses"
	"dev.azure.com/fee-service/services/api"
)

type utxoBlockchain struct {
	AllUtxos             []responses.Utxo
	TotalBalance         int
	UsefulBalance        int
	SatoshiAmount        int
	CalcFee              func(int, int, int) int
	MinFeePerByte        int
	FeePerByte           int
	MinFee               int
	Fee                  int
	MinInputs            int
	Input                int
	Output               int
	LastIterationBalance int
	UsefulUtxos          []responses.Utxo
	UselessUtxos         []responses.Utxo
	DustUtxos            []responses.Utxo
	MaxAmount            int
	MaxUsefulAmount      int
	IsBadFee             bool
	IsEnough             bool
}

func GetBitcoinFee(address string, amount float64, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {
	utxos, apiErr := api.GetBitcoinUtxo(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	fee, apiErr := calcUtxoFee(utxos.Utxo, amount, receiversCount, 10, calcBitcoinFee)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	return fee, responses.ResponseError{}, nil
}
