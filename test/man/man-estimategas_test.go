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
 * @file eth-estimategas_test.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */

package test

import (
	"testing"

	"math/big"
	"github.com/matrix/go-AIMan/common"
	"github.com/matrix/go-matrix/common/hexutil"
	"encoding/json"
	"github.com/matrix/go-AIMan/manager"
)

func TestEstimateGas(t *testing.T) {


	coinbase, err := manager.Tom_Manager.Man.GetCoinbase()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	transaction := new(common.ManCallArgs)
	//	transaction.Data = "test"
	transaction.From = coinbase
	transaction.To = &coinbase
	transaction.Value = hexutil.Big(*big.NewInt(10))
	transaction.Gas = hexutil.Uint64(40000)

	params := make([]*common.ManCallArgs, 1)

	params[0] = transaction
	buff,_ := json.Marshal(params)
	t.Log(string(buff))
	err = json.Unmarshal(buff,&params)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	gas, err := manager.Tom_Manager.Man.EstimateGas(transaction)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(gas)

}
