// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicall

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MulticallABI is the input ABI used to generate the binding from.
const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getDenormalizedWeights\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"inTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"outTokens\",\"type\":\"address[]\"}],\"name\":\"getSpotPrices\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTotalSupplies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getUsedBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MulticallBin is the compiled bytecode used for deploying new contracts.
var MulticallBin = "0x608060405234801561001057600080fd5b50611385806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632dd43d89146100675780636a385ae91461009857806395d7710e146100c9578063a826c114146100fa578063db16d8ac1461012b578063f7a6b6b61461015c575b600080fd5b610081600480360381019061007c9190610dad565b61018e565b60405161008f92919061106e565b60405180910390f35b6100b260048036038101906100ad9190610cda565b610334565b6040516100c092919061106e565b60405180910390f35b6100e360048036038101906100de9190610cda565b6104e7565b6040516100f192919061106e565b60405180910390f35b610114600480360381019061010f9190610dad565b61069a565b6040516101229291906110a5565b60405180910390f35b61014560048036038101906101409190610cda565b610848565b60405161015392919061106e565b60405180910390f35b61017660048036038101906101719190610d2e565b6109fb565b60405161018593929190611022565b60405180910390f35b6060806000835167ffffffffffffffff8111156101d4577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156102025781602001602082028036833780820191505090505b50905060005b84518110156103275784818151811061024a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561029757600080fd5b505afa1580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf9190610dee565b828281518110610308577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050808061031f90611252565b915050610208565b5083819250925050915091565b6060806000835167ffffffffffffffff81111561037a577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156103a85781602001602082028036833780820191505090505b50905060005b84518110156104d8578573ffffffffffffffffffffffffffffffffffffffff1663f8b2cb4f86838151811061040c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016104309190610fde565b60206040518083038186803b15801561044857600080fd5b505afa15801561045c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104809190610dee565b8282815181106104b9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101818152505080806104d090611252565b9150506103ae565b50838192509250509250929050565b6060806000835167ffffffffffffffff81111561052d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561055b5781602001602082028036833780820191505090505b50905060005b845181101561068b578573ffffffffffffffffffffffffffffffffffffffff16634aa4e0b58683815181106105bf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016105e39190610fde565b60206040518083038186803b1580156105fb57600080fd5b505afa15801561060f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106339190610dee565b82828151811061066c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050808061068390611252565b915050610561565b50838192509250509250929050565b6060806000835167ffffffffffffffff8111156106e0577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561070e5781602001602082028036833780820191505090505b50905060005b845181101561083b57848181518110610756577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156107a357600080fd5b505afa1580156107b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107db9190610e17565b828281518110610814577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019060ff16908160ff1681525050808061083390611252565b915050610714565b5083819250925050915091565b6060806000835167ffffffffffffffff81111561088e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156108bc5781602001602082028036833780820191505090505b50905060005b84518110156109ec578573ffffffffffffffffffffffffffffffffffffffff1663948d8ce6868381518110610920577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016109449190610fde565b60206040518083038186803b15801561095c57600080fd5b505afa158015610970573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109949190610dee565b8282815181106109cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101818152505080806109e490611252565b9150506108c2565b50838192509250509250929050565b60608060608351855114610a0e57600080fd5b6000855167ffffffffffffffff811115610a51577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610a7f5781602001602082028036833780820191505090505b50905060005b8651811015610bf1578773ffffffffffffffffffffffffffffffffffffffff166315e84af9888381518110610ae3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151888481518110610b24577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518363ffffffff1660e01b8152600401610b49929190610ff9565b60206040518083038186803b158015610b6157600080fd5b505afa158015610b75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b999190610dee565b828281518110610bd2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181815250508080610be990611252565b915050610a85565b508585829350935093505093509350939050565b6000610c18610c1384611101565b6110dc565b90508083825260208201905082856020860282011115610c3757600080fd5b60005b85811015610c675781610c4d8882610c71565b845260208401935060208301925050600181019050610c3a565b5050509392505050565b600081359050610c808161130a565b92915050565b600082601f830112610c9757600080fd5b8135610ca7848260208601610c05565b91505092915050565b600081519050610cbf81611321565b92915050565b600081519050610cd481611338565b92915050565b60008060408385031215610ced57600080fd5b6000610cfb85828601610c71565b925050602083013567ffffffffffffffff811115610d1857600080fd5b610d2485828601610c86565b9150509250929050565b600080600060608486031215610d4357600080fd5b6000610d5186828701610c71565b935050602084013567ffffffffffffffff811115610d6e57600080fd5b610d7a86828701610c86565b925050604084013567ffffffffffffffff811115610d9757600080fd5b610da386828701610c86565b9150509250925092565b600060208284031215610dbf57600080fd5b600082013567ffffffffffffffff811115610dd957600080fd5b610de584828501610c86565b91505092915050565b600060208284031215610e0057600080fd5b6000610e0e84828501610cb0565b91505092915050565b600060208284031215610e2957600080fd5b6000610e3784828501610cc5565b91505092915050565b6000610e4c8383610e88565b60208301905092915050565b6000610e648383610fc0565b60208301905092915050565b6000610e7c8383610fcf565b60208301905092915050565b610e91816111d8565b82525050565b610ea0816111d8565b82525050565b6000610eb18261115d565b610ebb81856111a5565b9350610ec68361112d565b8060005b83811015610ef7578151610ede8882610e40565b9750610ee98361117e565b925050600181019050610eca565b5085935050505092915050565b6000610f0f82611168565b610f1981856111b6565b9350610f248361113d565b8060005b83811015610f55578151610f3c8882610e58565b9750610f478361118b565b925050600181019050610f28565b5085935050505092915050565b6000610f6d82611173565b610f7781856111c7565b9350610f828361114d565b8060005b83811015610fb3578151610f9a8882610e70565b9750610fa583611198565b925050600181019050610f86565b5085935050505092915050565b610fc98161120a565b82525050565b610fd881611214565b82525050565b6000602082019050610ff36000830184610e97565b92915050565b600060408201905061100e6000830185610e97565b61101b6020830184610e97565b9392505050565b6000606082019050818103600083015261103c8186610ea6565b905081810360208301526110508185610ea6565b905081810360408301526110648184610f04565b9050949350505050565b600060408201905081810360008301526110888185610ea6565b9050818103602083015261109c8184610f04565b90509392505050565b600060408201905081810360008301526110bf8185610ea6565b905081810360208301526110d38184610f62565b90509392505050565b60006110e66110f7565b90506110f28282611221565b919050565b6000604051905090565b600067ffffffffffffffff82111561111c5761111b6112ca565b5b602082029050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006111e3826111ea565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b61122a826112f9565b810181811067ffffffffffffffff82111715611249576112486112ca565b5b80604052505050565b600061125d8261120a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156112905761128f61129b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b611313816111d8565b811461131e57600080fd5b50565b61132a8161120a565b811461133557600080fd5b50565b61134181611214565b811461134c57600080fd5b5056fea26469706673582212204692fb9f1ab26246b62838bd95d54efbea92be5818b98cfd2855b1d692bc1ed464736f6c63430008010033"

// DeployMulticall deploys a new Ethereum contract, binding an instance of Multicall to it.
func DeployMulticall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multicall, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MulticallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// Multicall is an auto generated Go binding around an Ethereum contract.
type Multicall struct {
	MulticallCaller     // Read-only binding to the contract
	MulticallTransactor // Write-only binding to the contract
	MulticallFilterer   // Log filterer for contract events
}

// MulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallSession struct {
	Contract     *Multicall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallCallerSession struct {
	Contract *MulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallTransactorSession struct {
	Contract     *MulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallRaw struct {
	Contract *Multicall // Generic contract binding to access the raw methods on
}

// MulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallCallerRaw struct {
	Contract *MulticallCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallTransactorRaw struct {
	Contract *MulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall creates a new instance of Multicall, bound to a specific deployed contract.
func NewMulticall(address common.Address, backend bind.ContractBackend) (*Multicall, error) {
	contract, err := bindMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// NewMulticallCaller creates a new read-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallCaller(address common.Address, caller bind.ContractCaller) (*MulticallCaller, error) {
	contract, err := bindMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallCaller{contract: contract}, nil
}

// NewMulticallTransactor creates a new write-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallTransactor, error) {
	contract, err := bindMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTransactor{contract: contract}, nil
}

// NewMulticallFilterer creates a new log filterer instance of Multicall, bound to a specific deployed contract.
func NewMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallFilterer, error) {
	contract, err := bindMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallFilterer{contract: contract}, nil
}

// bindMulticall binds a generic wrapper to an already deployed contract.
func bindMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.MulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transact(opts, method, params...)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetBalances(opts *bind.CallOpts, poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getBalances", poolAddress, tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetBalances(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetBalances(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetDecimals is a free data retrieval call binding the contract method 0xa826c114.
//
// Solidity: function getDecimals(address[] tokens) view returns(address[], uint8[])
func (_Multicall *MulticallCaller) GetDecimals(opts *bind.CallOpts, tokens []common.Address) ([]common.Address, []uint8, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getDecimals", tokens)

	if err != nil {
		return *new([]common.Address), *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return out0, out1, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xa826c114.
//
// Solidity: function getDecimals(address[] tokens) view returns(address[], uint8[])
func (_Multicall *MulticallSession) GetDecimals(tokens []common.Address) ([]common.Address, []uint8, error) {
	return _Multicall.Contract.GetDecimals(&_Multicall.CallOpts, tokens)
}

// GetDecimals is a free data retrieval call binding the contract method 0xa826c114.
//
// Solidity: function getDecimals(address[] tokens) view returns(address[], uint8[])
func (_Multicall *MulticallCallerSession) GetDecimals(tokens []common.Address) ([]common.Address, []uint8, error) {
	return _Multicall.Contract.GetDecimals(&_Multicall.CallOpts, tokens)
}

// GetDenormalizedWeights is a free data retrieval call binding the contract method 0xdb16d8ac.
//
// Solidity: function getDenormalizedWeights(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetDenormalizedWeights(opts *bind.CallOpts, poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getDenormalizedWeights", poolAddress, tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDenormalizedWeights is a free data retrieval call binding the contract method 0xdb16d8ac.
//
// Solidity: function getDenormalizedWeights(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetDenormalizedWeights(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetDenormalizedWeights(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetDenormalizedWeights is a free data retrieval call binding the contract method 0xdb16d8ac.
//
// Solidity: function getDenormalizedWeights(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetDenormalizedWeights(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetDenormalizedWeights(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetSpotPrices is a free data retrieval call binding the contract method 0xf7a6b6b6.
//
// Solidity: function getSpotPrices(address poolAddress, address[] inTokens, address[] outTokens) view returns(address[], address[], uint256[])
func (_Multicall *MulticallCaller) GetSpotPrices(opts *bind.CallOpts, poolAddress common.Address, inTokens []common.Address, outTokens []common.Address) ([]common.Address, []common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getSpotPrices", poolAddress, inTokens, outTokens)

	if err != nil {
		return *new([]common.Address), *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, err

}

// GetSpotPrices is a free data retrieval call binding the contract method 0xf7a6b6b6.
//
// Solidity: function getSpotPrices(address poolAddress, address[] inTokens, address[] outTokens) view returns(address[], address[], uint256[])
func (_Multicall *MulticallSession) GetSpotPrices(poolAddress common.Address, inTokens []common.Address, outTokens []common.Address) ([]common.Address, []common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetSpotPrices(&_Multicall.CallOpts, poolAddress, inTokens, outTokens)
}

// GetSpotPrices is a free data retrieval call binding the contract method 0xf7a6b6b6.
//
// Solidity: function getSpotPrices(address poolAddress, address[] inTokens, address[] outTokens) view returns(address[], address[], uint256[])
func (_Multicall *MulticallCallerSession) GetSpotPrices(poolAddress common.Address, inTokens []common.Address, outTokens []common.Address) ([]common.Address, []common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetSpotPrices(&_Multicall.CallOpts, poolAddress, inTokens, outTokens)
}

// GetTotalSupplies is a free data retrieval call binding the contract method 0x2dd43d89.
//
// Solidity: function getTotalSupplies(address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetTotalSupplies(opts *bind.CallOpts, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getTotalSupplies", tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetTotalSupplies is a free data retrieval call binding the contract method 0x2dd43d89.
//
// Solidity: function getTotalSupplies(address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetTotalSupplies(tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetTotalSupplies(&_Multicall.CallOpts, tokens)
}

// GetTotalSupplies is a free data retrieval call binding the contract method 0x2dd43d89.
//
// Solidity: function getTotalSupplies(address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetTotalSupplies(tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetTotalSupplies(&_Multicall.CallOpts, tokens)
}

// GetUsedBalances is a free data retrieval call binding the contract method 0x95d7710e.
//
// Solidity: function getUsedBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCaller) GetUsedBalances(opts *bind.CallOpts, poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getUsedBalances", poolAddress, tokens)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetUsedBalances is a free data retrieval call binding the contract method 0x95d7710e.
//
// Solidity: function getUsedBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallSession) GetUsedBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetUsedBalances(&_Multicall.CallOpts, poolAddress, tokens)
}

// GetUsedBalances is a free data retrieval call binding the contract method 0x95d7710e.
//
// Solidity: function getUsedBalances(address poolAddress, address[] tokens) view returns(address[], uint256[])
func (_Multicall *MulticallCallerSession) GetUsedBalances(poolAddress common.Address, tokens []common.Address) ([]common.Address, []*big.Int, error) {
	return _Multicall.Contract.GetUsedBalances(&_Multicall.CallOpts, poolAddress, tokens)
}
