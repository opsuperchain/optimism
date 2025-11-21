package frontend

import (
	ftypes "github.com/ethereum-optimism/optimism/op-faucet/faucet/backend/types"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type AdminBackend interface {
	EnableFaucet(id ftypes.FaucetID)
	DisableFaucet(id ftypes.FaucetID)
	Faucets() map[ftypes.FaucetID]eth.ChainID
	Defaults() map[eth.ChainID]ftypes.FaucetID
}

type AdminFrontend struct {
	AdminBackend
}

func NewAdminFrontend(b AdminBackend) *AdminFrontend {
	return &AdminFrontend{AdminBackend: b}
}
