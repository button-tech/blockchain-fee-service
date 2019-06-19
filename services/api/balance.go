package api

import (
	"dev.azure.com/fee-service/dto/fee/responses"
)

func GetBitcoinBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/btc/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetEthereumBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/eth/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetEthereumClassicBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/etc/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
