package helpers

import (
	"github.com/matrix/go-AIMan/dto"
	"math/big"
)

func GetTestTx(to string, from string, amount *big.Int) *dto.TransactionParameters {

	Tx := new(dto.TransactionParameters)

	transaction := new(dto.TransactionParameters)
	transaction.From = from
	transaction.To = to
	transaction.Value = amount
	transaction.Gas = big.NewInt(40000)

	return Tx
}
