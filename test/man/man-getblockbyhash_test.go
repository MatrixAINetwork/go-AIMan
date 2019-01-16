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
 *    Sigma Prime <sigmaprime.io>
 * @date 2018
 */

package test

import (
	"testing"

	"github.com/matrix/go-AIMan/manager"
)
func TestEthGetBlockByHash(t *testing.T) {

	var connection = manager.Tom_Manager

	blockNumber, err := connection.Man.GetBlockNumber()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	blockByNumber, err := connection.Man.GetBlockByNumber(blockNumber, false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	blockByHash, err := connection.Man.GetBlockByHash(blockByNumber.Hash, false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Ensure it's the same block
	if (blockByNumber.Number.ToInt().Cmp(blockByHash.Number.ToInt())) != 0 ||
		(blockByNumber.Miner != blockByHash.Miner) ||
		(blockByNumber.Hash != blockByHash.Hash) {
		t.Errorf("Not same block returned")
		t.FailNow()
		t.FailNow()
	}

	t.Log(blockByHash.Hash, blockByNumber.Hash)

	_, err = connection.Man.GetBlockByHash("0x1234", false)

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}

	_, err = connection.Man.GetBlockByHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0", false)

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}

	_, err = connection.Man.GetBlockByHash("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0", false)

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}

	blockByHash, err = connection.Man.GetBlockByHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false)

	if err == nil && len(blockByHash.Hash)>0 {
		t.Errorf("Found a block with incorrect hash?")
		t.FailNow()
	}
}

func TestEthGetBlockByHashFull(t *testing.T) {

	var connection = manager.Tom_Manager

	blockNumber, err := connection.Man.GetBlockNumber()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	blockByNumber, err := connection.Man.GetBlockByNumber(blockNumber, true)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	blockByHash, err := connection.Man.GetBlockByHash(blockByNumber.Hash, true)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Ensure it's the same block
	if (blockByNumber.Number.ToInt().Cmp(blockByHash.Number.ToInt())) != 0 ||
		(blockByNumber.Miner != blockByHash.Miner) ||
		(blockByNumber.Hash != blockByHash.Hash) {
		t.Errorf("Not same block returned")
		t.FailNow()
		t.FailNow()
	}

	t.Log(blockByHash.Hash, blockByNumber.Hash)

	_, err = connection.Man.GetBlockByHash("0x1234", true)

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}

	_, err = connection.Man.GetBlockByHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0", true)

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}

	_, err = connection.Man.GetBlockByHash("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0", true)

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}

	blockByHash, err = connection.Man.GetBlockByHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true)

	if err == nil && len(blockByHash.Hash)>0 {
		t.Errorf("Found a block with incorrect hash?")
		t.FailNow()
	}
}
