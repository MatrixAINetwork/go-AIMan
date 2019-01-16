package common

import (
	"github.com/matrix/go-matrix/common/hexutil"
	"bytes"
	"math/big"
	"github.com/matrix/go-matrix/params"
	"errors"
	"github.com/matrix/go-matrix/common"
	"encoding/json"
)

type RPCBalanceType struct {
	AccountType uint32       `json:"accountType"`
	Balance     *hexutil.Big `json:"balance"`
}
type ExtraTo_Mx1 struct {
	To2    *string        `json:"to"`
	Value2 *hexutil.Big   `json:"value"`
	Input2 *hexutil.Bytes `json:"input"`
}

// SendTxArgs represents the arguments to sumbit a new transaction into the transaction pool.
type SendTxArgs1 struct {
	From     string          `json:"from"`
	To       *string         `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons. "input" is the
	// newer name and should be preferred by clients.
	Data        *hexutil.Bytes `json:"data"`
	Input       *hexutil.Bytes `json:"input"`
	TxType      byte           `json:"txType"`     //
	LockHeight  uint64         `json:"lockHeight"` //
	IsEntrustTx byte           `json:"isEntrustTx"`
	ExtraTo     []*ExtraTo_Mx1 `json:"extra_to"` //
}
func (args *SendTxArgs1)SetValue(int2 *big.Int){
	args.Value = (*hexutil.Big)(int2)
}
func (args *SendTxArgs1)SetData(input []byte){
	args.Data = (*hexutil.Bytes)(&input)
}
func (args *SendTxArgs1) SetDefaults() error {
	if args.Gas == nil {
		args.Gas = new(hexutil.Uint64)
		//
		if len(args.ExtraTo) > 0 && args.LockHeight > 0 && args.TxType > 0 {
			*(*uint64)(args.Gas) = 21000*uint64(len(args.ExtraTo)) + 21000
		} else {
			*(*uint64)(args.Gas) = 21000
		}
	}
	if args.GasPrice == nil {
		price := (new(big.Int).SetUint64(params.TxGasPrice))
		args.GasPrice = (*hexutil.Big)(price)
	}
	if args.Gas == nil {
		gas := hexutil.Uint64(400000)
		args.Gas = &gas
	}
	if args.Value == nil {
		args.Value = new(hexutil.Big)
	}
//	if args.Nonce == nil {
//		nonce := uint64(0)
//		args.Nonce = (*hexutil.Uint64)(&nonce)
//	}
	if args.Data != nil && args.Input != nil && !bytes.Equal(*args.Data, *args.Input) {
		return errors.New(`Both "data" and "input" are set and not equal. Please use "input" to pass transaction call data.`)
	}
	if args.Data == nil && args.Input == nil {
		args.Data = new(hexutil.Bytes)
	}
	if args.To == nil {
		// Contract creation
		var input []byte
		if args.Data != nil {
			input = *args.Data
		} else if args.Input != nil {
			input = *args.Input
		}
		if len(input) == 0 {
			return errors.New(`contract creation without any data provided`)
		}
	}
	return nil
}
type ManCallArgs struct {
	From     string         `json:"from"`
	To       *string        `json:"to"`
	Gas      hexutil.Uint64 `json:"gas"`
	GasPrice hexutil.Big    `json:"gasPrice"`
	Value    hexutil.Big    `json:"value"`
	Data     hexutil.Bytes  `json:"data"`
}
type RPCTransaction struct {
	BlockHash        common.Hash    `json:"blockHash"`
	BlockNumber      *hexutil.Big   `json:"blockNumber"`
	From             string         `json:"from"`
	Gas              hexutil.Uint64 `json:"gas"`
	GasPrice         *hexutil.Big   `json:"gasPrice"`
	Hash             common.Hash    `json:"hash"`
	Input            hexutil.Bytes  `json:"input"`
	Nonce            hexutil.Uint64 `json:"nonce"`
	To               *string        `json:"to"`
	TransactionIndex hexutil.Uint   `json:"transactionIndex"`
	Value            *hexutil.Big   `json:"value"`
	V                *hexutil.Big   `json:"v"`
	R                *hexutil.Big   `json:"r"`
	S                *hexutil.Big   `json:"s"`
	TxEnterType      byte           `json:"TxEnterType"`
	IsEntrustTx      bool           `json:"IsEntrustTx"`
	Currency         string         `json:"Currency"`
	CommitTime       hexutil.Uint64 `json:"CommitTime"`
	MatrixType       byte           `json:"matrixType"`
	ExtraTo          []*ExtraTo_Mx1 `json:"extra_to"`
}
type Elect1 struct {
	Account string
	Stock   uint16
	Type    uint16
	VIP     uint16
}
type NetTopology1 struct {
	Type            uint8
	NetTopologyData []NetTopologyData1
}

//
type NetTopologyData1 struct {
	Account  string
	Position uint16
}
type BlockHeader struct {
	Number     *hexutil.Big `json:"number"`
	Hash       string   `json:"hash"`
	Leader	   string  `json:"leader"`
	Coinbase   string   `json:"miner"`
	ParentHash string   `json:"parentHash"`
	Author     string   `json:"author,omitempty"`
	Miner      string   `json:"miner,omitempty"`
	StateRoot      string   `json:"stateRoot,omitempty"`
	TransactionsRoot      string   `json:"transactionsRoot,omitempty"`
	ReceiptsRoot      string   `json:"receiptsRoot,omitempty"`
	Size       *hexutil.Big `json:"size"`
	GasUsed    *hexutil.Big `json:"gasUsed"`
	Nonce      string `json:"nonce"`
	Timestamp  *hexutil.Big `json:"timestamp"`
	Elect             *[]Elect1                             `json:"nextElect" gencodec:"required"`
	NetTopology       *NetTopology1                         `json:"nettopology"        gencodec:"required"`
	Signatures        *[]common.Signature                         `json:"signatures" gencodec:"required"`
	Version 	hexutil.Bytes `json:"version"`
	VrfValue 	hexutil.Bytes `json:"VrfValue"`
}
type FullBlock struct {
	BlockHeader
	Transactions []*RPCTransaction `json:"transactions"`
}
type HashBlock struct {
	BlockHeader
	TxHashs []string `json:"transactions"`
}
type Block struct {
	BlockHeader
	Transactions []*RPCTransaction `json:"transactions"`
	TxHashs []string `json:"transactions"`
}
func UnmarshalBlock(buff []byte,fullTx bool)(*Block,error){
	if fullTx {
		block := &FullBlock{}
		err := json.Unmarshal(buff,block)
		result := &Block{block.BlockHeader,block.Transactions,nil}
		return result,err
	}else{
		block := &HashBlock{}
		err := json.Unmarshal(buff,block)
		result := &Block{block.BlockHeader,nil,block.TxHashs}
		return result,err
	}

}