package test

import (
	"github.com/matrix/go-AIMan/providers"
	"github.com/matrix/go-AIMan/AIMan"
	"github.com/matrix/go-AIMan/transactions"
	"math/big"
	"github.com/matrix/go-matrix/core/types"
)

var Tom_connection = AIMan.NewAIMan(providers.NewHTTPProvider("api85.matrix.io", 100, false))
var Jerry_connection = AIMan.NewAIMan(providers.NewHTTPProvider("47.105.202.251:8341", 100, false))
var Local_connection = AIMan.NewAIMan(providers.NewHTTPProvider("localhost:8341", 100000, false))
func NewTestTrans(aiMan *AIMan.AIMan)(*types.Transaction,error){
	from := "MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww"
	err := transactions.Unlock(from,"R7c5Rsrj1Q7r4d5fp")
	if err != nil {
		return nil,err
	}
	nonce,err := aiMan.Man.GetTransactionCount(from,"latest")
	if err != nil {
		return nil,err
	}
	trans := transactions.NewTransaction(nonce.Uint64(),from,big.NewInt(1000),200000,big.NewInt(18e9),
		nil,0,0)
	return trans,nil
}