package feeCalculator

import "dev.azure.com/fee-service/dto"

func CalcStellarFee(balance string, amount string, fee int) dto.GetWavesAndStellarFeeResponse {
	minRequiredBalance := 100000000
	bal := stringAmountToSatoshi(balance)
	val := stringAmountToSatoshi(amount)
	activeBalance := bal - minRequiredBalance
	f := &dto.GetWavesAndStellarFeeResponse{SharedApiResp: &dto.SharedApiResp{
		Balance: uint64(bal),
		Fee:     fee,
	}}
	balanceWithoutFee := activeBalance - fee
	if balanceWithoutFee <= 0 {
		f.MaxAmountWithOptimalFee = 0
	} else {
		f.MaxAmountWithOptimalFee = uint64(balanceWithoutFee)
	}

	if balanceWithoutFee-val >= 0 && val >= minRequiredBalance {
		f.IsEnough = true
	}

	return *f
}
