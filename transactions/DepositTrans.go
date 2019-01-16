package transactions

import (
	"strings"
	"github.com/matrix/go-matrix/accounts/abi"
	"github.com/matrix/go-matrix/base58"
)

var (
	DepositDef = ` [{"constant": true,"inputs": [],"name": "getDepositList","outputs": [{"name": "","type": "address[]"}],"payable": false,"stateMutability": "view","type": "function"},
			{"constant": true,"inputs": [{"name": "addr","type": "address"}],"name": "getDepositInfo","outputs": [{"name": "","type": "uint256"},{"name": "","type": "address"},{"name": "","type": "uint256"}, {"name": "","type": "uint256"}],"payable": false,"stateMutability": "view","type": "function"},
    		{"constant": false,"inputs": [{"name": "address","type": "address"}],"name": "valiDeposit","outputs": [],"payable": true,"stateMutability": "payable","type": "function"},
    		{"constant": false,"inputs": [{"name": "address","type": "address"}],"name": "minerDeposit","outputs": [],"payable": true,"stateMutability": "payable","type": "function"},
    		{"constant": false,"inputs": [],"name": "withdraw","outputs": [],"payable": false,"stateMutability": "nonpayable","type": "function"},
    		{"constant": false,"inputs": [],"name": "refund","outputs": [],"payable": false,"stateMutability": "nonpayable","type": "function"},
			{"constant": false,"inputs": [{"name": "addr","type": "address"}],"name": "interestAdd","outputs": [],"payable": true,"stateMutability": "payable","type": "function"},
			{"constant": false,"inputs": [{"name": "addr","type": "address"}],"name": "getinterest","outputs": [],"payable": false,"stateMutability": "payable","type": "function"}]`

	DepContract,abiErr = abi.JSON(strings.NewReader(DepositDef))
	DepositAddr = "MAN.1111111111111111111B8"
)
func ValiDeposit(address string)([]byte,error){
	addr := base58.Base58DecodeToAddress(address)
	return DepContract.Pack("valiDeposit",addr)
}
func MinerDeposit(address string)([]byte,error){
	addr := base58.Base58DecodeToAddress(address)
	return DepContract.Pack("minerDeposit",addr)
}
func Withdraw()([]byte,error){
	return DepContract.Pack("withdraw")
}
func Refund()([]byte,error){
	return DepContract.Pack("refund")
}