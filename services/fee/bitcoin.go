package fee

import (
	"dev.azure.com/fee-service/dto/fee/responses"
)

func GetBitcoinFee() responses.Response {
	call := apiCall("GET", "https://bitcoinfees.earn.com", "/v2/customers/me", nil)
	var responseToClient responses.BitcoinFeeResponse
	return call.response(responseToClient)
}
