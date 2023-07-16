// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MIPSMetaData contains all meta data concerning the MIPS contract.
var MIPSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BRK_START\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIPreimageOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"step\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611c33806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063155633fe146100465780637dc0d1d014610067578063f8e0cb9614610098575b600080fd5b61004e61016c565b6040805163ffffffff9092168252519081900360200190f35b61006f610174565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61015a600480360360408110156100ae57600080fd5b8101906020810181356401000000008111156100c957600080fd5b8201836020820111156100db57600080fd5b803590602001918460018302840111640100000000831117156100fd57600080fd5b91939092909160208101903564010000000081111561011b57600080fd5b82018360208201111561012d57600080fd5b8035906020019184600183028401116401000000008311171561014f57600080fd5b509092509050610190565b60408051918252519081900360200190f35b634000000081565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600061019a611b79565b608081146101a757600080fd5b604051610600146101b757600080fd5b606486146101c457600080fd5b61016684146101d257600080fd5b6101ef565b8035602084810360031b9190911c8352920192910190565b8560806101fe602082846101d7565b9150915061020e602082846101d7565b9150915061021e600482846101d7565b9150915061022e600482846101d7565b9150915061023e600482846101d7565b9150915061024e600482846101d7565b9150915061025e600482846101d7565b9150915061026e600482846101d7565b9150915061027e600182846101d7565b9150915061028e600182846101d7565b9150915061029e600882846101d7565b6020810190819052909250905060005b60208110156102d0576102c3600483856101d7565b90935091506001016102ae565b505050806101200151156102ee576102e6610710565b915050610708565b6101408101805160010167ffffffffffffffff1690526060810151600090610316908261081e565b9050603f601a82901c16600281148061033557508063ffffffff166003145b15610382576103788163ffffffff1660021461035257601f610355565b60005b60ff16600261036b856303ffffff16601a6108e9565b63ffffffff16901b61095c565b9350505050610708565b6101608301516000908190601f601086901c81169190601587901c16602081106103a857fe5b602002015192508063ffffffff851615806103c957508463ffffffff16601c145b156103fa578661016001518263ffffffff16602081106103e557fe5b6020020151925050601f600b86901c166104b1565b60208563ffffffff16101561045d578463ffffffff16600c148061042457508463ffffffff16600d145b8061043557508463ffffffff16600e145b15610446578561ffff169250610458565b6104558661ffff1660106108e9565b92505b6104b1565b60288563ffffffff1610158061047957508463ffffffff166022145b8061048a57508463ffffffff166026145b156104b1578661016001518263ffffffff16602081106104a657fe5b602002015192508190505b60048563ffffffff16101580156104ce575060088563ffffffff16105b806104df57508463ffffffff166001145b156104fe576104f085878487610a50565b975050505050505050610708565b63ffffffff60006020878316106105635761051e8861ffff1660106108e9565b9095019463fffffffc861661053481600161081e565b915060288863ffffffff161015801561055457508763ffffffff16603014155b1561056157809250600093505b505b600061057189888885610c5f565b63ffffffff9081169150603f8a16908916158015610596575060088163ffffffff1610155b80156105a85750601c8163ffffffff16105b15610687578063ffffffff16600814806105c857508063ffffffff166009145b156105ff576105ed8163ffffffff166008146105e457856105e7565b60005b8961095c565b9b505050505050505050505050610708565b8063ffffffff16600a1415610620576105ed858963ffffffff8a1615611325565b8063ffffffff16600b1415610642576105ed858963ffffffff8a161515611325565b8063ffffffff16600c1415610659576105ed61140a565b60108163ffffffff16101580156106765750601c8163ffffffff16105b15610687576105ed81898988611887565b8863ffffffff1660381480156106a2575063ffffffff861615155b156106d15760018b61016001518763ffffffff16602081106106c057fe5b63ffffffff90921660209290920201525b8363ffffffff1663ffffffff146106ee576106ee84600184611a6b565b6106fa85836001611325565b9b5050505050505050505050505b949350505050565b6000610728565b602083810382015183520192910190565b60806040518061073a60208285610717565b9150925061074a60208285610717565b9150925061075a60048285610717565b9150925061076a60048285610717565b9150925061077a60048285610717565b9150925061078a60048285610717565b9150925061079a60048285610717565b915092506107aa60048285610717565b915092506107ba60018285610717565b915092506107ca60018285610717565b915092506107da60088285610717565b60209091019350905060005b6020811015610808576107fb60048386610717565b90945091506001016107e6565b506000815281810382a081900390209150505b90565b60008061082a83611b07565b9050600384161561083a57600080fd5b602081019035610857565b60009081526020919091526040902090565b8460051c8160005b601b8110156108af5760208501943583821c60011680156108875760018114610898576108a5565b6108918285610845565b93506108a5565b6108a28483610845565b93505b505060010161085f565b5060805191508181146108ca57630badf00d60005260206000fd5b5050601f94909416601c0360031b9390931c63ffffffff169392505050565b600063ffffffff8381167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80850183169190911c821615159160016020869003821681901b830191861691821b92911b0182610946576000610948565b815b90861663ffffffff16179250505092915050565b6000610966611b79565b60809050806060015160040163ffffffff16816080015163ffffffff16146109ef57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6a756d7020696e2064656c617920736c6f740000000000000000000000000000604482015290519081900360640190fd5b60608101805160808301805163ffffffff908116909352858316905290851615610a3f57806008018261016001518663ffffffff1660208110610a2e57fe5b63ffffffff90921660209290920201525b610a47610710565b95945050505050565b6000610a5a611b79565b608090506000816060015160040163ffffffff16826080015163ffffffff1614610ae557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6272616e636820696e2064656c617920736c6f74000000000000000000000000604482015290519081900360640190fd5b8663ffffffff1660041480610b0057508663ffffffff166005145b15610b765760008261016001518663ffffffff1660208110610b1e57fe5b602002015190508063ffffffff168563ffffffff16148015610b4657508763ffffffff166004145b80610b6e57508063ffffffff168563ffffffff1614158015610b6e57508763ffffffff166005145b915050610bf3565b8663ffffffff1660061415610b945760008460030b13159050610bf3565b8663ffffffff1660071415610bb15760008460030b139050610bf3565b8663ffffffff1660011415610bf357601f601087901c1680610bd75760008560030b1291505b8063ffffffff1660011415610bf15760008560030b121591505b505b606082018051608084015163ffffffff169091528115610c39576002610c1e8861ffff1660106108e9565b63ffffffff90811690911b8201600401166080840152610c4b565b60808301805160040163ffffffff1690525b610c53610710565b98975050505050505050565b6000603f601a86901c811690861660208210156110335760088263ffffffff1610158015610c935750600f8263ffffffff16105b15610d3a578163ffffffff1660081415610caf57506020610d35565b8163ffffffff1660091415610cc657506021610d35565b8163ffffffff16600a1415610cdd5750602a610d35565b8163ffffffff16600b1415610cf45750602b610d35565b8163ffffffff16600c1415610d0b57506024610d35565b8163ffffffff16600d1415610d2257506025610d35565b8163ffffffff16600e1415610d35575060265b600091505b63ffffffff8216610f8357601f600688901c16602063ffffffff83161015610e575760088263ffffffff1610610d7557869350505050610708565b63ffffffff8216610d955763ffffffff86811691161b9250610708915050565b8163ffffffff1660021415610db95763ffffffff86811691161c9250610708915050565b8163ffffffff1660031415610de4576103788163ffffffff168763ffffffff16901c826020036108e9565b8163ffffffff1660041415610e08575050505063ffffffff8216601f84161b610708565b8163ffffffff1660061415610e2c575050505063ffffffff8216601f84161c610708565b8163ffffffff1660071415610e57576103788763ffffffff168763ffffffff16901c886020036108e9565b8163ffffffff1660201480610e7257508163ffffffff166021145b15610e84578587019350505050610708565b8163ffffffff1660221480610e9f57508163ffffffff166023145b15610eb1578587039350505050610708565b8163ffffffff1660241415610ecd578587169350505050610708565b8163ffffffff1660251415610ee9578587179350505050610708565b8163ffffffff1660261415610f05578587189350505050610708565b8163ffffffff1660271415610f21575050505082821719610708565b8163ffffffff16602a1415610f54578560030b8760030b12610f44576000610f47565b60015b60ff169350505050610708565b8163ffffffff16602b1415610f7d578563ffffffff168763ffffffff1610610f44576000610f47565b5061102e565b8163ffffffff16600f1415610fa65760108563ffffffff16901b92505050610708565b8163ffffffff16601c141561102e578063ffffffff1660021415610fcf57505050828202610708565b8063ffffffff1660201480610fea57508063ffffffff166021145b1561102e578063ffffffff1660201415611002579419945b60005b6380000000871615611024576401fffffffe600197881b169601611005565b9250610708915050565b6112be565b60288263ffffffff16101561119d578163ffffffff1660201415611080576110778660031660080260180363ffffffff168563ffffffff16901c60ff1660086108e9565b92505050610708565b8163ffffffff16602114156110b6576110778660021660080260100363ffffffff168563ffffffff16901c61ffff1660106108e9565b8163ffffffff16602214156110e75750505063ffffffff60086003851602811681811b198416918316901b17610708565b8163ffffffff1660231415611100578392505050610708565b8163ffffffff1660241415611134578560031660080260180363ffffffff168463ffffffff16901c60ff1692505050610708565b8163ffffffff1660251415611169578560021660080260100363ffffffff168463ffffffff16901c61ffff1692505050610708565b8163ffffffff166026141561102e5750505063ffffffff60086003851602601803811681811c198416918316901c17610708565b8163ffffffff16602814156111d55750505060ff63ffffffff60086003861602601803811682811b9091188316918416901b17610708565b8163ffffffff166029141561120e5750505061ffff63ffffffff60086002861602601003811682811b9091188316918416901b17610708565b8163ffffffff16602a141561123f5750505063ffffffff60086003851602811681811c198316918416901c17610708565b8163ffffffff16602b1415611258578492505050610708565b8163ffffffff16602e141561128c5750505063ffffffff60086003851602601803811681811b198316918416901b17610708565b8163ffffffff16603014156112a5578392505050610708565b8163ffffffff16603814156112be578492505050610708565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e76616c696420696e737472756374696f6e00000000000000000000000000604482015290519081900360640190fd5b600061132f611b79565b506080602063ffffffff8616106113a757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f76616c6964207265676973746572000000000000000000000000000000000000604482015290519081900360640190fd5b63ffffffff8516158015906113b95750825b156113e757838161016001518663ffffffff16602081106113d657fe5b63ffffffff90921660209290920201525b60808101805163ffffffff80821660608501526004909101169052610a47610710565b6000611414611b79565b506101e051604081015160808083015160a084015160c09094015191936000928392919063ffffffff8616610ffa141561148c5781610fff81161561145e57610fff811661100003015b63ffffffff84166114825760e08801805163ffffffff838201169091529550611486565b8395505b5061183a565b8563ffffffff16610fcd14156114a8576340000000945061183a565b8563ffffffff1661101814156114c1576001945061183a565b8563ffffffff1661109614156114f957600161012088015260ff83166101008801526114eb610710565b97505050505050505061081b565b8563ffffffff16610fa314156116b85763ffffffff8316611519576116b3565b63ffffffff83166005141561169057600061153b8363fffffffc16600161081e565b6000805460208b01516040808d015181517fe03110e1000000000000000000000000000000000000000000000000000000008152600481019390935263ffffffff16602483015280519495509293849373ffffffffffffffffffffffffffffffffffffffff9093169263e03110e19260448082019391829003018186803b1580156115c557600080fd5b505afa1580156115d9573d6000803e3d6000fd5b505050506040513d60408110156115ef57600080fd5b508051602090910151909250905060038516600481900382811015611612578092505b508185101561161f578491505b8260088302610100031c9250826008828460040303021b9250600180600883600403021b036001806008858560040303021b039150811981169050838119861617945050506116768563fffffffc16600185611a6b565b60408a018051820163ffffffff16905296506116b3915050565b63ffffffff8316600314156116a7578094506116b3565b63ffffffff9450600993505b61183a565b8563ffffffff16610fa4141561178c5763ffffffff8316600114806116e3575063ffffffff83166002145b806116f4575063ffffffff83166004145b15611701578094506116b3565b63ffffffff8316600614156116a75760006117238363fffffffc16600161081e565b6020890151909150600384166004038381101561173e578093505b83900360089081029290921c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600193850293841b0116911b176020880152600060408801529350836116b3565b8563ffffffff16610fd7141561183a578163ffffffff166003141561182e5763ffffffff831615806117c4575063ffffffff83166005145b806117d5575063ffffffff83166003145b156117e357600094506116b3565b63ffffffff8316600114806117fe575063ffffffff83166002145b8061180f575063ffffffff83166006145b80611820575063ffffffff83166004145b156116a757600194506116b3565b63ffffffff9450601693505b6101608701805163ffffffff808816604090920191909152905185821660e09091015260808801805180831660608b0152600401909116905261187b610710565b97505050505050505090565b6000611891611b79565b5060806000601063ffffffff881614156118b0575060c0810151611a08565b8663ffffffff16601114156118d05763ffffffff861660c0830152611a08565b8663ffffffff16601214156118ea575060a0810151611a08565b8663ffffffff166013141561190a5763ffffffff861660a0830152611a08565b8663ffffffff166018141561193f5763ffffffff600387810b9087900b02602081901c821660c08501521660a0830152611a08565b8663ffffffff16601914156119715763ffffffff86811681871602602081901c821660c08501521660a0830152611a08565b8663ffffffff16601a14156119bc578460030b8660030b8161198f57fe5b0763ffffffff1660c0830152600385810b9087900b816119ab57fe5b0563ffffffff1660a0830152611a08565b8663ffffffff16601b1415611a08578463ffffffff168663ffffffff16816119e057fe5b0663ffffffff90811660c0840152858116908716816119fb57fe5b0463ffffffff1660a08301525b63ffffffff841615611a3d57808261016001518563ffffffff1660208110611a2c57fe5b63ffffffff90921660209290920201525b60808201805163ffffffff80821660608601526004909101169052611a60610710565b979650505050505050565b6000611a7683611b07565b90506003841615611a8657600080fd5b6020810190601f8516601c0360031b83811b913563ffffffff90911b1916178460051c60005b601b811015611afc5760208401933582821c6001168015611ad45760018114611ae557611af2565b611ade8286610845565b9450611af2565b611aef8583610845565b94505b5050600101611aac565b505060805250505050565b60ff81166103800261016681019036906104e601811015611b73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611c046023913960400191505060405180910390fd5b50919050565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101919091526101608101611bdf611be4565b905290565b604051806104000160405280602090602082028036833750919291505056fe636865636b207468617420746865726520697320656e6f7567682063616c6c64617461a164736f6c6343000706000a",
}

// MIPSABI is the input ABI used to generate the binding from.
// Deprecated: Use MIPSMetaData.ABI instead.
var MIPSABI = MIPSMetaData.ABI

// MIPSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MIPSMetaData.Bin instead.
var MIPSBin = MIPSMetaData.Bin

// DeployMIPS deploys a new Ethereum contract, binding an instance of MIPS to it.
func DeployMIPS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MIPS, error) {
	parsed, err := MIPSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MIPSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MIPS{MIPSCaller: MIPSCaller{contract: contract}, MIPSTransactor: MIPSTransactor{contract: contract}, MIPSFilterer: MIPSFilterer{contract: contract}}, nil
}

// MIPS is an auto generated Go binding around an Ethereum contract.
type MIPS struct {
	MIPSCaller     // Read-only binding to the contract
	MIPSTransactor // Write-only binding to the contract
	MIPSFilterer   // Log filterer for contract events
}

// MIPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type MIPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MIPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MIPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MIPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MIPSSession struct {
	Contract     *MIPS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MIPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MIPSCallerSession struct {
	Contract *MIPSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MIPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MIPSTransactorSession struct {
	Contract     *MIPSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MIPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type MIPSRaw struct {
	Contract *MIPS // Generic contract binding to access the raw methods on
}

// MIPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MIPSCallerRaw struct {
	Contract *MIPSCaller // Generic read-only contract binding to access the raw methods on
}

// MIPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MIPSTransactorRaw struct {
	Contract *MIPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMIPS creates a new instance of MIPS, bound to a specific deployed contract.
func NewMIPS(address common.Address, backend bind.ContractBackend) (*MIPS, error) {
	contract, err := bindMIPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MIPS{MIPSCaller: MIPSCaller{contract: contract}, MIPSTransactor: MIPSTransactor{contract: contract}, MIPSFilterer: MIPSFilterer{contract: contract}}, nil
}

// NewMIPSCaller creates a new read-only instance of MIPS, bound to a specific deployed contract.
func NewMIPSCaller(address common.Address, caller bind.ContractCaller) (*MIPSCaller, error) {
	contract, err := bindMIPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MIPSCaller{contract: contract}, nil
}

// NewMIPSTransactor creates a new write-only instance of MIPS, bound to a specific deployed contract.
func NewMIPSTransactor(address common.Address, transactor bind.ContractTransactor) (*MIPSTransactor, error) {
	contract, err := bindMIPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MIPSTransactor{contract: contract}, nil
}

// NewMIPSFilterer creates a new log filterer instance of MIPS, bound to a specific deployed contract.
func NewMIPSFilterer(address common.Address, filterer bind.ContractFilterer) (*MIPSFilterer, error) {
	contract, err := bindMIPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MIPSFilterer{contract: contract}, nil
}

// bindMIPS binds a generic wrapper to an already deployed contract.
func bindMIPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MIPSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MIPS *MIPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MIPS.Contract.MIPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MIPS *MIPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MIPS.Contract.MIPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MIPS *MIPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MIPS.Contract.MIPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MIPS *MIPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MIPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MIPS *MIPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MIPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MIPS *MIPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MIPS.Contract.contract.Transact(opts, method, params...)
}

// BRKSTART is a free data retrieval call binding the contract method 0x155633fe.
//
// Solidity: function BRK_START() view returns(uint32)
func (_MIPS *MIPSCaller) BRKSTART(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MIPS.contract.Call(opts, &out, "BRK_START")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BRKSTART is a free data retrieval call binding the contract method 0x155633fe.
//
// Solidity: function BRK_START() view returns(uint32)
func (_MIPS *MIPSSession) BRKSTART() (uint32, error) {
	return _MIPS.Contract.BRKSTART(&_MIPS.CallOpts)
}

// BRKSTART is a free data retrieval call binding the contract method 0x155633fe.
//
// Solidity: function BRK_START() view returns(uint32)
func (_MIPS *MIPSCallerSession) BRKSTART() (uint32, error) {
	return _MIPS.Contract.BRKSTART(&_MIPS.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_MIPS *MIPSCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MIPS.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_MIPS *MIPSSession) Oracle() (common.Address, error) {
	return _MIPS.Contract.Oracle(&_MIPS.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_MIPS *MIPSCallerSession) Oracle() (common.Address, error) {
	return _MIPS.Contract.Oracle(&_MIPS.CallOpts)
}

// Step is a paid mutator transaction binding the contract method 0xf8e0cb96.
//
// Solidity: function step(bytes stateData, bytes proof) returns(bytes32)
func (_MIPS *MIPSTransactor) Step(opts *bind.TransactOpts, stateData []byte, proof []byte) (*types.Transaction, error) {
	return _MIPS.contract.Transact(opts, "step", stateData, proof)
}

// Step is a paid mutator transaction binding the contract method 0xf8e0cb96.
//
// Solidity: function step(bytes stateData, bytes proof) returns(bytes32)
func (_MIPS *MIPSSession) Step(stateData []byte, proof []byte) (*types.Transaction, error) {
	return _MIPS.Contract.Step(&_MIPS.TransactOpts, stateData, proof)
}

// Step is a paid mutator transaction binding the contract method 0xf8e0cb96.
//
// Solidity: function step(bytes stateData, bytes proof) returns(bytes32)
func (_MIPS *MIPSTransactorSession) Step(stateData []byte, proof []byte) (*types.Transaction, error) {
	return _MIPS.Contract.Step(&_MIPS.TransactOpts, stateData, proof)
}
