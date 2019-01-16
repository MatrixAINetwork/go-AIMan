/********************************************************************************
   This file is part of go-AIMan.
   go-AIMan is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-AIMan is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-AIMan.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file eth-getblocktransactioncountbyhash_test.go
 * @authors:
 *   Sigma Prime <sigmaprime.io>
 * @date 2018
 */

package test

import (
	"testing"
	"time"
	"github.com/matrix/go-AIMan/test"
	"github.com/matrix/go-AIMan/transactions"
	"github.com/matrix/go-AIMan/waiting"
	"github.com/matrix/go-AIMan/AIMan"
	"math/big"
	"errors"
)
func SendTransactionCheck(connect *AIMan.AIMan,to string,amount *big.Int,from string,passphrase string)error{
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
	trans := transactions.NewTransaction(nonce.Uint64(),to,amount,200000,big.NewInt(18e9),
		nil,0,0)
	raw,err := transactions.SignTx(trans,from)
	if err != nil{
		return err
	}
	txID, err := connect.Man.SendRawTransaction(raw)
	if err != nil {
		return err
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connect,blockNumber.Uint64()+10),
		waiting.NewWaitTime(20*time.Second),
		waiting.NewWaitTxReceipt(connect,txID))
	index := wait3.Waiting()
	if index != 2{
		return errors.New("Time Out")
	}
	return nil
}
func TestSendTransactionCheck(t *testing.T){
	err := SendTransactionCheck(test.Jerry_connection,"MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww",
		big.NewInt(1000),"MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww","R7c5Rsrj1Q7r4d5fp")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func SendBatchTransactionCheck(connect *AIMan.AIMan,to string,amount *big.Int,from string,passphrase string)error{
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
	txHashAry:=make([]string,0,100)
	for i:=uint64(0);i<100;i++{
		trans := transactions.NewTransaction(nonce.Uint64()+i,to,amount,200000,big.NewInt(18e9),
			nil,0,0)
		raw,err := transactions.SignTx(trans,from)
		if err != nil{
			return err
		}
		txID, err := connect.Man.SendRawTransaction(raw)
		if err != nil {
			return err
		}
		txHashAry = append(txHashAry, txID)
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connect,blockNumber.Uint64()+10),
		waiting.NewWaitTime(60*time.Second),
		waiting.NewWaitTxReceiptAry(connect,txHashAry))
	index := wait3.Waiting()
	if index == 1{
		return errors.New("Time Out")
	}
	if index == 0{
		return errors.New("Exceed Block Height")
	}
	return nil
}
func TestBatchSendTransactionCheck(t *testing.T){
	err := SendBatchTransactionCheck(test.Jerry_connection,"MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww",
		big.NewInt(1000),"MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww","R7c5Rsrj1Q7r4d5fp")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestGetBlockTransactionCountByHash(t *testing.T) {

	var connection = test.Tom_connection

	blockNumber, err := connection.Man.GetBlockNumber()

	t.Log(blockNumber.Uint64())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	wait2 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connection,blockNumber.Uint64()+3),
		waiting.NewWaitTime(100*time.Second),
			waiting.NewWaitTxReceipt(connection,"0xb66e18011fc3cdd10d46d9bdd9ac68b4da6a45283759216336ace3f7aff9a150"))
	index := wait2.Waiting()
	t.Log(index)
	blockNumber, err = connection.Man.GetBlockNumber()
	t.Log(blockNumber.Uint64())
	block, err := connection.Man.GetBlockByNumber(blockNumber, false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	txCount, err := connection.Man.GetBlockTransactionCountByHash(block.Hash)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	trans,err := test.NewTestTrans(connection)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	raw,err := transactions.SignTx(trans,"MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww")
	if err != nil{
		t.Error(err)
		t.FailNow()
	}
	txID, err := connection.Man.SendRawTransaction(raw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(txID)
	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connection,blockNumber.Uint64()+3),
		waiting.NewWaitTime(100*time.Second),
		waiting.NewWaitTxReceipt(connection,txID))
	index = wait3.Waiting()
	t.Log(index)
	tx, err := connection.Man.GetTransactionByHash(txID)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	txCount, err = connection.Man.GetBlockTransactionCountByHash(tx.BlockHash.String())

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if txCount.Int64() < 1 {
		t.Error("invalid block transaction count")
		t.FailNow()
	}
	txReceipt, err := connection.Man.GetTransactionReceipt(txID)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(txReceipt)

}
