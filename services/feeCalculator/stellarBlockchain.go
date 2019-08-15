package feeCalculator

import (
	"github.com/button-tech/blockchain-fee-service/dto"
	"math"
	"math/big"
)

func CalcStellarFee(balance string, amount string, fee int) dto.GetWavesAndStellarFeeResponse {
	minRequiredBalance := 10000000
	bal := stringAmountToLumens(balance)
	val := stringAmountToLumens(amount)
	activeBalance := bal - minRequiredBalance
	f := &dto.GetWavesAndStellarFeeResponse{SharedApiResp: &dto.SharedApiResp{
		Balance: uint64(bal),
		Fee:     fee,
	}}
	balanceWithoutFee := activeBalance - fee - 1
	if balanceWithoutFee <= 0 {
		f.MaxAmountWithOptimalFee = 0
	} else {
		f.MaxAmountWithOptimalFee = uint64(balanceWithoutFee)
		f.MaxAmount = uint64(balanceWithoutFee)
	}

	if balanceWithoutFee-val >= 0 {
		f.IsEnough = true
	}

	return *f
}

func stringAmountToLumens(amount string) int {
	bigA, _ := new(big.Float).SetString(amount)
	multiplier := new(big.Float).SetFloat64(math.Pow(10, 7))
	bigA.Mul(bigA, multiplier)
	i, _ := bigA.Int64()
	return int(i)
}
