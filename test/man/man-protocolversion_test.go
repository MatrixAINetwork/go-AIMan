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
 * @file eth-protocolversion_test.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */

package test

import (
	"fmt"
	"testing"

	"github.com/matrix/go-AIMan/manager"
)

func TestEthGetProtocolVersion(t *testing.T) {

	var connection = manager.Tom_Manager

	version, err := connection.Man.GetProtocolVersion()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(version)

}
