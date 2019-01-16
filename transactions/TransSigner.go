package transactions

import (
	"github.com/matrix/go-AIMan/Accounts"
	"os"
	"encoding/json"
	"path/filepath"
	"github.com/matrix/go-matrix/core/types"
)

var KeystorePath = "/home/cranelv/work/src/github.com/matrix/go-AIMan/keystore/"
var ChainID = int64(3)
var Global_Keystore *Accounts.KeystoreManager = nil
func init(){
	keystoreFile := filepath.Join(KeystorePath,"keystore.json")
	if len(keystoreFile)>0 {
		file, err := os.Open(keystoreFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		ks := make([]interface{}, 0)
		if err := json.NewDecoder(file).Decode(&ks); err != nil {
			panic(err)
		}
		Global_Keystore = Accounts.NewKeystoreManager(KeystorePath, ChainID)
	}
}
func Unlock(from string,passphrase string)error {
	return  Global_Keystore.Unlock(from,passphrase)
}
func SignTx(transaction types.SelfTransaction,from string)([]byte,error){
	return Global_Keystore.SignTx(transaction,from)
}