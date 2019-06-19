package feeCalculator

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/dto/fee/requests"
	"dev.azure.com/fee-service/dto/fee/responses"
	"dev.azure.com/fee-service/services/api"
)

type feeCalculator struct {
	CalcFee       func(int, int, int) int
	MinFeePerByte int
	FeePerByte    int
}

func GetBitcoinFee(address string, amount string, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {
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

func GetLitecoinFee(address string, amount string, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {
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

func GetBitcoinCashFee(address string, amount string, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {
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

func GetEthereumFee(address string, amount string) (dto.GetEthFeeResponse, responses.ResponseError, error) {
	fee, apiErr := api.GetEthereumFee()
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, apiErr, nil
	}
	balance, apiErr := api.GetEthereumBalance(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, apiErr, nil
	}
	fr, err := CalculateEthBasedFee(balance.Balance, fee.GasPrice, 21000, amount)
	if err != nil {
		return dto.GetEthFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetEthereumClassicFee(address string, amount string) (dto.GetEthFeeResponse, responses.ResponseError, error) {
	fee, apiErr := api.GetEthereumClassicFee()
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, apiErr, nil
	}
	balance, apiErr := api.GetEthereumClassicBalance(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, apiErr, nil
	}
	fr, err := CalculateEthBasedFee(balance.Balance, fee.GasPrice, 21000, amount)
	if err != nil {
		return dto.GetEthFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetTokenFee(address, tokenAddress, amount string) (dto.GetTokenFeeResponse, responses.ResponseError, error) {
	gasLimit, apiErr := api.GetTokenGasLimit(requests.TokenGasLimitRequest{
		TokenAddress: tokenAddress,
		ToAddress:    address,
		Amount:       amount,
	})
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, apiErr, nil
	}
	fee, apiErr := api.GetEthereumFee()
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, apiErr, nil
	}
	balance, apiErr := api.GetEthereumBalance(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, apiErr, nil
	}
	tokenBalance, apiErr := api.GetTokenBalance(address, tokenAddress)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, apiErr, nil
	}
	fr, err := CalculateTokenFee(balance.Balance, tokenBalance.Balance, fee.GasPrice, gasLimit.GasLimit, amount)
	if err != nil {
		return dto.GetTokenFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetWavesFee(address string, amount string) (dto.GetWavesAndStellarFeeResponse, responses.ResponseError, error) {
	balance, apiErr := api.GetWavesBalance(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetWavesAndStellarFeeResponse{}, apiErr, nil
	}
	maxAmount := balance.Balance - 300000
	isEnough := true
	if stringAmountToSatoshi(amount) > maxAmount {
		isEnough = false
		if maxAmount < 0 {
			maxAmount = 0
		}
	}
	return dto.GetWavesAndStellarFeeResponse{
		Balance:                 balance.Balance,
		Fee:                     300000,
		MaxAmountWithOptimalFee: maxAmount,
		IsBadFee:                false,
		IsEnough:                isEnough,
	}, responses.ResponseError{}, nil
}

func GetStellarFee(address string, amount string) (dto.GetWavesAndStellarFeeResponse, responses.ResponseError, error) {
	balance, apiErr := api.GetStellarBalance(address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetWavesAndStellarFeeResponse{}, apiErr, nil
	}
	fr := CalcStellarFee(balance.Balance, amount, 100)
	return fr, responses.ResponseError{}, nil
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
