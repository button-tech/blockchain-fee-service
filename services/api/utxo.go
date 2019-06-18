package api

import (
	"dev.azure.com/fee-service/dto/fee/responses"
)

func GetBitcoinUtxo(address string) (responses.BitcoinUtxoResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/btc/utxo/"+address, nil)
	var responseToClient responses.BitcoinUtxoResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetLitecoinUtxo(address string) (responses.BitcoinUtxoResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/ltc/utxo/"+address, nil)
	var responseToClient responses.BitcoinUtxoResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetBitcoinCashUtxo(address string) (responses.BitcoinUtxoResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/bch/utxo/"+address, nil)
	var responseToClient responses.BitcoinUtxoResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
