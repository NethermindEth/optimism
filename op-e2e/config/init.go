package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

var (
	L1Allocs      *state.Dump
	L1Deployments Deployments
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

// ReadDeployments
func ReadDeployments(filename string) (map[string]common.Address, error) {
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
	var deployments map[string]common.Address
	if err := decoder.Decode(&deployments); err != nil {
		return nil, err
	}
	return deployments, nil
}

// TODO? this is twice?
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
	L1Deployments, err = ReadDeployments(l1DeploymentsPath)
	if err != nil {
		panic(err)
	}
	DeployConfig, err = genesis.NewDeployConfig(deployConfigPath)
	if err != nil {
		panic(err)
	}

	if l1StandardBridgeProxy, err := L1Deployments.Get("L1StandardBridgeProxy"); err == nil {
		fmt.Printf("Using L1StandardBridgeProxy: %s\n", l1StandardBridgeProxy)
		DeployConfig.L1StandardBridgeProxy = l1StandardBridgeProxy
	}
	if l1CrossDomainMessengerProxy, err := L1Deployments.Get("L1CrossDomainMessengerProxy"); err == nil {
		fmt.Printf("Using L1CrossDomainMessengerProxy: %s\n", l1CrossDomainMessengerProxy)
		DeployConfig.L1CrossDomainMessengerProxy = l1CrossDomainMessengerProxy
	}
	if l1ERC721BridgeProxy, err := L1Deployments.Get("L1ERC721BridgeProxy"); err == nil {
		fmt.Printf("Using L1ERC721BridgeProxy: %s\n", l1ERC721BridgeProxy)
		DeployConfig.L1ERC721BridgeProxy = l1ERC721BridgeProxy
	}
	if optimismPortalProxy, err := L1Deployments.Get("OptimismPortalProxy"); err == nil {
		fmt.Printf("Using OptimismPortalProxy: %s\n", optimismPortalProxy)
		DeployConfig.OptimismPortalProxy = optimismPortalProxy
	}
	if systemConfigProxy, err := L1Deployments.Get("SystemConfigProxy"); err == nil {
		fmt.Printf("Using SystemConfigProxy: %s\n", systemConfigProxy)
		DeployConfig.SystemConfigProxy = systemConfigProxy
	}

	// TODO: keep this logging?
	for k, v := range L1Deployments {
		fmt.Printf("%s: %s\n", k, v)
	}

	// TODO: completely remove this
	DeployConfig.CliqueSignerAddress = common.Address{}
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
