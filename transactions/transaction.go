package transactions

import (
	"github.com/MatrixAINetwork/go-AIMan/common"
	"github.com/MatrixAINetwork/go-matrix/common/hexutil"
	"math/big"
	"encoding/hex"
	"encoding/json"
)

func NewTransaction(nonce uint64, to string, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, typ byte, isEntrustTx byte, CommitTime uint64) *common.SendTxArgs1 {
	currency := "MAN"
	return &common.SendTxArgs1{To: &to, Gas: (*hexutil.Uint64)(&gasLimit), GasPrice: (*hexutil.Big)(gasPrice),
		Value: (*hexutil.Big)(amount), Nonce: (*hexutil.Uint64)(&nonce), Data: (*hexutil.Bytes)(&data), Currency: &currency, TxType: typ, IsEntrustTx: isEntrustTx, CommitTime: CommitTime}
}
func NewTransactions(nonce uint64, to string, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, typ byte, isEntrustTx byte, CommitTime uint64, extrato []*common.ExtraTo_Mx1) *common.SendTxArgs1 {
	currency := "MAN"
	return &common.SendTxArgs1{To: &to, Gas: (*hexutil.Uint64)(&gasLimit), GasPrice: (*hexutil.Big)(gasPrice),
		Value: (*hexutil.Big)(amount), Nonce: (*hexutil.Uint64)(&nonce), Data: (*hexutil.Bytes)(&data), Currency: &currency, TxType: typ, IsEntrustTx: isEntrustTx, CommitTime: CommitTime, ExtraTo: extrato}
}

func SendTxArgs1ToString(sendtx *common.SendTxArgs1) (string,error) {
	endata,err := json.Marshal(*sendtx)
	if err != nil{
		return "",err
	}
	return hex.EncodeToString(endata),nil
}

func StringToSendTxArgs1(str string) (sendtx *common.SendTxArgs1,err error) {
	data,err := hex.DecodeString(str)
	if err != nil{
		return nil,err
	}

	err = json.Unmarshal(data,&sendtx)
	return sendtx,err
}

//func NewContractTransaction(nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, typ byte, isEntrustTx byte) *common.SendTxArgs1 {
//	//tx:=types.NewContractCreation(nonce, amount, gasLimit, gasPrice, data, nil,nil,nil,typ,isEntrustTx,"MAN",0)
//	//return nil
//	currency := "MAN"
//	return &common.SendTxArgs1{Gas: (*hexutil.Uint64)(&gasLimit), GasPrice: (*hexutil.Big)(gasPrice),
//		Value: (*hexutil.Big)(amount), Nonce: (*hexutil.Uint64)(&nonce), Data: (*hexutil.Bytes)(&data), Currency: &currency, TxType: typ, IsEntrustTx: isEntrustTx}
//}
