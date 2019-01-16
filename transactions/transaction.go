package transactions

import (
	"math/big"
	"github.com/matrix/go-matrix/core/types"
	"github.com/matrix/go-AIMan/Accounts"
)
func NewTransaction(nonce uint64,to string,amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, typ byte, isEntrustTx byte) *types.Transaction {
	addrTo := Accounts.ManAddressToEthAddress(to)
	tx := types.NewTransaction(nonce, addrTo, amount, gasLimit, gasPrice, data, typ, isEntrustTx)
	tx.Currency = "MAN"
	return tx
}
