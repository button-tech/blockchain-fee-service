package feeCalculator

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/dto/fee/requests"
	"dev.azure.com/fee-service/dto/fee/responses"
	"dev.azure.com/fee-service/services/api"
	"sync"
)

type feeCalculator struct {
	CalcFee       func(int, int, int) int
	MinFeePerByte int
	FeePerByte    int
}

func GetBitcoinFee(address string, amount string, receiversCount int) (dto.GetFeeResponse, responses.ResponseError, error) {

	var utxos responses.UtxoResponse
	var apiUtxoErr, apiFeeErr responses.ResponseError
	var feePerByte responses.BitcoinFeeResponse

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		utxos, apiUtxoErr = api.GetUtxo(address, "btc")
	}()

	go func() {
		defer wg.Done()
		feePerByte, apiFeeErr = api.GetBitcoinFee()
	}()
	wg.Wait()

	if apiFeeErr.Error != nil || apiFeeErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiFeeErr, nil
	} else if apiUtxoErr.Error != nil || apiUtxoErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiUtxoErr, nil
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

	var utxos responses.UtxoResponse
	var apiUtxoErr, apiFeeErr responses.ResponseError
	var feePerByte responses.LitecoinFeeResponse

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		utxos, apiUtxoErr = api.GetUtxo(address, "ltc")
	}()

	go func() {
		defer wg.Done()
		feePerByte, apiFeeErr = api.GetLitecoinFee()
	}()
	wg.Wait()

	if apiFeeErr.Error != nil || apiFeeErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiFeeErr, nil
	} else if apiUtxoErr.Error != nil || apiUtxoErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiUtxoErr, nil
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
	utxos, apiErr := api.GetUtxo(address, "bch")
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

	var fee responses.EthereumFeeResponse
	var feeErr, balanceErr responses.ResponseError
	var balance responses.CurrencyBalanceResponse

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fee, feeErr = api.GetEthereumFee()
	}()

	go func() {
		defer wg.Done()
		balance, balanceErr = api.GetEthereumBalance(address)
	}()
	wg.Wait()

	if feeErr.Error != nil || feeErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, feeErr, nil
	} else if balanceErr.Error != nil || balanceErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, balanceErr, nil
	}

	fr, err := CalculateEthBasedFee(balance.Balance, fee.GasPrice, 21000, amount)
	if err != nil {
		return dto.GetEthFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetEthereumClassicFee(address string, amount string) (dto.GetEthFeeResponse, responses.ResponseError, error) {
	var fee responses.EthereumFeeResponse
	var feeErr, balanceErr responses.ResponseError
	var balance responses.CurrencyBalanceResponse

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fee, feeErr = api.GetEthereumClassicFee()
	}()

	go func() {
		defer wg.Done()
		balance, balanceErr = api.GetEthereumClassicBalance(address)
	}()
	wg.Wait()

	if feeErr.Error != nil || feeErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, feeErr, nil
	} else if balanceErr.Error != nil || balanceErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, balanceErr, nil
	}

	fr, err := CalculateEthBasedFee(balance.Balance, fee.GasPrice, 21000, amount)
	if err != nil {
		return dto.GetEthFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetTokenFee(address, tokenAddress, amount string) (dto.GetTokenFeeResponse, responses.ResponseError, error) {
	var gasLimitErr, feeErr, balanceErr, tokenBalanceErr responses.ResponseError
	var gasLimit responses.TokenFeeResponse
	var fee responses.EthereumFeeResponse
	var balance, tokenBalance responses.CurrencyBalanceResponse

	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()
		gasLimit, gasLimitErr = api.GetTokenGasLimit(requests.TokenGasLimitRequest{
			TokenAddress: tokenAddress,
			ToAddress:    address,
			Amount:       amount,
		})
	}()

	go func() {
		defer wg.Done()
		fee, feeErr = api.GetEthereumFee()
	}()

	go func() {
		defer wg.Done()
		balance, balanceErr = api.GetEthereumBalance(address)
	}()

	go func() {
		defer wg.Done()
		tokenBalance, tokenBalanceErr = api.GetTokenBalance(address, tokenAddress)
	}()
	wg.Wait()

	if gasLimitErr.Error != nil || gasLimitErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, gasLimitErr, nil
	} else if feeErr.Error != nil || feeErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, feeErr, nil
	} else if balanceErr.Error != nil || balanceErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, balanceErr, nil
	} else if tokenBalanceErr.Error != nil || tokenBalanceErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, tokenBalanceErr, nil
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
	return dto.GetWavesAndStellarFeeResponse{SharedApiResp: &dto.SharedApiResp{
		Balance:                 uint64(balance.Balance),
		Fee:                     300000,
		MaxAmountWithOptimalFee: uint64(maxAmount),
		IsBadFee:                false,
		IsEnough:                isEnough,
	}}, responses.ResponseError{}, nil
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
