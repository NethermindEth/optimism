package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

var (
	L1Allocs      *state.Dump
	L1Deployments *genesis.L1Deployments
	DeployConfig  *genesis.DeployConfig
)

type Deployments map[string]common.Address

func (d Deployments) Get(name string) (common.Address, error) {
	addr, ok := d[name]
	if !ok {
		return common.Address{}, fmt.Errorf("%s not found", name)
	}
	return addr, nil
}

// ReadAllocs
func ReadAllocs(filename string) (*state.Dump, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var dump state.Dump
	if err := decoder.Decode(&dump); err != nil {
		return nil, err
	}
	return &dump, nil
}

// Init testing to enable test flags
var _ = func() bool {
	testing.Init()
	return true
}()

func init() {
	var l1AllocsPath, l1DeploymentsPath, deployConfigPath string

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	root, err := findMonorepoRoot(cwd)
	if err != nil {
		panic(err)
	}

	defaultL1AllocsPath := filepath.Join(root, ".devnet", "allocs-l1.json")
	defaultL1DeploymentsPath := filepath.Join(root, "packages", "contracts-bedrock", "deployments", "devnetL1", ".deploy")
	defaultDeployConfigPath := filepath.Join(root, "packages", "contracts-bedrock", "deploy-config", "devnetL1.json")

	flag.StringVar(&l1AllocsPath, "l1-allocs", defaultL1AllocsPath, "")
	flag.StringVar(&l1DeploymentsPath, "l1-deployments", defaultL1DeploymentsPath, "")
	flag.StringVar(&deployConfigPath, "deploy-config", defaultDeployConfigPath, "")
	flag.Parse()

	L1Allocs, err = ReadAllocs(l1AllocsPath)
	if err != nil {
		panic(err)
	}
	L1Deployments, err = genesis.NewL1Deployments(l1DeploymentsPath)
	if err != nil {
		panic(err)
	}
	DeployConfig, err = genesis.NewDeployConfig(deployConfigPath)
	if err != nil {
		panic(err)
	}

	// Get the storage layout
	account := L1Allocs.Accounts[L1Deployments.OptimismPortalProxy]
	if len(account.Code) == 0 {
		panic("portal proxy doesn't exist")
	}

	layout, err := bindings.GetStorageLayout("OptimismPortal")
	if err != nil {
		panic(err)
	}
	entry, err := layout.GetStorageLayoutEntry("params")
	if err != nil {
		panic(err)
	}
	slot := common.BigToHash(big.NewInt(int64(entry.Slot)))
	account.Storage[slot] = slot.String()
	L1Allocs.Accounts[L1Deployments.OptimismPortalProxy] = account

	DeployConfig.L1StandardBridgeProxy = L1Deployments.L1StandardBridgeProxy
	DeployConfig.L1CrossDomainMessengerProxy = L1Deployments.L1CrossDomainMessengerProxy
	DeployConfig.L1ERC721BridgeProxy = L1Deployments.L1ERC721BridgeProxy
	DeployConfig.OptimismPortalProxy = L1Deployments.OptimismPortalProxy
	DeployConfig.SystemConfigProxy = L1Deployments.SystemConfigProxy
	// TODO: perhaps log L1Deployments?
}

// findMonorepoRoot will recursively search upwards for a go.mod file.
// This depends on the structure of the monorepo having a go.mod file at the root.
func findMonorepoRoot(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}

	for {
		modulePath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(modulePath); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		// Check if we reached the filesystem root
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("monorepo root not found")
}
