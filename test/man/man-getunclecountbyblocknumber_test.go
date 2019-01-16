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
 * @file eth-getunclecountbyblocknumber_test.go
 * @authors:
 * 		Sigma Prime <sigmaprime.io>
 * @date 2018
 */

package test

import (
	"math/big"
	"testing"
	"github.com/matrix/go-AIMan/test"
)

func TestGetUncleCountByBlockNumber(t *testing.T) {

	var connection = test.Tom_connection

	blockNumber, err := connection.Man.GetBlockNumber()

	uncleByNumber, err := connection.Man.GetUncleCountByBlockNumber(blockNumber)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(uncleByNumber.Int64())

	if uncleByNumber.Int64() != 0 {
		t.Errorf("Returned uncle for block with no uncle.")
		t.FailNow()
	}

	// should return err with negative number?
	uncleByNumber, err = connection.Man.GetUncleCountByBlockNumber(big.NewInt(-1))

	if err == nil {
		t.Error(err)
		t.FailNow()
	}
}
