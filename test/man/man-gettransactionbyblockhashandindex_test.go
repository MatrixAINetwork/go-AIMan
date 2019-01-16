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
 * @file eth-gettransactionbyblockhashandindex_test.go
 * @authors:
 *      Sigma Prime <sigmaprime.io>
 * @date 2018
 */
package test

import (
	"math/big"
	"testing"
	"time"
	"github.com/matrix/go-AIMan/test"
	"github.com/matrix/go-AIMan/transactions"
	"github.com/matrix/go-AIMan/waiting"
)

func TestGetTransactionByBlockHashAndIndex(t *testing.T) {

	var connection = test.Tom_connection

	blockNumber, err := connection.Man.GetBlockNumber()
	t.Log(blockNumber.Uint64())
	// submit a transaction, wait for the block and there should be 1 tx.
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
	index := wait3.Waiting()
	t.Log(index)

	txFromHash, err := connection.Man.GetTransactionByHash(txID)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	tx, err := connection.Man.GetTransactionByBlockHashAndIndex(txFromHash.BlockHash.String(), uint64(txFromHash.TransactionIndex))

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	from := "MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww"
	if tx.From != from || *tx.To != from || trans.Value().Cmp((*big.Int)(tx.Value)) != 0 || tx.Hash.String() != txID {
		t.Errorf("Incorrect transaction from hash and index")
		t.FailNow()
	}
	// test removing the 0x

	tx, err = connection.Man.GetTransactionByBlockHashAndIndex(txFromHash.BlockHash.String(), uint64(txFromHash.TransactionIndex))

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if tx.From != from || *tx.To != from || trans.Value().Cmp((*big.Int)(tx.Value)) != 0 || tx.Hash.String() != txID {
		t.Errorf("Incorrect transaction from hash and index")
		t.FailNow()
	}
}
