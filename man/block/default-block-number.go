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
 * @file default-block-number.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */

package block

// NUMBER - An integer block number
//func NUMBER(blocknumber *big.Int) string {
//	return utils.IntToHex(blocknumber)
//}

const (
	// EARLIEST - Earliest block
	EARLIEST string = "earliest"
	// LATEST - latest block
	LATEST string = "latest"
	// PENDING - Pending block
	PENDING string = "pending"
)
