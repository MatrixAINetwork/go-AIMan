package Deposit

import (
	"testing"
	"math/big"
	"github.com/matrix/go-AIMan/transactions"
	"github.com/matrix/go-AIMan/waiting"
	"github.com/matrix/go-AIMan/AIMan"
	"time"
	"errors"
	"github.com/matrix/go-AIMan/test"
)

func SendDepositTrans(connect *AIMan.AIMan,depAddr string,amount *big.Int,from string,passphrase string)error{
	err := transactions.Unlock(from,passphrase)
	if err != nil {
		return err
	}
	blockNumber, err := connect.Man.GetBlockNumber()
	if err != nil {
		return err
	}
	nonce,err := connect.Man.GetTransactionCount(from,"latest")
	if err != nil {
		return err
	}
	data,err := transactions.MinerDeposit(depAddr)
	if err != nil {
		return err
	}
	trans := transactions.NewTransaction(nonce.Uint64(),transactions.DepositAddr,amount,400000,big.NewInt(18e9),
		data,0,0)
	raw,err := transactions.SignTx(trans,from)
	if err != nil{
		return err
	}
	txID, err := connect.Man.SendRawTransaction(raw)
	if err != nil {
		return err
	}
	ReceiptWait := waiting.NewWaitTxReceipt(connect,txID)
	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connect,blockNumber.Uint64()+10),
		waiting.NewWaitTime(20*time.Second),
		ReceiptWait)
	index := wait3.Waiting()
	if index != 2{
		return errors.New("Time Out")
	}
	if !ReceiptWait.Receipt.Status {
		return errors.New("Contract Run Error")
	}
	return nil
}
func TestSendDepositCheck(t *testing.T){
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount,Man1)
	err := SendDepositTrans(test.Jerry_connection,"MAN.rp9DVBpuVHwPQGYGeLDHCtRnAA3n",
		amount,"MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww","R7c5Rsrj1Q7r4d5fp")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
