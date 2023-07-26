package chainconfig

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/params"
)

var (
	TestChainID = uint64(math.MaxUint64)

	testRollupConfigPath = "fpp_rollup_config.json"
	testL2GenesisPath    = "fpp_l2_genesis.json"

	testRollupConfig *rollup.Config
	testL2Genesis    *params.ChainConfig
)

func SetTestConfig(config *rollup.Config, l2Genesis *params.ChainConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal test rollup config: %w", err)
	}
	err = os.WriteFile(testRollupConfigPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write test rollup config: %w", err)
	}

	data, err = json.Marshal(l2Genesis)
	if err != nil {
		return fmt.Errorf("failed to marshal test l2 genesis: %w", err)
	}
	err = os.WriteFile(testL2GenesisPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write test l2 genesis: %w", err)
	}
	return nil
}

func TestRollupConfig() *rollup.Config {
	var config rollup.Config
	data, err := os.ReadFile(testRollupConfigPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic("failed to unmarshal test rollup config")
	}
	testRollupConfig = &config
	return &config
}

func TestL2Genesis() *params.ChainConfig {
	var config params.ChainConfig
	data, err := os.ReadFile(testL2GenesisPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic("failed to unmarshal test l2 genesis")
	}
	testL2Genesis = &config
	return &config
}
