package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/MatrixAINetwork/go-matrix/base58"
	"github.com/MatrixAINetwork/go-matrix/common"
	"github.com/MatrixAINetwork/go-matrix/common/hexutil"
	"github.com/MatrixAINetwork/go-matrix/core/types"
	"github.com/MatrixAINetwork/go-matrix/params"
	"math/big"
	"strings"
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
type ExtraTo_Mx struct {
	To2    *common.Address `json:"to"`
	Value2 *hexutil.Big    `json:"value"`
	Input2 *hexutil.Bytes  `json:"input"`
}
type SendTxArgs struct {
	From     common.Address  `json:"from"`
	Currency string          `json:"currency"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons. "input" is the
	// newer name and should be preferred by clients.
	Data        *hexutil.Bytes `json:"data"`
	Input       *hexutil.Bytes `json:"input"`
	V           *hexutil.Big   `json:"v"`
	R           *hexutil.Big   `json:"r"`
	S           *hexutil.Big   `json:"s"`
	TxType      byte           `json:"txType"`     //
	LockHeight  uint64         `json:"lockHeight"` //
	IsEntrustTx byte           `json:"isEntrustTx"`
	CommitTime  uint64         `json:"commitTime"`
	ExtraTo     []*ExtraTo_Mx  `json:"extra_to"` //
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
	V           *hexutil.Big   `json:"v"`
	R           *hexutil.Big   `json:"r"`
	S           *hexutil.Big   `json:"s"`
	Currency    *string        `json:"currency"`
	TxType      byte           `json:"txType"`     //
	LockHeight  uint64         `json:"lockHeight"` //
	IsEntrustTx byte           `json:"isEntrustTx"`
	CommitTime  uint64         `json:"commitTime"`
	ExtraTo     []*ExtraTo_Mx1 `json:"extra_to"` //
}

func StrArgsToByteArgs(args1 SendTxArgs1) (args SendTxArgs, err error) {
	if args1.From != "" {
		from := args1.From
		//err = CheckParams(from)
		//if err != nil {
		//	return SendTxArgs{}, err
		//}
		args.Currency = strings.Split(args1.From, ".")[0]
		args.From, _ = base58.Base58DecodeToAddress(from)
	}
	if args1.Currency != nil {
		args.Currency = *args1.Currency
	}
	if args1.To != nil {
		to := *args1.To
		//err = CheckParams(to)
		//if err != nil {
		//	return SendTxArgs{}, err
		//}
		args.To = new(common.Address)
		*args.To, _ = base58.Base58DecodeToAddress(to)
	}
	if args1.V != nil {
		args.V = args1.V
	}
	if args1.R != nil {
		args.R = args1.R
	}
	if args1.S != nil {
		args.S = args1.S
	}
	args.Gas = args1.Gas
	args.GasPrice = args1.GasPrice
	args.Value = args1.Value
	args.Nonce = args1.Nonce
	args.Data = args1.Data
	args.Input = args1.Input
	args.TxType = args1.TxType
	args.LockHeight = args1.LockHeight
	args.CommitTime = args1.CommitTime
	args.IsEntrustTx = args1.IsEntrustTx
	if len(args1.ExtraTo) > 0 { //扩展交易中的to属性不填写则删掉这个扩展交易
		extra := make([]*ExtraTo_Mx, 0)
		for _, ar := range args1.ExtraTo {
			if ar.To2 != nil {
				//extra = append(extra, ar)
				tmp := *ar.To2
				//err = CheckParams(tmp)
				//if err != nil {
				//	return SendTxArgs{}, err
				//}
				tmExtra := new(ExtraTo_Mx)
				tmExtra.To2 = new(common.Address)
				*tmExtra.To2, _ = base58.Base58DecodeToAddress(tmp)
				tmExtra.Input2 = ar.Input2
				tmExtra.Value2 = ar.Value2
				extra = append(extra, tmExtra)
			}
		}
		args.ExtraTo = extra
	}
	return args, nil
}

func (args *SendTxArgs1) SetValue(int2 *big.Int) {
	args.Value = (*hexutil.Big)(int2)
}
func (args *SendTxArgs1) SetData(input []byte) {
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

func (args1 *SendTxArgs1) ToTransaction() *types.Transaction {

	//from := common.Address{}
	currency := ""
	if args1.From != "" {
		currency = strings.Split(args1.From, ".")[0]
		//from = base58.Base58DecodeToAddress(args1.From)
	}
	if args1.Currency != nil {
		currency = *args1.Currency
	}
	var to *common.Address
	if args1.To != nil {
		to = new(common.Address)
		*to, _ = base58.Base58DecodeToAddress(*args1.To)
	}

	var input []byte
	if args1.Data != nil {
		input = *args1.Data
	} else if args1.Input != nil {
		input = *args1.Input
	}
	if to == nil {
		return types.NewContractCreation(uint64(*args1.Nonce), (*big.Int)(args1.Value), uint64(*args1.Gas), (*big.Int)(args1.GasPrice), input, (*big.Int)(args1.V), (*big.Int)(args1.R), (*big.Int)(args1.S), 0, args1.IsEntrustTx, *args1.Currency, args1.CommitTime)
	}

	extra := make([]*ExtraTo_Mx, 0)
	if len(args1.ExtraTo) > 0 { //扩展交易中的to属性不填写则删掉这个扩展交易
		for _, ar := range args1.ExtraTo {
			if ar.To2 != nil {
				//extra = append(extra, ar)
				tmp := *ar.To2
				//err = CheckParams(tmp)
				//if err != nil {
				//	return SendTxArgs{}, err
				//}
				tmExtra := new(ExtraTo_Mx)
				tmExtra.To2 = new(common.Address)
				*tmExtra.To2, _ = base58.Base58DecodeToAddress(tmp)
				tmExtra.Input2 = ar.Input2
				tmExtra.Value2 = ar.Value2
				extra = append(extra, tmExtra)
			}
		}
	}

	//
	txtr := make([]*types.ExtraTo_tr, 0)
	if len(extra) > 0 {
		for _, extra := range args1.ExtraTo {
			tmp := new(types.ExtraTo_tr)
			va := extra.Value2
			if va == nil {
				va = (*hexutil.Big)(big.NewInt(0))
			}
			if extra.To2 != nil {
				tmp.To_tr = new(common.Address)
				*tmp.To_tr, _ = base58.Base58DecodeToAddress(*extra.To2)
			}

			tmp.Value_tr = va
			tmp.Input_tr = extra.Input2
			txtr = append(txtr, tmp)
		}
	}
	return types.NewTransactions(uint64(*args1.Nonce), *to, (*big.Int)(args1.Value), uint64(*args1.Gas), (*big.Int)(args1.GasPrice), input, (*big.Int)(args1.V), (*big.Int)(args1.R), (*big.Int)(args1.S), txtr, args1.LockHeight, args1.TxType, args1.IsEntrustTx, currency, args1.CommitTime)

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
type DepositDetail struct {
	Address     string
	SignAddress string
	Deposit     *big.Int
	WithdrawH   *big.Int
	OnlineTime  *big.Int
	Role        *big.Int
}
type Elect1 struct {
	Account string
	Stock   uint64
	Type    common.RoleType
	VIP     common.VIPRoleType
}
type NetTopology1 struct {
	Type            uint8
	NetTopologyData []NetTopologyData1
}

//
type NodeInfo struct {
	Account  string `json:"account"`
	Online   bool   `json:"online"`
	Position uint16 `json:"position"`
}

type TopologyStatus struct {
	LeaderReelect         bool       `json:"leader_reelect"`
	Validators            []NodeInfo `json:"validators"`
	BackupValidators      []NodeInfo `json:"backup_validators"`
	Miners                []NodeInfo `json:"miners"`
	ElectValidators       []NodeInfo `json:"elect_validators"`
	ElectBackupValidators []NodeInfo `json:"elect_backup_validators"`
}
type NetTopologyData1 struct {
	Account  string
	Position uint16
}
type BlockHeader struct {
	Number   *hexutil.Big `json:"number"`
	Hash     string       `json:"hash"`
	SignHash string       `json:"signHash"`
	Leader   string       `json:"leader"`
	//	Coinbase   string   `json:"miner"`
	ParentHash       string              `json:"parentHash"`
	Author           string              `json:"author,omitempty"`
	Miner            string              `json:"miner,omitempty"`
	StateRoot        string              `json:"stateRoot,omitempty"`
	TransactionsRoot string              `json:"transactionsRoot,omitempty"`
	ReceiptsRoot     string              `json:"receiptsRoot,omitempty"`
	Size             *hexutil.Big        `json:"size"`
	GasUsed          *hexutil.Big        `json:"gasUsed"`
	Nonce            string              `json:"nonce"`
	Timestamp        *hexutil.Big        `json:"timestamp"`
	Elect            *[]Elect1           `json:"elect" gencodec:"required"`
	NetTopology      *NetTopology1       `json:"nettopology"        gencodec:"required"`
	Signatures       *[]common.Signature `json:"signatures" gencodec:"required"`
	Version          hexutil.Bytes       `json:"version"`
	VrfValue         hexutil.Bytes       `json:"VrfValue"`
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
	TxHashs      []string          `json:"transactions"`
}

func UnmarshalBlock(buff []byte, fullTx bool) (*Block, error) {
	if fullTx {
		block := &FullBlock{}
		err := json.Unmarshal(buff, block)
		result := &Block{block.BlockHeader, block.Transactions, nil}
		return result, err
	} else {
		block := &HashBlock{}
		err := json.Unmarshal(buff, block)
		result := &Block{block.BlockHeader, nil, block.TxHashs}
		return result, err
	}

}
