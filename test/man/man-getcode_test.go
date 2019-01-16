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
 * @file eth-getcode_test.go
 * @authors:
 *   Junjie Chen <chuckjunjchen@gmail.com>
 * @date 2018
 */

package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/matrix/go-AIMan/man/block"

	"github.com/matrix/go-AIMan/dto"
	"github.com/matrix/go-AIMan/common"
	"github.com/matrix/go-AIMan/test"
)

func TestEthGetcode(t *testing.T) {

	content, err := ioutil.ReadFile("../resources/simple-token.json")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	type TruffleContract struct {
		Abi              string `json:"abi"`
		Bytecode         string `json:"bytecode"`
		DeployedBytecode string `json:"deployedBytecode"`
	}

	var unmarshalResponse TruffleContract

	json.Unmarshal(content, &unmarshalResponse)

	var connection = test.Tom_connection
	bytecode := unmarshalResponse.Bytecode
	deployedBytecode := unmarshalResponse.DeployedBytecode

	contract, err := connection.Man.NewContract(unmarshalResponse.Abi)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	transaction := new(common.SendTxArgs1)
	coinbase, err := connection.Man.GetCoinbase()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	transaction.From = coinbase
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

	address := receipt.ContractAddress
	code, err := connection.Man.GetCode(address, block.LATEST)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if deployedBytecode != code {
		t.Error("Contract code not expected")
		t.FailNow()
	}
}
