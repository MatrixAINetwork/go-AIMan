package manager

import (
	"github.com/MatrixAINetwork/go-AIMan/AIMan"
	"github.com/MatrixAINetwork/go-AIMan/Accounts"
)

//var (
//	KeystorePath = "keystore"
//	Tom_Manager  = &Manager{
//		AIMan.NewAIMan(providers.NewHTTPProvider("api85.matrix.io", 100, false)),
//		Accounts.NewKeystoreManager(KeystorePath, 1),
//	}
//)
//
type Manager struct {
	*AIMan.AIMan
	*Accounts.KeystoreManager
}
