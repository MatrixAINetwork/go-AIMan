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
 * @file aiMan.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */

package AIMan

import (
	"github.com/matrix/go-AIMan/dto"
	"github.com/matrix/go-AIMan/man"
	"github.com/matrix/go-AIMan/net"
	"github.com/matrix/go-AIMan/personal"
	"github.com/matrix/go-AIMan/providers"
	"github.com/matrix/go-AIMan/utils"
)

// Coin - Ethereum value unity value
const (
	Coin float64 = 1000000000000000000
)

// AIMan - The AIMan Module
type AIMan struct {
	Provider providers.ProviderInterface
	Man      *man.Man
	Net      *net.Net
	Personal *personal.Personal
	Utils    *utils.Utils
}

// NewAIMan - AIMan Module constructor to set the default provider, Eth, Net and Personal
func NewAIMan(provider providers.ProviderInterface) *AIMan {
	aiMan := new(AIMan)
	aiMan.Provider = provider
	aiMan.Man = man.NewMan(provider)
	aiMan.Net = net.NewNet(provider)
	aiMan.Personal = personal.NewPersonal(provider)
	aiMan.Utils = utils.NewUtils(provider)
	return aiMan
}

// ClientVersion - Returns the current client version.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#web3_clientversion
// Parameters:
//    - none
// Returns:
// 	  - String - The current client version
func (web AIMan) ClientVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := web.Provider.SendRequest(pointer, "web3_clientVersion", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
