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
 * @file personal-gettransactionbyhash_test.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */
package test

import (
	"math/big"
	"testing"
	"time"
	"github.com/matrix/go-AIMan/common"
	"github.com/matrix/go-AIMan/test"
)

func TestGetTransactionByHash(t *testing.T) {

	var connection = test.Tom_connection

	coinbase, err := connection.Man.GetCoinbase()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	transaction := new(common.SendTxArgs1)
	transaction.From = coinbase
	transaction.To = &coinbase
	transaction.SetValue(big.NewInt(10))
//	transaction.Gas = big.NewInt(40000)

	txID, err := connection.Man.SendTransaction(transaction,"xxx")

	// Wait for a block
	time.Sleep(time.Second)

	if err != nil {
		t.Errorf("Failed SendTransaction")
		t.Error(err)
		t.FailNow()
	}

	time.Sleep(time.Second)

	tx, err := connection.Man.GetTransactionByHash(txID)

	if err != nil {
		t.Errorf("Failed GetTransactionByHash")
		t.Error(err)
		t.FailNow()
	}

	t.Log(tx.BlockNumber)

}
