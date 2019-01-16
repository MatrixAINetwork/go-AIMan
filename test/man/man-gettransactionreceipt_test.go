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
 * @file eth-gettransactionreceipt.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2018
 */

package test

import (
	"encoding/json"
	"github.com/matrix/go-AIMan/dto"
	"io/ioutil"
	"math/big"
	"testing"
	"github.com/matrix/go-AIMan/common"
	"github.com/matrix/go-AIMan/test"
)

func TestEthGetTransactionReceipt(t *testing.T) {

	content, err := ioutil.ReadFile("../resources/simple-token.json")

	type TruffleContract struct {
		Abi      string `json:"abi"`
		Bytecode string `json:"bytecode"`
	}

	var unmarshalResponse TruffleContract

	json.Unmarshal(content, &unmarshalResponse)

	var connection = test.Tom_connection
	bytecode := unmarshalResponse.Bytecode
	contract, err := connection.Man.NewContract(unmarshalResponse.Abi)

	transaction := new(common.SendTxArgs1)
	coinbase, err := connection.Man.GetCoinbase()
	transaction.From = coinbase
//	transaction.Gas = big.NewInt(4000000)

	hash, err := contract.Deploy(transaction, bytecode, nil)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	var receipt *dto.TransactionReceipt

	for receipt == nil {
		receipt, err = connection.Man.GetTransactionReceipt(hash)
	}

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

    if len(receipt.ContractAddress) == 0{
        t.Error("No contract address")
        t.FailNow()
    }

    if len(receipt.TransactionHash) == 0{
        t.Error("No transaction hash")
        t.FailNow()
    }

    if receipt.TransactionIndex == nil{
        t.Error("No transaction index")
        t.FailNow()
    }

    if len(receipt.BlockHash) == 0{
        t.Error("No block hash")
        t.FailNow()
    }

    if (receipt.BlockNumber == nil || receipt.BlockNumber.Cmp(big.NewInt(0)) == 0){
        t.Error("No block number")
        t.FailNow()
    }

    if (receipt.Logs == nil || len(receipt.Logs) == 0){
        t.Error("No logs")
        t.FailNow()
    }

    if (!receipt.Status){
        t.Error("False status")
        t.FailNow()
    }


}
