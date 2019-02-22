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
 * @file man.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */

package man

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/MatrixAINetwork/go-AIMan/common"
	"github.com/MatrixAINetwork/go-AIMan/dto"
	"github.com/MatrixAINetwork/go-AIMan/providers"
	"github.com/MatrixAINetwork/go-AIMan/utils"
	"github.com/MatrixAINetwork/go-matrix/common/hexutil"
	"github.com/MatrixAINetwork/go-matrix/core/types"
	"math/big"
)

// Man - The Man Module
type Man struct {
	provider providers.ProviderInterface
}

// NewMan - Man Module constructor to set the default provider
func NewMan(provider providers.ProviderInterface) *Man {
	man := new(Man)
	man.provider = provider
	return man
}

// GetBlockByNumber - Returns the information about a block requested by number.
// Parameters:
//    - number, QUANTITY - number of block
//    - transactionDetails, bool - indicate if we should have or not the details of the transactions of the block
// Returns:
//    1. Object - A block object, or null when no transaction was found
//    2. error
func (man *Man) GetBlockByNumber(number *big.Int, transactionDetails bool) (*common.Block, error) {

	params := make([]interface{}, 2)
	params[0] = utils.IntToHex(number)
	params[1] = transactionDetails

	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "man_getBlockByNumber", params)

	if err != nil {
		return nil, err
	}
	return common.UnmarshalBlock(pointer.Result, transactionDetails)
}

// GetBlockNumber - Returns the number of most recent block.
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the current block number the client is on.
func (man *Man) GetBlockNumber() (*big.Int, error) {

	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "man_blockNumber", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetTransactionCount -  Returns the number of transactions sent from an address.
// Parameters:
//    - DATA, 20 Bytes - address to check for balance.
// Returns:
// 	  - QUANTITY - integer of the number of transactions sent from this address
func (man *Man) GetTransactionCount(address string, defaultBlockParameter string) (*big.Int, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "man_getTransactionCount", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

func (man *Man) SendRawTransaction(rawTx interface{}) (string, error) {

	params := make([]interface{}, 1)
	params[0] = rawTx

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(&pointer, "man_sendRawTransaction", params)

	if err != nil {
		return "", err
	}
	return pointer.ToString()
}



// GetTransactionReceipt - Returns compiled solidity code.
// Parameters:
//    1. DATA, 32 Bytes - hash of a transaction.
// Returns:
//	  1. Object - A transaction receipt object, or null when no receipt was found:
//    - transactionHash: 		DATA, 32 Bytes - hash of the transaction.
//    - transactionIndex: 		QUANTITY - integer of the transactions index position in the block.
//    - blockHash: 				DATA, 32 Bytes - hash of the block where this transaction was in.
//    - blockNumber:			QUANTITY - block number where this transaction was in.
//    - cumulativeGasUsed: 		QUANTITY - The total amount of gas used when this transaction was executed in the block.
//    - gasUsed: 				QUANTITY - The amount of gas used by this specific transaction alone.
//    - contractAddress: 		DATA, 20 Bytes - The contract address created, if the transaction was a contract creation, otherwise null.
//    - logs: 					Array - Array of log objects, which this transaction generated.
func (man *Man) GetTransactionReceipt(hash string) (*dto.TransactionReceipt, error) {

	params := make([]string, 1)
	params[0] = hash

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "man_getTransactionReceipt", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToTransactionReceipt()

}

func (man *Man) GetTopologyStatusByNumber(number *big.Int) (*common.TopologyStatus, error) {
	params := make([]interface{}, 1)
	params[0] = utils.IntToHex(number)

	pointer := &dto.RequestResult1{}
	err := man.provider.SendRequest(pointer, "man_getTopologyStatusByNumber", params)
	if err != nil {
		return nil, err
	}

	status := &common.TopologyStatus{}
	if err := json.Unmarshal(pointer.Result, status); err != nil {
		return nil, err
	}
	return status, nil
}

func (man *Man) GetBalance(address string, defaultBlockParameter string) ([]common.RPCBalanceType, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = defaultBlockParameter
	//
	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "man_getBalance", params)

	if err != nil {
		return nil, err
	}
	var balance []common.RPCBalanceType
	err = pointer.UnMarshalResult(&balance)
	if err != nil {
		return nil, err
	}
	return balance, nil
}


func (man *Man) SignTxByPrivate(sendTX *common.SendTxArgs1, from string,Privatekey *ecdsa.PrivateKey,ChainId *big.Int)(*common.SendTxArgs1,error) {
	tx1 := sendTX.ToTransaction()
	tx,err:=types.SignTx(tx1, types.NewEIP155Signer(ChainId),Privatekey)
	if err!=nil {
		return sendTX,err
	}

	sendTX.R = (*hexutil.Big)(tx.GetTxR())
	sendTX.S = (*hexutil.Big)(tx.GetTxS())
	sendTX.V = (*hexutil.Big)(tx.GetTxV())
	return sendTX,nil
}

func (man *Man) GetGasPrice() (*big.Int, error) {
	keystring := "man_TxpoolGasLimitCfg"
	key := make([]string,1)
	key[0] = keystring
	m,err := man.GetCfgDataByState(key)
	if err != nil{
		return nil,err
	}
	bytes, err := json.Marshal(m[keystring])
	if err != nil{
		return nil,err
	}
	var b big.Int
	err = json.Unmarshal(bytes, &b)
	if err != nil{
		return nil,err
	}
	return &b,nil
}

func (man *Man) GetCfgDataByState(key []string) (m map[string]interface{}, e error) {

	params := make([][]string, 1)
	params[0] = key
	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "eth_getCfgDataByState", params)

	if err != nil {
		return m, err
	}
	err = json.Unmarshal(pointer.Result, &m)

	return m, err
}