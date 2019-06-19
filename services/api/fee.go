package api

import (
	"dev.azure.com/fee-service/dto/fee/requests"
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

func GetEthereumFee() (responses.EthereumFeeResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/eth/gasPrice", nil)
	var responseToClient responses.EthereumFeeResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetEthereumClassicFee() (responses.EthereumFeeResponse, responses.ResponseError) {
	call := apiCall("GET", "https://node.buttonwallet.com", "/etc/gasPrice", nil)
	var responseToClient responses.EthereumFeeResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}

func GetTokenGasLimit(data requests.TokenGasLimitRequest) (responses.TokenFeeResponse, responses.ResponseError) {
	call := apiCall("POST", "https://node.buttonwallet.com", "/eth/estimateGas", data)
	var responseToClient responses.TokenFeeResponse
	errors := call.response(&responseToClient)
	return responseToClient, errors
}
