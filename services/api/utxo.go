package api

import (
	"github.com/button-tech/blockchain-fee-service/dto/fee/responses"
)

func GetUtxo(address, currency string) (responses.UtxoResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com/", currency+"/utxo/"+address, nil)
	var responseToClient responses.UtxoResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
