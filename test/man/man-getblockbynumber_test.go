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
 * @file eth-getBlockByNumber_test.go
 * @authors:
 *   Jérôme Laurens <jeromelaurens@gmail.com>
 * @date 2017
 */

package test

import (
	"testing"

	"math/big"
	"github.com/matrix/go-AIMan/manager"
)
func TestEthGetGenesisBlock(t *testing.T) {

	var connection = manager.Tom_Manager

	block, err := connection.Man.GetBlockByNumber(big.NewInt(0), false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if block == nil {
		t.Error("Block returned is nil")
		t.FailNow()
	}

	if len(block.Hash) == 0 {
		t.Error("Block not found")
		t.FailNow()
	}

}

func TestEthGetBlockByNumber(t *testing.T) {

	var connection = manager.Tom_Manager

	blockNumber, err := connection.Man.GetBlockNumber()

	block, err := connection.Man.GetBlockByNumber(blockNumber, false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if block == nil {
		t.Error("Block returned is nil")
		t.FailNow()
	}

	if block.Number.ToInt().Int64() == 0 {
		t.Error("Block not found")
		t.FailNow()
	}

}
