package api

import (
	"github.com/button-tech/blockchain-fee-service/dto/fee/responses"
)


func GetEthereumBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", nodeUrl, "/eth/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetTokenBalance(address, tokenAddress string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", nodeUrl, "/eth/tokenBalance/"+tokenAddress+"/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetEthereumClassicBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", nodeUrl, "/etc/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetWavesBalance(address string) (responses.WavesBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://nodes.wavesnodes.com", "/addresses/balance/"+address, nil)
	var responseToClient responses.WavesBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetStellarBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", nodeUrl, "/xlm/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetZilliqaBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", nodeUrl, "/zilliqa/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
