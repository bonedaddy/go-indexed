// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governoralpha

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

// GovernorAlphaReceipt is an auto generated low-level Go binding around an user-defined struct.
type GovernorAlphaReceipt struct {
	HasVoted bool
	Support  bool
	Votes    *big.Int
}

// GovernoralphaABI is the input ABI used to generate the binding from.
const GovernoralphaABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"timelock_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ndx_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"setVotingPeriodAfter_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getActions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structGovernorAlpha.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestProposalIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ndx\",\"outputs\":[{\"internalType\":\"contractNdxInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permanentVotingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalMaxOperations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setPermanentVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setVotingPeriodAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumGovernorAlpha.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"contractITimelock\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Governoralpha is an auto generated Go binding around an Ethereum contract.
type Governoralpha struct {
	GovernoralphaCaller     // Read-only binding to the contract
	GovernoralphaTransactor // Write-only binding to the contract
	GovernoralphaFilterer   // Log filterer for contract events
}

// GovernoralphaCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernoralphaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernoralphaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernoralphaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernoralphaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernoralphaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernoralphaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernoralphaSession struct {
	Contract     *Governoralpha    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernoralphaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernoralphaCallerSession struct {
	Contract *GovernoralphaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GovernoralphaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernoralphaTransactorSession struct {
	Contract     *GovernoralphaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GovernoralphaRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernoralphaRaw struct {
	Contract *Governoralpha // Generic contract binding to access the raw methods on
}

// GovernoralphaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernoralphaCallerRaw struct {
	Contract *GovernoralphaCaller // Generic read-only contract binding to access the raw methods on
}

// GovernoralphaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernoralphaTransactorRaw struct {
	Contract *GovernoralphaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernoralpha creates a new instance of Governoralpha, bound to a specific deployed contract.
func NewGovernoralpha(address common.Address, backend bind.ContractBackend) (*Governoralpha, error) {
	contract, err := bindGovernoralpha(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governoralpha{GovernoralphaCaller: GovernoralphaCaller{contract: contract}, GovernoralphaTransactor: GovernoralphaTransactor{contract: contract}, GovernoralphaFilterer: GovernoralphaFilterer{contract: contract}}, nil
}

// NewGovernoralphaCaller creates a new read-only instance of Governoralpha, bound to a specific deployed contract.
func NewGovernoralphaCaller(address common.Address, caller bind.ContractCaller) (*GovernoralphaCaller, error) {
	contract, err := bindGovernoralpha(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernoralphaCaller{contract: contract}, nil
}

// NewGovernoralphaTransactor creates a new write-only instance of Governoralpha, bound to a specific deployed contract.
func NewGovernoralphaTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernoralphaTransactor, error) {
	contract, err := bindGovernoralpha(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernoralphaTransactor{contract: contract}, nil
}

// NewGovernoralphaFilterer creates a new log filterer instance of Governoralpha, bound to a specific deployed contract.
func NewGovernoralphaFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernoralphaFilterer, error) {
	contract, err := bindGovernoralpha(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernoralphaFilterer{contract: contract}, nil
}

// bindGovernoralpha binds a generic wrapper to an already deployed contract.
func bindGovernoralpha(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernoralphaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governoralpha *GovernoralphaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governoralpha.Contract.GovernoralphaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governoralpha *GovernoralphaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governoralpha.Contract.GovernoralphaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governoralpha *GovernoralphaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governoralpha.Contract.GovernoralphaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governoralpha *GovernoralphaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governoralpha.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governoralpha *GovernoralphaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governoralpha.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governoralpha *GovernoralphaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governoralpha.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Governoralpha *GovernoralphaCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Governoralpha *GovernoralphaSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _Governoralpha.Contract.BALLOTTYPEHASH(&_Governoralpha.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_Governoralpha *GovernoralphaCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _Governoralpha.Contract.BALLOTTYPEHASH(&_Governoralpha.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Governoralpha *GovernoralphaCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Governoralpha *GovernoralphaSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Governoralpha.Contract.DOMAINTYPEHASH(&_Governoralpha.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Governoralpha *GovernoralphaCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Governoralpha.Contract.DOMAINTYPEHASH(&_Governoralpha.CallOpts)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governoralpha *GovernoralphaCaller) GetActions(opts *bind.CallOpts, proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "getActions", proposalId)

	outstruct := new(struct {
		Targets    []common.Address
		Values     []*big.Int
		Signatures []string
		Calldatas  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Targets = out[0].([]common.Address)
	outstruct.Values = out[1].([]*big.Int)
	outstruct.Signatures = out[2].([]string)
	outstruct.Calldatas = out[3].([][]byte)

	return *outstruct, err

}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governoralpha *GovernoralphaSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _Governoralpha.Contract.GetActions(&_Governoralpha.CallOpts, proposalId)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governoralpha *GovernoralphaCallerSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _Governoralpha.Contract.GetActions(&_Governoralpha.CallOpts, proposalId)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,bool,uint96))
func (_Governoralpha *GovernoralphaCaller) GetReceipt(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (GovernorAlphaReceipt, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "getReceipt", proposalId, voter)

	if err != nil {
		return *new(GovernorAlphaReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(GovernorAlphaReceipt)).(*GovernorAlphaReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,bool,uint96))
func (_Governoralpha *GovernoralphaSession) GetReceipt(proposalId *big.Int, voter common.Address) (GovernorAlphaReceipt, error) {
	return _Governoralpha.Contract.GetReceipt(&_Governoralpha.CallOpts, proposalId, voter)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,bool,uint96))
func (_Governoralpha *GovernoralphaCallerSession) GetReceipt(proposalId *big.Int, voter common.Address) (GovernorAlphaReceipt, error) {
	return _Governoralpha.Contract.GetReceipt(&_Governoralpha.CallOpts, proposalId, voter)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_Governoralpha *GovernoralphaCaller) LatestProposalIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "latestProposalIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_Governoralpha *GovernoralphaSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _Governoralpha.Contract.LatestProposalIds(&_Governoralpha.CallOpts, arg0)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _Governoralpha.Contract.LatestProposalIds(&_Governoralpha.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Governoralpha *GovernoralphaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Governoralpha *GovernoralphaSession) Name() (string, error) {
	return _Governoralpha.Contract.Name(&_Governoralpha.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Governoralpha *GovernoralphaCallerSession) Name() (string, error) {
	return _Governoralpha.Contract.Name(&_Governoralpha.CallOpts)
}

// Ndx is a free data retrieval call binding the contract method 0x79521b2c.
//
// Solidity: function ndx() view returns(address)
func (_Governoralpha *GovernoralphaCaller) Ndx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "ndx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ndx is a free data retrieval call binding the contract method 0x79521b2c.
//
// Solidity: function ndx() view returns(address)
func (_Governoralpha *GovernoralphaSession) Ndx() (common.Address, error) {
	return _Governoralpha.Contract.Ndx(&_Governoralpha.CallOpts)
}

// Ndx is a free data retrieval call binding the contract method 0x79521b2c.
//
// Solidity: function ndx() view returns(address)
func (_Governoralpha *GovernoralphaCallerSession) Ndx() (common.Address, error) {
	return _Governoralpha.Contract.Ndx(&_Governoralpha.CallOpts)
}

// PermanentVotingPeriod is a free data retrieval call binding the contract method 0xec1e0e03.
//
// Solidity: function permanentVotingPeriod() view returns(uint256)
func (_Governoralpha *GovernoralphaCaller) PermanentVotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "permanentVotingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermanentVotingPeriod is a free data retrieval call binding the contract method 0xec1e0e03.
//
// Solidity: function permanentVotingPeriod() view returns(uint256)
func (_Governoralpha *GovernoralphaSession) PermanentVotingPeriod() (*big.Int, error) {
	return _Governoralpha.Contract.PermanentVotingPeriod(&_Governoralpha.CallOpts)
}

// PermanentVotingPeriod is a free data retrieval call binding the contract method 0xec1e0e03.
//
// Solidity: function permanentVotingPeriod() view returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) PermanentVotingPeriod() (*big.Int, error) {
	return _Governoralpha.Contract.PermanentVotingPeriod(&_Governoralpha.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Governoralpha *GovernoralphaCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Governoralpha *GovernoralphaSession) ProposalCount() (*big.Int, error) {
	return _Governoralpha.Contract.ProposalCount(&_Governoralpha.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) ProposalCount() (*big.Int, error) {
	return _Governoralpha.Contract.ProposalCount(&_Governoralpha.CallOpts)
}

// ProposalMaxOperations is a free data retrieval call binding the contract method 0x7bdbe4d0.
//
// Solidity: function proposalMaxOperations() pure returns(uint256)
func (_Governoralpha *GovernoralphaCaller) ProposalMaxOperations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "proposalMaxOperations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalMaxOperations is a free data retrieval call binding the contract method 0x7bdbe4d0.
//
// Solidity: function proposalMaxOperations() pure returns(uint256)
func (_Governoralpha *GovernoralphaSession) ProposalMaxOperations() (*big.Int, error) {
	return _Governoralpha.Contract.ProposalMaxOperations(&_Governoralpha.CallOpts)
}

// ProposalMaxOperations is a free data retrieval call binding the contract method 0x7bdbe4d0.
//
// Solidity: function proposalMaxOperations() pure returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) ProposalMaxOperations() (*big.Int, error) {
	return _Governoralpha.Contract.ProposalMaxOperations(&_Governoralpha.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Governoralpha *GovernoralphaCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Governoralpha *GovernoralphaSession) ProposalThreshold() (*big.Int, error) {
	return _Governoralpha.Contract.ProposalThreshold(&_Governoralpha.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) ProposalThreshold() (*big.Int, error) {
	return _Governoralpha.Contract.ProposalThreshold(&_Governoralpha.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, bool canceled, bool executed)
func (_Governoralpha *GovernoralphaCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Eta          *big.Int
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
}, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Id           *big.Int
		Proposer     common.Address
		Eta          *big.Int
		StartBlock   *big.Int
		EndBlock     *big.Int
		ForVotes     *big.Int
		AgainstVotes *big.Int
		Canceled     bool
		Executed     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = out[0].(*big.Int)
	outstruct.Proposer = out[1].(common.Address)
	outstruct.Eta = out[2].(*big.Int)
	outstruct.StartBlock = out[3].(*big.Int)
	outstruct.EndBlock = out[4].(*big.Int)
	outstruct.ForVotes = out[5].(*big.Int)
	outstruct.AgainstVotes = out[6].(*big.Int)
	outstruct.Canceled = out[7].(bool)
	outstruct.Executed = out[8].(bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, bool canceled, bool executed)
func (_Governoralpha *GovernoralphaSession) Proposals(arg0 *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Eta          *big.Int
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
}, error) {
	return _Governoralpha.Contract.Proposals(&_Governoralpha.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, bool canceled, bool executed)
func (_Governoralpha *GovernoralphaCallerSession) Proposals(arg0 *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Eta          *big.Int
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
}, error) {
	return _Governoralpha.Contract.Proposals(&_Governoralpha.CallOpts, arg0)
}

// QuorumVotes is a free data retrieval call binding the contract method 0x24bc1a64.
//
// Solidity: function quorumVotes() pure returns(uint256)
func (_Governoralpha *GovernoralphaCaller) QuorumVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "quorumVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumVotes is a free data retrieval call binding the contract method 0x24bc1a64.
//
// Solidity: function quorumVotes() pure returns(uint256)
func (_Governoralpha *GovernoralphaSession) QuorumVotes() (*big.Int, error) {
	return _Governoralpha.Contract.QuorumVotes(&_Governoralpha.CallOpts)
}

// QuorumVotes is a free data retrieval call binding the contract method 0x24bc1a64.
//
// Solidity: function quorumVotes() pure returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) QuorumVotes() (*big.Int, error) {
	return _Governoralpha.Contract.QuorumVotes(&_Governoralpha.CallOpts)
}

// SetVotingPeriodAfter is a free data retrieval call binding the contract method 0x1a5b534e.
//
// Solidity: function setVotingPeriodAfter() view returns(uint256)
func (_Governoralpha *GovernoralphaCaller) SetVotingPeriodAfter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "setVotingPeriodAfter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SetVotingPeriodAfter is a free data retrieval call binding the contract method 0x1a5b534e.
//
// Solidity: function setVotingPeriodAfter() view returns(uint256)
func (_Governoralpha *GovernoralphaSession) SetVotingPeriodAfter() (*big.Int, error) {
	return _Governoralpha.Contract.SetVotingPeriodAfter(&_Governoralpha.CallOpts)
}

// SetVotingPeriodAfter is a free data retrieval call binding the contract method 0x1a5b534e.
//
// Solidity: function setVotingPeriodAfter() view returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) SetVotingPeriodAfter() (*big.Int, error) {
	return _Governoralpha.Contract.SetVotingPeriodAfter(&_Governoralpha.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governoralpha *GovernoralphaCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governoralpha *GovernoralphaSession) State(proposalId *big.Int) (uint8, error) {
	return _Governoralpha.Contract.State(&_Governoralpha.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governoralpha *GovernoralphaCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _Governoralpha.Contract.State(&_Governoralpha.CallOpts, proposalId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_Governoralpha *GovernoralphaCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_Governoralpha *GovernoralphaSession) Timelock() (common.Address, error) {
	return _Governoralpha.Contract.Timelock(&_Governoralpha.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_Governoralpha *GovernoralphaCallerSession) Timelock() (common.Address, error) {
	return _Governoralpha.Contract.Timelock(&_Governoralpha.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Governoralpha *GovernoralphaCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Governoralpha *GovernoralphaSession) VotingDelay() (*big.Int, error) {
	return _Governoralpha.Contract.VotingDelay(&_Governoralpha.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) VotingDelay() (*big.Int, error) {
	return _Governoralpha.Contract.VotingDelay(&_Governoralpha.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_Governoralpha *GovernoralphaCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governoralpha.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_Governoralpha *GovernoralphaSession) VotingPeriod() (*big.Int, error) {
	return _Governoralpha.Contract.VotingPeriod(&_Governoralpha.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_Governoralpha *GovernoralphaCallerSession) VotingPeriod() (*big.Int, error) {
	return _Governoralpha.Contract.VotingPeriod(&_Governoralpha.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_Governoralpha *GovernoralphaTransactor) Cancel(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.contract.Transact(opts, "cancel", proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_Governoralpha *GovernoralphaSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.Contract.Cancel(&_Governoralpha.TransactOpts, proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_Governoralpha *GovernoralphaTransactorSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.Contract.Cancel(&_Governoralpha.TransactOpts, proposalId)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_Governoralpha *GovernoralphaTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governoralpha.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_Governoralpha *GovernoralphaSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governoralpha.Contract.CastVote(&_Governoralpha.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_Governoralpha *GovernoralphaTransactorSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governoralpha.Contract.CastVote(&_Governoralpha.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x4634c61f.
//
// Solidity: function castVoteBySig(uint256 proposalId, bool support, uint8 v, bytes32 r, bytes32 s) returns()
func (_Governoralpha *GovernoralphaTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governoralpha.contract.Transact(opts, "castVoteBySig", proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x4634c61f.
//
// Solidity: function castVoteBySig(uint256 proposalId, bool support, uint8 v, bytes32 r, bytes32 s) returns()
func (_Governoralpha *GovernoralphaSession) CastVoteBySig(proposalId *big.Int, support bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governoralpha.Contract.CastVoteBySig(&_Governoralpha.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x4634c61f.
//
// Solidity: function castVoteBySig(uint256 proposalId, bool support, uint8 v, bytes32 r, bytes32 s) returns()
func (_Governoralpha *GovernoralphaTransactorSession) CastVoteBySig(proposalId *big.Int, support bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Governoralpha.Contract.CastVoteBySig(&_Governoralpha.TransactOpts, proposalId, support, v, r, s)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governoralpha *GovernoralphaTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.contract.Transact(opts, "execute", proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governoralpha *GovernoralphaSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.Contract.Execute(&_Governoralpha.TransactOpts, proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governoralpha *GovernoralphaTransactorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.Contract.Execute(&_Governoralpha.TransactOpts, proposalId)
}

// Propose is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_Governoralpha *GovernoralphaTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governoralpha.contract.Transact(opts, "propose", targets, values, signatures, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_Governoralpha *GovernoralphaSession) Propose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governoralpha.Contract.Propose(&_Governoralpha.TransactOpts, targets, values, signatures, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_Governoralpha *GovernoralphaTransactorSession) Propose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _Governoralpha.Contract.Propose(&_Governoralpha.TransactOpts, targets, values, signatures, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governoralpha *GovernoralphaTransactor) Queue(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.contract.Transact(opts, "queue", proposalId)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governoralpha *GovernoralphaSession) Queue(proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.Contract.Queue(&_Governoralpha.TransactOpts, proposalId)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governoralpha *GovernoralphaTransactorSession) Queue(proposalId *big.Int) (*types.Transaction, error) {
	return _Governoralpha.Contract.Queue(&_Governoralpha.TransactOpts, proposalId)
}

// SetPermanentVotingPeriod is a paid mutator transaction binding the contract method 0x7fd9a857.
//
// Solidity: function setPermanentVotingPeriod() returns()
func (_Governoralpha *GovernoralphaTransactor) SetPermanentVotingPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governoralpha.contract.Transact(opts, "setPermanentVotingPeriod")
}

// SetPermanentVotingPeriod is a paid mutator transaction binding the contract method 0x7fd9a857.
//
// Solidity: function setPermanentVotingPeriod() returns()
func (_Governoralpha *GovernoralphaSession) SetPermanentVotingPeriod() (*types.Transaction, error) {
	return _Governoralpha.Contract.SetPermanentVotingPeriod(&_Governoralpha.TransactOpts)
}

// SetPermanentVotingPeriod is a paid mutator transaction binding the contract method 0x7fd9a857.
//
// Solidity: function setPermanentVotingPeriod() returns()
func (_Governoralpha *GovernoralphaTransactorSession) SetPermanentVotingPeriod() (*types.Transaction, error) {
	return _Governoralpha.Contract.SetPermanentVotingPeriod(&_Governoralpha.TransactOpts)
}

// GovernoralphaProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the Governoralpha contract.
type GovernoralphaProposalCanceledIterator struct {
	Event *GovernoralphaProposalCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernoralphaProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernoralphaProposalCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernoralphaProposalCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernoralphaProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernoralphaProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernoralphaProposalCanceled represents a ProposalCanceled event raised by the Governoralpha contract.
type GovernoralphaProposalCanceled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 id)
func (_Governoralpha *GovernoralphaFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*GovernoralphaProposalCanceledIterator, error) {

	logs, sub, err := _Governoralpha.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &GovernoralphaProposalCanceledIterator{contract: _Governoralpha.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 id)
func (_Governoralpha *GovernoralphaFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *GovernoralphaProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _Governoralpha.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernoralphaProposalCanceled)
				if err := _Governoralpha.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 id)
func (_Governoralpha *GovernoralphaFilterer) ParseProposalCanceled(log types.Log) (*GovernoralphaProposalCanceled, error) {
	event := new(GovernoralphaProposalCanceled)
	if err := _Governoralpha.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernoralphaProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Governoralpha contract.
type GovernoralphaProposalCreatedIterator struct {
	Event *GovernoralphaProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernoralphaProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernoralphaProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernoralphaProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernoralphaProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernoralphaProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernoralphaProposalCreated represents a ProposalCreated event raised by the Governoralpha contract.
type GovernoralphaProposalCreated struct {
	Id          *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	StartBlock  *big.Int
	EndBlock    *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_Governoralpha *GovernoralphaFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GovernoralphaProposalCreatedIterator, error) {

	logs, sub, err := _Governoralpha.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GovernoralphaProposalCreatedIterator{contract: _Governoralpha.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_Governoralpha *GovernoralphaFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GovernoralphaProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Governoralpha.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernoralphaProposalCreated)
				if err := _Governoralpha.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_Governoralpha *GovernoralphaFilterer) ParseProposalCreated(log types.Log) (*GovernoralphaProposalCreated, error) {
	event := new(GovernoralphaProposalCreated)
	if err := _Governoralpha.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernoralphaProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Governoralpha contract.
type GovernoralphaProposalExecutedIterator struct {
	Event *GovernoralphaProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernoralphaProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernoralphaProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernoralphaProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernoralphaProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernoralphaProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernoralphaProposalExecuted represents a ProposalExecuted event raised by the Governoralpha contract.
type GovernoralphaProposalExecuted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 id)
func (_Governoralpha *GovernoralphaFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*GovernoralphaProposalExecutedIterator, error) {

	logs, sub, err := _Governoralpha.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &GovernoralphaProposalExecutedIterator{contract: _Governoralpha.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 id)
func (_Governoralpha *GovernoralphaFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernoralphaProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _Governoralpha.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernoralphaProposalExecuted)
				if err := _Governoralpha.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 id)
func (_Governoralpha *GovernoralphaFilterer) ParseProposalExecuted(log types.Log) (*GovernoralphaProposalExecuted, error) {
	event := new(GovernoralphaProposalExecuted)
	if err := _Governoralpha.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernoralphaProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the Governoralpha contract.
type GovernoralphaProposalQueuedIterator struct {
	Event *GovernoralphaProposalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernoralphaProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernoralphaProposalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernoralphaProposalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernoralphaProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernoralphaProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernoralphaProposalQueued represents a ProposalQueued event raised by the Governoralpha contract.
type GovernoralphaProposalQueued struct {
	Id  *big.Int
	Eta *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 id, uint256 eta)
func (_Governoralpha *GovernoralphaFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*GovernoralphaProposalQueuedIterator, error) {

	logs, sub, err := _Governoralpha.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &GovernoralphaProposalQueuedIterator{contract: _Governoralpha.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 id, uint256 eta)
func (_Governoralpha *GovernoralphaFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *GovernoralphaProposalQueued) (event.Subscription, error) {

	logs, sub, err := _Governoralpha.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernoralphaProposalQueued)
				if err := _Governoralpha.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 id, uint256 eta)
func (_Governoralpha *GovernoralphaFilterer) ParseProposalQueued(log types.Log) (*GovernoralphaProposalQueued, error) {
	event := new(GovernoralphaProposalQueued)
	if err := _Governoralpha.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernoralphaVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Governoralpha contract.
type GovernoralphaVoteCastIterator struct {
	Event *GovernoralphaVoteCast // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernoralphaVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernoralphaVoteCast)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernoralphaVoteCast)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernoralphaVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernoralphaVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernoralphaVoteCast represents a VoteCast event raised by the Governoralpha contract.
type GovernoralphaVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    bool
	Votes      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x877856338e13f63d0c36822ff0ef736b80934cd90574a3a5bc9262c39d217c46.
//
// Solidity: event VoteCast(address voter, uint256 proposalId, bool support, uint256 votes)
func (_Governoralpha *GovernoralphaFilterer) FilterVoteCast(opts *bind.FilterOpts) (*GovernoralphaVoteCastIterator, error) {

	logs, sub, err := _Governoralpha.contract.FilterLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return &GovernoralphaVoteCastIterator{contract: _Governoralpha.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x877856338e13f63d0c36822ff0ef736b80934cd90574a3a5bc9262c39d217c46.
//
// Solidity: event VoteCast(address voter, uint256 proposalId, bool support, uint256 votes)
func (_Governoralpha *GovernoralphaFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *GovernoralphaVoteCast) (event.Subscription, error) {

	logs, sub, err := _Governoralpha.contract.WatchLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernoralphaVoteCast)
				if err := _Governoralpha.contract.UnpackLog(event, "VoteCast", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCast is a log parse operation binding the contract event 0x877856338e13f63d0c36822ff0ef736b80934cd90574a3a5bc9262c39d217c46.
//
// Solidity: event VoteCast(address voter, uint256 proposalId, bool support, uint256 votes)
func (_Governoralpha *GovernoralphaFilterer) ParseVoteCast(log types.Log) (*GovernoralphaVoteCast, error) {
	event := new(GovernoralphaVoteCast)
	if err := _Governoralpha.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
