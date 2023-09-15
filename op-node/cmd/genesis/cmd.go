package genesis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-bindings/hardhat"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
)

var Subcommands = cli.Commands{
	{
		Name:  "l1",
		Usage: "Generates a L1 genesis state file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "deploy-config",
				Usage:    "Path to hardhat deploy config file",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "l1-allocs",
				Usage: "Path to L1 genesis state dump",
			},
			&cli.StringFlag{
				Name:  "l1-deployments",
				Usage: "Path to L1 deployments file",
			},
			&cli.StringFlag{
				Name:  "outfile.l1",
				Usage: "Path to L1 genesis output file",
			},
		},
		Action: func(ctx *cli.Context) error {
			deployConfig := ctx.String("deploy-config")
			config, err := genesis.NewDeployConfig(deployConfig)
			if err != nil {
				return err
			}

			var deployments *genesis.L1Deployments
			if l1Deployments := ctx.String("l1-deployments"); l1Deployments != "" {
				deployments, err = genesis.NewL1Deployments(l1Deployments)
				if err != nil {
					return err
				}
			}

			if deployments != nil {
				config.SetDeployments(deployments)
			}

			if err := config.Check(); err != nil {
				return fmt.Errorf("deploy config at %s invalid: %w", deployConfig, err)
			}

			// Check the addresses after setting the deployments
			if err := config.CheckAddresses(); err != nil {
				return fmt.Errorf("deploy config at %s invalid: %w", deployConfig, err)
			}

			var dump *state.Dump
			if l1Allocs := ctx.String("l1-allocs"); l1Allocs != "" {
				dump, err = genesis.NewStateDump(l1Allocs)
				if err != nil {
					return err
				}
			}

			l1Genesis, err := genesis.BuildL1DeveloperGenesis(config, dump, deployments, true)
			if err != nil {
				return err
			}

			return writeGenesisFile(ctx.String("outfile.l1"), l1Genesis)
		},
	},
	{
		Name:  "l2",
		Usage: "Generates an L2 genesis file and rollup config suitable for a deployed network",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "l1-rpc",
				Usage: "L1 RPC URL",
			},
			&cli.StringFlag{
				Name:  "deploy-config",
				Usage: "Path to deploy config file",
			},
			&cli.StringFlag{
				Name:  "deployment-dir",
				Usage: "Path to network deployment directory",
			},
			&cli.StringFlag{
				Name:  "outfile.l2",
				Usage: "Path to L2 genesis output file",
			},
			&cli.StringFlag{
				Name:  "outfile.rollup",
				Usage: "Path to rollup output file",
			},
		},
		Action: func(ctx *cli.Context) error {
			deployConfig := ctx.String("deploy-config")
			log.Info("Deploy config", "path", deployConfig)
			config, err := genesis.NewDeployConfig(deployConfig)
			if err != nil {
				return err
			}

			deployDir := ctx.String("deployment-dir")
			if deployDir == "" {
				return errors.New("Must specify --deployment-dir")
			}

			log.Info("Deployment directory", "path", deployDir)
			depPath, network := filepath.Split(deployDir)
			hh, err := hardhat.New(network, nil, []string{depPath})
			if err != nil {
				return err
			}

			// Read the appropriate deployment addresses from disk
			if err := config.GetDeployedAddresses(hh); err != nil {
				return err
			}

			client, err := ethclient.Dial(ctx.String("l1-rpc"))
			if err != nil {
				return fmt.Errorf("cannot dial %s: %w", ctx.String("l1-rpc"), err)
			}

			var l1StartBlock *types.Block
			if config.L1StartingBlockTag == nil {
				l1StartBlock, err = client.BlockByNumber(context.Background(), nil)
				tag := rpc.BlockNumberOrHashWithHash(l1StartBlock.Hash(), true)
				config.L1StartingBlockTag = (*genesis.MarshalableRPCBlockNumberOrHash)(&tag)
			} else if config.L1StartingBlockTag.BlockHash != nil {
				l1StartBlock, err = client.BlockByHash(context.Background(), *config.L1StartingBlockTag.BlockHash)
			} else if config.L1StartingBlockTag.BlockNumber != nil {
				l1StartBlock, err = client.BlockByNumber(context.Background(), big.NewInt(config.L1StartingBlockTag.BlockNumber.Int64()))
			}
			if err != nil {
				return fmt.Errorf("error getting l1 start block: %w", err)
			}

			// Sanity check the config. Do this after filling in the L1StartingBlockTag
			// if it is not defined.
			if err := config.Check(); err != nil {
				return err
			}

			log.Info("Using L1 Start Block", "number", l1StartBlock.Number(), "hash", l1StartBlock.Hash().Hex())

			// Build the L2 genesis block
			l2Genesis, err := genesis.BuildL2Genesis(config, l1StartBlock)
			if err != nil {
				return fmt.Errorf("error creating l2 genesis: %w", err)
			}

			l2GenesisBlock := l2Genesis.ToBlock()
			rollupConfig, err := config.RollupConfig(l1StartBlock, l2GenesisBlock.Hash(), l2GenesisBlock.Number().Uint64())
			if err != nil {
				return err
			}
			if err := rollupConfig.Check(); err != nil {
				return fmt.Errorf("generated rollup config does not pass validation: %w", err)
			}

			if err := writeGenesisFile(fmt.Sprintf("%s-geth", ctx.String("outfile.l2")), l2Genesis); err != nil {
				return err
			}
			l2NethGenesis := toNethGenesis(l2Genesis)
			if err := writeGenesisFile(fmt.Sprintf("%s-neth", ctx.String("outfile.l2")), l2NethGenesis); err != nil {
				return fmt.Errorf("error writing neth genesis: %w", err)
			}
			return writeGenesisFile(ctx.String("outfile.rollup"), rollupConfig)
		},
	},
}

func writeGenesisFile(outfile string, input any) error {
	f, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(input)
}

type nethSpecEngine struct {
	Params nethSpecEngineParams `json:"params"`
}

type nethSpecEngineParams struct {
	RegolithTimestamp  hexutil.Uint64 `json:"regolithTimestamp"`
	BedrockBlockNumber *hexutil.Big   `json:"bedrockBlockNumber"`
}

type nethSpecParams struct {
	ChainID                            string         `json:"chainId"`
	// GasLimitBoundDivisor               hexutil.Uint   `json:"gasLimitBoundDivisor"`
	// AccountStartNonce                  hexutil.Uint   `json:"accountStartNonce"`
	// MaximumExtraDataSize               hexutil.Uint   `json:"maximumExtraDataSize"`
	// MinGasLimit                        hexutil.Uint   `json:"minGasLimit"`
	// ForkBlock                          hexutil.Uint   `json:"forkBlock"`
	MaxCodeSize                        hexutil.Uint   `json:"maxCodeSize"`
	MaxCodeSizeTransition              hexutil.Uint   `json:"maxCodeSizeTransition"`
	Eip150Transition                   hexutil.Uint   `json:"eip150Transition"`
	Eip160Transition                   hexutil.Uint   `json:"eip160Transition"`
	Eip161AbcTransition                hexutil.Uint   `json:"eip161abcTransition"`
	Eip161DTransition                  hexutil.Uint   `json:"eip161dTransition"`
	Eip155Transition                   hexutil.Uint   `json:"eip155Transition"`
	Eip140Transition                   hexutil.Uint   `json:"eip140Transition"`
	Eip211Transition                   hexutil.Uint   `json:"eip211Transition"`
	Eip214Transition                   hexutil.Uint   `json:"eip214Transition"`
	Eip658Transition                   hexutil.Uint   `json:"eip658Transition"`
	Eip145Transition                   hexutil.Uint   `json:"eip145Transition"`
	Eip1014Transition                  hexutil.Uint   `json:"eip1014Transition"`
	Eip1052Transition                  hexutil.Uint   `json:"eip1052Transition"`
	Eip1283Transition                  hexutil.Uint   `json:"eip1283Transition"`
	Eip1283DisableTransition           hexutil.Uint   `json:"eip1283DisableTransition"`
	Eip152Transition                   hexutil.Uint   `json:"eip152Transition"`
	Eip1108Transition                  hexutil.Uint   `json:"eip1108Transition"`
	Eip1344Transition                  hexutil.Uint   `json:"eip1344Transition"`
	Eip1884Transition                  hexutil.Uint   `json:"eip1884Transition"`
	Eip2028Transition                  hexutil.Uint   `json:"eip2028Transition"`
	Eip2200Transition                  hexutil.Uint   `json:"eip2200Transition"`
	Eip2565Transition                  hexutil.Uint   `json:"eip2565Transition"`
	Eip2929Transition                  hexutil.Uint   `json:"eip2929Transition"`
	Eip2930Transition                  hexutil.Uint   `json:"eip2930Transition"`
	Eip1559Transition                  hexutil.Uint   `json:"eip1559Transition"`
	Eip1559ElasticityMultiplier        hexutil.Uint   `json:"eip1559ElasticityMultiplier"`
	Eip1559BaseFeeMaxChangeDenominator hexutil.Uint   `json:"eip1559BaseFeeMaxChangeDenominator"`
	Eip1559FeeCollectorTransition      hexutil.Uint   `json:"eip1559FeeCollectorTransition"`
	Eip1559FeeCollector                common.Address `json:"eip1559FeeCollector"`
	Eip3198Transition                  hexutil.Uint   `json:"eip3198Transition"`
	Eip3529Transition                  hexutil.Uint   `json:"eip3529Transition"`
	Eip3541Transition                  hexutil.Uint   `json:"eip3541Transition"`
	// Eip4895TransitionTimestamp         hexutil.Uint   `json:"eip4895TransitionTimestamp"`
	// Eip3855TransitionTimestamp         hexutil.Uint   `json:"eip3855TransitionTimestamp"`
	// Eip3651TransitionTimestamp         hexutil.Uint   `json:"eip3651TransitionTimestamp"`
	// Eip3860TransitionTimestamp         hexutil.Uint   `json:"eip3860TransitionTimestamp"`
	TerminalTotalDifficulty            hexutil.Uint   `json:"terminalTotalDifficulty"`
}

type nethSpecGenesis struct {
	Number        hexutil.Uint        `json:"number"`
	Difficulty    hexutil.Big         `json:"difficulty"`
	Author        common.Address      `json:"author"`
	Timestamp     hexutil.Uint        `json:"timestamp"`
	ParentHash    common.Hash         `json:"parentHash"`
	ExtraData     hexutil.Bytes       `json:"extraData"`
	GasLimit      hexutil.Uint        `json:"gasLimit"`
	BaseFeePerGas hexutil.Big         `json:"baseFeePerGas"`
	Seal          nethSpecGenesisSeal `json:"seal"`

	// StateUnavailable bool           `json:"stateUnavailable"`
	// StateRoot        string         `json:"stateRoot"`
}

type nethSpecGenesisSeal struct {
	Ethereum ethereumSeal `json:"ethereum"`
}

type ethereumSeal struct {
	Nonce   hexutil.Uint64 `json:"nonce"`
	MixHash common.Hash    `json:"mixHash"`
}

type nethSpec struct {
	Name     string                    `json:"name"`
	Engine   map[string]nethSpecEngine `json:"engine"`
	Params   nethSpecParams            `json:"params"`
	Genesis  nethSpecGenesis           `json:"genesis"`
	Accounts map[string]any            `json:"accounts"`
}

func toNethGenesis(l2Genesis *core.Genesis) *nethSpec {
	spec := &nethSpec{
		Name: "OPDevnet",
		Engine: map[string]nethSpecEngine{
			"Optimism": {
				Params: nethSpecEngineParams{
					RegolithTimestamp:  hexutil.Uint64(*l2Genesis.Config.RegolithTime),
					BedrockBlockNumber: (*hexutil.Big)(l2Genesis.Config.BedrockBlock),
				},
			},
		},
		Genesis: nethSpecGenesis{
			Number:        hexutil.Uint(l2Genesis.Number),
			Difficulty:    hexutil.Big(*l2Genesis.Difficulty),
			Author:        l2Genesis.Coinbase,
			Timestamp:     hexutil.Uint(l2Genesis.Timestamp),
			ParentHash:    l2Genesis.ParentHash,
			ExtraData:     l2Genesis.ExtraData,
			GasLimit:      hexutil.Uint(l2Genesis.GasLimit),
			BaseFeePerGas: hexutil.Big(*l2Genesis.BaseFee),
			Seal: nethSpecGenesisSeal{
				Ethereum: ethereumSeal{
					Nonce:   hexutil.Uint64(l2Genesis.Nonce),
					MixHash: l2Genesis.Mixhash,
				},
			},
		},
		Params: nethSpecParams{
			ChainID: l2Genesis.Config.ChainID.String(),

			MaxCodeSize:           hexutil.Uint(0x6000),
			MaxCodeSizeTransition: hexutil.Uint(0),

			Eip150Transition:    hexutil.Uint(0),
			Eip160Transition:    hexutil.Uint(0),
			Eip161AbcTransition: hexutil.Uint(0),
			Eip161DTransition:   hexutil.Uint(0),
			Eip155Transition:    hexutil.Uint(0),
			Eip140Transition:    hexutil.Uint(0),
			Eip211Transition:    hexutil.Uint(0),
			Eip214Transition:    hexutil.Uint(0),
			Eip658Transition:    hexutil.Uint(0),
			Eip145Transition:    hexutil.Uint(0),
			Eip1014Transition:   hexutil.Uint(0),
			Eip1052Transition:   hexutil.Uint(0),
			Eip1283Transition:   hexutil.Uint(0),
			Eip152Transition:    hexutil.Uint(0),
			Eip1108Transition:   hexutil.Uint(0),
			Eip1344Transition:   hexutil.Uint(0),
			Eip1884Transition:   hexutil.Uint(0),
			Eip2028Transition:   hexutil.Uint(0),
			Eip2200Transition:   hexutil.Uint(0),
			Eip2565Transition:   hexutil.Uint(0),
			Eip2929Transition:   hexutil.Uint(0),
			Eip2930Transition:   hexutil.Uint(0),

			Eip1559Transition:                  hexutil.Uint(0),
			Eip1559FeeCollectorTransition:      hexutil.Uint(0),
			Eip1559FeeCollector:                params.OptimismBaseFeeRecipient,
			Eip1559ElasticityMultiplier:        hexutil.Uint(l2Genesis.Config.Optimism.EIP1559Elasticity),
			Eip1559BaseFeeMaxChangeDenominator: hexutil.Uint(l2Genesis.Config.Optimism.EIP1559Denominator),

			Eip3198Transition:          hexutil.Uint(0),
			Eip3529Transition:          hexutil.Uint(0),
			Eip3541Transition:          hexutil.Uint(0),
			// Eip4895TransitionTimestamp: hexutil.Uint(0),
			// Eip3855TransitionTimestamp: hexutil.Uint(0),
			// Eip3651TransitionTimestamp: hexutil.Uint(0),
			// Eip3860TransitionTimestamp: hexutil.Uint(0),

			TerminalTotalDifficulty: hexutil.Uint(0),
		},
		Accounts: map[string]any{
			"0x0000000000000000000000000000000000000001": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "ecrecover",
					"pricing": map[string]interface{}{
						"linear": map[string]interface{}{
							"base": 3000,
							"word": 0,
						},
					},
				},
			},
			"0x0000000000000000000000000000000000000002": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "sha256",
					"pricing": map[string]interface{}{
						"linear": map[string]interface{}{
							"base": 60,
							"word": 12,
						},
					},
				},
			},
			"0x0000000000000000000000000000000000000003": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "ripemd160",
					"pricing": map[string]interface{}{
						"linear": map[string]interface{}{
							"base": 600,
							"word": 120,
						},
					},
				},
			},
			"0x0000000000000000000000000000000000000004": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "identity",
					"pricing": map[string]interface{}{
						"linear": map[string]interface{}{
							"base": 15,
							"word": 3,
						},
					},
				},
			},
			"0x0000000000000000000000000000000000000005": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "modexp",
					// "activate_at": "0x42ae50",
					"activate_at": "0x0",
					"pricing": map[string]interface{}{
						"modexp": map[string]interface{}{
							"divisor": 20,
						},
					},
				},
			},
			"0x0000000000000000000000000000000000000006": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "alt_bn128_add",
					// "activate_at": "0x42ae50",
					"activate_at": "0x0",
					"pricing": map[string]interface{}{
						"linear": map[string]interface{}{
							"base": 500,
							"word": 0,
						},
					},
				},
			},
			"0x0000000000000000000000000000000000000007": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "alt_bn128_mul",
					// "activate_at": "0x42ae50",
					"activate_at": "0x0",
					"pricing": map[string]interface{}{
						"linear": map[string]interface{}{
							"base": 40000,
							"word": 0,
						},
					},
				},
			},
			"0x0000000000000000000000000000000000000008": map[string]interface{}{
				"builtin": map[string]interface{}{
					"name": "alt_bn128_pairing",
					// "activate_at": "0x42ae50",
					"activate_at": "0x0",
					"pricing": map[string]interface{}{
						"alt_bn128_pairing": map[string]interface{}{
							"base": 100000,
							"pair": 80000,
						},
					},
				},
			},
		},
	}

	for addr, alloc := range l2Genesis.Alloc {
		nethAlloc := map[string]interface{}{}
		if alloc.Code != nil && len(alloc.Code) > 0 {
			nethAlloc["code"]= hexutil.Bytes(alloc.Code)
		}
		if alloc.Storage != nil && len(alloc.Storage) > 0 {
			storage := make(map[string]string)
			for k, v := range alloc.Storage {
				storage[k.Hex()] = v.Hex()
			}
			nethAlloc["storage"] = storage
		}
		if alloc.Balance != nil {
			nethAlloc["balance"] = hexutil.EncodeBig(alloc.Balance)
		}
		if alloc.Nonce > 0 {
			nethAlloc["nonce"] = hexutil.Uint64(alloc.Nonce)
		}
		spec.Accounts[addr.Hex()] = nethAlloc
	}

	return spec
}
