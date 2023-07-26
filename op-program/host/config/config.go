package config

import (
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum-optimism/optimism/op-program/chainconfig"
	"github.com/ethereum-optimism/optimism/op-program/host/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

var (
	ErrMissingL2ChainID    = errors.New("missing l2 chain ID")
	ErrInvalidL1Head       = errors.New("invalid l1 head")
	ErrInvalidL2Head       = errors.New("invalid l2 head")
	ErrInvalidL2OutputRoot = errors.New("invalid l2 output root")
	ErrL1AndL2Inconsistent = errors.New("l1 and l2 options must be specified together or both omitted")
	ErrInvalidL2Claim      = errors.New("invalid l2 claim")
	ErrInvalidL2ClaimBlock = errors.New("invalid l2 claim block number")
	ErrDataDirRequired     = errors.New("datadir must be specified when in non-fetching mode")
	ErrNoExecInServerMode  = errors.New("exec command must not be set when in server mode")
)

type Config struct {
	// DataDir is the directory to read/write pre-image data from/to.
	//If not set, an in-memory key-value store is used and fetching data must be enabled
	DataDir string

	// L1Head is the block has of the L1 chain head block
	L1Head     common.Hash
	L1URL      string
	L1TrustRPC bool
	L1RPCKind  sources.RPCProviderKind

	// L2Head is the l2 block hash contained in the L2 Output referenced by the L2OutputRoot
	// TODO(inphi): This can be made optional with hardcoded rollup configs and output oracle addresses by searching the oracle for the l2 output root
	L2Head common.Hash
	// L2OutputRoot is the agreed L2 output root to start derivation from
	L2OutputRoot common.Hash
	L2URL        string
	// L2Claim is the claimed L2 output root to verify
	L2Claim common.Hash
	// L2ClaimBlockNumber is the block number the claimed L2 output root is from
	// Must be above 0 and to be a valid claim needs to be above the L2Head block.
	L2ClaimBlockNumber uint64

	// L2ChainID is the chain ID of the L2 execution engine
	// For testing purposes, this can be set to MaxUint64 to use a preset chain configuration
	L2ChainID uint64

	// ExecCmd specifies the client program to execute in a separate process.
	// If unset, the fault proof client is run in the same process.
	ExecCmd string

	// ServerMode indicates that the program should run in pre-image server mode and wait for requests.
	// No client program is run.
	ServerMode bool
}

func (c *Config) Check() error {
	if c.L1Head == (common.Hash{}) {
		return ErrInvalidL1Head
	}
	if c.L2Head == (common.Hash{}) {
		return ErrInvalidL2Head
	}
	if c.L2OutputRoot == (common.Hash{}) {
		return ErrInvalidL2OutputRoot
	}
	if c.L2Claim == (common.Hash{}) {
		return ErrInvalidL2Claim
	}
	if c.L2ClaimBlockNumber == 0 {
		return ErrInvalidL2ClaimBlock
	}
	if c.L2ChainID == 0 {
		return ErrMissingL2ChainID
	}
	if (c.L1URL != "") != (c.L2URL != "") {
		return ErrL1AndL2Inconsistent
	}
	if !c.FetchingEnabled() && c.DataDir == "" {
		return ErrDataDirRequired
	}
	if c.ServerMode && c.ExecCmd != "" {
		return ErrNoExecInServerMode
	}
	return nil
}

func (c *Config) FetchingEnabled() bool {
	return c.L1URL != "" && c.L2URL != ""
}

// NewConfig creates a Config with all optional values set to the CLI default value
func NewConfig(
	l2ChainID uint64,
	l1Head common.Hash,
	l2Head common.Hash,
	l2OutputRoot common.Hash,
	l2Claim common.Hash,
	l2ClaimBlockNum uint64,
) *Config {
	return &Config{
		L2ChainID:          l2ChainID,
		L1Head:             l1Head,
		L2Head:             l2Head,
		L2OutputRoot:       l2OutputRoot,
		L2Claim:            l2Claim,
		L2ClaimBlockNumber: l2ClaimBlockNum,
		L1RPCKind:          sources.RPCKindBasic,
	}
}

func NewConfigFromCLI(log log.Logger, ctx *cli.Context) (*Config, error) {
	if err := flags.CheckRequired(ctx); err != nil {
		return nil, err
	}
	l2Head := common.HexToHash(ctx.String(flags.L2Head.Name))
	if l2Head == (common.Hash{}) {
		return nil, ErrInvalidL2Head
	}
	l2OutputRoot := common.HexToHash(ctx.String(flags.L2OutputRoot.Name))
	if l2OutputRoot == (common.Hash{}) {
		return nil, ErrInvalidL2OutputRoot
	}
	l2Claim := common.HexToHash(ctx.String(flags.L2Claim.Name))
	if l2Claim == (common.Hash{}) {
		return nil, ErrInvalidL2Claim
	}
	l2ClaimBlockNum := ctx.Uint64(flags.L2BlockNumber.Name)
	l1Head := common.HexToHash(ctx.String(flags.L1Head.Name))
	if l1Head == (common.Hash{}) {
		return nil, ErrInvalidL1Head
	}
	l2ChainID := ctx.Uint64(flags.L2ChainID.Name)
	if l2ChainID == 0 {
		config := chainconfig.L2ChainConfigsByName[ctx.String(flags.Network.Name)]
		if config == nil {
			return nil, fmt.Errorf("invalid network %s", ctx.String(flags.Network.Name))
		}
		l2ChainID = config.ChainID.Uint64()
	}
	if l2ChainID == 0 {
		return nil, fmt.Errorf("unknown chainID for network %s", ctx.String(flags.Network.Name))
	}
	return &Config{
		L2ChainID:          l2ChainID,
		DataDir:            ctx.String(flags.DataDir.Name),
		L2URL:              ctx.String(flags.L2NodeAddr.Name),
		L2Head:             l2Head,
		L2OutputRoot:       l2OutputRoot,
		L2Claim:            l2Claim,
		L2ClaimBlockNumber: l2ClaimBlockNum,
		L1Head:             l1Head,
		L1URL:              ctx.String(flags.L1NodeAddr.Name),
		L1TrustRPC:         ctx.Bool(flags.L1TrustRPC.Name),
		L1RPCKind:          sources.RPCProviderKind(ctx.String(flags.L1RPCProviderKind.Name)),
		ExecCmd:            ctx.String(flags.Exec.Name),
		ServerMode:         ctx.Bool(flags.Server.Name),
	}, nil
}
