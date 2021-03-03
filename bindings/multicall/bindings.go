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
const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getDenormalizedWeights\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"inTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"outTokens\",\"type\":\"address[]\"}],\"name\":\"getSpotPrices\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTotalSupplies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getUsedBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"poolTokensFor\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MulticallBin is the compiled bytecode used for deploying new contracts.
var MulticallBin = "0x608060405234801561001057600080fd5b50611a39806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063a826c1141161005b578063a826c11414610115578063db16d8ac14610146578063f7a6b6b614610177578063f8504f57146101a95761007d565b80632dd43d89146100825780636a385ae9146100b357806395d7710e146100e4575b600080fd5b61009c600480360381019061009791906111b0565b6101da565b6040516100aa92919061160f565b60405180910390f35b6100cd60048036038101906100c891906110dd565b610380565b6040516100db92919061160f565b60405180910390f35b6100fe60048036038101906100f991906110dd565b610533565b60405161010c92919061160f565b60405180910390f35b61012f600480360381019061012a91906111b0565b6106e6565b60405161013d929190611646565b60405180910390f35b610160600480360381019061015b91906110dd565b610894565b60405161016e92919061160f565b60405180910390f35b610191600480360381019061018c9190611131565b610a47565b6040516101a09392919061158c565b60405180910390f35b6101c360048036038101906101be91906110b4565b610c51565b6040516101d19291906115d8565b60405180910390f35b6060806000835167ffffffffffffffff811115610220577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561024e5781602001602082028036833780820191505090505b50905060005b845181101561037357848181518110610296577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156102e357600080fd5b505afa1580156102f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031b9190611273565b828281518110610354577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050808061036b906118dd565b915050610254565b5083819250925050915091565b6060806000835167ffffffffffffffff8111156103c6577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156103f45781602001602082028036833780820191505090505b50905060005b8451811015610524578573ffffffffffffffffffffffffffffffffffffffff1663f8b2cb4f868381518110610458577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040161047c9190611548565b60206040518083038186803b15801561049457600080fd5b505afa1580156104a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cc9190611273565b828281518110610505577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018181525050808061051c906118dd565b9150506103fa565b50838192509250509250929050565b6060806000835167ffffffffffffffff811115610579577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156105a75781602001602082028036833780820191505090505b50905060005b84518110156106d7578573ffffffffffffffffffffffffffffffffffffffff16634aa4e0b586838151811061060b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040161062f9190611548565b60206040518083038186803b15801561064757600080fd5b505afa15801561065b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067f9190611273565b8282815181106106b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101818152505080806106cf906118dd565b9150506105ad565b50838192509250509250929050565b6060806000835167ffffffffffffffff81111561072c577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561075a5781602001602082028036833780820191505090505b50905060005b8451811015610887578481815181106107a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156107ef57600080fd5b505afa158015610803573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610827919061129c565b828281518110610860577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019060ff16908160ff1681525050808061087f906118dd565b915050610760565b5083819250925050915091565b6060806000835167ffffffffffffffff8111156108da577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156109085781602001602082028036833780820191505090505b50905060005b8451811015610a38578573ffffffffffffffffffffffffffffffffffffffff1663948d8ce686838151811061096c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016109909190611548565b60206040518083038186803b1580156109a857600080fd5b505afa1580156109bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e09190611273565b828281518110610a19577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181815250508080610a30906118dd565b91505061090e565b50838192509250509250929050565b60608060608351855114610a5a57600080fd5b6000855167ffffffffffffffff811115610a9d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610acb5781602001602082028036833780820191505090505b50905060005b8651811015610c3d578773ffffffffffffffffffffffffffffffffffffffff166315e84af9888381518110610b2f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151888481518110610b70577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518363ffffffff1660e01b8152600401610b95929190611563565b60206040518083038186803b158015610bad57600080fd5b505afa158015610bc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be59190611273565b828281518110610c1e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181815250508080610c35906118dd565b915050610ad1565b508585829350935093505093509350939050565b60608060008373ffffffffffffffffffffffffffffffffffffffff1663cc77828d6040518163ffffffff1660e01b815260040160006040518083038186803b158015610c9c57600080fd5b505afa158015610cb0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610cd991906111f1565b90506000815111610d1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d169061167d565b60405180910390fd5b6000815167ffffffffffffffff811115610d62577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610d9557816020015b6060815260200190600190039081610d805790505b50905060005b8251811015610ebe57828181518110610ddd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015610e2a57600080fd5b505afa158015610e3e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610e679190611232565b828281518110610ea0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052508080610eb6906118dd565b915050610d9b565b508181935093505050915091565b6000610edf610eda846116c2565b61169d565b90508083825260208201905082856020860282011115610efe57600080fd5b60005b85811015610f2e5781610f148882610fe2565b845260208401935060208301925050600181019050610f01565b5050509392505050565b6000610f4b610f46846116c2565b61169d565b90508083825260208201905082856020860282011115610f6a57600080fd5b60005b85811015610f9a5781610f808882610ff7565b845260208401935060208301925050600181019050610f6d565b5050509392505050565b6000610fb7610fb2846116ee565b61169d565b905082815260208101848484011115610fcf57600080fd5b610fda848285611879565b509392505050565b600081359050610ff1816119be565b92915050565b600081519050611006816119be565b92915050565b600082601f83011261101d57600080fd5b813561102d848260208601610ecc565b91505092915050565b600082601f83011261104757600080fd5b8151611057848260208601610f38565b91505092915050565b600082601f83011261107157600080fd5b8151611081848260208601610fa4565b91505092915050565b600081519050611099816119d5565b92915050565b6000815190506110ae816119ec565b92915050565b6000602082840312156110c657600080fd5b60006110d484828501610fe2565b91505092915050565b600080604083850312156110f057600080fd5b60006110fe85828601610fe2565b925050602083013567ffffffffffffffff81111561111b57600080fd5b6111278582860161100c565b9150509250929050565b60008060006060848603121561114657600080fd5b600061115486828701610fe2565b935050602084013567ffffffffffffffff81111561117157600080fd5b61117d8682870161100c565b925050604084013567ffffffffffffffff81111561119a57600080fd5b6111a68682870161100c565b9150509250925092565b6000602082840312156111c257600080fd5b600082013567ffffffffffffffff8111156111dc57600080fd5b6111e88482850161100c565b91505092915050565b60006020828403121561120357600080fd5b600082015167ffffffffffffffff81111561121d57600080fd5b61122984828501611036565b91505092915050565b60006020828403121561124457600080fd5b600082015167ffffffffffffffff81111561125e57600080fd5b61126a84828501611060565b91505092915050565b60006020828403121561128557600080fd5b60006112938482850161108a565b91505092915050565b6000602082840312156112ae57600080fd5b60006112bc8482850161109f565b91505092915050565b60006112d18383611321565b60208301905092915050565b60006112e983836114ce565b905092915050565b60006112fd838361152a565b60208301905092915050565b60006113158383611539565b60208301905092915050565b61132a81611830565b82525050565b61133981611830565b82525050565b600061134a8261175f565b61135481856117ca565b935061135f8361171f565b8060005b8381101561139057815161137788826112c5565b975061138283611796565b925050600181019050611363565b5085935050505092915050565b60006113a88261176a565b6113b281856117db565b9350836020820285016113c48561172f565b8060005b8581101561140057848403895281516113e185826112dd565b94506113ec836117a3565b925060208a019950506001810190506113c8565b50829750879550505050505092915050565b600061141d82611775565b61142781856117ec565b93506114328361173f565b8060005b8381101561146357815161144a88826112f1565b9750611455836117b0565b925050600181019050611436565b5085935050505092915050565b600061147b82611780565b61148581856117fd565b93506114908361174f565b8060005b838110156114c15781516114a88882611309565b97506114b3836117bd565b925050600181019050611494565b5085935050505092915050565b60006114d98261178b565b6114e3818561180e565b93506114f3818560208601611879565b6114fc81611984565b840191505092915050565b6000611514600e8361181f565b915061151f82611995565b602082019050919050565b61153381611862565b82525050565b6115428161186c565b82525050565b600060208201905061155d6000830184611330565b92915050565b60006040820190506115786000830185611330565b6115856020830184611330565b9392505050565b600060608201905081810360008301526115a6818661133f565b905081810360208301526115ba818561133f565b905081810360408301526115ce8184611412565b9050949350505050565b600060408201905081810360008301526115f2818561133f565b90508181036020830152611606818461139d565b90509392505050565b60006040820190508181036000830152611629818561133f565b9050818103602083015261163d8184611412565b90509392505050565b60006040820190508181036000830152611660818561133f565b905081810360208301526116748184611470565b90509392505050565b6000602082019050818103600083015261169681611507565b9050919050565b60006116a76116b8565b90506116b382826118ac565b919050565b6000604051905090565b600067ffffffffffffffff8211156116dd576116dc611955565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561170957611708611955565b5b61171282611984565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061183b82611842565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561189757808201518184015260208101905061187c565b838111156118a6576000848401525b50505050565b6118b582611984565b810181811067ffffffffffffffff821117156118d4576118d3611955565b5b80604052505050565b60006118e882611862565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561191b5761191a611926565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f6e6f20706f6f6c20746f6b656e73000000000000000000000000000000000000600082015250565b6119c781611830565b81146119d257600080fd5b50565b6119de81611862565b81146119e957600080fd5b50565b6119f58161186c565b8114611a0057600080fd5b5056fea2646970667358221220ab8c0203a7944720b911a79153e9cde0ed213814da778c88075faaf71aadb90c64736f6c63430008010033"

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

// PoolTokensFor is a free data retrieval call binding the contract method 0xf8504f57.
//
// Solidity: function poolTokensFor(address poolAddress) view returns(address[], string[])
func (_Multicall *MulticallCaller) PoolTokensFor(opts *bind.CallOpts, poolAddress common.Address) ([]common.Address, []string, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "poolTokensFor", poolAddress)

	if err != nil {
		return *new([]common.Address), *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]string)).(*[]string)

	return out0, out1, err

}

// PoolTokensFor is a free data retrieval call binding the contract method 0xf8504f57.
//
// Solidity: function poolTokensFor(address poolAddress) view returns(address[], string[])
func (_Multicall *MulticallSession) PoolTokensFor(poolAddress common.Address) ([]common.Address, []string, error) {
	return _Multicall.Contract.PoolTokensFor(&_Multicall.CallOpts, poolAddress)
}

// PoolTokensFor is a free data retrieval call binding the contract method 0xf8504f57.
//
// Solidity: function poolTokensFor(address poolAddress) view returns(address[], string[])
func (_Multicall *MulticallCallerSession) PoolTokensFor(poolAddress common.Address) ([]common.Address, []string, error) {
	return _Multicall.Contract.PoolTokensFor(&_Multicall.CallOpts, poolAddress)
}
