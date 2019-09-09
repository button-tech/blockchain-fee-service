package feeCalculator

import (
	"github.com/button-tech/blockchain-fee-service/dto"
	"github.com/button-tech/blockchain-fee-service/dto/fee/responses"
	"math"
	"math/big"
	"sort"
)

type utxoBlockchain struct {
	AllUtxos             []responses.Utxo
	SendingAmount        int
	CalcFee              func(int, int, int) int
	MinFeePerByte        int
	FeePerByte           int
	MinFee               int
	MinInputs            int
	LastIterationBalance int
	UsefulUtxos          []responses.Utxo
	UselessUtxos         []responses.Utxo
	DustUtxos            []responses.Utxo
	MaxAmount            int
	MaxUsefulAmount      int
}

func calcUtxoFee(utxos []responses.Utxo, amount string, receiversCount int, feeCalculator feeCalculator) (dto.GetFeeResponse, responses.ResponseError) {
	var result dto.GetFeeResponse

	totalBalance := calcTotalBalance(utxos)
	if totalBalance == 0 {
		return dto.GetFeeResponse{}, responses.ResponseError{}
	}

	satoshiAmount := stringAmountToSatoshi(amount)

	sortUtxo(utxos)

	ux := utxoBlockchain{
		AllUtxos:      utxos,
		SendingAmount: satoshiAmount,
		CalcFee:       feeCalculator.CalcFee,
		FeePerByte:    feeCalculator.FeePerByte,
		MinFeePerByte: feeCalculator.MinFeePerByte,
	}
	ux.setMinimalRequirements()

	result = dto.GetFeeResponse{SharedApiResp: &dto.SharedApiResp{
		Balance:                 uint64(totalBalance),
		MaxAmount:               uint64(ux.MaxAmount),
		MaxAmountWithOptimalFee: uint64(ux.MaxUsefulAmount),
	},
		Inputs: []responses.Utxo{},
		FeePerByte: feeCalculator.FeePerByte,
	}
	result.Input = ux.MinInputs - 1

	if ux.SendingAmount == 0 || ux.SendingAmount >= totalBalance {
		return result, responses.ResponseError{}
	}

	result.IsEnough = true
	result.Output = receiversCount
	iterationBalance := ux.LastIterationBalance
	maxIterations := len(ux.UsefulUtxos)+len(ux.UselessUtxos)

	for i := ux.MinInputs - 1; i < maxIterations; i++ {
		result.Input++
		iterationBalance += utxos[i].Satoshis

		feeWithoutReturningOutput := ux.CalcFee(result.Input, result.Output, feeCalculator.FeePerByte)
		fee := ux.CalcFee(result.Input, result.Output+1, feeCalculator.FeePerByte)

		currentValueOneOutput := feeWithoutReturningOutput + ux.SendingAmount
		currentValueTwoOutputs := fee + ux.SendingAmount

		isEnoughForMinFee := iterationBalance-ux.SendingAmount >= ux.CalcFee(i+1, result.Output, ux.MinFeePerByte)

		con0 := iterationBalance < currentValueOneOutput
		con1 := iterationBalance == currentValueOneOutput
		con2 := iterationBalance > currentValueOneOutput && iterationBalance < currentValueTwoOutputs
		con3 := iterationBalance == currentValueTwoOutputs
		con4 := iterationBalance > currentValueTwoOutputs

		if con1 || con2 || con3 {
			result.Output = 1
			result.Fee = iterationBalance - ux.SendingAmount
			break
		} else if con4 {
			result.Output = 2
			result.Fee = fee
			break
		} else if (i > len(ux.UsefulUtxos)-1 && isEnoughForMinFee) || (con0 && i == maxIterations-1 && isEnoughForMinFee) {
			result.Fee = iterationBalance - ux.SendingAmount
			result.FeePerByte = ux.MinFeePerByte
			result.Output = 1
			result.IsBadFee = true
			break
		} else if i == maxIterations-1 && !isEnoughForMinFee {
			result.IsEnough = false
			result.Input = 0
			break
		}
	}

	result.Inputs = utxos[:result.Input]

	return result, responses.ResponseError{}
}

func (ux *utxoBlockchain) setMinimalRequirements() {
	ux.setMinFee()
	ux.setUtxos()
	ux.setMinInputs()
	ux.setMaxAmounts()
}

func (ux *utxoBlockchain) setMaxAmounts() {
	workableBalance := 0
	for _, utxo := range ux.UsefulUtxos {
		workableBalance += utxo.Satoshis
	}
	useFulBalance := workableBalance
	for _, utxo := range ux.UselessUtxos {
		workableBalance += utxo.Satoshis
	}
	if len(ux.UsefulUtxos) > 0 {
		ux.MaxUsefulAmount = useFulBalance - ux.CalcFee(len(ux.UsefulUtxos), 1, ux.FeePerByte)
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
	utxos := append(ux.UsefulUtxos, ux.UselessUtxos...)
	for i := 0; i < len(utxos); i++ {
		iterationBalance += utxos[i].Satoshis
		ux.MinInputs++

		if iterationBalance > ux.SendingAmount {
			ux.LastIterationBalance = iterationBalance - utxos[i].Satoshis
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

func stringAmountToSatoshi(amount string) int {
	bigA, _ := new(big.Float).SetString(amount)
	multiplier := new(big.Float).SetFloat64(math.Pow(10, 8))
	bigA.Mul(bigA, multiplier)
	i, _ := bigA.Int64()
	return int(i)
}

func sortUtxo(utxos []responses.Utxo) {
	sort.Sort(UtxoSorter(utxos))
}

type UtxoSorter []responses.Utxo

func (a UtxoSorter) Len() int           { return len(a) }
func (a UtxoSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UtxoSorter) Less(i, j int) bool { return a[i].Satoshis > a[j].Satoshis }
