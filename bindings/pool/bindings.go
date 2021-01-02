// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolbindings

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

// IPoolRecord is an auto generated low-level Go binding around an user-defined struct.
type IPoolRecord struct {
	Bound            bool
	Ready            bool
	LastDenormUpdate *big.Int
	Denorm           *big.Int
	DesiredDenorm    *big.Int
	Index            uint8
	Balance          *big.Int
}

// PoolbindingsABI is the input ABI used to generate the binding from.
const PoolbindingsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDenorm\",\"type\":\"uint256\"}],\"name\":\"LOG_DENORM_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"desiredDenorm\",\"type\":\"uint256\"}],\"name\":\"LOG_DESIRED_DENORM_SET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_EXIT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"name\":\"LOG_JOIN\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPoolTokens\",\"type\":\"uint256\"}],\"name\":\"LOG_MAX_TOKENS_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBalance\",\"type\":\"uint256\"}],\"name\":\"LOG_MINIMUM_BALANCE_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LOG_PUBLIC_SWAP_ENABLED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP_FEE_UPDATED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"desiredDenorm\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBalance\",\"type\":\"uint256\"}],\"name\":\"LOG_TOKEN_ADDED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LOG_TOKEN_READY\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LOG_TOKEN_REMOVED\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"configure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"exitPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolAmountIn\",\"type\":\"uint256\"}],\"name\":\"exitswapExternAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"exitswapPoolAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extrapolatePoolValueFromToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDesiredTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxPoolTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getMinimumBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"bound\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"lastDenormUpdate\",\"type\":\"uint40\"},{\"internalType\":\"uint96\",\"name\":\"denorm\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"desiredDenorm\",\"type\":\"uint96\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.Record\",\"name\":\"record\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUsedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"gulp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint96[]\",\"name\":\"denorms\",\"type\":\"uint96[]\"},{\"internalType\":\"address\",\"name\":\"tokenProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unbindHandler\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"isBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPublicSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"}],\"name\":\"joinPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPoolAmountOut\",\"type\":\"uint256\"}],\"name\":\"joinswapExternAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"joinswapPoolAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint96[]\",\"name\":\"desiredDenorms\",\"type\":\"uint96[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minimumBalances\",\"type\":\"uint256[]\"}],\"name\":\"reindexTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint96[]\",\"name\":\"desiredDenorms\",\"type\":\"uint96[]\"}],\"name\":\"reweighTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxPoolTokens\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Poolbindings is an auto generated Go binding around an Ethereum contract.
type Poolbindings struct {
	PoolbindingsCaller     // Read-only binding to the contract
	PoolbindingsTransactor // Write-only binding to the contract
	PoolbindingsFilterer   // Log filterer for contract events
}

// PoolbindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolbindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolbindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolbindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolbindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolbindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolbindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolbindingsSession struct {
	Contract     *Poolbindings     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolbindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolbindingsCallerSession struct {
	Contract *PoolbindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PoolbindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolbindingsTransactorSession struct {
	Contract     *PoolbindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PoolbindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolbindingsRaw struct {
	Contract *Poolbindings // Generic contract binding to access the raw methods on
}

// PoolbindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolbindingsCallerRaw struct {
	Contract *PoolbindingsCaller // Generic read-only contract binding to access the raw methods on
}

// PoolbindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolbindingsTransactorRaw struct {
	Contract *PoolbindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolbindings creates a new instance of Poolbindings, bound to a specific deployed contract.
func NewPoolbindings(address common.Address, backend bind.ContractBackend) (*Poolbindings, error) {
	contract, err := bindPoolbindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poolbindings{PoolbindingsCaller: PoolbindingsCaller{contract: contract}, PoolbindingsTransactor: PoolbindingsTransactor{contract: contract}, PoolbindingsFilterer: PoolbindingsFilterer{contract: contract}}, nil
}

// NewPoolbindingsCaller creates a new read-only instance of Poolbindings, bound to a specific deployed contract.
func NewPoolbindingsCaller(address common.Address, caller bind.ContractCaller) (*PoolbindingsCaller, error) {
	contract, err := bindPoolbindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsCaller{contract: contract}, nil
}

// NewPoolbindingsTransactor creates a new write-only instance of Poolbindings, bound to a specific deployed contract.
func NewPoolbindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolbindingsTransactor, error) {
	contract, err := bindPoolbindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsTransactor{contract: contract}, nil
}

// NewPoolbindingsFilterer creates a new log filterer instance of Poolbindings, bound to a specific deployed contract.
func NewPoolbindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolbindingsFilterer, error) {
	contract, err := bindPoolbindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsFilterer{contract: contract}, nil
}

// bindPoolbindings binds a generic wrapper to an already deployed contract.
func bindPoolbindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolbindingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolbindings *PoolbindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolbindings.Contract.PoolbindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolbindings *PoolbindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolbindings.Contract.PoolbindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolbindings *PoolbindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolbindings.Contract.PoolbindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolbindings *PoolbindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolbindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolbindings *PoolbindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolbindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolbindings *PoolbindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolbindings.Contract.contract.Transact(opts, method, params...)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) VERSIONNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "VERSION_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Poolbindings *PoolbindingsSession) VERSIONNUMBER() (*big.Int, error) {
	return _Poolbindings.Contract.VERSIONNUMBER(&_Poolbindings.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) VERSIONNUMBER() (*big.Int, error) {
	return _Poolbindings.Contract.VERSIONNUMBER(&_Poolbindings.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "allowance", src, dst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Poolbindings *PoolbindingsSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.Allowance(&_Poolbindings.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.Allowance(&_Poolbindings.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "balanceOf", whom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Poolbindings *PoolbindingsSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.BalanceOf(&_Poolbindings.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.BalanceOf(&_Poolbindings.CallOpts, whom)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Poolbindings *PoolbindingsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Poolbindings *PoolbindingsSession) Decimals() (uint8, error) {
	return _Poolbindings.Contract.Decimals(&_Poolbindings.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Poolbindings *PoolbindingsCallerSession) Decimals() (uint8, error) {
	return _Poolbindings.Contract.Decimals(&_Poolbindings.CallOpts)
}

// ExtrapolatePoolValueFromToken is a free data retrieval call binding the contract method 0x98836f08.
//
// Solidity: function extrapolatePoolValueFromToken() view returns(address, uint256)
func (_Poolbindings *PoolbindingsCaller) ExtrapolatePoolValueFromToken(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "extrapolatePoolValueFromToken")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ExtrapolatePoolValueFromToken is a free data retrieval call binding the contract method 0x98836f08.
//
// Solidity: function extrapolatePoolValueFromToken() view returns(address, uint256)
func (_Poolbindings *PoolbindingsSession) ExtrapolatePoolValueFromToken() (common.Address, *big.Int, error) {
	return _Poolbindings.Contract.ExtrapolatePoolValueFromToken(&_Poolbindings.CallOpts)
}

// ExtrapolatePoolValueFromToken is a free data retrieval call binding the contract method 0x98836f08.
//
// Solidity: function extrapolatePoolValueFromToken() view returns(address, uint256)
func (_Poolbindings *PoolbindingsCallerSession) ExtrapolatePoolValueFromToken() (common.Address, *big.Int, error) {
	return _Poolbindings.Contract.ExtrapolatePoolValueFromToken(&_Poolbindings.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetBalance(&_Poolbindings.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetBalance(&_Poolbindings.CallOpts, token)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Poolbindings *PoolbindingsCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Poolbindings *PoolbindingsSession) GetController() (common.Address, error) {
	return _Poolbindings.Contract.GetController(&_Poolbindings.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Poolbindings *PoolbindingsCallerSession) GetController() (common.Address, error) {
	return _Poolbindings.Contract.GetController(&_Poolbindings.CallOpts)
}

// GetCurrentDesiredTokens is a free data retrieval call binding the contract method 0x039209af.
//
// Solidity: function getCurrentDesiredTokens() view returns(address[] tokens)
func (_Poolbindings *PoolbindingsCaller) GetCurrentDesiredTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getCurrentDesiredTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentDesiredTokens is a free data retrieval call binding the contract method 0x039209af.
//
// Solidity: function getCurrentDesiredTokens() view returns(address[] tokens)
func (_Poolbindings *PoolbindingsSession) GetCurrentDesiredTokens() ([]common.Address, error) {
	return _Poolbindings.Contract.GetCurrentDesiredTokens(&_Poolbindings.CallOpts)
}

// GetCurrentDesiredTokens is a free data retrieval call binding the contract method 0x039209af.
//
// Solidity: function getCurrentDesiredTokens() view returns(address[] tokens)
func (_Poolbindings *PoolbindingsCallerSession) GetCurrentDesiredTokens() ([]common.Address, error) {
	return _Poolbindings.Contract.GetCurrentDesiredTokens(&_Poolbindings.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Poolbindings *PoolbindingsCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getCurrentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Poolbindings *PoolbindingsSession) GetCurrentTokens() ([]common.Address, error) {
	return _Poolbindings.Contract.GetCurrentTokens(&_Poolbindings.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Poolbindings *PoolbindingsCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _Poolbindings.Contract.GetCurrentTokens(&_Poolbindings.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getDenormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetDenormalizedWeight(&_Poolbindings.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetDenormalizedWeight(&_Poolbindings.CallOpts, token)
}

// GetMaxPoolTokens is a free data retrieval call binding the contract method 0x46ecfbd6.
//
// Solidity: function getMaxPoolTokens() view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetMaxPoolTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getMaxPoolTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPoolTokens is a free data retrieval call binding the contract method 0x46ecfbd6.
//
// Solidity: function getMaxPoolTokens() view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetMaxPoolTokens() (*big.Int, error) {
	return _Poolbindings.Contract.GetMaxPoolTokens(&_Poolbindings.CallOpts)
}

// GetMaxPoolTokens is a free data retrieval call binding the contract method 0x46ecfbd6.
//
// Solidity: function getMaxPoolTokens() view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetMaxPoolTokens() (*big.Int, error) {
	return _Poolbindings.Contract.GetMaxPoolTokens(&_Poolbindings.CallOpts)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0x91bfa2bf.
//
// Solidity: function getMinimumBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetMinimumBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getMinimumBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBalance is a free data retrieval call binding the contract method 0x91bfa2bf.
//
// Solidity: function getMinimumBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetMinimumBalance(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetMinimumBalance(&_Poolbindings.CallOpts, token)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0x91bfa2bf.
//
// Solidity: function getMinimumBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetMinimumBalance(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetMinimumBalance(&_Poolbindings.CallOpts, token)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetNumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getNumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetNumTokens() (*big.Int, error) {
	return _Poolbindings.Contract.GetNumTokens(&_Poolbindings.CallOpts)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetNumTokens() (*big.Int, error) {
	return _Poolbindings.Contract.GetNumTokens(&_Poolbindings.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getSpotPrice", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetSpotPrice(&_Poolbindings.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetSpotPrice(&_Poolbindings.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getSwapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetSwapFee() (*big.Int, error) {
	return _Poolbindings.Contract.GetSwapFee(&_Poolbindings.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetSwapFee() (*big.Int, error) {
	return _Poolbindings.Contract.GetSwapFee(&_Poolbindings.CallOpts)
}

// GetTokenRecord is a free data retrieval call binding the contract method 0x64c7d661.
//
// Solidity: function getTokenRecord(address token) view returns((bool,bool,uint40,uint96,uint96,uint8,uint256) record)
func (_Poolbindings *PoolbindingsCaller) GetTokenRecord(opts *bind.CallOpts, token common.Address) (IPoolRecord, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getTokenRecord", token)

	if err != nil {
		return *new(IPoolRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolRecord)).(*IPoolRecord)

	return out0, err

}

// GetTokenRecord is a free data retrieval call binding the contract method 0x64c7d661.
//
// Solidity: function getTokenRecord(address token) view returns((bool,bool,uint40,uint96,uint96,uint8,uint256) record)
func (_Poolbindings *PoolbindingsSession) GetTokenRecord(token common.Address) (IPoolRecord, error) {
	return _Poolbindings.Contract.GetTokenRecord(&_Poolbindings.CallOpts, token)
}

// GetTokenRecord is a free data retrieval call binding the contract method 0x64c7d661.
//
// Solidity: function getTokenRecord(address token) view returns((bool,bool,uint40,uint96,uint96,uint8,uint256) record)
func (_Poolbindings *PoolbindingsCallerSession) GetTokenRecord(token common.Address) (IPoolRecord, error) {
	return _Poolbindings.Contract.GetTokenRecord(&_Poolbindings.CallOpts, token)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getTotalDenormalizedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Poolbindings.Contract.GetTotalDenormalizedWeight(&_Poolbindings.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Poolbindings.Contract.GetTotalDenormalizedWeight(&_Poolbindings.CallOpts)
}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) GetUsedBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "getUsedBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsSession) GetUsedBalance(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetUsedBalance(&_Poolbindings.CallOpts, token)
}

// GetUsedBalance is a free data retrieval call binding the contract method 0x4aa4e0b5.
//
// Solidity: function getUsedBalance(address token) view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) GetUsedBalance(token common.Address) (*big.Int, error) {
	return _Poolbindings.Contract.GetUsedBalance(&_Poolbindings.CallOpts, token)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Poolbindings *PoolbindingsCaller) IsBound(opts *bind.CallOpts, t common.Address) (bool, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "isBound", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Poolbindings *PoolbindingsSession) IsBound(t common.Address) (bool, error) {
	return _Poolbindings.Contract.IsBound(&_Poolbindings.CallOpts, t)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Poolbindings *PoolbindingsCallerSession) IsBound(t common.Address) (bool, error) {
	return _Poolbindings.Contract.IsBound(&_Poolbindings.CallOpts, t)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Poolbindings *PoolbindingsCaller) IsPublicSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "isPublicSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Poolbindings *PoolbindingsSession) IsPublicSwap() (bool, error) {
	return _Poolbindings.Contract.IsPublicSwap(&_Poolbindings.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Poolbindings *PoolbindingsCallerSession) IsPublicSwap() (bool, error) {
	return _Poolbindings.Contract.IsPublicSwap(&_Poolbindings.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Poolbindings *PoolbindingsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Poolbindings *PoolbindingsSession) Name() (string, error) {
	return _Poolbindings.Contract.Name(&_Poolbindings.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Poolbindings *PoolbindingsCallerSession) Name() (string, error) {
	return _Poolbindings.Contract.Name(&_Poolbindings.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Poolbindings *PoolbindingsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Poolbindings *PoolbindingsSession) Symbol() (string, error) {
	return _Poolbindings.Contract.Symbol(&_Poolbindings.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Poolbindings *PoolbindingsCallerSession) Symbol() (string, error) {
	return _Poolbindings.Contract.Symbol(&_Poolbindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Poolbindings *PoolbindingsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolbindings.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Poolbindings *PoolbindingsSession) TotalSupply() (*big.Int, error) {
	return _Poolbindings.Contract.TotalSupply(&_Poolbindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Poolbindings *PoolbindingsCallerSession) TotalSupply() (*big.Int, error) {
	return _Poolbindings.Contract.TotalSupply(&_Poolbindings.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.Approve(&_Poolbindings.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.Approve(&_Poolbindings.TransactOpts, dst, amt)
}

// Configure is a paid mutator transaction binding the contract method 0x19f0f849.
//
// Solidity: function configure(address controller, string name, string symbol) returns()
func (_Poolbindings *PoolbindingsTransactor) Configure(opts *bind.TransactOpts, controller common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "configure", controller, name, symbol)
}

// Configure is a paid mutator transaction binding the contract method 0x19f0f849.
//
// Solidity: function configure(address controller, string name, string symbol) returns()
func (_Poolbindings *PoolbindingsSession) Configure(controller common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Poolbindings.Contract.Configure(&_Poolbindings.TransactOpts, controller, name, symbol)
}

// Configure is a paid mutator transaction binding the contract method 0x19f0f849.
//
// Solidity: function configure(address controller, string name, string symbol) returns()
func (_Poolbindings *PoolbindingsTransactorSession) Configure(controller common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Poolbindings.Contract.Configure(&_Poolbindings.TransactOpts, controller, name, symbol)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.DecreaseApproval(&_Poolbindings.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.DecreaseApproval(&_Poolbindings.TransactOpts, dst, amt)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Poolbindings *PoolbindingsTransactor) ExitPool(opts *bind.TransactOpts, poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "exitPool", poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Poolbindings *PoolbindingsSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ExitPool(&_Poolbindings.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Poolbindings *PoolbindingsTransactorSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ExitPool(&_Poolbindings.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256)
func (_Poolbindings *PoolbindingsTransactor) ExitswapExternAmountOut(opts *bind.TransactOpts, tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "exitswapExternAmountOut", tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256)
func (_Poolbindings *PoolbindingsSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ExitswapExternAmountOut(&_Poolbindings.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256)
func (_Poolbindings *PoolbindingsTransactorSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ExitswapExternAmountOut(&_Poolbindings.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256)
func (_Poolbindings *PoolbindingsTransactor) ExitswapPoolAmountIn(opts *bind.TransactOpts, tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "exitswapPoolAmountIn", tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256)
func (_Poolbindings *PoolbindingsSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ExitswapPoolAmountIn(&_Poolbindings.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256)
func (_Poolbindings *PoolbindingsTransactorSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ExitswapPoolAmountIn(&_Poolbindings.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// FlashBorrow is a paid mutator transaction binding the contract method 0x3043ffc9.
//
// Solidity: function flashBorrow(address recipient, address token, uint256 amount, bytes data) returns()
func (_Poolbindings *PoolbindingsTransactor) FlashBorrow(opts *bind.TransactOpts, recipient common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "flashBorrow", recipient, token, amount, data)
}

// FlashBorrow is a paid mutator transaction binding the contract method 0x3043ffc9.
//
// Solidity: function flashBorrow(address recipient, address token, uint256 amount, bytes data) returns()
func (_Poolbindings *PoolbindingsSession) FlashBorrow(recipient common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Poolbindings.Contract.FlashBorrow(&_Poolbindings.TransactOpts, recipient, token, amount, data)
}

// FlashBorrow is a paid mutator transaction binding the contract method 0x3043ffc9.
//
// Solidity: function flashBorrow(address recipient, address token, uint256 amount, bytes data) returns()
func (_Poolbindings *PoolbindingsTransactorSession) FlashBorrow(recipient common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Poolbindings.Contract.FlashBorrow(&_Poolbindings.TransactOpts, recipient, token, amount, data)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Poolbindings *PoolbindingsTransactor) Gulp(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "gulp", token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Poolbindings *PoolbindingsSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Poolbindings.Contract.Gulp(&_Poolbindings.TransactOpts, token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Poolbindings *PoolbindingsTransactorSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Poolbindings.Contract.Gulp(&_Poolbindings.TransactOpts, token)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.IncreaseApproval(&_Poolbindings.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.IncreaseApproval(&_Poolbindings.TransactOpts, dst, amt)
}

// Initialize is a paid mutator transaction binding the contract method 0xa20cfa75.
//
// Solidity: function initialize(address[] tokens, uint256[] balances, uint96[] denorms, address tokenProvider, address unbindHandler) returns()
func (_Poolbindings *PoolbindingsTransactor) Initialize(opts *bind.TransactOpts, tokens []common.Address, balances []*big.Int, denorms []*big.Int, tokenProvider common.Address, unbindHandler common.Address) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "initialize", tokens, balances, denorms, tokenProvider, unbindHandler)
}

// Initialize is a paid mutator transaction binding the contract method 0xa20cfa75.
//
// Solidity: function initialize(address[] tokens, uint256[] balances, uint96[] denorms, address tokenProvider, address unbindHandler) returns()
func (_Poolbindings *PoolbindingsSession) Initialize(tokens []common.Address, balances []*big.Int, denorms []*big.Int, tokenProvider common.Address, unbindHandler common.Address) (*types.Transaction, error) {
	return _Poolbindings.Contract.Initialize(&_Poolbindings.TransactOpts, tokens, balances, denorms, tokenProvider, unbindHandler)
}

// Initialize is a paid mutator transaction binding the contract method 0xa20cfa75.
//
// Solidity: function initialize(address[] tokens, uint256[] balances, uint96[] denorms, address tokenProvider, address unbindHandler) returns()
func (_Poolbindings *PoolbindingsTransactorSession) Initialize(tokens []common.Address, balances []*big.Int, denorms []*big.Int, tokenProvider common.Address, unbindHandler common.Address) (*types.Transaction, error) {
	return _Poolbindings.Contract.Initialize(&_Poolbindings.TransactOpts, tokens, balances, denorms, tokenProvider, unbindHandler)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Poolbindings *PoolbindingsTransactor) JoinPool(opts *bind.TransactOpts, poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "joinPool", poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Poolbindings *PoolbindingsSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.JoinPool(&_Poolbindings.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Poolbindings *PoolbindingsTransactorSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.JoinPool(&_Poolbindings.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256)
func (_Poolbindings *PoolbindingsTransactor) JoinswapExternAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "joinswapExternAmountIn", tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256)
func (_Poolbindings *PoolbindingsSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.JoinswapExternAmountIn(&_Poolbindings.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256)
func (_Poolbindings *PoolbindingsTransactorSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.JoinswapExternAmountIn(&_Poolbindings.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256)
func (_Poolbindings *PoolbindingsTransactor) JoinswapPoolAmountOut(opts *bind.TransactOpts, tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "joinswapPoolAmountOut", tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256)
func (_Poolbindings *PoolbindingsSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.JoinswapPoolAmountOut(&_Poolbindings.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256)
func (_Poolbindings *PoolbindingsTransactorSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.JoinswapPoolAmountOut(&_Poolbindings.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// ReindexTokens is a paid mutator transaction binding the contract method 0xc3f46810.
//
// Solidity: function reindexTokens(address[] tokens, uint96[] desiredDenorms, uint256[] minimumBalances) returns()
func (_Poolbindings *PoolbindingsTransactor) ReindexTokens(opts *bind.TransactOpts, tokens []common.Address, desiredDenorms []*big.Int, minimumBalances []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "reindexTokens", tokens, desiredDenorms, minimumBalances)
}

// ReindexTokens is a paid mutator transaction binding the contract method 0xc3f46810.
//
// Solidity: function reindexTokens(address[] tokens, uint96[] desiredDenorms, uint256[] minimumBalances) returns()
func (_Poolbindings *PoolbindingsSession) ReindexTokens(tokens []common.Address, desiredDenorms []*big.Int, minimumBalances []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ReindexTokens(&_Poolbindings.TransactOpts, tokens, desiredDenorms, minimumBalances)
}

// ReindexTokens is a paid mutator transaction binding the contract method 0xc3f46810.
//
// Solidity: function reindexTokens(address[] tokens, uint96[] desiredDenorms, uint256[] minimumBalances) returns()
func (_Poolbindings *PoolbindingsTransactorSession) ReindexTokens(tokens []common.Address, desiredDenorms []*big.Int, minimumBalances []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ReindexTokens(&_Poolbindings.TransactOpts, tokens, desiredDenorms, minimumBalances)
}

// ReweighTokens is a paid mutator transaction binding the contract method 0x5d5e8ce7.
//
// Solidity: function reweighTokens(address[] tokens, uint96[] desiredDenorms) returns()
func (_Poolbindings *PoolbindingsTransactor) ReweighTokens(opts *bind.TransactOpts, tokens []common.Address, desiredDenorms []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "reweighTokens", tokens, desiredDenorms)
}

// ReweighTokens is a paid mutator transaction binding the contract method 0x5d5e8ce7.
//
// Solidity: function reweighTokens(address[] tokens, uint96[] desiredDenorms) returns()
func (_Poolbindings *PoolbindingsSession) ReweighTokens(tokens []common.Address, desiredDenorms []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ReweighTokens(&_Poolbindings.TransactOpts, tokens, desiredDenorms)
}

// ReweighTokens is a paid mutator transaction binding the contract method 0x5d5e8ce7.
//
// Solidity: function reweighTokens(address[] tokens, uint96[] desiredDenorms) returns()
func (_Poolbindings *PoolbindingsTransactorSession) ReweighTokens(tokens []common.Address, desiredDenorms []*big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.ReweighTokens(&_Poolbindings.TransactOpts, tokens, desiredDenorms)
}

// SetMaxPoolTokens is a paid mutator transaction binding the contract method 0x534d4f3d.
//
// Solidity: function setMaxPoolTokens(uint256 maxPoolTokens) returns()
func (_Poolbindings *PoolbindingsTransactor) SetMaxPoolTokens(opts *bind.TransactOpts, maxPoolTokens *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "setMaxPoolTokens", maxPoolTokens)
}

// SetMaxPoolTokens is a paid mutator transaction binding the contract method 0x534d4f3d.
//
// Solidity: function setMaxPoolTokens(uint256 maxPoolTokens) returns()
func (_Poolbindings *PoolbindingsSession) SetMaxPoolTokens(maxPoolTokens *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SetMaxPoolTokens(&_Poolbindings.TransactOpts, maxPoolTokens)
}

// SetMaxPoolTokens is a paid mutator transaction binding the contract method 0x534d4f3d.
//
// Solidity: function setMaxPoolTokens(uint256 maxPoolTokens) returns()
func (_Poolbindings *PoolbindingsTransactorSession) SetMaxPoolTokens(maxPoolTokens *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SetMaxPoolTokens(&_Poolbindings.TransactOpts, maxPoolTokens)
}

// SetMinimumBalance is a paid mutator transaction binding the contract method 0xa49c44d7.
//
// Solidity: function setMinimumBalance(address token, uint256 minimumBalance) returns()
func (_Poolbindings *PoolbindingsTransactor) SetMinimumBalance(opts *bind.TransactOpts, token common.Address, minimumBalance *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "setMinimumBalance", token, minimumBalance)
}

// SetMinimumBalance is a paid mutator transaction binding the contract method 0xa49c44d7.
//
// Solidity: function setMinimumBalance(address token, uint256 minimumBalance) returns()
func (_Poolbindings *PoolbindingsSession) SetMinimumBalance(token common.Address, minimumBalance *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SetMinimumBalance(&_Poolbindings.TransactOpts, token, minimumBalance)
}

// SetMinimumBalance is a paid mutator transaction binding the contract method 0xa49c44d7.
//
// Solidity: function setMinimumBalance(address token, uint256 minimumBalance) returns()
func (_Poolbindings *PoolbindingsTransactorSession) SetMinimumBalance(token common.Address, minimumBalance *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SetMinimumBalance(&_Poolbindings.TransactOpts, token, minimumBalance)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Poolbindings *PoolbindingsTransactor) SetSwapFee(opts *bind.TransactOpts, swapFee *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "setSwapFee", swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Poolbindings *PoolbindingsSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SetSwapFee(&_Poolbindings.TransactOpts, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Poolbindings *PoolbindingsTransactorSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SetSwapFee(&_Poolbindings.TransactOpts, swapFee)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Poolbindings *PoolbindingsTransactor) SwapExactAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "swapExactAmountIn", tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Poolbindings *PoolbindingsSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SwapExactAmountIn(&_Poolbindings.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Poolbindings *PoolbindingsTransactorSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SwapExactAmountIn(&_Poolbindings.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Poolbindings *PoolbindingsTransactor) SwapExactAmountOut(opts *bind.TransactOpts, tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "swapExactAmountOut", tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Poolbindings *PoolbindingsSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SwapExactAmountOut(&_Poolbindings.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256, uint256)
func (_Poolbindings *PoolbindingsTransactorSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.SwapExactAmountOut(&_Poolbindings.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.Transfer(&_Poolbindings.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.Transfer(&_Poolbindings.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.TransferFrom(&_Poolbindings.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Poolbindings *PoolbindingsTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Poolbindings.Contract.TransferFrom(&_Poolbindings.TransactOpts, src, dst, amt)
}

// PoolbindingsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Poolbindings contract.
type PoolbindingsApprovalIterator struct {
	Event *PoolbindingsApproval // Event containing the contract specifics and raw log

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
func (it *PoolbindingsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsApproval)
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
		it.Event = new(PoolbindingsApproval)
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
func (it *PoolbindingsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsApproval represents a Approval event raised by the Poolbindings contract.
type PoolbindingsApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Poolbindings *PoolbindingsFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*PoolbindingsApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsApprovalIterator{contract: _Poolbindings.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Poolbindings *PoolbindingsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolbindingsApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsApproval)
				if err := _Poolbindings.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Poolbindings *PoolbindingsFilterer) ParseApproval(log types.Log) (*PoolbindingsApproval, error) {
	event := new(PoolbindingsApproval)
	if err := _Poolbindings.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGDENORMUPDATEDIterator is returned from FilterLOGDENORMUPDATED and is used to iterate over the raw logs and unpacked data for LOGDENORMUPDATED events raised by the Poolbindings contract.
type PoolbindingsLOGDENORMUPDATEDIterator struct {
	Event *PoolbindingsLOGDENORMUPDATED // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGDENORMUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGDENORMUPDATED)
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
		it.Event = new(PoolbindingsLOGDENORMUPDATED)
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
func (it *PoolbindingsLOGDENORMUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGDENORMUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGDENORMUPDATED represents a LOGDENORMUPDATED event raised by the Poolbindings contract.
type PoolbindingsLOGDENORMUPDATED struct {
	Token     common.Address
	NewDenorm *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLOGDENORMUPDATED is a free log retrieval operation binding the contract event 0x21b12aed5d425f5675450ffeeae01039085e5323974c3099e1828155d9b51e77.
//
// Solidity: event LOG_DENORM_UPDATED(address indexed token, uint256 newDenorm)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGDENORMUPDATED(opts *bind.FilterOpts, token []common.Address) (*PoolbindingsLOGDENORMUPDATEDIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_DENORM_UPDATED", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGDENORMUPDATEDIterator{contract: _Poolbindings.contract, event: "LOG_DENORM_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGDENORMUPDATED is a free log subscription operation binding the contract event 0x21b12aed5d425f5675450ffeeae01039085e5323974c3099e1828155d9b51e77.
//
// Solidity: event LOG_DENORM_UPDATED(address indexed token, uint256 newDenorm)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGDENORMUPDATED(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGDENORMUPDATED, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_DENORM_UPDATED", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGDENORMUPDATED)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_DENORM_UPDATED", log); err != nil {
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

// ParseLOGDENORMUPDATED is a log parse operation binding the contract event 0x21b12aed5d425f5675450ffeeae01039085e5323974c3099e1828155d9b51e77.
//
// Solidity: event LOG_DENORM_UPDATED(address indexed token, uint256 newDenorm)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGDENORMUPDATED(log types.Log) (*PoolbindingsLOGDENORMUPDATED, error) {
	event := new(PoolbindingsLOGDENORMUPDATED)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_DENORM_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGDESIREDDENORMSETIterator is returned from FilterLOGDESIREDDENORMSET and is used to iterate over the raw logs and unpacked data for LOGDESIREDDENORMSET events raised by the Poolbindings contract.
type PoolbindingsLOGDESIREDDENORMSETIterator struct {
	Event *PoolbindingsLOGDESIREDDENORMSET // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGDESIREDDENORMSETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGDESIREDDENORMSET)
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
		it.Event = new(PoolbindingsLOGDESIREDDENORMSET)
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
func (it *PoolbindingsLOGDESIREDDENORMSETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGDESIREDDENORMSETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGDESIREDDENORMSET represents a LOGDESIREDDENORMSET event raised by the Poolbindings contract.
type PoolbindingsLOGDESIREDDENORMSET struct {
	Token         common.Address
	DesiredDenorm *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGDESIREDDENORMSET is a free log retrieval operation binding the contract event 0xc7ea88f3376e27ce6ebc2025310023327f743a8377d438258c36b166dd8b2983.
//
// Solidity: event LOG_DESIRED_DENORM_SET(address indexed token, uint256 desiredDenorm)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGDESIREDDENORMSET(opts *bind.FilterOpts, token []common.Address) (*PoolbindingsLOGDESIREDDENORMSETIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_DESIRED_DENORM_SET", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGDESIREDDENORMSETIterator{contract: _Poolbindings.contract, event: "LOG_DESIRED_DENORM_SET", logs: logs, sub: sub}, nil
}

// WatchLOGDESIREDDENORMSET is a free log subscription operation binding the contract event 0xc7ea88f3376e27ce6ebc2025310023327f743a8377d438258c36b166dd8b2983.
//
// Solidity: event LOG_DESIRED_DENORM_SET(address indexed token, uint256 desiredDenorm)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGDESIREDDENORMSET(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGDESIREDDENORMSET, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_DESIRED_DENORM_SET", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGDESIREDDENORMSET)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_DESIRED_DENORM_SET", log); err != nil {
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

// ParseLOGDESIREDDENORMSET is a log parse operation binding the contract event 0xc7ea88f3376e27ce6ebc2025310023327f743a8377d438258c36b166dd8b2983.
//
// Solidity: event LOG_DESIRED_DENORM_SET(address indexed token, uint256 desiredDenorm)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGDESIREDDENORMSET(log types.Log) (*PoolbindingsLOGDESIREDDENORMSET, error) {
	event := new(PoolbindingsLOGDESIREDDENORMSET)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_DESIRED_DENORM_SET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGEXITIterator is returned from FilterLOGEXIT and is used to iterate over the raw logs and unpacked data for LOGEXIT events raised by the Poolbindings contract.
type PoolbindingsLOGEXITIterator struct {
	Event *PoolbindingsLOGEXIT // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGEXITIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGEXIT)
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
		it.Event = new(PoolbindingsLOGEXIT)
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
func (it *PoolbindingsLOGEXITIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGEXITIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGEXIT represents a LOGEXIT event raised by the Poolbindings contract.
type PoolbindingsLOGEXIT struct {
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGEXIT is a free log retrieval operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGEXIT(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address) (*PoolbindingsLOGEXITIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGEXITIterator{contract: _Poolbindings.contract, event: "LOG_EXIT", logs: logs, sub: sub}, nil
}

// WatchLOGEXIT is a free log subscription operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGEXIT(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGEXIT, caller []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGEXIT)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
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

// ParseLOGEXIT is a log parse operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGEXIT(log types.Log) (*PoolbindingsLOGEXIT, error) {
	event := new(PoolbindingsLOGEXIT)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGJOINIterator is returned from FilterLOGJOIN and is used to iterate over the raw logs and unpacked data for LOGJOIN events raised by the Poolbindings contract.
type PoolbindingsLOGJOINIterator struct {
	Event *PoolbindingsLOGJOIN // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGJOINIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGJOIN)
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
		it.Event = new(PoolbindingsLOGJOIN)
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
func (it *PoolbindingsLOGJOINIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGJOINIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGJOIN represents a LOGJOIN event raised by the Poolbindings contract.
type PoolbindingsLOGJOIN struct {
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGJOIN is a free log retrieval operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGJOIN(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address) (*PoolbindingsLOGJOINIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGJOINIterator{contract: _Poolbindings.contract, event: "LOG_JOIN", logs: logs, sub: sub}, nil
}

// WatchLOGJOIN is a free log subscription operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGJOIN(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGJOIN, caller []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGJOIN)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
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

// ParseLOGJOIN is a log parse operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGJOIN(log types.Log) (*PoolbindingsLOGJOIN, error) {
	event := new(PoolbindingsLOGJOIN)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGMAXTOKENSUPDATEDIterator is returned from FilterLOGMAXTOKENSUPDATED and is used to iterate over the raw logs and unpacked data for LOGMAXTOKENSUPDATED events raised by the Poolbindings contract.
type PoolbindingsLOGMAXTOKENSUPDATEDIterator struct {
	Event *PoolbindingsLOGMAXTOKENSUPDATED // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGMAXTOKENSUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGMAXTOKENSUPDATED)
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
		it.Event = new(PoolbindingsLOGMAXTOKENSUPDATED)
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
func (it *PoolbindingsLOGMAXTOKENSUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGMAXTOKENSUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGMAXTOKENSUPDATED represents a LOGMAXTOKENSUPDATED event raised by the Poolbindings contract.
type PoolbindingsLOGMAXTOKENSUPDATED struct {
	MaxPoolTokens *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGMAXTOKENSUPDATED is a free log retrieval operation binding the contract event 0x65492266ae9a1f46497efddd24b6946862569680a511c543590d4587ca800d05.
//
// Solidity: event LOG_MAX_TOKENS_UPDATED(uint256 maxPoolTokens)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGMAXTOKENSUPDATED(opts *bind.FilterOpts) (*PoolbindingsLOGMAXTOKENSUPDATEDIterator, error) {

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_MAX_TOKENS_UPDATED")
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGMAXTOKENSUPDATEDIterator{contract: _Poolbindings.contract, event: "LOG_MAX_TOKENS_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGMAXTOKENSUPDATED is a free log subscription operation binding the contract event 0x65492266ae9a1f46497efddd24b6946862569680a511c543590d4587ca800d05.
//
// Solidity: event LOG_MAX_TOKENS_UPDATED(uint256 maxPoolTokens)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGMAXTOKENSUPDATED(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGMAXTOKENSUPDATED) (event.Subscription, error) {

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_MAX_TOKENS_UPDATED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGMAXTOKENSUPDATED)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_MAX_TOKENS_UPDATED", log); err != nil {
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

// ParseLOGMAXTOKENSUPDATED is a log parse operation binding the contract event 0x65492266ae9a1f46497efddd24b6946862569680a511c543590d4587ca800d05.
//
// Solidity: event LOG_MAX_TOKENS_UPDATED(uint256 maxPoolTokens)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGMAXTOKENSUPDATED(log types.Log) (*PoolbindingsLOGMAXTOKENSUPDATED, error) {
	event := new(PoolbindingsLOGMAXTOKENSUPDATED)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_MAX_TOKENS_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGMINIMUMBALANCEUPDATEDIterator is returned from FilterLOGMINIMUMBALANCEUPDATED and is used to iterate over the raw logs and unpacked data for LOGMINIMUMBALANCEUPDATED events raised by the Poolbindings contract.
type PoolbindingsLOGMINIMUMBALANCEUPDATEDIterator struct {
	Event *PoolbindingsLOGMINIMUMBALANCEUPDATED // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGMINIMUMBALANCEUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGMINIMUMBALANCEUPDATED)
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
		it.Event = new(PoolbindingsLOGMINIMUMBALANCEUPDATED)
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
func (it *PoolbindingsLOGMINIMUMBALANCEUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGMINIMUMBALANCEUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGMINIMUMBALANCEUPDATED represents a LOGMINIMUMBALANCEUPDATED event raised by the Poolbindings contract.
type PoolbindingsLOGMINIMUMBALANCEUPDATED struct {
	Token          common.Address
	MinimumBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGMINIMUMBALANCEUPDATED is a free log retrieval operation binding the contract event 0x000c7a55677231b335e6dea005fa240ac2aeafbd62f188372a7d66892b722c52.
//
// Solidity: event LOG_MINIMUM_BALANCE_UPDATED(address token, uint256 minimumBalance)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGMINIMUMBALANCEUPDATED(opts *bind.FilterOpts) (*PoolbindingsLOGMINIMUMBALANCEUPDATEDIterator, error) {

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_MINIMUM_BALANCE_UPDATED")
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGMINIMUMBALANCEUPDATEDIterator{contract: _Poolbindings.contract, event: "LOG_MINIMUM_BALANCE_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGMINIMUMBALANCEUPDATED is a free log subscription operation binding the contract event 0x000c7a55677231b335e6dea005fa240ac2aeafbd62f188372a7d66892b722c52.
//
// Solidity: event LOG_MINIMUM_BALANCE_UPDATED(address token, uint256 minimumBalance)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGMINIMUMBALANCEUPDATED(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGMINIMUMBALANCEUPDATED) (event.Subscription, error) {

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_MINIMUM_BALANCE_UPDATED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGMINIMUMBALANCEUPDATED)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_MINIMUM_BALANCE_UPDATED", log); err != nil {
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

// ParseLOGMINIMUMBALANCEUPDATED is a log parse operation binding the contract event 0x000c7a55677231b335e6dea005fa240ac2aeafbd62f188372a7d66892b722c52.
//
// Solidity: event LOG_MINIMUM_BALANCE_UPDATED(address token, uint256 minimumBalance)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGMINIMUMBALANCEUPDATED(log types.Log) (*PoolbindingsLOGMINIMUMBALANCEUPDATED, error) {
	event := new(PoolbindingsLOGMINIMUMBALANCEUPDATED)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_MINIMUM_BALANCE_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGPUBLICSWAPENABLEDIterator is returned from FilterLOGPUBLICSWAPENABLED and is used to iterate over the raw logs and unpacked data for LOGPUBLICSWAPENABLED events raised by the Poolbindings contract.
type PoolbindingsLOGPUBLICSWAPENABLEDIterator struct {
	Event *PoolbindingsLOGPUBLICSWAPENABLED // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGPUBLICSWAPENABLEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGPUBLICSWAPENABLED)
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
		it.Event = new(PoolbindingsLOGPUBLICSWAPENABLED)
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
func (it *PoolbindingsLOGPUBLICSWAPENABLEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGPUBLICSWAPENABLEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGPUBLICSWAPENABLED represents a LOGPUBLICSWAPENABLED event raised by the Poolbindings contract.
type PoolbindingsLOGPUBLICSWAPENABLED struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLOGPUBLICSWAPENABLED is a free log retrieval operation binding the contract event 0x183bed17d33ee5be9a47bb997aee6a152d84a34309cdb66f76301de0a40a1389.
//
// Solidity: event LOG_PUBLIC_SWAP_ENABLED()
func (_Poolbindings *PoolbindingsFilterer) FilterLOGPUBLICSWAPENABLED(opts *bind.FilterOpts) (*PoolbindingsLOGPUBLICSWAPENABLEDIterator, error) {

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_PUBLIC_SWAP_ENABLED")
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGPUBLICSWAPENABLEDIterator{contract: _Poolbindings.contract, event: "LOG_PUBLIC_SWAP_ENABLED", logs: logs, sub: sub}, nil
}

// WatchLOGPUBLICSWAPENABLED is a free log subscription operation binding the contract event 0x183bed17d33ee5be9a47bb997aee6a152d84a34309cdb66f76301de0a40a1389.
//
// Solidity: event LOG_PUBLIC_SWAP_ENABLED()
func (_Poolbindings *PoolbindingsFilterer) WatchLOGPUBLICSWAPENABLED(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGPUBLICSWAPENABLED) (event.Subscription, error) {

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_PUBLIC_SWAP_ENABLED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGPUBLICSWAPENABLED)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_PUBLIC_SWAP_ENABLED", log); err != nil {
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

// ParseLOGPUBLICSWAPENABLED is a log parse operation binding the contract event 0x183bed17d33ee5be9a47bb997aee6a152d84a34309cdb66f76301de0a40a1389.
//
// Solidity: event LOG_PUBLIC_SWAP_ENABLED()
func (_Poolbindings *PoolbindingsFilterer) ParseLOGPUBLICSWAPENABLED(log types.Log) (*PoolbindingsLOGPUBLICSWAPENABLED, error) {
	event := new(PoolbindingsLOGPUBLICSWAPENABLED)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_PUBLIC_SWAP_ENABLED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the Poolbindings contract.
type PoolbindingsLOGSWAPIterator struct {
	Event *PoolbindingsLOGSWAP // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGSWAP)
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
		it.Event = new(PoolbindingsLOGSWAP)
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
func (it *PoolbindingsLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGSWAP represents a LOGSWAP event raised by the Poolbindings contract.
type PoolbindingsLOGSWAP struct {
	Caller         common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	TokenAmountIn  *big.Int
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGSWAP is a free log retrieval operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*PoolbindingsLOGSWAPIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGSWAPIterator{contract: _Poolbindings.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGSWAP)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
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

// ParseLOGSWAP is a log parse operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGSWAP(log types.Log) (*PoolbindingsLOGSWAP, error) {
	event := new(PoolbindingsLOGSWAP)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGSWAPFEEUPDATEDIterator is returned from FilterLOGSWAPFEEUPDATED and is used to iterate over the raw logs and unpacked data for LOGSWAPFEEUPDATED events raised by the Poolbindings contract.
type PoolbindingsLOGSWAPFEEUPDATEDIterator struct {
	Event *PoolbindingsLOGSWAPFEEUPDATED // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGSWAPFEEUPDATEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGSWAPFEEUPDATED)
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
		it.Event = new(PoolbindingsLOGSWAPFEEUPDATED)
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
func (it *PoolbindingsLOGSWAPFEEUPDATEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGSWAPFEEUPDATEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGSWAPFEEUPDATED represents a LOGSWAPFEEUPDATED event raised by the Poolbindings contract.
type PoolbindingsLOGSWAPFEEUPDATED struct {
	SwapFee *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLOGSWAPFEEUPDATED is a free log retrieval operation binding the contract event 0xccfe595973efc7c1f6c29e31974d380470b9431d7770290185b7129419c7e63e.
//
// Solidity: event LOG_SWAP_FEE_UPDATED(uint256 swapFee)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGSWAPFEEUPDATED(opts *bind.FilterOpts) (*PoolbindingsLOGSWAPFEEUPDATEDIterator, error) {

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_SWAP_FEE_UPDATED")
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGSWAPFEEUPDATEDIterator{contract: _Poolbindings.contract, event: "LOG_SWAP_FEE_UPDATED", logs: logs, sub: sub}, nil
}

// WatchLOGSWAPFEEUPDATED is a free log subscription operation binding the contract event 0xccfe595973efc7c1f6c29e31974d380470b9431d7770290185b7129419c7e63e.
//
// Solidity: event LOG_SWAP_FEE_UPDATED(uint256 swapFee)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGSWAPFEEUPDATED(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGSWAPFEEUPDATED) (event.Subscription, error) {

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_SWAP_FEE_UPDATED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGSWAPFEEUPDATED)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_SWAP_FEE_UPDATED", log); err != nil {
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

// ParseLOGSWAPFEEUPDATED is a log parse operation binding the contract event 0xccfe595973efc7c1f6c29e31974d380470b9431d7770290185b7129419c7e63e.
//
// Solidity: event LOG_SWAP_FEE_UPDATED(uint256 swapFee)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGSWAPFEEUPDATED(log types.Log) (*PoolbindingsLOGSWAPFEEUPDATED, error) {
	event := new(PoolbindingsLOGSWAPFEEUPDATED)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_SWAP_FEE_UPDATED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGTOKENADDEDIterator is returned from FilterLOGTOKENADDED and is used to iterate over the raw logs and unpacked data for LOGTOKENADDED events raised by the Poolbindings contract.
type PoolbindingsLOGTOKENADDEDIterator struct {
	Event *PoolbindingsLOGTOKENADDED // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGTOKENADDEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGTOKENADDED)
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
		it.Event = new(PoolbindingsLOGTOKENADDED)
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
func (it *PoolbindingsLOGTOKENADDEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGTOKENADDEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGTOKENADDED represents a LOGTOKENADDED event raised by the Poolbindings contract.
type PoolbindingsLOGTOKENADDED struct {
	Token          common.Address
	DesiredDenorm  *big.Int
	MinimumBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGTOKENADDED is a free log retrieval operation binding the contract event 0xb2daf560899f6307b318aecfb57eb2812c488da4a4c1cad2019b482fa63294ed.
//
// Solidity: event LOG_TOKEN_ADDED(address indexed token, uint256 desiredDenorm, uint256 minimumBalance)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGTOKENADDED(opts *bind.FilterOpts, token []common.Address) (*PoolbindingsLOGTOKENADDEDIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_TOKEN_ADDED", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGTOKENADDEDIterator{contract: _Poolbindings.contract, event: "LOG_TOKEN_ADDED", logs: logs, sub: sub}, nil
}

// WatchLOGTOKENADDED is a free log subscription operation binding the contract event 0xb2daf560899f6307b318aecfb57eb2812c488da4a4c1cad2019b482fa63294ed.
//
// Solidity: event LOG_TOKEN_ADDED(address indexed token, uint256 desiredDenorm, uint256 minimumBalance)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGTOKENADDED(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGTOKENADDED, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_TOKEN_ADDED", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGTOKENADDED)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_TOKEN_ADDED", log); err != nil {
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

// ParseLOGTOKENADDED is a log parse operation binding the contract event 0xb2daf560899f6307b318aecfb57eb2812c488da4a4c1cad2019b482fa63294ed.
//
// Solidity: event LOG_TOKEN_ADDED(address indexed token, uint256 desiredDenorm, uint256 minimumBalance)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGTOKENADDED(log types.Log) (*PoolbindingsLOGTOKENADDED, error) {
	event := new(PoolbindingsLOGTOKENADDED)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_TOKEN_ADDED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGTOKENREADYIterator is returned from FilterLOGTOKENREADY and is used to iterate over the raw logs and unpacked data for LOGTOKENREADY events raised by the Poolbindings contract.
type PoolbindingsLOGTOKENREADYIterator struct {
	Event *PoolbindingsLOGTOKENREADY // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGTOKENREADYIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGTOKENREADY)
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
		it.Event = new(PoolbindingsLOGTOKENREADY)
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
func (it *PoolbindingsLOGTOKENREADYIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGTOKENREADYIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGTOKENREADY represents a LOGTOKENREADY event raised by the Poolbindings contract.
type PoolbindingsLOGTOKENREADY struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLOGTOKENREADY is a free log retrieval operation binding the contract event 0xf7bb8e57ffdfd9a31e7580ee84f68757f44fb4a8a913f44520d22f2da1c955e5.
//
// Solidity: event LOG_TOKEN_READY(address indexed token)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGTOKENREADY(opts *bind.FilterOpts, token []common.Address) (*PoolbindingsLOGTOKENREADYIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_TOKEN_READY", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGTOKENREADYIterator{contract: _Poolbindings.contract, event: "LOG_TOKEN_READY", logs: logs, sub: sub}, nil
}

// WatchLOGTOKENREADY is a free log subscription operation binding the contract event 0xf7bb8e57ffdfd9a31e7580ee84f68757f44fb4a8a913f44520d22f2da1c955e5.
//
// Solidity: event LOG_TOKEN_READY(address indexed token)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGTOKENREADY(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGTOKENREADY, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_TOKEN_READY", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGTOKENREADY)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_TOKEN_READY", log); err != nil {
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

// ParseLOGTOKENREADY is a log parse operation binding the contract event 0xf7bb8e57ffdfd9a31e7580ee84f68757f44fb4a8a913f44520d22f2da1c955e5.
//
// Solidity: event LOG_TOKEN_READY(address indexed token)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGTOKENREADY(log types.Log) (*PoolbindingsLOGTOKENREADY, error) {
	event := new(PoolbindingsLOGTOKENREADY)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_TOKEN_READY", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsLOGTOKENREMOVEDIterator is returned from FilterLOGTOKENREMOVED and is used to iterate over the raw logs and unpacked data for LOGTOKENREMOVED events raised by the Poolbindings contract.
type PoolbindingsLOGTOKENREMOVEDIterator struct {
	Event *PoolbindingsLOGTOKENREMOVED // Event containing the contract specifics and raw log

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
func (it *PoolbindingsLOGTOKENREMOVEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsLOGTOKENREMOVED)
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
		it.Event = new(PoolbindingsLOGTOKENREMOVED)
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
func (it *PoolbindingsLOGTOKENREMOVEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsLOGTOKENREMOVEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsLOGTOKENREMOVED represents a LOGTOKENREMOVED event raised by the Poolbindings contract.
type PoolbindingsLOGTOKENREMOVED struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLOGTOKENREMOVED is a free log retrieval operation binding the contract event 0x12a8262eb28ee8a8c11e6cf411b3af6ce5bea42abb36e051bf0a65ae602d52ec.
//
// Solidity: event LOG_TOKEN_REMOVED(address token)
func (_Poolbindings *PoolbindingsFilterer) FilterLOGTOKENREMOVED(opts *bind.FilterOpts) (*PoolbindingsLOGTOKENREMOVEDIterator, error) {

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "LOG_TOKEN_REMOVED")
	if err != nil {
		return nil, err
	}
	return &PoolbindingsLOGTOKENREMOVEDIterator{contract: _Poolbindings.contract, event: "LOG_TOKEN_REMOVED", logs: logs, sub: sub}, nil
}

// WatchLOGTOKENREMOVED is a free log subscription operation binding the contract event 0x12a8262eb28ee8a8c11e6cf411b3af6ce5bea42abb36e051bf0a65ae602d52ec.
//
// Solidity: event LOG_TOKEN_REMOVED(address token)
func (_Poolbindings *PoolbindingsFilterer) WatchLOGTOKENREMOVED(opts *bind.WatchOpts, sink chan<- *PoolbindingsLOGTOKENREMOVED) (event.Subscription, error) {

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "LOG_TOKEN_REMOVED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsLOGTOKENREMOVED)
				if err := _Poolbindings.contract.UnpackLog(event, "LOG_TOKEN_REMOVED", log); err != nil {
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

// ParseLOGTOKENREMOVED is a log parse operation binding the contract event 0x12a8262eb28ee8a8c11e6cf411b3af6ce5bea42abb36e051bf0a65ae602d52ec.
//
// Solidity: event LOG_TOKEN_REMOVED(address token)
func (_Poolbindings *PoolbindingsFilterer) ParseLOGTOKENREMOVED(log types.Log) (*PoolbindingsLOGTOKENREMOVED, error) {
	event := new(PoolbindingsLOGTOKENREMOVED)
	if err := _Poolbindings.contract.UnpackLog(event, "LOG_TOKEN_REMOVED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolbindingsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Poolbindings contract.
type PoolbindingsTransferIterator struct {
	Event *PoolbindingsTransfer // Event containing the contract specifics and raw log

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
func (it *PoolbindingsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolbindingsTransfer)
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
		it.Event = new(PoolbindingsTransfer)
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
func (it *PoolbindingsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolbindingsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolbindingsTransfer represents a Transfer event raised by the Poolbindings contract.
type PoolbindingsTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Poolbindings *PoolbindingsFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*PoolbindingsTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Poolbindings.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &PoolbindingsTransferIterator{contract: _Poolbindings.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Poolbindings *PoolbindingsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolbindingsTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Poolbindings.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolbindingsTransfer)
				if err := _Poolbindings.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Poolbindings *PoolbindingsFilterer) ParseTransfer(log types.Log) (*PoolbindingsTransfer, error) {
	event := new(PoolbindingsTransfer)
	if err := _Poolbindings.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
