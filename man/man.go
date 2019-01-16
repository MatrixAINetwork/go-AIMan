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
	"errors"
	"github.com/matrix/go-AIMan/complex/types"
	"github.com/matrix/go-AIMan/dto"
	"github.com/matrix/go-AIMan/man/block"
	"github.com/matrix/go-AIMan/providers"
	"github.com/matrix/go-AIMan/utils"
	"math/big"
	"strings"
	"github.com/matrix/go-AIMan/common"
	"github.com/matrix/go-matrix/base58"
	manCommon "github.com/matrix/go-matrix/common"
	"encoding/json"
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

func (man *Man) Contract(jsonInterface string) (*Contract, error) {
	return man.NewContract(jsonInterface)
}

// GetProtocolVersion - Returns the current ethereum protocol version.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_protocolversion
// Parameters:
//    - none
// Returns:
// 	  - String - The current ethereum protocol version
func (man *Man) GetProtocolVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_protocolVersion", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// IsSyncing - Returns an object with data about the sync status or false.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_syncing
// Parameters:
//    - none
// Returns:
// 	  - Object|Boolean, An object with sync status data or FALSE, when not syncing:
//    	- startingBlock: 	QUANTITY - The block at which the import started (will only be reset, after the sync reached his head)
//    	- currentBlock: 	QUANTITY - The current block, same as eth_blockNumber
//    	- highestBlock: 	QUANTITY - The estimated highest block
func (man *Man) IsSyncing() (*dto.SyncingResponse, error) {

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_syncing", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToSyncingResponse()

}

// GetCoinbase - Returns the client coinbase address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_coinbase
// Parameters:
//    - none
// Returns:
// 	  - DATA, 20 bytes - the current coinbase address.
func (man *Man) GetCoinbase() (string, error) {

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_coinbase", nil)

	if err != nil {
		return "", err
	}
	addr,_ := pointer.ToString()
	return base58.Base58EncodeToString("MAN",manCommon.HexToAddress(addr)),err

}

// IsMining - Returns true if client is actively mining new blocks.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_mining
// Parameters:
//    - none
// Returns:
// 	  - Boolean - returns true of the client is mining, otherwise false.
func (man *Man) IsMining() (bool, error) {

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_mining", nil)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}

// GetHashRate - Returns the number of hashes per second that the node is mining with.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_hashrate
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - number of hashes per second.
func (man *Man) GetHashRate() (*big.Int, error) {

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_hashrate", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetGasPrice - Returns the current price per gas in wei.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gasprice
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the current gas price in wei.
func (man *Man) GetGasPrice() (*big.Int, error) {

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_gasPrice", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// ListAccounts - Returns a list of addresses owned by client.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_accounts
// Parameters:
//    - none
// Returns:
//    - Array of DATA, 20 Bytes - addresses owned by the client.
func (man *Man) ListAccounts() ([]string, error) {

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_accounts", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToStringArray()

}

// GetBlockNumber - Returns the number of most recent block.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_blocknumber
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the current block number the client is on.
func (man *Man) GetBlockNumber() (*big.Int, error) {

	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "eth_blockNumber", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetBalance - Returns the balance of the account of given address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getbalance
// Parameters:
//    - DATA, 20 Bytes - address to check for balance.
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
// Returns:
// 	  - QUANTITY - integer of the current balance in wei.

func (man *Man) GetBalance(address string, defaultBlockParameter string) ([]common.RPCBalanceType, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = defaultBlockParameter
//
	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "eth_getBalance", params)

	if err != nil {
		return nil, err
	}
	var balance []common.RPCBalanceType
	err = pointer.UnMarshalResult(&balance)
	if err != nil{
		return nil, err
	}
	return balance,nil
}

// GetTransactionCount -  Returns the number of transactions sent from an address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionaccount
// Parameters:
//    - DATA, 20 Bytes - address to check for balance.
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
// Returns:
// 	  - QUANTITY - integer of the number of transactions sent from this address
func (man *Man) GetTransactionCount(address string, defaultBlockParameter string) (*big.Int, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_getTransactionCount", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetStorageAt - Returns the value from a storage position at a given address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getstorageat
// Parameters:
//    - DATA, 20 Bytes - address of the storage.
//	  - QUANTITY - integer of the position in the storage.
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter.
// Returns:
// 	  - DATA - the value at this storage position.
func (man *Man) GetStorageAt(address string, position *big.Int, defaultBlockParameter string) (string, error) {

	params := make([]string, 3)
	params[0] = address
	params[1] = utils.IntToHex(position)
	params[2] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_getstorageat", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()
}

// EstimateGas - Makes a call or transaction, which won't be added to the blockchain and returns the used gas, which can be used for estimating the used gas.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_estimategas
// Parameters:
//    - See eth_call parameters, expect that all properties are optional. If no gas limit is specified geth uses the block gas limit from the pending block as an
// 		upper bound. As a result the returned estimate might not be enough to executed the call/transaction when the amount of gas is higher than the pending block gas limit.
// Returns:
//    - QUANTITY - the amount of gas used.
func (man *Man) EstimateGas(transaction *common.ManCallArgs) (*big.Int, error) {

	params := make([]*common.ManCallArgs, 1)

	params[0] = transaction

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(&pointer, "eth_estimateGas", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetTransactionByHash - Returns the information about a transaction requested by transaction hash.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash
// Parameters:
//    - DATA, 32 Bytes - hash of a transaction
// Returns:
//    1. Object - A transaction object, or null when no transaction was found
//    - hash: DATA, 32 Bytes - hash of the transaction.
//    - nonce: QUANTITY - the number of transactions made by the sender prior to this one.
//    - blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
//    - blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
//    - transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
//    - from: DATA, 20 Bytes - address of the sender.
//    - to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
//    - value: QUANTITY - value transferred in Wei.
//    - gasPrice: QUANTITY - gas price provided by the sender in Wei.
//    - gas: QUANTITY - gas provided by the sender.
//    - input: DATA - the data send along with the transaction.
func (man *Man) GetTransactionByHash(hash string) (*common.RPCTransaction, error) {

	params := make([]string, 1)
	params[0] = hash

	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "eth_getTransactionByHash", params)

	if err != nil {
		return nil, err
	}
	trans := common.RPCTransaction{}
	err = json.Unmarshal(pointer.Result,&trans)
	return &trans,err

}

// GetTransactionByBlockHashAndIndex - Returns the information about a transaction requested by block hash.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getTransactionByBlockNumberAndIndex
// Parameters:
//    - DATA, 32 Bytes - hash of a block
//    - QUANTITY, number - index of the transaction position
// Returns:
//    1. Object - A transaction object, or null when no transaction was found
//    - hash: DATA, 32 Bytes - hash of the transaction.
//    - nonce: QUANTITY - the number of transactions made by the sender prior to this one.
//    - blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
//    - blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
//    - transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
//    - from: DATA, 20 Bytes - address of the sender.
//    - to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
//    - value: QUANTITY - value transferred in Wei.
//    - gasPrice: QUANTITY - gas price provided by the sender in Wei.
//    - gas: QUANTITY - gas provided by the sender.
//    - input: DATA - the data send along with the transaction.
func (man *Man) GetTransactionByBlockHashAndIndex(hash string, index uint64) (*common.RPCTransaction, error) {

	// ensure that the hash is correctlyformatted
	if strings.HasPrefix(hash, "0x") {
		if len(hash) != 66 {
			return nil, errors.New("malformed block hash")
		}
	} else {
		if len(hash) != 64 {
			return nil, errors.New("malformed block hash")
		}

		hash = "0x" + hash
	}

	params := make([]string, 2)
	params[0] = hash
	params[1] = utils.UIntToHex(index)

	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "eth_getTransactionByBlockHashAndIndex", params)

	if err != nil {
		return nil, err
	}
	var trans common.RPCTransaction
	err = json.Unmarshal(pointer.Result,&trans)
	return &trans,err
}

// GetTransactionByBlockNumberAndIndex - Returns the information about a transaction requested by block index.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getTransactionByBlockNumberAndIndex
// Parameters:
//    - QUANTITY, number - block number
//    - QUANTITY, number - transaction index in block
// Returns:
//    1. Object - A transaction object, or null when no transaction was found
//    - hash: DATA, 32 Bytes - hash of the transaction.
//    - nonce: QUANTITY - the number of transactions made by the sender prior to this one.
//    - blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
//    - blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
//    - transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
//    - from: DATA, 20 Bytes - address of the sender.
//    - to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
//    - value: QUANTITY - value transferred in Wei.
//    - gasPrice: QUANTITY - gas price provided by the sender in Wei.
//    - gas: QUANTITY - gas provided by the sender.
//    - input: DATA - the data send along with the transaction.
func (man *Man) GetTransactionByBlockNumberAndIndex(blockIndex *big.Int, index *big.Int) (*common.RPCTransaction, error) {

	params := make([]string, 2)
	params[0] = utils.IntToHex(blockIndex)
	params[1] = utils.IntToHex(index)

	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "eth_getTransactionByBlockNumberAndIndex", params)

	if err != nil {
		return nil, err
	}
	var trans common.RPCTransaction
	err = json.Unmarshal(pointer.Result,&trans)
	return &trans,err

}

// SendTransaction - Creates new message call transaction or a contract creation, if the data field contains code.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
// Parameters:
//    1. Object - The transaction object
//    - from: 		DATA, 20 Bytes - The address the transaction is send from.
//    - to: 		DATA, 20 Bytes - (optional when creating new contract) The address the transaction is directed to.
//    - gas: 		QUANTITY - (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
//    - gasPrice: 	QUANTITY - (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas
//    - value: 		QUANTITY - (optional) Integer of the value send with this transaction
//    - data: 		DATA - The compiled code of a contract OR the hash of the invoked method signature and encoded parameters. For details see Manereum Contract ABI (https://github.com/ethereum/wiki/wiki/Manereum-Contract-ABI)
//    - nonce: 		QUANTITY - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
// Returns:
//	  - DATA, 32 Bytes - the transaction hash, or the zero hash if the transaction is not yet available.
// Use eth_getTransactionReceipt to get the contract address, after the transaction was mined, when you created a contract.
func (man *Man) SendTransaction(transaction *common.SendTxArgs1,password string) (string, error) {

	params := make([]interface{}, 2)
	transaction.SetDefaults()
	params[0] = transaction
	params[1] = password

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(&pointer, "eth_sendTransaction", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
func (man *Man) SendRawTransaction(rawTx []byte) (string, error) {

	params := make([]interface{}, 1)
	params[0] = "0x" + manCommon.Bytes2Hex(rawTx)

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(&pointer, "eth_sendRawTransaction", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// SignTransaction - Signs transactions without dispatching it to the network. It can be later submitted using eth_sendRawTransaction.
// Reference: https://wiki.parity.io/JSONRPC-man-module.html#eth_signtransaction
// Parameters:
//    1. Object - The transaction call object
//    - from: 		DATA, 20 Bytes - The address the transaction is send from.
//    - to: 		DATA, 20 Bytes - (optional when creating new contract) The address the transaction is directed to.
//    - gas: 		QUANTITY - (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
//    - gasPrice: 	QUANTITY - (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas
//    - value: 		QUANTITY - (optional) Integer of the value send with this transaction
//    - data: 		DATA - The compiled code of a contract OR the hash of the invoked method signature and encoded parameters. For details see Ethereum Contract ABI (https://github.com/ethereum/wiki/wiki/Ethereum-Contract-ABI)
//    - nonce: 		QUANTITY - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
// Returns:
//    1. Object - A transaction sign result object
//    - raw: DATA - The signed, RLP encoded transaction.
//    - tx: Object - A transaction object
//      - hash: DATA, 32 Bytes - hash of the transaction.
//      - nonce: QUANTITY - the number of transactions made by the sender prior to this one.
//      - blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
//      - blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
//      - transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
//      - from: DATA, 20 Bytes - address of the sender.
//      - to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
//      - value: QUANTITY - value transferred in Wei.
//      - gasPrice: QUANTITY - gas price provided by the sender in Wei.
//      - gas: QUANTITY - gas provided by the sender.
//      - input: DATA - the data send along with the transaction.
// Use eth_sendRawTransaction to submit the transaction after it was signed.
func (man *Man) SignTransaction(transaction *dto.TransactionParameters) (*dto.SignTransactionResponse, error) {
	params := make([]*dto.RequestTransactionParameters, 1)
	params[0] = transaction.Transform()

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(&pointer, "eth_signTransaction", params)

	if err != nil {
		return &dto.SignTransactionResponse{}, err
	}

	return pointer.ToSignTransactionResponse()
}

// Call - Executes a new message call immediately without creating a transaction on the block chain.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_call
// Parameters:
//    1. Object - The transaction call object
//    - from: 		DATA, 20 Bytes - The address the transaction is send from.
//    - to: 		DATA, 20 Bytes - (optional when creating new contract) The address the transaction is directed to.
//    - gas: 		QUANTITY - (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
//    - gasPrice: 	QUANTITY - (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas
//    - value: 		QUANTITY - (optional) Integer of the value send with this transaction
//    - data: 		DATA - The compiled code of a contract OR the hash of the invoked method signature and encoded parameters. For details see Ethereum Contract ABI (https://github.com/ethereum/wiki/wiki/Ethereum-Contract-ABI)
//	  2. QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
// Returns:
//	  - DATA - the return value of executed contract.
func (man *Man) Call(transaction *common.SendTxArgs1) (*dto.RequestResult, error) {

	params := make([]interface{}, 2)
	transaction.SetDefaults()
	params[0] = transaction
	params[1] = block.LATEST

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(&pointer, "eth_call", params)

	if err != nil {
		return nil, err
	}

	return pointer, err

}

// CompileSolidity - Returns compiled solidity code.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_compilesolidity
// Parameters:
//    1. String - The source code.
// Returns:
//	  - DATA - The compiled source code.
func (man *Man) CompileSolidity(sourceCode string) (types.ComplexString, error) {

	params := make([]string, 1)
	params[0] = sourceCode

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_compileSolidity", params)

	if err != nil {
		return "", err
	}

	return pointer.ToComplexString()

}

// GetTransactionReceipt - Returns compiled solidity code.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionreceipt
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

	err := man.provider.SendRequest(pointer, "eth_getTransactionReceipt", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToTransactionReceipt()

}

// GetBlockByNumber - Returns the information about a block requested by number.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbynumber
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

	err := man.provider.SendRequest(pointer, "eth_getBlockByNumber", params)

	if err != nil {
		return nil, err
	}
	return common.UnmarshalBlock(pointer.Result,transactionDetails)

}

// GetBlockTransactionCountByHash
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblocktransactioncountbyhash
// Parameters:
//    - DATA, 32 bytes - block hash
// Returns:
//    1. QUANTITY, number - number of transactions in the block
//    2. error
func (man *Man) GetBlockTransactionCountByHash(hash string) (*big.Int, error) {
	// ensure that the hash is correctlyformatted
	if strings.HasPrefix(hash, "0x") {
		if len(hash) != 66 {
			return nil, errors.New("malformed block hash")
		}
	} else {
		if len(hash) != 64 {
			return nil, errors.New("malformed block hash")
		}
		hash = "0x" + hash
	}

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_getBlockTransactionCountByHash", []string{hash})

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetBlockTransactionCountByNumber - Returns the number of transactions in a block matching the given block number
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblocktransactioncountbynumber
// Parameters:
//    - QUANTITY|TAG - integer of a block number, or the string "earliest", "latest" or "pending", as in the default block parameter
// Returns:
//    - QUANTITY - integer of the number of transactions in this block
func (man *Man) GetBlockTransactionCountByNumber(defaultBlockParameter string) (*big.Int, error) {

	params := make([]string, 1)
	params[0] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_getBlockTransactionCountByNumber", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetBlockByHash - Returns information about a block by hash.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash
// Parameters:
//    - DATA, 32 bytes - Hash of a block
//    - transactionDetails, bool - indicate if we should have or not the details of the transactions of the block
// Returns:
//    1. Object - A block object, or null when no transaction was found
//    2. error
func (man *Man) GetBlockByHash(hash string, transactionDetails bool) (*common.Block, error) {
	// ensure that the hash is correctlyformatted
	if strings.HasPrefix(hash, "0x") {
		if len(hash) != 66 {
			return nil, errors.New("malformed block hash")
		}
	} else {
		hash = "0x" + hash
		if len(hash) != 62 {
			return nil, errors.New("malformed block hash")
		}
	}

	params := make([]interface{}, 2)
	params[0] = hash
	params[1] = transactionDetails

	pointer := &dto.RequestResult1{}

	err := man.provider.SendRequest(pointer, "eth_getBlockByHash", params)

	if err != nil {
		return nil, err
	}
	return common.UnmarshalBlock(pointer.Result,transactionDetails)
}

// GetUncleCountByBlockHash - Returns the number of uncles in a block from a block matching the given block hash.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getunclecountbyblockhash
// Parameters:
//    - DATA, 32 bytes - Hash of a block
// Returns:
//    - QUANTITY, number - integer of the number of uncles in this block
//    - error
func (man *Man) GetUncleCountByBlockHash(hash string) (*big.Int, error) {
	// ensure that the hash has been correctly formatted
	if strings.HasPrefix(hash, "0x") {
		if len(hash) != 66 {
			return nil, errors.New("malformed block hash")
		}
	} else {
		if len(hash) != 64 {
			return nil, errors.New("malformed block hash")
		}
		hash = "0x" + hash
	}

	params := make([]string, 1)
	params[0] = hash

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_getUncleCountByBlockHash", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetUncleCountByBlockNumber - Returns the number of uncles in a block from a block matching the given block number.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getunclecountbyblocknumber
// Parameters:
//    - QUANTITY, number - integer of a block number
// Returns:
//    - QUANTITY, number - integer of the number of uncles in this block
//    - error
func (man *Man) GetUncleCountByBlockNumber(quantity *big.Int) (*big.Int, error) {
	// ensure that the hash has been correctly formatted

	params := make([]string, 1)
	params[0] = utils.IntToHex(quantity)

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_getUncleCountByBlockNumber", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBigInt()
}

// GetCode - Returns code at a given address
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getcode
// Parameters:
//    - DATA, 20 Bytes - address
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
// Returns:
//    - DATA - the code from the given address.
func (man *Man) GetCode(address string, defaultBlockParameter string) (string, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := man.provider.SendRequest(pointer, "eth_getCode", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()
}
