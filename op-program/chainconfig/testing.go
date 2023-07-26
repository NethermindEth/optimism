package chainconfig

import (
	"encoding/json"
	"io/ioutil"
	"math"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/params"
)

var (
	TestChainID = uint64(math.MaxUint64)

	testRollupConfigPath = "fpp_rollup_config.json"
	testL2GenesisPath    = "fpp_l2_genesis.json"
)

func SetTestConfig(config *rollup.Config, l2Genesis *params.ChainConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(testRollupConfigPath, data, 0644)
	if err != nil {
		return err
	}

	data, err = json.Marshal(l2Genesis)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(testL2GenesisPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func TestRollupConfig() *rollup.Config {
	var config rollup.Config
	data, err := ioutil.ReadFile(testRollupConfigPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic("failed to unmarshal test rollup config")
	}
	// TODO(inphi): cache this
	return &config
}

func TestL2Genesis() *params.ChainConfig {
	var config params.ChainConfig
	data, err := ioutil.ReadFile(testL2GenesisPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic("failed to unmarshal test l2 genesis")
	}
	return &config
}
