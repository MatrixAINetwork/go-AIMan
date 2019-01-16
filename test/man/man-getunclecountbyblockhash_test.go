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
 * @file eth-getunclecountbyblockhash_test.go
 * @authors:
 * 		Sigma Prime <sigmaprime.io>
 * @date 2018
 */

package test

import (
	"testing"
	"github.com/matrix/go-AIMan/manager"
)

func TestGetUncleCountByBlockHash(t *testing.T) {

	var connection = manager.Tom_Manager

	blockNumber, err := connection.Man.GetBlockNumber()

	blockByNumber, err := connection.Man.GetBlockByNumber(blockNumber, false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	uncleByHash, err := connection.Man.GetUncleCountByBlockHash(blockByNumber.Hash)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(uncleByHash.Int64())

	if uncleByHash.Int64() != 0 {
		t.Errorf("Returned uncle for block with no uncle.")
		t.FailNow()
	}

	_, err = connection.Man.GetUncleCountByBlockHash("0x1234")

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}

	_, err = connection.Man.GetUncleCountByBlockHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0")

	if err == nil {
		t.Errorf("Invalid hash not rejected")
		t.FailNow()
	}
}
