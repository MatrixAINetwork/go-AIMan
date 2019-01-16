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
 * @file eth-sendtransaction_test.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */
package test

import (
	"testing"

	"github.com/matrix/go-AIMan/complex/types"
	"github.com/matrix/go-AIMan/dto"
	"math/big"
	"github.com/matrix/go-AIMan/test"
)

func TestGetTransactionByBlockNumberAndIndex(t *testing.T) {

	var connection = test.Tom_connection

	coinbase, err := connection.Man.GetCoinbase()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	transaction := new(dto.TransactionParameters)
	transaction.From = coinbase
	transaction.To = coinbase
	transaction.Value = big.NewInt(0).Mul(big.NewInt(500), big.NewInt(1E18))
	transaction.Gas = big.NewInt(40000)
	transaction.Data = types.ComplexString("p2p transaction")

	//txID, err := connection.Man.SendTransaction(transaction)

	//t.Log(txID)

	blockNumber, err := connection.Man.GetBlockNumber()

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	tx, err := connection.Man.GetTransactionByBlockNumberAndIndex(blockNumber, big.NewInt(0))

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(tx.Hash)
	t.Log(tx.BlockHash)
	t.Log(tx.BlockNumber)
	t.Log(tx.TransactionIndex)
}
