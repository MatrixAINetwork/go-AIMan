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
 * @file eth-gettransactioncount_test.go
 * @authors:
 * 		Sigma Prime <sigmaprime.io>
 * @date 2017
 */

package test

import (
	"fmt"
	"github.com/matrix/go-AIMan/man/block"
	"math/big"
	"testing"
	"time"
	"github.com/matrix/go-AIMan/common"
	"github.com/matrix/go-AIMan/manager"
)

func TestEthGetTransactionCount(t *testing.T) {

	var connection = manager.Tom_Manager

	coinbase, _ := connection.Man.GetCoinbase()

	count, err := connection.Man.GetTransactionCount(coinbase, block.LATEST)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	countTwo, err := connection.Man.GetTransactionCount(coinbase, block.LATEST)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// count should not change
	if count.Cmp(countTwo) != 0 {
		t.Errorf("Count incorrect, changed between calls")
		t.FailNow()
	}
	// send a transaction and the count should increase

	t.Log("Starting Count:", count)
	transaction := new(common.SendTxArgs1)
	transaction.From = coinbase
	transaction.To = &coinbase
	transaction.SetValue(big.NewInt(0).Mul(big.NewInt(500), big.NewInt(1E18)))
//	transaction.Gas = big.NewInt(40000)
	transaction.SetData([]byte("p2p transaction"))

	_, err = connection.Man.SendTransaction(transaction,"xxx")

	if err != nil {
		t.Errorf("Failed to send tx")
		t.FailNow()
	}

	time.Sleep(time.Second)

	newCount, err := connection.Man.GetTransactionCount(coinbase, block.LATEST)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if newCount.Int64() != (countTwo.Int64() + 1) {
		t.Errorf(fmt.Sprintf("Incorrect count retrieved; [Expected %d | Got %d]", countTwo.Int64()+1, newCount))
		t.FailNow()
	}

	t.Log("Final Count: ", newCount)
}
