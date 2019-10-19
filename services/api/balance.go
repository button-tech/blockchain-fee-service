package api

import (
	"github.com/button-tech/blockchain-fee-service/dto/fee/responses"
)

func GetEthereumBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/eth/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetTokenBalance(address, tokenAddress string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/eth/tokenBalance/"+tokenAddress+"/"+address, nil)
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

func GetWavesBalance(address string) (responses.WavesBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://nodes.wavesnodes.com", "/addresses/balance/"+address, nil)
	var responseToClient responses.WavesBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetStellarBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/xlm/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetZilliqaBalance(address string) (responses.CurrencyBalanceResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/zilliqa/balance/"+address, nil)
	var responseToClient responses.CurrencyBalanceResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
