package feeCalculator

import (
	"github.com/button-tech/blockchain-fee-service/dto"
	"github.com/button-tech/blockchain-fee-service/dto/fee/requests"
	"github.com/button-tech/blockchain-fee-service/dto/fee/responses"
	"github.com/button-tech/blockchain-fee-service/services/api"
	"sync"
)

type Params struct {
	Address        string
	Amount         string
	ReceiversCount int
	Speed          string
	TokenAddress   string
}

type feeCalculator struct {
	CalcFee       func(int, int, int) int
	MinFeePerByte int
	FeePerByte    int
}

func GetBitcoinFee(params *Params) (dto.GetFeeResponse, responses.ResponseError, error) {

	var utxos responses.UtxoResponse
	var apiUtxoErr, apiFeeErr responses.ResponseError
	var feePerByte responses.BitcoinFeeResponse

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		utxos, apiUtxoErr = api.GetUtxo(params.Address, "btc")
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

	feeCounted, _ := speedControl(feePerByte, params.Speed)

	fee, apiErr := calcUtxoFee(utxos.Utxo, params.Amount, params.ReceiversCount, *feeCounted)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}

	return fee, responses.ResponseError{}, nil
}

func GetLitecoinFee(params *Params) (dto.GetFeeResponse, responses.ResponseError, error) {

	var utxos responses.UtxoResponse
	var apiUtxoErr, apiFeeErr responses.ResponseError
	var feePerByte responses.LitecoinFeeResponse

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		utxos, apiUtxoErr = api.GetUtxo(params.Address, "ltc")
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

	feeCounted, _ := speedControl(feePerByte, params.Speed)
	// todo : complete
	feeCounted.MinFeePerByte = 8

	fee, apiErr := calcUtxoFee(utxos.Utxo, params.Amount, params.ReceiversCount, *feeCounted)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	return fee, responses.ResponseError{}, nil
}

func GetBitcoinCashFee(params *Params) (dto.GetFeeResponse, responses.ResponseError, error) {
	utxos, apiErr := api.GetUtxo(params.Address, "bch")
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	fee, apiErr := calcUtxoFee(utxos.Utxo, params.Amount, params.ReceiversCount, feeCalculator{
		CalcFee:       calcBitcoinCashFee,
		MinFeePerByte: 1,
		FeePerByte:    3,
	})
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr, nil
	}
	return fee, responses.ResponseError{}, nil
}

func GetEthereumFee(params *Params) (dto.GetEthFeeResponse, responses.ResponseError, error) {

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
		balance, balanceErr = api.GetEthereumBalance(params.Address)
	}()
	wg.Wait()

	if feeErr.Error != nil || feeErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, feeErr, nil
	} else if balanceErr.Error != nil || balanceErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, balanceErr, nil
	}

	_, gasPrice := speedControl(fee, params.Speed)

	fr, err := CalculateEthBasedFee(balance.Balance, gasPrice, 21000, params.Amount)
	if err != nil {
		return dto.GetEthFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetEthereumClassicFee(params *Params) (dto.GetEthFeeResponse, responses.ResponseError, error) {
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
		balance, balanceErr = api.GetEthereumClassicBalance(params.Address)
	}()
	wg.Wait()

	if feeErr.Error != nil || feeErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, feeErr, nil
	} else if balanceErr.Error != nil || balanceErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, balanceErr, nil
	}

	_, gasPrice := speedControl(fee, params.Speed)

	fr, err := CalculateEthBasedFee(balance.Balance, gasPrice, 21000, params.Amount)
	if err != nil {
		return dto.GetEthFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetTokenFee(params *Params) (dto.GetTokenFeeResponse, responses.ResponseError, error) {
	var gasLimitErr, feeErr, balanceErr, tokenBalanceErr responses.ResponseError
	var gasLimit responses.TokenFeeResponse
	var fee responses.EthereumFeeResponse
	var balance, tokenBalance responses.CurrencyBalanceResponse

	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()
		gasLimit, gasLimitErr = api.GetTokenGasLimit(requests.TokenGasLimitRequest{
			TokenAddress: params.TokenAddress,
			ToAddress:    params.Address,
			Amount:       params.Amount,
		})
	}()

	go func() {
		defer wg.Done()
		fee, feeErr = api.GetEthereumFee()
	}()

	go func() {
		defer wg.Done()
		balance, balanceErr = api.GetEthereumBalance(params.Address)
	}()

	go func() {
		defer wg.Done()
		tokenBalance, tokenBalanceErr = api.GetTokenBalance(params.Address, params.TokenAddress)
		if tokenBalance.Balance == "" {
			tokenBalance.Balance = "0"
		}
	}()
	wg.Wait()

	if gasLimitErr.Error != nil || gasLimitErr.ApiError != nil {
		// May not be enough ethereum for transaction fee
		gasLimit = responses.TokenFeeResponse{
			GasLimit: 38000,
		}
	} else if feeErr.Error != nil || feeErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, feeErr, nil
	} else if balanceErr.Error != nil || balanceErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, balanceErr, nil
	} else if tokenBalanceErr.Error != nil || tokenBalanceErr.ApiError != nil {
		return dto.GetTokenFeeResponse{}, tokenBalanceErr, nil
	}

	_, gasPrice := speedControl(fee, params.Speed)

	fr, err := CalculateTokenFee(balance.Balance, tokenBalance.Balance, gasPrice, gasLimit.GasLimit, params.Amount)
	if err != nil {
		return dto.GetTokenFeeResponse{}, responses.ResponseError{}, err
	}
	return fr, responses.ResponseError{}, nil
}

func GetWavesFee(params *Params) (dto.GetWavesAndStellarFeeResponse, responses.ResponseError, error) {
	balance, apiErr := api.GetWavesBalance(params.Address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetWavesAndStellarFeeResponse{}, apiErr, nil
	}
	maxAmount := balance.Balance - 300000
	isEnough := true
	if stringAmountToSatoshi(params.Amount) > maxAmount {
		isEnough = false
		if maxAmount < 0 {
			maxAmount = 0
		}
	}
	return dto.GetWavesAndStellarFeeResponse{SharedApiResp: &dto.SharedApiResp{
		Balance:                 uint64(balance.Balance),
		Fee:                     300000,
		MaxAmountWithOptimalFee: uint64(maxAmount),
		MaxAmount:               uint64(maxAmount),
		IsBadFee:                false,
		IsEnough:                isEnough,
	}}, responses.ResponseError{}, nil
}

func GetStellarFee(params *Params) (dto.GetWavesAndStellarFeeResponse, responses.ResponseError, error) {
	balance, apiErr := api.GetStellarBalance(params.Address)
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetWavesAndStellarFeeResponse{}, apiErr, nil
	}
	fr := CalcStellarFee(balance.Balance, params.Amount, 100)
	return fr, responses.ResponseError{}, nil
}

func GetZilliqaFee(address string, amount string) (dto.GetEthFeeResponse, responses.ResponseError, error) {

	var balanceErr responses.ResponseError
	var balance responses.CurrencyBalanceResponse

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		balance, balanceErr = api.GetZilliqaBalance(address)
	}()
	wg.Wait()

	if balanceErr.Error != nil || balanceErr.ApiError != nil {
		return dto.GetEthFeeResponse{}, balanceErr, nil
	}

	fr, err := CalculateZilliqaFee(balance.Balance, 2000000000, 1, amount)
	if err != nil {
		return dto.GetEthFeeResponse{}, responses.ResponseError{}, err
	}
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

func speedControl(t interface{}, speed string) (*feeCalculator, int) {
	var (
		f        feeCalculator
		gasPrice int
	)

	switch t.(type) {
	case responses.BitcoinFeeResponse:
		typed := t.(responses.BitcoinFeeResponse)
		f.CalcFee = calcBitcoinFee
		switch speed {
		case "slow":
			f.FeePerByte = typed.HourFee
		case "average":
			f.FeePerByte = typed.HalfHourFee
		case "fast":
			f.FeePerByte = typed.FastestFee
		default:
			f.FeePerByte = typed.HourFee
		}
		if f.FeePerByte <= 10 {
			f.MinFeePerByte = f.FeePerByte - 1
		} else {
			f.MinFeePerByte = 10
		}
	case responses.LitecoinFeeResponse:
		typed := t.(responses.LitecoinFeeResponse)
		f.CalcFee = calcLitecoinFee
		switch speed {
		case "slow":
			f.FeePerByte = typed.LowFeePerKb
		case "average":
			f.FeePerByte = typed.MediumFeePerKb
		case "fast":
			f.FeePerByte = typed.HighFeePerKb
		default:
			f.FeePerByte = typed.MediumFeePerKb
		}
		f.FeePerByte = f.FeePerByte / 1024
		if f.FeePerByte <= 8 {
			f.MinFeePerByte = f.FeePerByte - 1
		} else {
			f.MinFeePerByte = 8
		}
	case responses.EthereumFeeResponse:
		typed := t.(responses.EthereumFeeResponse)
		switch speed {
		case "slow":
			gasPrice = typed.GasPrice * 10 / 8
		case "average":
			gasPrice = typed.GasPrice
		case "fast":
			gasPrice = typed.GasPrice + (typed.GasPrice * 10 / 5)
		default:
			gasPrice = typed.GasPrice
		}
	}

	return &f, gasPrice
}
