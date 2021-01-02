// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mcapscontroller

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

// McapscontrollerABI is the input ABI used to generate the binding from.
const McapscontrollerABI = "[{\"inputs\":[{\"internalType\":\"contractIIndexedUniswapV2Oracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"contractIPoolFactory\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"contractIDelegateCallProxyManager\",\"name\":\"proxyManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"CategoryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"CategorySorted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"NewPoolInitializer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unboundTokenSeller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categoryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"computeAverageMarketCap\",\"outputs\":[{\"internalType\":\"uint144\",\"name\":\"\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"computeAverageMarketCaps\",\"outputs\":[{\"internalType\":\"uint144[]\",\"name\":\"marketCaps\",\"type\":\"uint144[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"computeInitializerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initializerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"}],\"name\":\"computePoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"computeSellerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sellerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"createCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSellerPremium\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"name\":\"finishPreparedIndexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getCategoryMarketCaps\",\"outputs\":[{\"internalType\":\"uint144[]\",\"name\":\"marketCaps\",\"type\":\"uint144[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getCategoryTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"},{\"internalType\":\"uint144\",\"name\":\"wethValue\",\"type\":\"uint144\"}],\"name\":\"getInitialTokensAndBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"getLastCategoryUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getTopCategoryTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"hasCategory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenInCategory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIIndexedUniswapV2Oracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"orderCategoryTokensByMarketCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialWethValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"prepareIndexPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initializerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"reindexPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"reweighPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_defaultSellerPremium\",\"type\":\"uint8\"}],\"name\":\"setDefaultSellerPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolTokens\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"categoryID\",\"type\":\"uint256\"}],\"name\":\"updateCategoryPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"pricesUpdated\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIndexPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"updateMinimumBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSeller\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercent\",\"type\":\"uint8\"}],\"name\":\"updateSellerPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Mcapscontroller is an auto generated Go binding around an Ethereum contract.
type Mcapscontroller struct {
	McapscontrollerCaller     // Read-only binding to the contract
	McapscontrollerTransactor // Write-only binding to the contract
	McapscontrollerFilterer   // Log filterer for contract events
}

// McapscontrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type McapscontrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McapscontrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type McapscontrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McapscontrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type McapscontrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// McapscontrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type McapscontrollerSession struct {
	Contract     *Mcapscontroller  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// McapscontrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type McapscontrollerCallerSession struct {
	Contract *McapscontrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// McapscontrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type McapscontrollerTransactorSession struct {
	Contract     *McapscontrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// McapscontrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type McapscontrollerRaw struct {
	Contract *Mcapscontroller // Generic contract binding to access the raw methods on
}

// McapscontrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type McapscontrollerCallerRaw struct {
	Contract *McapscontrollerCaller // Generic read-only contract binding to access the raw methods on
}

// McapscontrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type McapscontrollerTransactorRaw struct {
	Contract *McapscontrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMcapscontroller creates a new instance of Mcapscontroller, bound to a specific deployed contract.
func NewMcapscontroller(address common.Address, backend bind.ContractBackend) (*Mcapscontroller, error) {
	contract, err := bindMcapscontroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mcapscontroller{McapscontrollerCaller: McapscontrollerCaller{contract: contract}, McapscontrollerTransactor: McapscontrollerTransactor{contract: contract}, McapscontrollerFilterer: McapscontrollerFilterer{contract: contract}}, nil
}

// NewMcapscontrollerCaller creates a new read-only instance of Mcapscontroller, bound to a specific deployed contract.
func NewMcapscontrollerCaller(address common.Address, caller bind.ContractCaller) (*McapscontrollerCaller, error) {
	contract, err := bindMcapscontroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &McapscontrollerCaller{contract: contract}, nil
}

// NewMcapscontrollerTransactor creates a new write-only instance of Mcapscontroller, bound to a specific deployed contract.
func NewMcapscontrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*McapscontrollerTransactor, error) {
	contract, err := bindMcapscontroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &McapscontrollerTransactor{contract: contract}, nil
}

// NewMcapscontrollerFilterer creates a new log filterer instance of Mcapscontroller, bound to a specific deployed contract.
func NewMcapscontrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*McapscontrollerFilterer, error) {
	contract, err := bindMcapscontroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &McapscontrollerFilterer{contract: contract}, nil
}

// bindMcapscontroller binds a generic wrapper to an already deployed contract.
func bindMcapscontroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(McapscontrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcapscontroller *McapscontrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mcapscontroller.Contract.McapscontrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcapscontroller *McapscontrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.McapscontrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcapscontroller *McapscontrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.McapscontrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mcapscontroller *McapscontrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mcapscontroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mcapscontroller *McapscontrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mcapscontroller *McapscontrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.contract.Transact(opts, method, params...)
}

// CategoryIndex is a free data retrieval call binding the contract method 0xea99fc04.
//
// Solidity: function categoryIndex() view returns(uint256)
func (_Mcapscontroller *McapscontrollerCaller) CategoryIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "categoryIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CategoryIndex is a free data retrieval call binding the contract method 0xea99fc04.
//
// Solidity: function categoryIndex() view returns(uint256)
func (_Mcapscontroller *McapscontrollerSession) CategoryIndex() (*big.Int, error) {
	return _Mcapscontroller.Contract.CategoryIndex(&_Mcapscontroller.CallOpts)
}

// CategoryIndex is a free data retrieval call binding the contract method 0xea99fc04.
//
// Solidity: function categoryIndex() view returns(uint256)
func (_Mcapscontroller *McapscontrollerCallerSession) CategoryIndex() (*big.Int, error) {
	return _Mcapscontroller.Contract.CategoryIndex(&_Mcapscontroller.CallOpts)
}

// ComputeAverageMarketCap is a free data retrieval call binding the contract method 0xfd371138.
//
// Solidity: function computeAverageMarketCap(address token) view returns(uint144)
func (_Mcapscontroller *McapscontrollerCaller) ComputeAverageMarketCap(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "computeAverageMarketCap", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeAverageMarketCap is a free data retrieval call binding the contract method 0xfd371138.
//
// Solidity: function computeAverageMarketCap(address token) view returns(uint144)
func (_Mcapscontroller *McapscontrollerSession) ComputeAverageMarketCap(token common.Address) (*big.Int, error) {
	return _Mcapscontroller.Contract.ComputeAverageMarketCap(&_Mcapscontroller.CallOpts, token)
}

// ComputeAverageMarketCap is a free data retrieval call binding the contract method 0xfd371138.
//
// Solidity: function computeAverageMarketCap(address token) view returns(uint144)
func (_Mcapscontroller *McapscontrollerCallerSession) ComputeAverageMarketCap(token common.Address) (*big.Int, error) {
	return _Mcapscontroller.Contract.ComputeAverageMarketCap(&_Mcapscontroller.CallOpts, token)
}

// ComputeAverageMarketCaps is a free data retrieval call binding the contract method 0x6bfc3d40.
//
// Solidity: function computeAverageMarketCaps(address[] tokens) view returns(uint144[] marketCaps)
func (_Mcapscontroller *McapscontrollerCaller) ComputeAverageMarketCaps(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "computeAverageMarketCaps", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ComputeAverageMarketCaps is a free data retrieval call binding the contract method 0x6bfc3d40.
//
// Solidity: function computeAverageMarketCaps(address[] tokens) view returns(uint144[] marketCaps)
func (_Mcapscontroller *McapscontrollerSession) ComputeAverageMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Mcapscontroller.Contract.ComputeAverageMarketCaps(&_Mcapscontroller.CallOpts, tokens)
}

// ComputeAverageMarketCaps is a free data retrieval call binding the contract method 0x6bfc3d40.
//
// Solidity: function computeAverageMarketCaps(address[] tokens) view returns(uint144[] marketCaps)
func (_Mcapscontroller *McapscontrollerCallerSession) ComputeAverageMarketCaps(tokens []common.Address) ([]*big.Int, error) {
	return _Mcapscontroller.Contract.ComputeAverageMarketCaps(&_Mcapscontroller.CallOpts, tokens)
}

// ComputeInitializerAddress is a free data retrieval call binding the contract method 0xfa7c0b5e.
//
// Solidity: function computeInitializerAddress(address poolAddress) view returns(address initializerAddress)
func (_Mcapscontroller *McapscontrollerCaller) ComputeInitializerAddress(opts *bind.CallOpts, poolAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "computeInitializerAddress", poolAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeInitializerAddress is a free data retrieval call binding the contract method 0xfa7c0b5e.
//
// Solidity: function computeInitializerAddress(address poolAddress) view returns(address initializerAddress)
func (_Mcapscontroller *McapscontrollerSession) ComputeInitializerAddress(poolAddress common.Address) (common.Address, error) {
	return _Mcapscontroller.Contract.ComputeInitializerAddress(&_Mcapscontroller.CallOpts, poolAddress)
}

// ComputeInitializerAddress is a free data retrieval call binding the contract method 0xfa7c0b5e.
//
// Solidity: function computeInitializerAddress(address poolAddress) view returns(address initializerAddress)
func (_Mcapscontroller *McapscontrollerCallerSession) ComputeInitializerAddress(poolAddress common.Address) (common.Address, error) {
	return _Mcapscontroller.Contract.ComputeInitializerAddress(&_Mcapscontroller.CallOpts, poolAddress)
}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 categoryID, uint256 indexSize) view returns(address poolAddress)
func (_Mcapscontroller *McapscontrollerCaller) ComputePoolAddress(opts *bind.CallOpts, categoryID *big.Int, indexSize *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "computePoolAddress", categoryID, indexSize)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 categoryID, uint256 indexSize) view returns(address poolAddress)
func (_Mcapscontroller *McapscontrollerSession) ComputePoolAddress(categoryID *big.Int, indexSize *big.Int) (common.Address, error) {
	return _Mcapscontroller.Contract.ComputePoolAddress(&_Mcapscontroller.CallOpts, categoryID, indexSize)
}

// ComputePoolAddress is a free data retrieval call binding the contract method 0x9dd19a52.
//
// Solidity: function computePoolAddress(uint256 categoryID, uint256 indexSize) view returns(address poolAddress)
func (_Mcapscontroller *McapscontrollerCallerSession) ComputePoolAddress(categoryID *big.Int, indexSize *big.Int) (common.Address, error) {
	return _Mcapscontroller.Contract.ComputePoolAddress(&_Mcapscontroller.CallOpts, categoryID, indexSize)
}

// ComputeSellerAddress is a free data retrieval call binding the contract method 0x582edc9b.
//
// Solidity: function computeSellerAddress(address poolAddress) view returns(address sellerAddress)
func (_Mcapscontroller *McapscontrollerCaller) ComputeSellerAddress(opts *bind.CallOpts, poolAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "computeSellerAddress", poolAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeSellerAddress is a free data retrieval call binding the contract method 0x582edc9b.
//
// Solidity: function computeSellerAddress(address poolAddress) view returns(address sellerAddress)
func (_Mcapscontroller *McapscontrollerSession) ComputeSellerAddress(poolAddress common.Address) (common.Address, error) {
	return _Mcapscontroller.Contract.ComputeSellerAddress(&_Mcapscontroller.CallOpts, poolAddress)
}

// ComputeSellerAddress is a free data retrieval call binding the contract method 0x582edc9b.
//
// Solidity: function computeSellerAddress(address poolAddress) view returns(address sellerAddress)
func (_Mcapscontroller *McapscontrollerCallerSession) ComputeSellerAddress(poolAddress common.Address) (common.Address, error) {
	return _Mcapscontroller.Contract.ComputeSellerAddress(&_Mcapscontroller.CallOpts, poolAddress)
}

// DefaultSellerPremium is a free data retrieval call binding the contract method 0x372accd2.
//
// Solidity: function defaultSellerPremium() view returns(uint8)
func (_Mcapscontroller *McapscontrollerCaller) DefaultSellerPremium(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "defaultSellerPremium")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultSellerPremium is a free data retrieval call binding the contract method 0x372accd2.
//
// Solidity: function defaultSellerPremium() view returns(uint8)
func (_Mcapscontroller *McapscontrollerSession) DefaultSellerPremium() (uint8, error) {
	return _Mcapscontroller.Contract.DefaultSellerPremium(&_Mcapscontroller.CallOpts)
}

// DefaultSellerPremium is a free data retrieval call binding the contract method 0x372accd2.
//
// Solidity: function defaultSellerPremium() view returns(uint8)
func (_Mcapscontroller *McapscontrollerCallerSession) DefaultSellerPremium() (uint8, error) {
	return _Mcapscontroller.Contract.DefaultSellerPremium(&_Mcapscontroller.CallOpts)
}

// GetCategoryMarketCaps is a free data retrieval call binding the contract method 0xf3cd354f.
//
// Solidity: function getCategoryMarketCaps(uint256 categoryID) view returns(uint144[] marketCaps)
func (_Mcapscontroller *McapscontrollerCaller) GetCategoryMarketCaps(opts *bind.CallOpts, categoryID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "getCategoryMarketCaps", categoryID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCategoryMarketCaps is a free data retrieval call binding the contract method 0xf3cd354f.
//
// Solidity: function getCategoryMarketCaps(uint256 categoryID) view returns(uint144[] marketCaps)
func (_Mcapscontroller *McapscontrollerSession) GetCategoryMarketCaps(categoryID *big.Int) ([]*big.Int, error) {
	return _Mcapscontroller.Contract.GetCategoryMarketCaps(&_Mcapscontroller.CallOpts, categoryID)
}

// GetCategoryMarketCaps is a free data retrieval call binding the contract method 0xf3cd354f.
//
// Solidity: function getCategoryMarketCaps(uint256 categoryID) view returns(uint144[] marketCaps)
func (_Mcapscontroller *McapscontrollerCallerSession) GetCategoryMarketCaps(categoryID *big.Int) ([]*big.Int, error) {
	return _Mcapscontroller.Contract.GetCategoryMarketCaps(&_Mcapscontroller.CallOpts, categoryID)
}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[] tokens)
func (_Mcapscontroller *McapscontrollerCaller) GetCategoryTokens(opts *bind.CallOpts, categoryID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "getCategoryTokens", categoryID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[] tokens)
func (_Mcapscontroller *McapscontrollerSession) GetCategoryTokens(categoryID *big.Int) ([]common.Address, error) {
	return _Mcapscontroller.Contract.GetCategoryTokens(&_Mcapscontroller.CallOpts, categoryID)
}

// GetCategoryTokens is a free data retrieval call binding the contract method 0xde105dbd.
//
// Solidity: function getCategoryTokens(uint256 categoryID) view returns(address[] tokens)
func (_Mcapscontroller *McapscontrollerCallerSession) GetCategoryTokens(categoryID *big.Int) ([]common.Address, error) {
	return _Mcapscontroller.Contract.GetCategoryTokens(&_Mcapscontroller.CallOpts, categoryID)
}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0x36176824.
//
// Solidity: function getInitialTokensAndBalances(uint256 categoryID, uint256 indexSize, uint144 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Mcapscontroller *McapscontrollerCaller) GetInitialTokensAndBalances(opts *bind.CallOpts, categoryID *big.Int, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "getInitialTokensAndBalances", categoryID, indexSize, wethValue)

	outstruct := new(struct {
		Tokens   []common.Address
		Balances []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = out[0].([]common.Address)
	outstruct.Balances = out[1].([]*big.Int)

	return *outstruct, err

}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0x36176824.
//
// Solidity: function getInitialTokensAndBalances(uint256 categoryID, uint256 indexSize, uint144 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Mcapscontroller *McapscontrollerSession) GetInitialTokensAndBalances(categoryID *big.Int, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	return _Mcapscontroller.Contract.GetInitialTokensAndBalances(&_Mcapscontroller.CallOpts, categoryID, indexSize, wethValue)
}

// GetInitialTokensAndBalances is a free data retrieval call binding the contract method 0x36176824.
//
// Solidity: function getInitialTokensAndBalances(uint256 categoryID, uint256 indexSize, uint144 wethValue) view returns(address[] tokens, uint256[] balances)
func (_Mcapscontroller *McapscontrollerCallerSession) GetInitialTokensAndBalances(categoryID *big.Int, indexSize *big.Int, wethValue *big.Int) (struct {
	Tokens   []common.Address
	Balances []*big.Int
}, error) {
	return _Mcapscontroller.Contract.GetInitialTokensAndBalances(&_Mcapscontroller.CallOpts, categoryID, indexSize, wethValue)
}

// GetLastCategoryUpdate is a free data retrieval call binding the contract method 0xf256a9d0.
//
// Solidity: function getLastCategoryUpdate(uint256 categoryID) view returns(uint256)
func (_Mcapscontroller *McapscontrollerCaller) GetLastCategoryUpdate(opts *bind.CallOpts, categoryID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "getLastCategoryUpdate", categoryID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCategoryUpdate is a free data retrieval call binding the contract method 0xf256a9d0.
//
// Solidity: function getLastCategoryUpdate(uint256 categoryID) view returns(uint256)
func (_Mcapscontroller *McapscontrollerSession) GetLastCategoryUpdate(categoryID *big.Int) (*big.Int, error) {
	return _Mcapscontroller.Contract.GetLastCategoryUpdate(&_Mcapscontroller.CallOpts, categoryID)
}

// GetLastCategoryUpdate is a free data retrieval call binding the contract method 0xf256a9d0.
//
// Solidity: function getLastCategoryUpdate(uint256 categoryID) view returns(uint256)
func (_Mcapscontroller *McapscontrollerCallerSession) GetLastCategoryUpdate(categoryID *big.Int) (*big.Int, error) {
	return _Mcapscontroller.Contract.GetLastCategoryUpdate(&_Mcapscontroller.CallOpts, categoryID)
}

// GetTopCategoryTokens is a free data retrieval call binding the contract method 0x030fb7f4.
//
// Solidity: function getTopCategoryTokens(uint256 categoryID, uint256 num) view returns(address[] tokens)
func (_Mcapscontroller *McapscontrollerCaller) GetTopCategoryTokens(opts *bind.CallOpts, categoryID *big.Int, num *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "getTopCategoryTokens", categoryID, num)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTopCategoryTokens is a free data retrieval call binding the contract method 0x030fb7f4.
//
// Solidity: function getTopCategoryTokens(uint256 categoryID, uint256 num) view returns(address[] tokens)
func (_Mcapscontroller *McapscontrollerSession) GetTopCategoryTokens(categoryID *big.Int, num *big.Int) ([]common.Address, error) {
	return _Mcapscontroller.Contract.GetTopCategoryTokens(&_Mcapscontroller.CallOpts, categoryID, num)
}

// GetTopCategoryTokens is a free data retrieval call binding the contract method 0x030fb7f4.
//
// Solidity: function getTopCategoryTokens(uint256 categoryID, uint256 num) view returns(address[] tokens)
func (_Mcapscontroller *McapscontrollerCallerSession) GetTopCategoryTokens(categoryID *big.Int, num *big.Int) ([]common.Address, error) {
	return _Mcapscontroller.Contract.GetTopCategoryTokens(&_Mcapscontroller.CallOpts, categoryID, num)
}

// HasCategory is a free data retrieval call binding the contract method 0x1ab66317.
//
// Solidity: function hasCategory(uint256 categoryID) view returns(bool)
func (_Mcapscontroller *McapscontrollerCaller) HasCategory(opts *bind.CallOpts, categoryID *big.Int) (bool, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "hasCategory", categoryID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCategory is a free data retrieval call binding the contract method 0x1ab66317.
//
// Solidity: function hasCategory(uint256 categoryID) view returns(bool)
func (_Mcapscontroller *McapscontrollerSession) HasCategory(categoryID *big.Int) (bool, error) {
	return _Mcapscontroller.Contract.HasCategory(&_Mcapscontroller.CallOpts, categoryID)
}

// HasCategory is a free data retrieval call binding the contract method 0x1ab66317.
//
// Solidity: function hasCategory(uint256 categoryID) view returns(bool)
func (_Mcapscontroller *McapscontrollerCallerSession) HasCategory(categoryID *big.Int) (bool, error) {
	return _Mcapscontroller.Contract.HasCategory(&_Mcapscontroller.CallOpts, categoryID)
}

// IsTokenInCategory is a free data retrieval call binding the contract method 0x1fbaf275.
//
// Solidity: function isTokenInCategory(uint256 categoryID, address token) view returns(bool)
func (_Mcapscontroller *McapscontrollerCaller) IsTokenInCategory(opts *bind.CallOpts, categoryID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "isTokenInCategory", categoryID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenInCategory is a free data retrieval call binding the contract method 0x1fbaf275.
//
// Solidity: function isTokenInCategory(uint256 categoryID, address token) view returns(bool)
func (_Mcapscontroller *McapscontrollerSession) IsTokenInCategory(categoryID *big.Int, token common.Address) (bool, error) {
	return _Mcapscontroller.Contract.IsTokenInCategory(&_Mcapscontroller.CallOpts, categoryID, token)
}

// IsTokenInCategory is a free data retrieval call binding the contract method 0x1fbaf275.
//
// Solidity: function isTokenInCategory(uint256 categoryID, address token) view returns(bool)
func (_Mcapscontroller *McapscontrollerCallerSession) IsTokenInCategory(categoryID *big.Int, token common.Address) (bool, error) {
	return _Mcapscontroller.Contract.IsTokenInCategory(&_Mcapscontroller.CallOpts, categoryID, token)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Mcapscontroller *McapscontrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Mcapscontroller *McapscontrollerSession) Oracle() (common.Address, error) {
	return _Mcapscontroller.Contract.Oracle(&_Mcapscontroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Mcapscontroller *McapscontrollerCallerSession) Oracle() (common.Address, error) {
	return _Mcapscontroller.Contract.Oracle(&_Mcapscontroller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mcapscontroller *McapscontrollerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mcapscontroller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mcapscontroller *McapscontrollerSession) Owner() (common.Address, error) {
	return _Mcapscontroller.Contract.Owner(&_Mcapscontroller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mcapscontroller *McapscontrollerCallerSession) Owner() (common.Address, error) {
	return _Mcapscontroller.Contract.Owner(&_Mcapscontroller.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 categoryID, address token) returns()
func (_Mcapscontroller *McapscontrollerTransactor) AddToken(opts *bind.TransactOpts, categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "addToken", categoryID, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 categoryID, address token) returns()
func (_Mcapscontroller *McapscontrollerSession) AddToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.AddToken(&_Mcapscontroller.TransactOpts, categoryID, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x57016b0a.
//
// Solidity: function addToken(uint256 categoryID, address token) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) AddToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.AddToken(&_Mcapscontroller.TransactOpts, categoryID, token)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 categoryID, address[] tokens) returns()
func (_Mcapscontroller *McapscontrollerTransactor) AddTokens(opts *bind.TransactOpts, categoryID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "addTokens", categoryID, tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 categoryID, address[] tokens) returns()
func (_Mcapscontroller *McapscontrollerSession) AddTokens(categoryID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.AddTokens(&_Mcapscontroller.TransactOpts, categoryID, tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0x6b5b8808.
//
// Solidity: function addTokens(uint256 categoryID, address[] tokens) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) AddTokens(categoryID *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.AddTokens(&_Mcapscontroller.TransactOpts, categoryID, tokens)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x6696bf11.
//
// Solidity: function createCategory(bytes32 metadataHash) returns()
func (_Mcapscontroller *McapscontrollerTransactor) CreateCategory(opts *bind.TransactOpts, metadataHash [32]byte) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "createCategory", metadataHash)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x6696bf11.
//
// Solidity: function createCategory(bytes32 metadataHash) returns()
func (_Mcapscontroller *McapscontrollerSession) CreateCategory(metadataHash [32]byte) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.CreateCategory(&_Mcapscontroller.TransactOpts, metadataHash)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x6696bf11.
//
// Solidity: function createCategory(bytes32 metadataHash) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) CreateCategory(metadataHash [32]byte) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.CreateCategory(&_Mcapscontroller.TransactOpts, metadataHash)
}

// FinishPreparedIndexPool is a paid mutator transaction binding the contract method 0x6a7a93ed.
//
// Solidity: function finishPreparedIndexPool(address poolAddress, address[] tokens, uint256[] balances) returns()
func (_Mcapscontroller *McapscontrollerTransactor) FinishPreparedIndexPool(opts *bind.TransactOpts, poolAddress common.Address, tokens []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "finishPreparedIndexPool", poolAddress, tokens, balances)
}

// FinishPreparedIndexPool is a paid mutator transaction binding the contract method 0x6a7a93ed.
//
// Solidity: function finishPreparedIndexPool(address poolAddress, address[] tokens, uint256[] balances) returns()
func (_Mcapscontroller *McapscontrollerSession) FinishPreparedIndexPool(poolAddress common.Address, tokens []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.FinishPreparedIndexPool(&_Mcapscontroller.TransactOpts, poolAddress, tokens, balances)
}

// FinishPreparedIndexPool is a paid mutator transaction binding the contract method 0x6a7a93ed.
//
// Solidity: function finishPreparedIndexPool(address poolAddress, address[] tokens, uint256[] balances) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) FinishPreparedIndexPool(poolAddress common.Address, tokens []common.Address, balances []*big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.FinishPreparedIndexPool(&_Mcapscontroller.TransactOpts, poolAddress, tokens, balances)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mcapscontroller *McapscontrollerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mcapscontroller *McapscontrollerSession) Initialize() (*types.Transaction, error) {
	return _Mcapscontroller.Contract.Initialize(&_Mcapscontroller.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Mcapscontroller.Contract.Initialize(&_Mcapscontroller.TransactOpts)
}

// OrderCategoryTokensByMarketCap is a paid mutator transaction binding the contract method 0xb25f86f9.
//
// Solidity: function orderCategoryTokensByMarketCap(uint256 categoryID) returns()
func (_Mcapscontroller *McapscontrollerTransactor) OrderCategoryTokensByMarketCap(opts *bind.TransactOpts, categoryID *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "orderCategoryTokensByMarketCap", categoryID)
}

// OrderCategoryTokensByMarketCap is a paid mutator transaction binding the contract method 0xb25f86f9.
//
// Solidity: function orderCategoryTokensByMarketCap(uint256 categoryID) returns()
func (_Mcapscontroller *McapscontrollerSession) OrderCategoryTokensByMarketCap(categoryID *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.OrderCategoryTokensByMarketCap(&_Mcapscontroller.TransactOpts, categoryID)
}

// OrderCategoryTokensByMarketCap is a paid mutator transaction binding the contract method 0xb25f86f9.
//
// Solidity: function orderCategoryTokensByMarketCap(uint256 categoryID) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) OrderCategoryTokensByMarketCap(categoryID *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.OrderCategoryTokensByMarketCap(&_Mcapscontroller.TransactOpts, categoryID)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0x74a58783.
//
// Solidity: function prepareIndexPool(uint256 categoryID, uint256 indexSize, uint256 initialWethValue, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Mcapscontroller *McapscontrollerTransactor) PrepareIndexPool(opts *bind.TransactOpts, categoryID *big.Int, indexSize *big.Int, initialWethValue *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "prepareIndexPool", categoryID, indexSize, initialWethValue, name, symbol)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0x74a58783.
//
// Solidity: function prepareIndexPool(uint256 categoryID, uint256 indexSize, uint256 initialWethValue, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Mcapscontroller *McapscontrollerSession) PrepareIndexPool(categoryID *big.Int, indexSize *big.Int, initialWethValue *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.PrepareIndexPool(&_Mcapscontroller.TransactOpts, categoryID, indexSize, initialWethValue, name, symbol)
}

// PrepareIndexPool is a paid mutator transaction binding the contract method 0x74a58783.
//
// Solidity: function prepareIndexPool(uint256 categoryID, uint256 indexSize, uint256 initialWethValue, string name, string symbol) returns(address poolAddress, address initializerAddress)
func (_Mcapscontroller *McapscontrollerTransactorSession) PrepareIndexPool(categoryID *big.Int, indexSize *big.Int, initialWethValue *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.PrepareIndexPool(&_Mcapscontroller.TransactOpts, categoryID, indexSize, initialWethValue, name, symbol)
}

// ReindexPool is a paid mutator transaction binding the contract method 0x50b1e342.
//
// Solidity: function reindexPool(address poolAddress) returns()
func (_Mcapscontroller *McapscontrollerTransactor) ReindexPool(opts *bind.TransactOpts, poolAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "reindexPool", poolAddress)
}

// ReindexPool is a paid mutator transaction binding the contract method 0x50b1e342.
//
// Solidity: function reindexPool(address poolAddress) returns()
func (_Mcapscontroller *McapscontrollerSession) ReindexPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.ReindexPool(&_Mcapscontroller.TransactOpts, poolAddress)
}

// ReindexPool is a paid mutator transaction binding the contract method 0x50b1e342.
//
// Solidity: function reindexPool(address poolAddress) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) ReindexPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.ReindexPool(&_Mcapscontroller.TransactOpts, poolAddress)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 categoryID, address token) returns()
func (_Mcapscontroller *McapscontrollerTransactor) RemoveToken(opts *bind.TransactOpts, categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "removeToken", categoryID, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 categoryID, address token) returns()
func (_Mcapscontroller *McapscontrollerSession) RemoveToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.RemoveToken(&_Mcapscontroller.TransactOpts, categoryID, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x33fe5676.
//
// Solidity: function removeToken(uint256 categoryID, address token) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) RemoveToken(categoryID *big.Int, token common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.RemoveToken(&_Mcapscontroller.TransactOpts, categoryID, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mcapscontroller *McapscontrollerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mcapscontroller *McapscontrollerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mcapscontroller.Contract.RenounceOwnership(&_Mcapscontroller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mcapscontroller.Contract.RenounceOwnership(&_Mcapscontroller.TransactOpts)
}

// ReweighPool is a paid mutator transaction binding the contract method 0xe9819daf.
//
// Solidity: function reweighPool(address poolAddress) returns()
func (_Mcapscontroller *McapscontrollerTransactor) ReweighPool(opts *bind.TransactOpts, poolAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "reweighPool", poolAddress)
}

// ReweighPool is a paid mutator transaction binding the contract method 0xe9819daf.
//
// Solidity: function reweighPool(address poolAddress) returns()
func (_Mcapscontroller *McapscontrollerSession) ReweighPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.ReweighPool(&_Mcapscontroller.TransactOpts, poolAddress)
}

// ReweighPool is a paid mutator transaction binding the contract method 0xe9819daf.
//
// Solidity: function reweighPool(address poolAddress) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) ReweighPool(poolAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.ReweighPool(&_Mcapscontroller.TransactOpts, poolAddress)
}

// SetDefaultSellerPremium is a paid mutator transaction binding the contract method 0xdf8e238a.
//
// Solidity: function setDefaultSellerPremium(uint8 _defaultSellerPremium) returns()
func (_Mcapscontroller *McapscontrollerTransactor) SetDefaultSellerPremium(opts *bind.TransactOpts, _defaultSellerPremium uint8) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "setDefaultSellerPremium", _defaultSellerPremium)
}

// SetDefaultSellerPremium is a paid mutator transaction binding the contract method 0xdf8e238a.
//
// Solidity: function setDefaultSellerPremium(uint8 _defaultSellerPremium) returns()
func (_Mcapscontroller *McapscontrollerSession) SetDefaultSellerPremium(_defaultSellerPremium uint8) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.SetDefaultSellerPremium(&_Mcapscontroller.TransactOpts, _defaultSellerPremium)
}

// SetDefaultSellerPremium is a paid mutator transaction binding the contract method 0xdf8e238a.
//
// Solidity: function setDefaultSellerPremium(uint8 _defaultSellerPremium) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) SetDefaultSellerPremium(_defaultSellerPremium uint8) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.SetDefaultSellerPremium(&_Mcapscontroller.TransactOpts, _defaultSellerPremium)
}

// SetMaxPoolTokens is a paid mutator transaction binding the contract method 0x51531cce.
//
// Solidity: function setMaxPoolTokens(address poolAddress, uint256 maxPoolTokens) returns()
func (_Mcapscontroller *McapscontrollerTransactor) SetMaxPoolTokens(opts *bind.TransactOpts, poolAddress common.Address, maxPoolTokens *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "setMaxPoolTokens", poolAddress, maxPoolTokens)
}

// SetMaxPoolTokens is a paid mutator transaction binding the contract method 0x51531cce.
//
// Solidity: function setMaxPoolTokens(address poolAddress, uint256 maxPoolTokens) returns()
func (_Mcapscontroller *McapscontrollerSession) SetMaxPoolTokens(poolAddress common.Address, maxPoolTokens *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.SetMaxPoolTokens(&_Mcapscontroller.TransactOpts, poolAddress, maxPoolTokens)
}

// SetMaxPoolTokens is a paid mutator transaction binding the contract method 0x51531cce.
//
// Solidity: function setMaxPoolTokens(address poolAddress, uint256 maxPoolTokens) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) SetMaxPoolTokens(poolAddress common.Address, maxPoolTokens *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.SetMaxPoolTokens(&_Mcapscontroller.TransactOpts, poolAddress, maxPoolTokens)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x991991c7.
//
// Solidity: function setSwapFee(address poolAddress, uint256 swapFee) returns()
func (_Mcapscontroller *McapscontrollerTransactor) SetSwapFee(opts *bind.TransactOpts, poolAddress common.Address, swapFee *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "setSwapFee", poolAddress, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x991991c7.
//
// Solidity: function setSwapFee(address poolAddress, uint256 swapFee) returns()
func (_Mcapscontroller *McapscontrollerSession) SetSwapFee(poolAddress common.Address, swapFee *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.SetSwapFee(&_Mcapscontroller.TransactOpts, poolAddress, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x991991c7.
//
// Solidity: function setSwapFee(address poolAddress, uint256 swapFee) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) SetSwapFee(poolAddress common.Address, swapFee *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.SetSwapFee(&_Mcapscontroller.TransactOpts, poolAddress, swapFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mcapscontroller *McapscontrollerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mcapscontroller *McapscontrollerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.TransferOwnership(&_Mcapscontroller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.TransferOwnership(&_Mcapscontroller.TransactOpts, newOwner)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns(bool[] pricesUpdated)
func (_Mcapscontroller *McapscontrollerTransactor) UpdateCategoryPrices(opts *bind.TransactOpts, categoryID *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "updateCategoryPrices", categoryID)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns(bool[] pricesUpdated)
func (_Mcapscontroller *McapscontrollerSession) UpdateCategoryPrices(categoryID *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.UpdateCategoryPrices(&_Mcapscontroller.TransactOpts, categoryID)
}

// UpdateCategoryPrices is a paid mutator transaction binding the contract method 0x6866ada8.
//
// Solidity: function updateCategoryPrices(uint256 categoryID) returns(bool[] pricesUpdated)
func (_Mcapscontroller *McapscontrollerTransactorSession) UpdateCategoryPrices(categoryID *big.Int) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.UpdateCategoryPrices(&_Mcapscontroller.TransactOpts, categoryID)
}

// UpdateMinimumBalance is a paid mutator transaction binding the contract method 0x034b904e.
//
// Solidity: function updateMinimumBalance(address pool, address tokenAddress) returns()
func (_Mcapscontroller *McapscontrollerTransactor) UpdateMinimumBalance(opts *bind.TransactOpts, pool common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "updateMinimumBalance", pool, tokenAddress)
}

// UpdateMinimumBalance is a paid mutator transaction binding the contract method 0x034b904e.
//
// Solidity: function updateMinimumBalance(address pool, address tokenAddress) returns()
func (_Mcapscontroller *McapscontrollerSession) UpdateMinimumBalance(pool common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.UpdateMinimumBalance(&_Mcapscontroller.TransactOpts, pool, tokenAddress)
}

// UpdateMinimumBalance is a paid mutator transaction binding the contract method 0x034b904e.
//
// Solidity: function updateMinimumBalance(address pool, address tokenAddress) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) UpdateMinimumBalance(pool common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.UpdateMinimumBalance(&_Mcapscontroller.TransactOpts, pool, tokenAddress)
}

// UpdateSellerPremium is a paid mutator transaction binding the contract method 0xdcdf7e16.
//
// Solidity: function updateSellerPremium(address tokenSeller, uint8 premiumPercent) returns()
func (_Mcapscontroller *McapscontrollerTransactor) UpdateSellerPremium(opts *bind.TransactOpts, tokenSeller common.Address, premiumPercent uint8) (*types.Transaction, error) {
	return _Mcapscontroller.contract.Transact(opts, "updateSellerPremium", tokenSeller, premiumPercent)
}

// UpdateSellerPremium is a paid mutator transaction binding the contract method 0xdcdf7e16.
//
// Solidity: function updateSellerPremium(address tokenSeller, uint8 premiumPercent) returns()
func (_Mcapscontroller *McapscontrollerSession) UpdateSellerPremium(tokenSeller common.Address, premiumPercent uint8) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.UpdateSellerPremium(&_Mcapscontroller.TransactOpts, tokenSeller, premiumPercent)
}

// UpdateSellerPremium is a paid mutator transaction binding the contract method 0xdcdf7e16.
//
// Solidity: function updateSellerPremium(address tokenSeller, uint8 premiumPercent) returns()
func (_Mcapscontroller *McapscontrollerTransactorSession) UpdateSellerPremium(tokenSeller common.Address, premiumPercent uint8) (*types.Transaction, error) {
	return _Mcapscontroller.Contract.UpdateSellerPremium(&_Mcapscontroller.TransactOpts, tokenSeller, premiumPercent)
}

// McapscontrollerCategoryAddedIterator is returned from FilterCategoryAdded and is used to iterate over the raw logs and unpacked data for CategoryAdded events raised by the Mcapscontroller contract.
type McapscontrollerCategoryAddedIterator struct {
	Event *McapscontrollerCategoryAdded // Event containing the contract specifics and raw log

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
func (it *McapscontrollerCategoryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McapscontrollerCategoryAdded)
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
		it.Event = new(McapscontrollerCategoryAdded)
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
func (it *McapscontrollerCategoryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McapscontrollerCategoryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McapscontrollerCategoryAdded represents a CategoryAdded event raised by the Mcapscontroller contract.
type McapscontrollerCategoryAdded struct {
	CategoryID   *big.Int
	MetadataHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCategoryAdded is a free log retrieval operation binding the contract event 0x087cd6515ab9d8ad2d847c4b093743cd73acf162a4260bd0ea2429ec6709a632.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash)
func (_Mcapscontroller *McapscontrollerFilterer) FilterCategoryAdded(opts *bind.FilterOpts) (*McapscontrollerCategoryAddedIterator, error) {

	logs, sub, err := _Mcapscontroller.contract.FilterLogs(opts, "CategoryAdded")
	if err != nil {
		return nil, err
	}
	return &McapscontrollerCategoryAddedIterator{contract: _Mcapscontroller.contract, event: "CategoryAdded", logs: logs, sub: sub}, nil
}

// WatchCategoryAdded is a free log subscription operation binding the contract event 0x087cd6515ab9d8ad2d847c4b093743cd73acf162a4260bd0ea2429ec6709a632.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash)
func (_Mcapscontroller *McapscontrollerFilterer) WatchCategoryAdded(opts *bind.WatchOpts, sink chan<- *McapscontrollerCategoryAdded) (event.Subscription, error) {

	logs, sub, err := _Mcapscontroller.contract.WatchLogs(opts, "CategoryAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McapscontrollerCategoryAdded)
				if err := _Mcapscontroller.contract.UnpackLog(event, "CategoryAdded", log); err != nil {
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

// ParseCategoryAdded is a log parse operation binding the contract event 0x087cd6515ab9d8ad2d847c4b093743cd73acf162a4260bd0ea2429ec6709a632.
//
// Solidity: event CategoryAdded(uint256 categoryID, bytes32 metadataHash)
func (_Mcapscontroller *McapscontrollerFilterer) ParseCategoryAdded(log types.Log) (*McapscontrollerCategoryAdded, error) {
	event := new(McapscontrollerCategoryAdded)
	if err := _Mcapscontroller.contract.UnpackLog(event, "CategoryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// McapscontrollerCategorySortedIterator is returned from FilterCategorySorted and is used to iterate over the raw logs and unpacked data for CategorySorted events raised by the Mcapscontroller contract.
type McapscontrollerCategorySortedIterator struct {
	Event *McapscontrollerCategorySorted // Event containing the contract specifics and raw log

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
func (it *McapscontrollerCategorySortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McapscontrollerCategorySorted)
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
		it.Event = new(McapscontrollerCategorySorted)
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
func (it *McapscontrollerCategorySortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McapscontrollerCategorySortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McapscontrollerCategorySorted represents a CategorySorted event raised by the Mcapscontroller contract.
type McapscontrollerCategorySorted struct {
	CategoryID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCategorySorted is a free log retrieval operation binding the contract event 0x60d2f5fd1812c906f738651cc42bfcb9c52908b88265a1ec508ce32367af72c3.
//
// Solidity: event CategorySorted(uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) FilterCategorySorted(opts *bind.FilterOpts) (*McapscontrollerCategorySortedIterator, error) {

	logs, sub, err := _Mcapscontroller.contract.FilterLogs(opts, "CategorySorted")
	if err != nil {
		return nil, err
	}
	return &McapscontrollerCategorySortedIterator{contract: _Mcapscontroller.contract, event: "CategorySorted", logs: logs, sub: sub}, nil
}

// WatchCategorySorted is a free log subscription operation binding the contract event 0x60d2f5fd1812c906f738651cc42bfcb9c52908b88265a1ec508ce32367af72c3.
//
// Solidity: event CategorySorted(uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) WatchCategorySorted(opts *bind.WatchOpts, sink chan<- *McapscontrollerCategorySorted) (event.Subscription, error) {

	logs, sub, err := _Mcapscontroller.contract.WatchLogs(opts, "CategorySorted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McapscontrollerCategorySorted)
				if err := _Mcapscontroller.contract.UnpackLog(event, "CategorySorted", log); err != nil {
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

// ParseCategorySorted is a log parse operation binding the contract event 0x60d2f5fd1812c906f738651cc42bfcb9c52908b88265a1ec508ce32367af72c3.
//
// Solidity: event CategorySorted(uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) ParseCategorySorted(log types.Log) (*McapscontrollerCategorySorted, error) {
	event := new(McapscontrollerCategorySorted)
	if err := _Mcapscontroller.contract.UnpackLog(event, "CategorySorted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// McapscontrollerNewPoolInitializerIterator is returned from FilterNewPoolInitializer and is used to iterate over the raw logs and unpacked data for NewPoolInitializer events raised by the Mcapscontroller contract.
type McapscontrollerNewPoolInitializerIterator struct {
	Event *McapscontrollerNewPoolInitializer // Event containing the contract specifics and raw log

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
func (it *McapscontrollerNewPoolInitializerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McapscontrollerNewPoolInitializer)
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
		it.Event = new(McapscontrollerNewPoolInitializer)
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
func (it *McapscontrollerNewPoolInitializerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McapscontrollerNewPoolInitializerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McapscontrollerNewPoolInitializer represents a NewPoolInitializer event raised by the Mcapscontroller contract.
type McapscontrollerNewPoolInitializer struct {
	Pool        common.Address
	Initializer common.Address
	CategoryID  *big.Int
	IndexSize   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPoolInitializer is a free log retrieval operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize)
func (_Mcapscontroller *McapscontrollerFilterer) FilterNewPoolInitializer(opts *bind.FilterOpts) (*McapscontrollerNewPoolInitializerIterator, error) {

	logs, sub, err := _Mcapscontroller.contract.FilterLogs(opts, "NewPoolInitializer")
	if err != nil {
		return nil, err
	}
	return &McapscontrollerNewPoolInitializerIterator{contract: _Mcapscontroller.contract, event: "NewPoolInitializer", logs: logs, sub: sub}, nil
}

// WatchNewPoolInitializer is a free log subscription operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize)
func (_Mcapscontroller *McapscontrollerFilterer) WatchNewPoolInitializer(opts *bind.WatchOpts, sink chan<- *McapscontrollerNewPoolInitializer) (event.Subscription, error) {

	logs, sub, err := _Mcapscontroller.contract.WatchLogs(opts, "NewPoolInitializer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McapscontrollerNewPoolInitializer)
				if err := _Mcapscontroller.contract.UnpackLog(event, "NewPoolInitializer", log); err != nil {
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

// ParseNewPoolInitializer is a log parse operation binding the contract event 0x7ad23833dba658b2bdc6f260fb60a3240b3868086ad60b4e42276f5eeba73e6a.
//
// Solidity: event NewPoolInitializer(address pool, address initializer, uint256 categoryID, uint256 indexSize)
func (_Mcapscontroller *McapscontrollerFilterer) ParseNewPoolInitializer(log types.Log) (*McapscontrollerNewPoolInitializer, error) {
	event := new(McapscontrollerNewPoolInitializer)
	if err := _Mcapscontroller.contract.UnpackLog(event, "NewPoolInitializer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// McapscontrollerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mcapscontroller contract.
type McapscontrollerOwnershipTransferredIterator struct {
	Event *McapscontrollerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *McapscontrollerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McapscontrollerOwnershipTransferred)
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
		it.Event = new(McapscontrollerOwnershipTransferred)
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
func (it *McapscontrollerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McapscontrollerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McapscontrollerOwnershipTransferred represents a OwnershipTransferred event raised by the Mcapscontroller contract.
type McapscontrollerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mcapscontroller *McapscontrollerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*McapscontrollerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mcapscontroller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &McapscontrollerOwnershipTransferredIterator{contract: _Mcapscontroller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mcapscontroller *McapscontrollerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *McapscontrollerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mcapscontroller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McapscontrollerOwnershipTransferred)
				if err := _Mcapscontroller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mcapscontroller *McapscontrollerFilterer) ParseOwnershipTransferred(log types.Log) (*McapscontrollerOwnershipTransferred, error) {
	event := new(McapscontrollerOwnershipTransferred)
	if err := _Mcapscontroller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// McapscontrollerPoolInitializedIterator is returned from FilterPoolInitialized and is used to iterate over the raw logs and unpacked data for PoolInitialized events raised by the Mcapscontroller contract.
type McapscontrollerPoolInitializedIterator struct {
	Event *McapscontrollerPoolInitialized // Event containing the contract specifics and raw log

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
func (it *McapscontrollerPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McapscontrollerPoolInitialized)
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
		it.Event = new(McapscontrollerPoolInitialized)
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
func (it *McapscontrollerPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McapscontrollerPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McapscontrollerPoolInitialized represents a PoolInitialized event raised by the Mcapscontroller contract.
type McapscontrollerPoolInitialized struct {
	Pool               common.Address
	UnboundTokenSeller common.Address
	CategoryID         *big.Int
	IndexSize          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolInitialized is a free log retrieval operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 categoryID, uint256 indexSize)
func (_Mcapscontroller *McapscontrollerFilterer) FilterPoolInitialized(opts *bind.FilterOpts) (*McapscontrollerPoolInitializedIterator, error) {

	logs, sub, err := _Mcapscontroller.contract.FilterLogs(opts, "PoolInitialized")
	if err != nil {
		return nil, err
	}
	return &McapscontrollerPoolInitializedIterator{contract: _Mcapscontroller.contract, event: "PoolInitialized", logs: logs, sub: sub}, nil
}

// WatchPoolInitialized is a free log subscription operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 categoryID, uint256 indexSize)
func (_Mcapscontroller *McapscontrollerFilterer) WatchPoolInitialized(opts *bind.WatchOpts, sink chan<- *McapscontrollerPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _Mcapscontroller.contract.WatchLogs(opts, "PoolInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McapscontrollerPoolInitialized)
				if err := _Mcapscontroller.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
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

// ParsePoolInitialized is a log parse operation binding the contract event 0x9ff2050ac7faae9dc192e1cc9abe73b18a9b849f9b43f509914d80fa104d5903.
//
// Solidity: event PoolInitialized(address pool, address unboundTokenSeller, uint256 categoryID, uint256 indexSize)
func (_Mcapscontroller *McapscontrollerFilterer) ParsePoolInitialized(log types.Log) (*McapscontrollerPoolInitialized, error) {
	event := new(McapscontrollerPoolInitialized)
	if err := _Mcapscontroller.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// McapscontrollerTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Mcapscontroller contract.
type McapscontrollerTokenAddedIterator struct {
	Event *McapscontrollerTokenAdded // Event containing the contract specifics and raw log

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
func (it *McapscontrollerTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McapscontrollerTokenAdded)
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
		it.Event = new(McapscontrollerTokenAdded)
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
func (it *McapscontrollerTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McapscontrollerTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McapscontrollerTokenAdded represents a TokenAdded event raised by the Mcapscontroller contract.
type McapscontrollerTokenAdded struct {
	Token      common.Address
	CategoryID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*McapscontrollerTokenAddedIterator, error) {

	logs, sub, err := _Mcapscontroller.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &McapscontrollerTokenAddedIterator{contract: _Mcapscontroller.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *McapscontrollerTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Mcapscontroller.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McapscontrollerTokenAdded)
				if err := _Mcapscontroller.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address token, uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) ParseTokenAdded(log types.Log) (*McapscontrollerTokenAdded, error) {
	event := new(McapscontrollerTokenAdded)
	if err := _Mcapscontroller.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// McapscontrollerTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Mcapscontroller contract.
type McapscontrollerTokenRemovedIterator struct {
	Event *McapscontrollerTokenRemoved // Event containing the contract specifics and raw log

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
func (it *McapscontrollerTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(McapscontrollerTokenRemoved)
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
		it.Event = new(McapscontrollerTokenRemoved)
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
func (it *McapscontrollerTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *McapscontrollerTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// McapscontrollerTokenRemoved represents a TokenRemoved event raised by the Mcapscontroller contract.
type McapscontrollerTokenRemoved struct {
	Token      common.Address
	CategoryID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*McapscontrollerTokenRemovedIterator, error) {

	logs, sub, err := _Mcapscontroller.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &McapscontrollerTokenRemovedIterator{contract: _Mcapscontroller.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *McapscontrollerTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _Mcapscontroller.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(McapscontrollerTokenRemoved)
				if err := _Mcapscontroller.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address token, uint256 categoryID)
func (_Mcapscontroller *McapscontrollerFilterer) ParseTokenRemoved(log types.Log) (*McapscontrollerTokenRemoved, error) {
	event := new(McapscontrollerTokenRemoved)
	if err := _Mcapscontroller.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
