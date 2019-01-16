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
 * @file request-result.go
 * @authors:
 *   Reginaldo Costa <matrix@gmail.com>
 * @date 2017
 */

package dto

import (
	"errors"
	"strconv"
	"strings"

	"github.com/matrix/go-AIMan/complex/types"
	"github.com/matrix/go-AIMan/constants"

	"encoding/json"
	"fmt"
	"math/big"
	"github.com/matrix/go-matrix/common/hexutil"
)

type RequestResult struct {
	ID      int         `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   *Error      `json:"error,omitempty"`
	Data    string      `json:"data,omitempty"`
}
type RequestResult1 struct {
	ID      int         `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  json.RawMessage	    `json:"result"`
	Error   *Error      `json:"error,omitempty"`
	Data    string      `json:"data,omitempty"`
}
func NewRequestResult(result interface{})*RequestResult{
	value := &RequestResult{}
	value.Result = result
	return value
}
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
func (pointer *RequestResult1)UnMarshalResult(result interface{})error{
	return json.Unmarshal(pointer.Result,result)
}
func (pointer *RequestResult1)ToBigInt()(*big.Int,error){
	value := new(hexutil.Big)
	err := json.Unmarshal(pointer.Result,&value)
	return value.ToInt(),err
}
func (pointer *RequestResult) ToStringArray() ([]string, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.([]interface{})

	new := make([]string, len(result))
	for i, v := range result {
		new[i] = v.(string)
	}

	return new, nil

}

func (pointer *RequestResult) ToComplexString() (types.ComplexString, error) {

	if err := pointer.checkResponse(); err != nil {
		return "", err
	}

	result := (pointer).Result.(interface{})

	return types.ComplexString(result.(string)), nil

}

func (pointer *RequestResult) ToString() (string, error) {

	if err := pointer.checkResponse(); err != nil {
		return "", err
	}

	result := (pointer).Result.(interface{})

	return result.(string), nil

}

func (pointer *RequestResult) ToInt() (int64, error) {

	if err := pointer.checkResponse(); err != nil {
		return 0, err
	}

	result := (pointer).Result.(interface{})

	hex := result.(string)

	numericResult, err := strconv.ParseInt(hex, 16, 64)

	return numericResult, err

}
func (pointer *RequestResult) GetResult() (interface{}, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	res := (pointer).Result.(interface{})

	return res, nil
}
func (pointer *RequestResult) ToBigInt() (*big.Int, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	res := (pointer).Result.(interface{})

	ret, success := big.NewInt(0).SetString(res.(string)[2:], 16)

	if !success {
		return nil, errors.New(fmt.Sprintf("Failed to convert %s to BigInt", res.(string)))
	}

	return ret, nil
}

func (pointer *RequestResult) ToComplexIntResponse() (types.ComplexIntResponse, error) {

	if err := pointer.checkResponse(); err != nil {
		return types.ComplexIntResponse(0), err
	}

	result := (pointer).Result.(interface{})

	var hex string

	switch v := result.(type) {
	//Testrpc returns a float64
	case float64:
		hex = strconv.FormatFloat(v, 'E', 16, 64)
		break
	default:
		hex = result.(string)
	}

	cleaned := strings.TrimPrefix(hex, "0x")

	return types.ComplexIntResponse(cleaned), nil

}

func (pointer *RequestResult) ToBoolean() (bool, error) {

	if err := pointer.checkResponse(); err != nil {
		return false, err
	}

	result := (pointer).Result.(interface{})

	return result.(bool), nil

}

func (pointer *RequestResult) ToSignTransactionResponse() (*SignTransactionResponse, error) {
	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.(map[string]interface{})

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	signTransactionResponse := &SignTransactionResponse{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	err = json.Unmarshal([]byte(marshal), signTransactionResponse)

	return signTransactionResponse, err
}


func (pointer *RequestResult) ToTransactionReceipt() (*TransactionReceipt, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	result := (pointer).Result.(map[string]interface{})

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	transactionReceipt := &TransactionReceipt{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	err = json.Unmarshal([]byte(marshal), transactionReceipt)

	return transactionReceipt, err

}


func (pointer *RequestResult) ToSyncingResponse() (*SyncingResponse, error) {

	if err := pointer.checkResponse(); err != nil {
		return nil, err
	}

	var result map[string]interface{}

	switch (pointer).Result.(type) {
	case bool:
		return &SyncingResponse{}, nil
	case map[string]interface{}:
		result = (pointer).Result.(map[string]interface{})
	default:
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	if len(result) == 0 {
		return nil, customerror.EMPTYRESPONSE
	}

	syncingResponse := &SyncingResponse{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, customerror.UNPARSEABLEINTERFACE
	}

	json.Unmarshal([]byte(marshal), syncingResponse)

	return syncingResponse, nil

}

// To avoid a conversion of a nil interface
func (pointer *RequestResult) checkResponse() error {

	if pointer.Error != nil {
		return errors.New(pointer.Error.Message)
	}

	if pointer.Result == nil {
		return customerror.EMPTYRESPONSE
	}

	return nil

}
