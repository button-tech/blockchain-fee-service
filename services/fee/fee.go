package fee

import (
	"dev.azure.com/fee-service/dto/fee/responses"
)

func GetBitcoinFee() (responses.BitcoinFeeResponse, responses.ResponseError) {
	call := apiCall("GET", "https://bitcoinfees.earn.com", "/api/v1/fees/recommended", nil)
	var responseToClient responses.BitcoinFeeResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
