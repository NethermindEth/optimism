package genesis

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	gstate "github.com/ethereum/go-ethereum/core/state"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-chain-ops/state"
)

var (
	// uint128Max is type(uint128).max and is set in the init function.
	uint128Max = new(big.Int)
	// The default values for the ResourceConfig, used as part of
	// an EIP-1559 curve for deposit gas.
	defaultResourceConfig = bindings.ResourceMeteringResourceConfig{
		MaxResourceLimit:            20_000_000,
		ElasticityMultiplier:        10,
		BaseFeeMaxChangeDenominator: 8,
		MinimumBaseFee:              params.GWei,
		SystemTxMaxGas:              1_000_000,
	}
)

func init() {
	var ok bool
	uint128Max, ok = new(big.Int).SetString("ffffffffffffffffffffffffffffffff", 16)
	if !ok {
		panic("bad uint128Max")
	}
	// Set the maximum base fee on the default config.
	defaultResourceConfig.MaximumBaseFee = uint128Max
}

// BuildL1DeveloperGenesis will create a L1 genesis block after creating
// all of the state required for an Optimism network to function.
func BuildL1DeveloperGenesis(config *DeployConfig, dump *gstate.Dump) (*core.Genesis, error) {
	if config.L2OutputOracleStartingTimestamp != -1 {
		return nil, errors.New("l2oo starting timestamp must be -1")
	}

	if config.L1GenesisBlockTimestamp == 0 {
		return nil, errors.New("must specify l1 genesis block timestamp")
	}

	genesis, err := NewL1Genesis(config)
	if err != nil {
		return nil, err
	}

	memDB := state.NewMemoryStateDB(genesis)

	FundDevAccounts(memDB)
	SetPrecompileBalances(memDB)

	if dump != nil {
		log.Info("Building developer L1 genesis block")
		for address, account := range dump.Accounts {
			log.Info("Setting account", "address", address.Hex())
			memDB.CreateAccount(address)
			memDB.SetNonce(address, account.Nonce)

			balance, ok := new(big.Int).SetString(account.Balance, 10)
			if !ok {
				return nil, fmt.Errorf("failed to parse balance for %s", address)
			}
			memDB.AddBalance(address, balance)
			memDB.SetCode(address, account.Code)
			for key, value := range account.Storage {
				memDB.SetState(address, key, common.HexToHash(value))
			}
		}
	}
	return memDB.Genesis(), nil
}
