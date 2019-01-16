package Accounts

import (
	"testing"
	"os"
	"encoding/json"
	"path/filepath"
)
var KeystorePath = "../keystore"
func TestUnlockKeystore(t *testing.T){
	keystoreFile := filepath.Join(KeystorePath,"keystore.json")
	if len(keystoreFile)>0{
		file, err := os.Open(keystoreFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		ks := make([]interface{},0)
		if err := json.NewDecoder(file).Decode(&ks); err != nil {
			t.Error(err)
		}
		keystore := NewKeystoreManager(KeystorePath,1)
		for _,item := range ks {
			account := item.(map[string]interface{})
			err := keystore.Unlock(account["address"].(string),account["password"].(string))
			if err != nil {
				t.Error(err)
			}
		}
	}
}