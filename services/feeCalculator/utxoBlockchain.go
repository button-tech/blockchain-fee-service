package feeCalculator

import (
	"dev.azure.com/fee-service/dto"
	"dev.azure.com/fee-service/dto/fee/responses"
	"dev.azure.com/fee-service/services/api"
	"math"
	"math/big"
	"sort"
	"strconv"
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

func calcUtxoFee(utxos []responses.Utxo, amount float64, receiversCount, MinFeePerByte int, calcFee func(int, int, int) int) (dto.GetFeeResponse, responses.ResponseError) {
	totalBalance := calcTotalBalance(utxos)
	if totalBalance == 0 {
		return dto.GetFeeResponse{}, responses.ResponseError{}
	}

	feePerByte, apiErr := api.GetBitcoinFee()
	if apiErr.Error != nil || apiErr.ApiError != nil {
		return dto.GetFeeResponse{}, apiErr
	}

	satoshiAmount, err := floatAmountToSatoshi(amount)
	if err != nil {
		return dto.GetFeeResponse{}, responses.ResponseError{Error: err}
	}

	sortUtxo(utxos)

	ux := utxoBlockchain{
		AllUtxos:      utxos,
		TotalBalance:  totalBalance,
		SatoshiAmount: satoshiAmount,
		CalcFee:       calcBitcoinFee,
		FeePerByte:    feePerByte.HalfHourFee,
		MinFeePerByte: MinFeePerByte,
		Output:        receiversCount,
	}
	ux.setMinimalRequirements()
	ux.Input = ux.MinInputs - 1

	iterationBalance := ux.LastIterationBalance
	for i := ux.MinInputs - 1; i < len(ux.UsefulUtxos)+len(ux.UselessUtxos); i++ {
		ux.Input++
		iterationBalance += utxos[i].Satoshis
		if iterationBalance > satoshiAmount {
			feeWithoutReturningOutput := calcFee(ux.Input, ux.Output, feePerByte.HalfHourFee)
			fee := calcFee(ux.Input, ux.Output+1, feePerByte.HalfHourFee)
			currentValueOneOutput := feeWithoutReturningOutput + satoshiAmount
			currentValueTwoOutputs := fee + satoshiAmount
			isEnoughForMinFee := iterationBalance-satoshiAmount >= calcFee(i+1, ux.Output, ux.MinFeePerByte)
			con0 := iterationBalance < currentValueOneOutput
			con1 := iterationBalance == currentValueOneOutput
			con2 := iterationBalance > currentValueOneOutput && iterationBalance < currentValueTwoOutputs
			con3 := iterationBalance >= currentValueTwoOutputs
			if i > len(ux.UsefulUtxos)-1 && isEnoughForMinFee {
				ux.Fee = iterationBalance - satoshiAmount
				ux.Output = 1
				ux.IsBadFee = true
				ux.IsEnough = true
				break
			}
			if con1 {
				ux.Fee = feeWithoutReturningOutput
				ux.IsEnough = true
				break
			} else if con2 {
				ux.Fee = iterationBalance - satoshiAmount
				ux.IsEnough = true
				break
			} else if con3 {
				ux.Fee = fee
				ux.Output = 2
				if iterationBalance-currentValueTwoOutputs < ux.MinFee {
					ux.Fee += iterationBalance - currentValueTwoOutputs
					ux.Output = 1
				}
				ux.IsEnough = true
				break
			} else if con0 && i == len(utxos)-1 && isEnoughForMinFee {
				ux.Fee = totalBalance - satoshiAmount
				ux.Output = 1
				ux.IsBadFee = true
				ux.IsEnough = true
			}
		}
	}

	return dto.GetFeeResponse{
		Fee:                     ux.Fee,
		Input:                   ux.Input,
		Output:                  ux.Output,
		Balance:                 ux.TotalBalance,
		MaxAmount:               ux.MaxAmount,
		MaxAmountWithOptimalFee: ux.MaxUsefulAmount,
		IsEnough:                ux.IsEnough,
		IsBadFee:                ux.IsBadFee,
	}, responses.ResponseError{}
}

func (ux *utxoBlockchain) setMinimalRequirements() {
	ux.setMinFee()
	ux.setMinInputs()
	ux.setUtxos()
	ux.setMaxAmounts()
}

func (ux *utxoBlockchain) setMaxAmounts() {
	workableBalance := 0
	for _, utxo := range ux.UsefulUtxos {
		workableBalance += utxo.Satoshis
	}
	ux.UsefulBalance = workableBalance
	for _, utxo := range ux.UselessUtxos {
		workableBalance += utxo.Satoshis
	}
	if len(ux.UsefulUtxos) > 0 {
		ux.MaxUsefulAmount = ux.UsefulBalance - ux.CalcFee(len(ux.UsefulUtxos), 1, ux.FeePerByte)
	}
	ux.MaxAmount = workableBalance - ux.CalcFee(len(ux.UsefulUtxos)+len(ux.UselessUtxos), 1, ux.MinFeePerByte)
}

func (ux *utxoBlockchain) setUtxos() {
	ux.UsefulUtxos = ux.AllUtxos
	for i, utxo := range ux.AllUtxos {
		avarageFee := ux.CalcFee(i+1, 1, ux.FeePerByte)
		if utxo.Satoshis <= avarageFee {
			ux.UsefulUtxos = ux.AllUtxos[:i]
			ux.UselessUtxos = ux.AllUtxos[i:]
			break
		}
	}
	ux.setDustUtxo()
}

func (ux *utxoBlockchain) setDustUtxo() {
	for i, utxo := range ux.UselessUtxos {
		if utxo.Satoshis < ux.MinFee {
			ux.DustUtxos = ux.UselessUtxos[i:]
			ux.UselessUtxos = ux.UselessUtxos[:i]
			break
		}
	}
}

func (ux *utxoBlockchain) setMinInputs() {
	iterationBalance := 0
	for _, utxo := range ux.AllUtxos {
		ux.MinInputs++
		iterationBalance += utxo.Satoshis
		if iterationBalance > ux.SatoshiAmount {
			ux.LastIterationBalance = iterationBalance - utxo.Satoshis
			break
		}
	}
}

func (ux *utxoBlockchain) setMinFee() {
	ux.MinFee = ux.CalcFee(1, 1, ux.MinFeePerByte)
}

func calcTotalBalance(utxos []responses.Utxo) int {
	totalBalance := 0
	for _, utxo := range utxos {
		totalBalance += utxo.Satoshis
	}
	return totalBalance
}

func calcBitcoinFee(inputCount, outputCount, feePerByte int) int {
	return (inputCount*148 + outputCount*34 + 10) * feePerByte
}

func floatAmountToSatoshi(amount float64) (int, error) {
	bigA := new(big.Float).SetFloat64(amount)
	multiplier := new(big.Float).SetFloat64(math.Pow(10, 8))
	bigA.Mul(bigA, multiplier)
	return strconv.Atoi(bigA.String())

}

func sortUtxo(utxos []responses.Utxo) {
	sort.Sort(UtxoSorter(utxos))
}

type UtxoSorter []responses.Utxo

func (a UtxoSorter) Len() int           { return len(a) }
func (a UtxoSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UtxoSorter) Less(i, j int) bool { return a[i].Satoshis > a[j].Satoshis }
