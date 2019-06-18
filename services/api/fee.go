package api

import (
	"dev.azure.com/fee-service/dto/fee/responses"
)

func GetBitcoinFee() (responses.BitcoinFeeResponse, responses.ResponseError) {
	call := apiCall("GET", "https://bitcoinfees.earn.com", "/api/v1/fees/recommended", nil)
	var responseToClient responses.BitcoinFeeResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetLitecoinFee() (responses.LitecoinFeeResponse, responses.ResponseError) {
	call := apiCall("GET", "https://api.blockcypher.com", "/v1/ltc/main", nil)
	var responseToClient responses.LitecoinFeeResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
