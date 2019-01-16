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
 * @file eth-getblocktransactioncountbynumber.go
 * @authors:
 *   Junjie CHen <chuckjunjchen@gmail.com>
 * @date 2018
 */

package test

import (
	"testing"
	"time"

	"github.com/matrix/go-AIMan/waiting"
	"github.com/matrix/go-AIMan/manager"
)

func TestGetBlockTransactionCountByNumber(t *testing.T) {

	var connection = manager.Tom_Manager

	blockNumber, err := connection.Man.GetBlockNumber()
	t.Log(blockNumber.Uint64())
	// submit a transaction, wait for the block and there should be 1 tx.
	trans,err := NewTestTrans(connection)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	raw,err := connection.SignTx(trans,"MAN.4BRmmxsC9iPPDyr8CRpRKUcp7GAww")

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

	tx, err := connection.Man.GetTransactionByHash(txID)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	blockNumber1 := tx.BlockNumber.String()

	txCount, err := connection.Man.GetBlockTransactionCountByNumber(blockNumber1)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if txCount.Int64() < 1 {
		t.Error("invalid block transaction count")
		t.FailNow()
	}

	txCount, err = connection.Man.GetBlockTransactionCountByNumber("latest")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

}
