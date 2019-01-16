package Accounts

import (
	"github.com/matrix/go-matrix/accounts/keystore"
	"github.com/matrix/go-matrix/core/types"
	"github.com/matrix/go-matrix/accounts"
	"io/ioutil"
	"fmt"
	"github.com/matrix/go-matrix/common"
	"math/big"
	"github.com/matrix/go-matrix/base58"
	"time"
	"github.com/matrix/go-matrix/rlp"
)

func EthAddressToManAddress(address common.Address)string{
	return base58.Base58EncodeToString("MAN",address)
}
func ManAddressToEthAddress(manAddr string) common.Address{
	return base58.Base58DecodeToAddress(manAddr)
}
type KeystoreManager struct {
	Keystore *keystore.KeyStore
	ChainID *big.Int
	Signer types.Signer
}
func NewKeystoreManager(keystoreDir string,chainId int64)* KeystoreManager{
	ks := &KeystoreManager{Signer:types.NewEIP155Signer(big.NewInt(chainId))}
	ks.ChainID = big.NewInt(chainId)
	ks.Keystore = keystore.NewKeyStore(keystoreDir,keystore.StandardScryptN,keystore.StandardScryptP)
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
func (ks *KeystoreManager) Unlock(manAddr string,password string)error{
	return ks.TimedUnlock(manAddr,password,0)
}
func (ks *KeystoreManager) TimedUnlock(manAddr string,password string,timeout time.Duration)error{
	addr := ManAddressToEthAddress(manAddr)
	acc := accounts.Account{Address:addr}
	return ks.Keystore.TimedUnlock(acc,password,timeout)
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
func (ks *KeystoreManager)SignTx(transaction types.SelfTransaction,from string)([]byte,error){
	addr := ManAddressToEthAddress(from)
	acc := accounts.Account{Address:addr}
	tx,err := ks.Keystore.SignTx(acc,transaction,ks.ChainID)
	tx.(*types.Transaction).Currency = "MAN"
	if err != nil {
		return nil,err
	}
	return tx.(*types.Transaction).EncodeForRawTransaction()
}
func (ks *KeystoreManager)SignTxWithPassphrase(transaction types.SelfTransaction,from string, passphrase string)([]byte,error){
	addr := ManAddressToEthAddress(from)
	acc := accounts.Account{Address:addr}
	tx,err := ks.Keystore.SignTxWithPassphrase(acc,passphrase,transaction,ks.ChainID)
	if err != nil {
		return nil,err
	}
	return rlp.EncodeToBytes(tx)
}