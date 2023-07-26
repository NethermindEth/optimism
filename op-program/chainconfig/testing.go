package chainconfig

import (
	"math"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/params"
)

var (
	TestChainID = uint64(math.MaxUint64)

	testConf *testConfig
)

func SetTestConfigs(config *rollup.Config, l2Genesis *params.ChainConfig) {
	testConf.rollupConfig = config
	testConf.l2Genesis = l2Genesis
}

func init() {
	testConf = &testConfig{}
}

type testConfig struct {
	rollupConfig *rollup.Config
	l2Genesis    *params.ChainConfig
}

func (t *testConfig) initialized() bool {
	return t.rollupConfig != nil && t.l2Genesis != nil
}
