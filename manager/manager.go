package manager

import (
	"github.com/matrix/go-AIMan/AIMan"
	"github.com/matrix/go-AIMan/Accounts"
	"github.com/matrix/go-AIMan/providers"
)
var (
	KeystorePath = "/home/cranelv/work/src/github.com/matrix/go-AIMan/keystore/"
	Tom_Manager = &Manager{
		AIMan.NewAIMan(providers.NewHTTPProvider("api85.matrix.io", 100, false)),
		Accounts.NewKeystoreManager(KeystorePath, 1),
	}
	Jerry_Manager = &Manager{
		AIMan.NewAIMan(providers.NewHTTPProvider("47.105.202.251:8341", 100, false)),
		Accounts.NewKeystoreManager(KeystorePath, 3),
	}
	Local_Manager = &Manager{
		AIMan.NewAIMan(providers.NewHTTPProvider("localhost:8341", 100000, false)),
		Accounts.NewKeystoreManager(KeystorePath, 1),
	}
)
type Manager struct {
	*AIMan.AIMan
	*Accounts.KeystoreManager
}
