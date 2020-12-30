// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uv2oraclebindings

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

// FixedPointuq112x112 is an auto generated low-level Go binding around an user-defined struct.
type FixedPointuq112x112 struct {
	X *big.Int
}

// PriceLibraryPriceObservation is an auto generated low-level Go binding around an user-defined struct.
type PriceLibraryPriceObservation struct {
	Timestamp              uint32
	PriceCumulativeLast    *big.Int
	EthPriceCumulativeLast *big.Int
}

// PriceLibraryTwoWayAveragePrice is an auto generated low-level Go binding around an user-defined struct.
type PriceLibraryTwoWayAveragePrice struct {
	PriceAverage    *big.Int
	EthPriceAverage *big.Int
}

// Uv2oraclebindingsABI is the input ABI used to generate the binding from.
const Uv2oraclebindingsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"canUpdatePrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"canUpdatePrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageEthForTokens\",\"outputs\":[{\"internalType\":\"uint144[]\",\"name\":\"\",\"type\":\"uint144[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageEthForTokens\",\"outputs\":[{\"internalType\":\"uint144\",\"name\":\"\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageEthPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"internalType\":\"structFixedPoint.uq112x112\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageEthPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"internalType\":\"structFixedPoint.uq112x112[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageTokenPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"internalType\":\"structFixedPoint.uq112x112\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageTokenPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"internalType\":\"structFixedPoint.uq112x112[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"wethAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageTokensForEth\",\"outputs\":[{\"internalType\":\"uint144[]\",\"name\":\"\",\"type\":\"uint144[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeAverageTokensForEth\",\"outputs\":[{\"internalType\":\"uint144\",\"name\":\"\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeTwoWayAveragePrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"priceAverage\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"ethPriceAverage\",\"type\":\"uint224\"}],\"internalType\":\"structPriceLibrary.TwoWayAveragePrice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimeElapsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeElapsed\",\"type\":\"uint256\"}],\"name\":\"computeTwoWayAveragePrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"priceAverage\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"ethPriceAverage\",\"type\":\"uint224\"}],\"internalType\":\"structPriceLibrary.TwoWayAveragePrice[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"priceKey\",\"type\":\"uint256\"}],\"name\":\"getPriceObservationInWindow\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint224\",\"name\":\"priceCumulativeLast\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"ethPriceCumulativeLast\",\"type\":\"uint224\"}],\"internalType\":\"structPriceLibrary.PriceObservation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeTo\",\"type\":\"uint256\"}],\"name\":\"getPriceObservationsInRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint224\",\"name\":\"priceCumulativeLast\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"ethPriceCumulativeLast\",\"type\":\"uint224\"}],\"internalType\":\"structPriceLibrary.PriceObservation[]\",\"name\":\"prices\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"priceKey\",\"type\":\"uint256\"}],\"name\":\"hasPriceObservationInWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"updatePrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"updatePrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Uv2oraclebindings is an auto generated Go binding around an Ethereum contract.
type Uv2oraclebindings struct {
	Uv2oraclebindingsCaller     // Read-only binding to the contract
	Uv2oraclebindingsTransactor // Write-only binding to the contract
	Uv2oraclebindingsFilterer   // Log filterer for contract events
}

// Uv2oraclebindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uv2oraclebindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uv2oraclebindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uv2oraclebindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uv2oraclebindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uv2oraclebindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uv2oraclebindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uv2oraclebindingsSession struct {
	Contract     *Uv2oraclebindings // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Uv2oraclebindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uv2oraclebindingsCallerSession struct {
	Contract *Uv2oraclebindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Uv2oraclebindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uv2oraclebindingsTransactorSession struct {
	Contract     *Uv2oraclebindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Uv2oraclebindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uv2oraclebindingsRaw struct {
	Contract *Uv2oraclebindings // Generic contract binding to access the raw methods on
}

// Uv2oraclebindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uv2oraclebindingsCallerRaw struct {
	Contract *Uv2oraclebindingsCaller // Generic read-only contract binding to access the raw methods on
}

// Uv2oraclebindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uv2oraclebindingsTransactorRaw struct {
	Contract *Uv2oraclebindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUv2oraclebindings creates a new instance of Uv2oraclebindings, bound to a specific deployed contract.
func NewUv2oraclebindings(address common.Address, backend bind.ContractBackend) (*Uv2oraclebindings, error) {
	contract, err := bindUv2oraclebindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uv2oraclebindings{Uv2oraclebindingsCaller: Uv2oraclebindingsCaller{contract: contract}, Uv2oraclebindingsTransactor: Uv2oraclebindingsTransactor{contract: contract}, Uv2oraclebindingsFilterer: Uv2oraclebindingsFilterer{contract: contract}}, nil
}

// NewUv2oraclebindingsCaller creates a new read-only instance of Uv2oraclebindings, bound to a specific deployed contract.
func NewUv2oraclebindingsCaller(address common.Address, caller bind.ContractCaller) (*Uv2oraclebindingsCaller, error) {
	contract, err := bindUv2oraclebindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uv2oraclebindingsCaller{contract: contract}, nil
}

// NewUv2oraclebindingsTransactor creates a new write-only instance of Uv2oraclebindings, bound to a specific deployed contract.
func NewUv2oraclebindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*Uv2oraclebindingsTransactor, error) {
	contract, err := bindUv2oraclebindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uv2oraclebindingsTransactor{contract: contract}, nil
}

// NewUv2oraclebindingsFilterer creates a new log filterer instance of Uv2oraclebindings, bound to a specific deployed contract.
func NewUv2oraclebindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*Uv2oraclebindingsFilterer, error) {
	contract, err := bindUv2oraclebindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uv2oraclebindingsFilterer{contract: contract}, nil
}

// bindUv2oraclebindings binds a generic wrapper to an already deployed contract.
func bindUv2oraclebindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uv2oraclebindingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uv2oraclebindings *Uv2oraclebindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uv2oraclebindings.Contract.Uv2oraclebindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uv2oraclebindings *Uv2oraclebindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.Uv2oraclebindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uv2oraclebindings *Uv2oraclebindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.Uv2oraclebindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uv2oraclebindings *Uv2oraclebindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uv2oraclebindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uv2oraclebindings *Uv2oraclebindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uv2oraclebindings *Uv2oraclebindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.contract.Transact(opts, method, params...)
}

// CanUpdatePrice is a free data retrieval call binding the contract method 0x5acd73d7.
//
// Solidity: function canUpdatePrice(address token) view returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) CanUpdatePrice(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "canUpdatePrice", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanUpdatePrice is a free data retrieval call binding the contract method 0x5acd73d7.
//
// Solidity: function canUpdatePrice(address token) view returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsSession) CanUpdatePrice(token common.Address) (bool, error) {
	return _Uv2oraclebindings.Contract.CanUpdatePrice(&_Uv2oraclebindings.CallOpts, token)
}

// CanUpdatePrice is a free data retrieval call binding the contract method 0x5acd73d7.
//
// Solidity: function canUpdatePrice(address token) view returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) CanUpdatePrice(token common.Address) (bool, error) {
	return _Uv2oraclebindings.Contract.CanUpdatePrice(&_Uv2oraclebindings.CallOpts, token)
}

// CanUpdatePrices is a free data retrieval call binding the contract method 0x88a14c3c.
//
// Solidity: function canUpdatePrices(address[] tokens) view returns(bool[])
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) CanUpdatePrices(opts *bind.CallOpts, tokens []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "canUpdatePrices", tokens)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CanUpdatePrices is a free data retrieval call binding the contract method 0x88a14c3c.
//
// Solidity: function canUpdatePrices(address[] tokens) view returns(bool[])
func (_Uv2oraclebindings *Uv2oraclebindingsSession) CanUpdatePrices(tokens []common.Address) ([]bool, error) {
	return _Uv2oraclebindings.Contract.CanUpdatePrices(&_Uv2oraclebindings.CallOpts, tokens)
}

// CanUpdatePrices is a free data retrieval call binding the contract method 0x88a14c3c.
//
// Solidity: function canUpdatePrices(address[] tokens) view returns(bool[])
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) CanUpdatePrices(tokens []common.Address) ([]bool, error) {
	return _Uv2oraclebindings.Contract.CanUpdatePrices(&_Uv2oraclebindings.CallOpts, tokens)
}

// ComputeAverageEthForTokens is a free data retrieval call binding the contract method 0x7c906a2a.
//
// Solidity: function computeAverageEthForTokens(address[] tokens, uint256[] tokenAmounts, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144[])
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageEthForTokens(opts *bind.CallOpts, tokens []common.Address, tokenAmounts []*big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageEthForTokens", tokens, tokenAmounts, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ComputeAverageEthForTokens is a free data retrieval call binding the contract method 0x7c906a2a.
//
// Solidity: function computeAverageEthForTokens(address[] tokens, uint256[] tokenAmounts, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144[])
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageEthForTokens(tokens []common.Address, tokenAmounts []*big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthForTokens(&_Uv2oraclebindings.CallOpts, tokens, tokenAmounts, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageEthForTokens is a free data retrieval call binding the contract method 0x7c906a2a.
//
// Solidity: function computeAverageEthForTokens(address[] tokens, uint256[] tokenAmounts, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144[])
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageEthForTokens(tokens []common.Address, tokenAmounts []*big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthForTokens(&_Uv2oraclebindings.CallOpts, tokens, tokenAmounts, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageEthForTokens0 is a free data retrieval call binding the contract method 0xbf6dd82e.
//
// Solidity: function computeAverageEthForTokens(address token, uint256 tokenAmount, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144)
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageEthForTokens0(opts *bind.CallOpts, token common.Address, tokenAmount *big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageEthForTokens0", token, tokenAmount, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeAverageEthForTokens0 is a free data retrieval call binding the contract method 0xbf6dd82e.
//
// Solidity: function computeAverageEthForTokens(address token, uint256 tokenAmount, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144)
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageEthForTokens0(token common.Address, tokenAmount *big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthForTokens0(&_Uv2oraclebindings.CallOpts, token, tokenAmount, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageEthForTokens0 is a free data retrieval call binding the contract method 0xbf6dd82e.
//
// Solidity: function computeAverageEthForTokens(address token, uint256 tokenAmount, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144)
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageEthForTokens0(token common.Address, tokenAmount *big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthForTokens0(&_Uv2oraclebindings.CallOpts, token, tokenAmount, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageEthPrice is a free data retrieval call binding the contract method 0x9691b116.
//
// Solidity: function computeAverageEthPrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageEthPrice(opts *bind.CallOpts, token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (FixedPointuq112x112, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageEthPrice", token, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new(FixedPointuq112x112), err
	}

	out0 := *abi.ConvertType(out[0], new(FixedPointuq112x112)).(*FixedPointuq112x112)

	return out0, err

}

// ComputeAverageEthPrice is a free data retrieval call binding the contract method 0x9691b116.
//
// Solidity: function computeAverageEthPrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageEthPrice(token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthPrice(&_Uv2oraclebindings.CallOpts, token, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageEthPrice is a free data retrieval call binding the contract method 0x9691b116.
//
// Solidity: function computeAverageEthPrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageEthPrice(token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthPrice(&_Uv2oraclebindings.CallOpts, token, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageEthPrices is a free data retrieval call binding the contract method 0x2d726197.
//
// Solidity: function computeAverageEthPrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageEthPrices(opts *bind.CallOpts, tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]FixedPointuq112x112, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageEthPrices", tokens, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new([]FixedPointuq112x112), err
	}

	out0 := *abi.ConvertType(out[0], new([]FixedPointuq112x112)).(*[]FixedPointuq112x112)

	return out0, err

}

// ComputeAverageEthPrices is a free data retrieval call binding the contract method 0x2d726197.
//
// Solidity: function computeAverageEthPrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageEthPrices(tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthPrices(&_Uv2oraclebindings.CallOpts, tokens, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageEthPrices is a free data retrieval call binding the contract method 0x2d726197.
//
// Solidity: function computeAverageEthPrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageEthPrices(tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageEthPrices(&_Uv2oraclebindings.CallOpts, tokens, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokenPrice is a free data retrieval call binding the contract method 0x3a9d6535.
//
// Solidity: function computeAverageTokenPrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageTokenPrice(opts *bind.CallOpts, token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (FixedPointuq112x112, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageTokenPrice", token, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new(FixedPointuq112x112), err
	}

	out0 := *abi.ConvertType(out[0], new(FixedPointuq112x112)).(*FixedPointuq112x112)

	return out0, err

}

// ComputeAverageTokenPrice is a free data retrieval call binding the contract method 0x3a9d6535.
//
// Solidity: function computeAverageTokenPrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageTokenPrice(token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokenPrice(&_Uv2oraclebindings.CallOpts, token, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokenPrice is a free data retrieval call binding the contract method 0x3a9d6535.
//
// Solidity: function computeAverageTokenPrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageTokenPrice(token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokenPrice(&_Uv2oraclebindings.CallOpts, token, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokenPrices is a free data retrieval call binding the contract method 0xcdee7f6e.
//
// Solidity: function computeAverageTokenPrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageTokenPrices(opts *bind.CallOpts, tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]FixedPointuq112x112, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageTokenPrices", tokens, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new([]FixedPointuq112x112), err
	}

	out0 := *abi.ConvertType(out[0], new([]FixedPointuq112x112)).(*[]FixedPointuq112x112)

	return out0, err

}

// ComputeAverageTokenPrices is a free data retrieval call binding the contract method 0xcdee7f6e.
//
// Solidity: function computeAverageTokenPrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageTokenPrices(tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokenPrices(&_Uv2oraclebindings.CallOpts, tokens, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokenPrices is a free data retrieval call binding the contract method 0xcdee7f6e.
//
// Solidity: function computeAverageTokenPrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageTokenPrices(tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]FixedPointuq112x112, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokenPrices(&_Uv2oraclebindings.CallOpts, tokens, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokensForEth is a free data retrieval call binding the contract method 0x45060060.
//
// Solidity: function computeAverageTokensForEth(address[] tokens, uint256[] wethAmounts, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144[])
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageTokensForEth(opts *bind.CallOpts, tokens []common.Address, wethAmounts []*big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageTokensForEth", tokens, wethAmounts, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ComputeAverageTokensForEth is a free data retrieval call binding the contract method 0x45060060.
//
// Solidity: function computeAverageTokensForEth(address[] tokens, uint256[] wethAmounts, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144[])
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageTokensForEth(tokens []common.Address, wethAmounts []*big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokensForEth(&_Uv2oraclebindings.CallOpts, tokens, wethAmounts, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokensForEth is a free data retrieval call binding the contract method 0x45060060.
//
// Solidity: function computeAverageTokensForEth(address[] tokens, uint256[] wethAmounts, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144[])
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageTokensForEth(tokens []common.Address, wethAmounts []*big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokensForEth(&_Uv2oraclebindings.CallOpts, tokens, wethAmounts, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokensForEth0 is a free data retrieval call binding the contract method 0x4f2acedf.
//
// Solidity: function computeAverageTokensForEth(address token, uint256 wethAmount, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144)
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeAverageTokensForEth0(opts *bind.CallOpts, token common.Address, wethAmount *big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeAverageTokensForEth0", token, wethAmount, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeAverageTokensForEth0 is a free data retrieval call binding the contract method 0x4f2acedf.
//
// Solidity: function computeAverageTokensForEth(address token, uint256 wethAmount, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144)
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeAverageTokensForEth0(token common.Address, wethAmount *big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokensForEth0(&_Uv2oraclebindings.CallOpts, token, wethAmount, minTimeElapsed, maxTimeElapsed)
}

// ComputeAverageTokensForEth0 is a free data retrieval call binding the contract method 0x4f2acedf.
//
// Solidity: function computeAverageTokensForEth(address token, uint256 wethAmount, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns(uint144)
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeAverageTokensForEth0(token common.Address, wethAmount *big.Int, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (*big.Int, error) {
	return _Uv2oraclebindings.Contract.ComputeAverageTokensForEth0(&_Uv2oraclebindings.CallOpts, token, wethAmount, minTimeElapsed, maxTimeElapsed)
}

// ComputeTwoWayAveragePrice is a free data retrieval call binding the contract method 0x57653615.
//
// Solidity: function computeTwoWayAveragePrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224,uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeTwoWayAveragePrice(opts *bind.CallOpts, token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (PriceLibraryTwoWayAveragePrice, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeTwoWayAveragePrice", token, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new(PriceLibraryTwoWayAveragePrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceLibraryTwoWayAveragePrice)).(*PriceLibraryTwoWayAveragePrice)

	return out0, err

}

// ComputeTwoWayAveragePrice is a free data retrieval call binding the contract method 0x57653615.
//
// Solidity: function computeTwoWayAveragePrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224,uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeTwoWayAveragePrice(token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (PriceLibraryTwoWayAveragePrice, error) {
	return _Uv2oraclebindings.Contract.ComputeTwoWayAveragePrice(&_Uv2oraclebindings.CallOpts, token, minTimeElapsed, maxTimeElapsed)
}

// ComputeTwoWayAveragePrice is a free data retrieval call binding the contract method 0x57653615.
//
// Solidity: function computeTwoWayAveragePrice(address token, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224,uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeTwoWayAveragePrice(token common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) (PriceLibraryTwoWayAveragePrice, error) {
	return _Uv2oraclebindings.Contract.ComputeTwoWayAveragePrice(&_Uv2oraclebindings.CallOpts, token, minTimeElapsed, maxTimeElapsed)
}

// ComputeTwoWayAveragePrices is a free data retrieval call binding the contract method 0x768d2c19.
//
// Solidity: function computeTwoWayAveragePrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224,uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) ComputeTwoWayAveragePrices(opts *bind.CallOpts, tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]PriceLibraryTwoWayAveragePrice, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "computeTwoWayAveragePrices", tokens, minTimeElapsed, maxTimeElapsed)

	if err != nil {
		return *new([]PriceLibraryTwoWayAveragePrice), err
	}

	out0 := *abi.ConvertType(out[0], new([]PriceLibraryTwoWayAveragePrice)).(*[]PriceLibraryTwoWayAveragePrice)

	return out0, err

}

// ComputeTwoWayAveragePrices is a free data retrieval call binding the contract method 0x768d2c19.
//
// Solidity: function computeTwoWayAveragePrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224,uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsSession) ComputeTwoWayAveragePrices(tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]PriceLibraryTwoWayAveragePrice, error) {
	return _Uv2oraclebindings.Contract.ComputeTwoWayAveragePrices(&_Uv2oraclebindings.CallOpts, tokens, minTimeElapsed, maxTimeElapsed)
}

// ComputeTwoWayAveragePrices is a free data retrieval call binding the contract method 0x768d2c19.
//
// Solidity: function computeTwoWayAveragePrices(address[] tokens, uint256 minTimeElapsed, uint256 maxTimeElapsed) view returns((uint224,uint224)[])
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) ComputeTwoWayAveragePrices(tokens []common.Address, minTimeElapsed *big.Int, maxTimeElapsed *big.Int) ([]PriceLibraryTwoWayAveragePrice, error) {
	return _Uv2oraclebindings.Contract.ComputeTwoWayAveragePrices(&_Uv2oraclebindings.CallOpts, tokens, minTimeElapsed, maxTimeElapsed)
}

// GetPriceObservationInWindow is a free data retrieval call binding the contract method 0xc735e38c.
//
// Solidity: function getPriceObservationInWindow(address token, uint256 priceKey) view returns((uint32,uint224,uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) GetPriceObservationInWindow(opts *bind.CallOpts, token common.Address, priceKey *big.Int) (PriceLibraryPriceObservation, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "getPriceObservationInWindow", token, priceKey)

	if err != nil {
		return *new(PriceLibraryPriceObservation), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceLibraryPriceObservation)).(*PriceLibraryPriceObservation)

	return out0, err

}

// GetPriceObservationInWindow is a free data retrieval call binding the contract method 0xc735e38c.
//
// Solidity: function getPriceObservationInWindow(address token, uint256 priceKey) view returns((uint32,uint224,uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsSession) GetPriceObservationInWindow(token common.Address, priceKey *big.Int) (PriceLibraryPriceObservation, error) {
	return _Uv2oraclebindings.Contract.GetPriceObservationInWindow(&_Uv2oraclebindings.CallOpts, token, priceKey)
}

// GetPriceObservationInWindow is a free data retrieval call binding the contract method 0xc735e38c.
//
// Solidity: function getPriceObservationInWindow(address token, uint256 priceKey) view returns((uint32,uint224,uint224))
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) GetPriceObservationInWindow(token common.Address, priceKey *big.Int) (PriceLibraryPriceObservation, error) {
	return _Uv2oraclebindings.Contract.GetPriceObservationInWindow(&_Uv2oraclebindings.CallOpts, token, priceKey)
}

// GetPriceObservationsInRange is a free data retrieval call binding the contract method 0xffd97b6f.
//
// Solidity: function getPriceObservationsInRange(address token, uint256 timeFrom, uint256 timeTo) view returns((uint32,uint224,uint224)[] prices)
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) GetPriceObservationsInRange(opts *bind.CallOpts, token common.Address, timeFrom *big.Int, timeTo *big.Int) ([]PriceLibraryPriceObservation, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "getPriceObservationsInRange", token, timeFrom, timeTo)

	if err != nil {
		return *new([]PriceLibraryPriceObservation), err
	}

	out0 := *abi.ConvertType(out[0], new([]PriceLibraryPriceObservation)).(*[]PriceLibraryPriceObservation)

	return out0, err

}

// GetPriceObservationsInRange is a free data retrieval call binding the contract method 0xffd97b6f.
//
// Solidity: function getPriceObservationsInRange(address token, uint256 timeFrom, uint256 timeTo) view returns((uint32,uint224,uint224)[] prices)
func (_Uv2oraclebindings *Uv2oraclebindingsSession) GetPriceObservationsInRange(token common.Address, timeFrom *big.Int, timeTo *big.Int) ([]PriceLibraryPriceObservation, error) {
	return _Uv2oraclebindings.Contract.GetPriceObservationsInRange(&_Uv2oraclebindings.CallOpts, token, timeFrom, timeTo)
}

// GetPriceObservationsInRange is a free data retrieval call binding the contract method 0xffd97b6f.
//
// Solidity: function getPriceObservationsInRange(address token, uint256 timeFrom, uint256 timeTo) view returns((uint32,uint224,uint224)[] prices)
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) GetPriceObservationsInRange(token common.Address, timeFrom *big.Int, timeTo *big.Int) ([]PriceLibraryPriceObservation, error) {
	return _Uv2oraclebindings.Contract.GetPriceObservationsInRange(&_Uv2oraclebindings.CallOpts, token, timeFrom, timeTo)
}

// HasPriceObservationInWindow is a free data retrieval call binding the contract method 0x10bc482a.
//
// Solidity: function hasPriceObservationInWindow(address token, uint256 priceKey) view returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsCaller) HasPriceObservationInWindow(opts *bind.CallOpts, token common.Address, priceKey *big.Int) (bool, error) {
	var out []interface{}
	err := _Uv2oraclebindings.contract.Call(opts, &out, "hasPriceObservationInWindow", token, priceKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPriceObservationInWindow is a free data retrieval call binding the contract method 0x10bc482a.
//
// Solidity: function hasPriceObservationInWindow(address token, uint256 priceKey) view returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsSession) HasPriceObservationInWindow(token common.Address, priceKey *big.Int) (bool, error) {
	return _Uv2oraclebindings.Contract.HasPriceObservationInWindow(&_Uv2oraclebindings.CallOpts, token, priceKey)
}

// HasPriceObservationInWindow is a free data retrieval call binding the contract method 0x10bc482a.
//
// Solidity: function hasPriceObservationInWindow(address token, uint256 priceKey) view returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsCallerSession) HasPriceObservationInWindow(token common.Address, priceKey *big.Int) (bool, error) {
	return _Uv2oraclebindings.Contract.HasPriceObservationInWindow(&_Uv2oraclebindings.CallOpts, token, priceKey)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x96e85ced.
//
// Solidity: function updatePrice(address token) returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsTransactor) UpdatePrice(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Uv2oraclebindings.contract.Transact(opts, "updatePrice", token)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x96e85ced.
//
// Solidity: function updatePrice(address token) returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsSession) UpdatePrice(token common.Address) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.UpdatePrice(&_Uv2oraclebindings.TransactOpts, token)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x96e85ced.
//
// Solidity: function updatePrice(address token) returns(bool)
func (_Uv2oraclebindings *Uv2oraclebindingsTransactorSession) UpdatePrice(token common.Address) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.UpdatePrice(&_Uv2oraclebindings.TransactOpts, token)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x7e5296e0.
//
// Solidity: function updatePrices(address[] tokens) returns(bool[])
func (_Uv2oraclebindings *Uv2oraclebindingsTransactor) UpdatePrices(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _Uv2oraclebindings.contract.Transact(opts, "updatePrices", tokens)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x7e5296e0.
//
// Solidity: function updatePrices(address[] tokens) returns(bool[])
func (_Uv2oraclebindings *Uv2oraclebindingsSession) UpdatePrices(tokens []common.Address) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.UpdatePrices(&_Uv2oraclebindings.TransactOpts, tokens)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x7e5296e0.
//
// Solidity: function updatePrices(address[] tokens) returns(bool[])
func (_Uv2oraclebindings *Uv2oraclebindingsTransactorSession) UpdatePrices(tokens []common.Address) (*types.Transaction, error) {
	return _Uv2oraclebindings.Contract.UpdatePrices(&_Uv2oraclebindings.TransactOpts, tokens)
}
