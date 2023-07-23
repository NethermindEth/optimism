package op_e2e

import (
	"fmt"

	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

type Deployments map[string]common.Address

func (d Deployments) Get(name string) (common.Address, error) {
	addr, ok := d[name]
	if !ok {
		return common.Address{}, fmt.Errorf("%s not found", name)
	}
	return addr, nil
}

var (
	err          error
	allocs       *state.Dump
	deployments  Deployments
	deployConfig *genesis.DeployConfig
)

func init() {
	allocs, err = e2eutils.ReadAllocs("")
	if err != nil {
		panic("Generate allocs")
	}

	deployments, err = e2eutils.ReadDeployments("")
	if err != nil {
		panic("Generate deployments")
	}
	// deployConfig = e2eutils.ReadDeployConfig("")
}
