package Accounts

import (
	"fmt"
	common1 "github.com/MatrixAINetwork/go-AIMan/common"
	"github.com/MatrixAINetwork/go-matrix/accounts"
	"github.com/MatrixAINetwork/go-matrix/accounts/keystore"
	"github.com/MatrixAINetwork/go-matrix/base58"
	"github.com/MatrixAINetwork/go-matrix/common"
	"github.com/MatrixAINetwork/go-matrix/common/hexutil"
	"github.com/MatrixAINetwork/go-matrix/core/types"
	"github.com/MatrixAINetwork/go-matrix/rlp"
	"io/ioutil"
	"math/big"
	"time"
	strings "strings"
	"github.com/MatrixAINetwork/go-matrix/crc8"
)

func EthAddressToManAddress(address common.Address) string {
	return base58.Base58EncodeToString("MAN", address)
}
func ManAddressToEthAddress(manAddr string) (common.Address, error) {
	return base58.Base58DecodeToAddress(manAddr)
}

type KeystoreManager struct {
	Keystore *keystore.KeyStore
	ChainID  *big.Int
	Signer   types.Signer
}

func NewKeystoreManager(keystoreDir string, chainId int64) *KeystoreManager {
	ks := &KeystoreManager{Signer: types.NewEIP155Signer(big.NewInt(chainId))}
	ks.ChainID = big.NewInt(chainId)
	ks.Keystore = keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks
}
func (ks *KeystoreManager) getDecryptedKey(a accounts.Account, auth string) (accounts.Account, *keystore.Key, error) {
	a, err := ks.Keystore.Find(a)
	if err != nil {
		return a, nil, err
	}
	key, err := ks.GetKey(a.Address, a.URL.Path, auth)
	return a, key, err
}
func (ks *KeystoreManager) Unlock(manAddr string, password string) error {
	return ks.TimedUnlock(manAddr, password, 0)
}
func (ks *KeystoreManager) TimedUnlock(manAddr string, password string, timeout time.Duration) error {
	addr, _ := ManAddressToEthAddress(manAddr)
	acc := accounts.Account{Address: addr}
	return ks.Keystore.TimedUnlock(acc, password, timeout)
}
func (ks *KeystoreManager) GetKey(addr common.Address, filename, auth string) (*keystore.Key, error) {
	// Load the key from the keystore and decrypt its contents
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	// Make sure we're really operating on the requested key (no swap attacks)
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, addr)
	}
	return key, nil
}

func (ks *KeystoreManager) SignTx(sendTx *common1.SendTxArgs1, from string) (*common1.SendTxArgs1, error) {
	addr, _ := ManAddressToEthAddress(from)
	acc := accounts.Account{Address: addr}
	tx1 := sendTx.ToTransaction()
	tx, err := ks.Keystore.SignTx(acc, tx1, ks.ChainID)
	if err != nil {
		return nil, err
	}
	sendTx.R = (*hexutil.Big)(tx.GetTxR())
	sendTx.S = (*hexutil.Big)(tx.GetTxS())
	sendTx.V = (*hexutil.Big)(tx.GetTxV())
	return sendTx, nil
}

func (ks *KeystoreManager) SignTxWithPassphrase(transaction types.SelfTransaction, from string, passphrase string) ([]byte, error) {
	addr, _ := ManAddressToEthAddress(from)
	acc := accounts.Account{Address: addr}
	tx, err := ks.Keystore.SignTxWithPassphrase(acc, passphrase, transaction, ks.ChainID)
	if err != nil {
		return nil, err
	}
	return rlp.EncodeToBytes(tx)
}


func checkCrc8(strData string) bool {
	Crc := strData[len(strData)-1 : len(strData)]
	reCrc := crc8.CalCRC8([]byte(strData[0 : len(strData)-1]))
	ModCrc := reCrc % 58
	ret := base58.EncodeInt(ModCrc)
	if Crc != ret {
		return false
	}
	return true
}
func checkCurrency(strData string) bool {
	currency := strings.Split(strData, ".")[0]
	return common.IsValidityManCurrency(currency)
}
func checkFormat(strData string) bool {
	if !strings.Contains(strData, ".") {
		return false
	}
	return true
}
func CheckIsManAddress(strData string) bool {
	strData = strings.TrimSpace(strData)
	if !checkFormat(strData) {
		return false
	}
	if !checkCrc8(strData) {
		return false
	}
	if !checkCurrency(strData) {
		return false
	}
	return true
}