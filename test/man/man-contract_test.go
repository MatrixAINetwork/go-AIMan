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
 * @file contract.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2018
 */

package test

import (
	"encoding/json"
	"fmt"
	"github.com/matrix/go-AIMan/dto"
	"io/ioutil"
	"math/big"
	"testing"
	"github.com/matrix/go-AIMan/common"
	"github.com/matrix/go-AIMan/test"
)

func TestEthContract(t *testing.T) {

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
	//transaction.Gas = big.NewInt(4000000)

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

	t.Log("Contract Address: ", receipt.ContractAddress)

	transaction.To = &receipt.ContractAddress

	result, err := contract.Call(transaction, "name")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if result != nil && err == nil {
		name, _ := result.ToComplexString()
		if name.ToString() != "SimpleToken" {
			t.Errorf(fmt.Sprintf("Name not expected; [Expected %s | Got %s]", "SimpleToken", name.ToString()))
			t.FailNow()
		}
	}

	result, err = contract.Call(transaction, "symbol")
	if result != nil && err == nil {
		symbol, _ := result.ToComplexString()
		if symbol.ToString() != "SIM" {
			t.Errorf("Symbol not expected")
			t.FailNow()
		}
	}

	result, err = contract.Call(transaction, "decimals")
	if result != nil && err == nil {
		decimals, _ := result.ToBigInt()
		if decimals.Int64() != 18 {
			t.Errorf("Decimals not expected")
			t.FailNow()
		}
	}

	bigInt, _ := new(big.Int).SetString("00000000000000000000000000000000000000000000021e19e0c9bab2400000", 16)

	result, err = contract.Call(transaction, "totalSupply")
	if result != nil && err == nil {
		total, _ := result.ToBigInt()
		if total.Cmp(bigInt) != 0 {
			t.Errorf("Total not expected")
			t.FailNow()
		}
	}

	result, err = contract.Call(transaction, "balanceOf", coinbase)
	if result != nil && err == nil {
		balance, _ := result.ToBigInt()
		if balance.Cmp(bigInt) != 0 {
			t.Errorf("Balance not expected")
			t.FailNow()
		}
	}

	hash, err = contract.Send(transaction, "approve", coinbase, big.NewInt(10))
	if err != nil {
		t.Errorf("Can't send approve transaction")
		t.FailNow()
	}

	t.Log(hash)

	reallyBigInt, _ := big.NewInt(0).SetString("20000000000000000000000000000000000000000000000000000000000000000", 16)
	_, err = contract.Send(transaction, "approve", coinbase, reallyBigInt)
	if err == nil {
		t.Errorf("Can't send approve transaction")
		t.FailNow()
	}

}
