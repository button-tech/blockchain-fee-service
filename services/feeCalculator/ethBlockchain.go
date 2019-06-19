package feeCalculator

import (
	"dev.azure.com/fee-service/dto"
	"math"
	"math/big"
)

func CalculateEthBasedFee(balance, gasPrice, gas int, amount string) (dto.GetEthFeeResponse, error) {
	wei := stringAmountToWei(amount)

	bigBalance := IntToBigInt(balance)
	if bigBalance.Cmp(wei) < 0 {
		return dto.GetEthFeeResponse{}, nil
	}
	fr := dto.GetEthFeeResponse{
		Balance: balance,
		Gas:     gas,
	}

	bigGasPrice := IntToBigInt(gasPrice)
	bigGas := IntToBigInt(gas)
	defaultFee := Mul(bigGasPrice, bigGas)
	bigOptimalGasPriceNotDivided := Mul(bigGasPrice, IntToBigInt(6))
	bigOptimalGasPrice := Div(bigOptimalGasPriceNotDivided, IntToBigInt(5))
	optimalFee := Mul(bigOptimalGasPrice, bigGas)

	defaultSendingAmount := Add(wei, defaultFee)
	optimalSendingAmount := Add(wei, optimalFee)

	con1 := bigBalance.Cmp(defaultSendingAmount) < 0
	con2 := bigBalance.Cmp(defaultSendingAmount) >= 0 && bigBalance.Cmp(optimalSendingAmount) < 0
	con3 := bigBalance.Cmp(optimalSendingAmount) >= 0

	fr.MaxAmount = int(Sub(bigBalance, defaultFee).Int64())
	fr.MaxAmountWithOptimalFee = int(Sub(bigBalance, optimalFee).Int64())

	if con1 {
		fr.Fee = int(defaultFee.Int64())
		fr.GasPrice = gasPrice
	} else if con2 {
		fr.Fee = int(defaultFee.Int64())
		fr.IsEnough = true
		fr.IsBadFee = true
		fr.GasPrice = gasPrice
	} else if con3 {
		fr.Fee = int(optimalFee.Int64())
		fr.IsEnough = true
		fr.GasPrice = int(bigOptimalGasPrice.Int64())
	}

	return fr, nil
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
