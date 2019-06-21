package feeCalculator

import (
	"github.com/button-tech/blockchain-fee-service/dto"
	"github.com/button-tech/blockchain-fee-service/dto/errors"
	"math"
	"math/big"
)

func CalculateEthBasedFee(balance string, gasPrice, gas int, amount string) (dto.GetEthFeeResponse, error) {
	wei := stringAmountToWei(amount)
	bal, ok := StringToBigInt(balance)
	if !ok {
		return dto.GetEthFeeResponse{}, errors.CustomError("failed to parse String to big.Int")
	}

	bigBalance := bal

	fr := &dto.GetEthFeeResponse{SharedApiResp: &dto.SharedApiResp{
		Balance: bal.Uint64(),
	}, SharedEthBasedResp: &dto.SharedEthBasedResp{}}

	bigGasPrice := IntToBigInt(gasPrice)
	bigGas := IntToBigInt(gas)

	fr.GasPrice = bigGasPrice.Uint64()
	fr.Gas = bigGas.Uint64()

	defaultFee := Mul(bigGasPrice, bigGas)
	bigOptimalGasPriceNotDivided := Mul(bigGasPrice, IntToBigInt(6))
	bigOptimalGasPrice := Div(bigOptimalGasPriceNotDivided, IntToBigInt(5))
	optimalFee := Mul(bigOptimalGasPrice, bigGas)

	if bigBalance.Cmp(wei) < 0 {
		fr.Fee = int(defaultFee.Int64())
		fr.Gas = bigOptimalGasPrice.Uint64()
		fr.GasPrice = bigOptimalGasPrice.Uint64()
		return *fr, nil
	}

	defaultSendingAmount := Add(wei, defaultFee)
	optimalSendingAmount := Add(wei, optimalFee)

	con1 := bigBalance.Cmp(defaultSendingAmount) < 0
	con2 := bigBalance.Cmp(defaultSendingAmount) >= 0 && bigBalance.Cmp(optimalSendingAmount) < 0
	con3 := bigBalance.Cmp(optimalSendingAmount) >= 0

	maxAmount := Sub(bigBalance, defaultFee).Uint64()
	maxAmountWithOptimalFee := Sub(bigBalance, optimalFee).Uint64()

	if maxAmount > 0 {
		fr.MaxAmount = maxAmount
	}

	if maxAmountWithOptimalFee > 0 {
		fr.MaxAmountWithOptimalFee = maxAmountWithOptimalFee
	}

	if con1 {
		fr.Fee = int(defaultFee.Int64())
	} else if con2 {
		fr.Fee = int(defaultFee.Int64())
		fr.IsEnough = true
		fr.IsBadFee = true
	} else if con3 {
		fr.Fee = int(optimalFee.Int64())
		fr.IsEnough = true
		fr.GasPrice = bigOptimalGasPrice.Uint64()
	}

	return *fr, nil
}

func CalculateTokenFee(ethBalance, tokenBalance string, gasPrice, gas int, amount string) (dto.GetTokenFeeResponse, error) {
	ethBal, ok := StringToBigInt(ethBalance)
	if !ok {
		return dto.GetTokenFeeResponse{}, errors.CustomError("failed to parse String to big.Int")
	}
	tokenBal, ok := StringToBigInt(tokenBalance)
	if !ok {
		return dto.GetTokenFeeResponse{}, errors.CustomError("failed to parse String to big.Int")
	}
	tokenVal := stringAmountToWei(amount)
	eth, err := CalculateEthBasedFee(ethBalance, gasPrice, 21000, "0")
	if err != nil {
		return dto.GetTokenFeeResponse{}, err
	}
	f := &dto.GetTokenFeeResponse{SharedApiResp: &dto.SharedApiResp{
		Balance:                 ethBal.Uint64(),
		IsBadFee:                eth.IsBadFee,
		IsEnough:                eth.IsEnough,
		MaxAmountWithOptimalFee: tokenBal.Uint64(),
		Fee:                     eth.Fee,
	}, SharedEthBasedResp: &dto.SharedEthBasedResp{
		GasPrice: eth.GasPrice,
		Gas:      uint64(gas),
	},
		TokenBalance: tokenBal.Uint64(),
	}
	if tokenBal.Cmp(tokenVal) < 0 {
		f.IsEnough = false
	}
	return *f, nil

}

func stringAmountToWei(amount string) *big.Int {
	bigA, _ := new(big.Float).SetString(amount)
	multiplier := new(big.Float).SetFloat64(math.Pow(10, 18))
	bigA.Mul(bigA, multiplier)
	i, _ := bigA.Int64()
	return new(big.Int).SetInt64(i)
}

func IntToBigInt(i int) *big.Int {
	return new(big.Int).SetInt64(int64(i))
}

func StringToBigInt(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 10)
}

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}
