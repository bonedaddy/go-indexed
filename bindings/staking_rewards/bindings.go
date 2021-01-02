// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingbindings

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

// StakingbindingsABI is the input ABI used to generate the binding from.
const StakingbindingsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardsDistribution_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardsDuration_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"setRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Stakingbindings is an auto generated Go binding around an Ethereum contract.
type Stakingbindings struct {
	StakingbindingsCaller     // Read-only binding to the contract
	StakingbindingsTransactor // Write-only binding to the contract
	StakingbindingsFilterer   // Log filterer for contract events
}

// StakingbindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingbindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingbindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingbindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingbindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingbindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingbindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingbindingsSession struct {
	Contract     *Stakingbindings  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingbindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingbindingsCallerSession struct {
	Contract *StakingbindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StakingbindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingbindingsTransactorSession struct {
	Contract     *StakingbindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StakingbindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingbindingsRaw struct {
	Contract *Stakingbindings // Generic contract binding to access the raw methods on
}

// StakingbindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingbindingsCallerRaw struct {
	Contract *StakingbindingsCaller // Generic read-only contract binding to access the raw methods on
}

// StakingbindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingbindingsTransactorRaw struct {
	Contract *StakingbindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingbindings creates a new instance of Stakingbindings, bound to a specific deployed contract.
func NewStakingbindings(address common.Address, backend bind.ContractBackend) (*Stakingbindings, error) {
	contract, err := bindStakingbindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakingbindings{StakingbindingsCaller: StakingbindingsCaller{contract: contract}, StakingbindingsTransactor: StakingbindingsTransactor{contract: contract}, StakingbindingsFilterer: StakingbindingsFilterer{contract: contract}}, nil
}

// NewStakingbindingsCaller creates a new read-only instance of Stakingbindings, bound to a specific deployed contract.
func NewStakingbindingsCaller(address common.Address, caller bind.ContractCaller) (*StakingbindingsCaller, error) {
	contract, err := bindStakingbindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingbindingsCaller{contract: contract}, nil
}

// NewStakingbindingsTransactor creates a new write-only instance of Stakingbindings, bound to a specific deployed contract.
func NewStakingbindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingbindingsTransactor, error) {
	contract, err := bindStakingbindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingbindingsTransactor{contract: contract}, nil
}

// NewStakingbindingsFilterer creates a new log filterer instance of Stakingbindings, bound to a specific deployed contract.
func NewStakingbindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingbindingsFilterer, error) {
	contract, err := bindStakingbindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingbindingsFilterer{contract: contract}, nil
}

// bindStakingbindings binds a generic wrapper to an already deployed contract.
func bindStakingbindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingbindingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingbindings *StakingbindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingbindings.Contract.StakingbindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingbindings *StakingbindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingbindings.Contract.StakingbindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingbindings *StakingbindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingbindings.Contract.StakingbindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingbindings *StakingbindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingbindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingbindings *StakingbindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingbindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingbindings *StakingbindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingbindings.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.BalanceOf(&_Stakingbindings.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.BalanceOf(&_Stakingbindings.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) Earned(account common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.Earned(&_Stakingbindings.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.Earned(&_Stakingbindings.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) GetRewardForDuration() (*big.Int, error) {
	return _Stakingbindings.Contract.GetRewardForDuration(&_Stakingbindings.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _Stakingbindings.Contract.GetRewardForDuration(&_Stakingbindings.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Stakingbindings.Contract.LastTimeRewardApplicable(&_Stakingbindings.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Stakingbindings.Contract.LastTimeRewardApplicable(&_Stakingbindings.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) LastUpdateTime() (*big.Int, error) {
	return _Stakingbindings.Contract.LastUpdateTime(&_Stakingbindings.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Stakingbindings.Contract.LastUpdateTime(&_Stakingbindings.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) PeriodFinish() (*big.Int, error) {
	return _Stakingbindings.Contract.PeriodFinish(&_Stakingbindings.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) PeriodFinish() (*big.Int, error) {
	return _Stakingbindings.Contract.PeriodFinish(&_Stakingbindings.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) RewardPerToken() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardPerToken(&_Stakingbindings.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) RewardPerToken() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardPerToken(&_Stakingbindings.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) RewardPerTokenStored() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardPerTokenStored(&_Stakingbindings.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardPerTokenStored(&_Stakingbindings.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) RewardRate() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardRate(&_Stakingbindings.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) RewardRate() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardRate(&_Stakingbindings.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.Rewards(&_Stakingbindings.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.Rewards(&_Stakingbindings.CallOpts, arg0)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_Stakingbindings *StakingbindingsCaller) RewardsDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "rewardsDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_Stakingbindings *StakingbindingsSession) RewardsDistribution() (common.Address, error) {
	return _Stakingbindings.Contract.RewardsDistribution(&_Stakingbindings.CallOpts)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_Stakingbindings *StakingbindingsCallerSession) RewardsDistribution() (common.Address, error) {
	return _Stakingbindings.Contract.RewardsDistribution(&_Stakingbindings.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) RewardsDuration() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardsDuration(&_Stakingbindings.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) RewardsDuration() (*big.Int, error) {
	return _Stakingbindings.Contract.RewardsDuration(&_Stakingbindings.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_Stakingbindings *StakingbindingsCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_Stakingbindings *StakingbindingsSession) RewardsToken() (common.Address, error) {
	return _Stakingbindings.Contract.RewardsToken(&_Stakingbindings.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_Stakingbindings *StakingbindingsCallerSession) RewardsToken() (common.Address, error) {
	return _Stakingbindings.Contract.RewardsToken(&_Stakingbindings.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Stakingbindings *StakingbindingsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Stakingbindings *StakingbindingsSession) StakingToken() (common.Address, error) {
	return _Stakingbindings.Contract.StakingToken(&_Stakingbindings.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Stakingbindings *StakingbindingsCallerSession) StakingToken() (common.Address, error) {
	return _Stakingbindings.Contract.StakingToken(&_Stakingbindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) TotalSupply() (*big.Int, error) {
	return _Stakingbindings.Contract.TotalSupply(&_Stakingbindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) TotalSupply() (*big.Int, error) {
	return _Stakingbindings.Contract.TotalSupply(&_Stakingbindings.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Stakingbindings *StakingbindingsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingbindings.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Stakingbindings *StakingbindingsSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.UserRewardPerTokenPaid(&_Stakingbindings.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Stakingbindings *StakingbindingsCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Stakingbindings.Contract.UserRewardPerTokenPaid(&_Stakingbindings.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingbindings *StakingbindingsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingbindings *StakingbindingsSession) Exit() (*types.Transaction, error) {
	return _Stakingbindings.Contract.Exit(&_Stakingbindings.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingbindings *StakingbindingsTransactorSession) Exit() (*types.Transaction, error) {
	return _Stakingbindings.Contract.Exit(&_Stakingbindings.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingbindings *StakingbindingsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingbindings *StakingbindingsSession) GetReward() (*types.Transaction, error) {
	return _Stakingbindings.Contract.GetReward(&_Stakingbindings.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingbindings *StakingbindingsTransactorSession) GetReward() (*types.Transaction, error) {
	return _Stakingbindings.Contract.GetReward(&_Stakingbindings.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address stakingToken_, uint256 rewardsDuration_) returns()
func (_Stakingbindings *StakingbindingsTransactor) Initialize(opts *bind.TransactOpts, stakingToken_ common.Address, rewardsDuration_ *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "initialize", stakingToken_, rewardsDuration_)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address stakingToken_, uint256 rewardsDuration_) returns()
func (_Stakingbindings *StakingbindingsSession) Initialize(stakingToken_ common.Address, rewardsDuration_ *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.Initialize(&_Stakingbindings.TransactOpts, stakingToken_, rewardsDuration_)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address stakingToken_, uint256 rewardsDuration_) returns()
func (_Stakingbindings *StakingbindingsTransactorSession) Initialize(stakingToken_ common.Address, rewardsDuration_ *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.Initialize(&_Stakingbindings.TransactOpts, stakingToken_, rewardsDuration_)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Stakingbindings *StakingbindingsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Stakingbindings *StakingbindingsSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.NotifyRewardAmount(&_Stakingbindings.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Stakingbindings *StakingbindingsTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.NotifyRewardAmount(&_Stakingbindings.TransactOpts, reward)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address tokenAddress, address recipient) returns()
func (_Stakingbindings *StakingbindingsTransactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "recoverERC20", tokenAddress, recipient)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address tokenAddress, address recipient) returns()
func (_Stakingbindings *StakingbindingsSession) RecoverERC20(tokenAddress common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Stakingbindings.Contract.RecoverERC20(&_Stakingbindings.TransactOpts, tokenAddress, recipient)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address tokenAddress, address recipient) returns()
func (_Stakingbindings *StakingbindingsTransactorSession) RecoverERC20(tokenAddress common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Stakingbindings.Contract.RecoverERC20(&_Stakingbindings.TransactOpts, tokenAddress, recipient)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_Stakingbindings *StakingbindingsTransactor) SetRewardsDuration(opts *bind.TransactOpts, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "setRewardsDuration", _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_Stakingbindings *StakingbindingsSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.SetRewardsDuration(&_Stakingbindings.TransactOpts, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_Stakingbindings *StakingbindingsTransactorSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.SetRewardsDuration(&_Stakingbindings.TransactOpts, _rewardsDuration)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingbindings *StakingbindingsTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingbindings *StakingbindingsSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.Stake(&_Stakingbindings.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingbindings *StakingbindingsTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.Stake(&_Stakingbindings.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingbindings *StakingbindingsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingbindings *StakingbindingsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.Withdraw(&_Stakingbindings.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingbindings *StakingbindingsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stakingbindings.Contract.Withdraw(&_Stakingbindings.TransactOpts, amount)
}

// StakingbindingsRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the Stakingbindings contract.
type StakingbindingsRecoveredIterator struct {
	Event *StakingbindingsRecovered // Event containing the contract specifics and raw log

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
func (it *StakingbindingsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingbindingsRecovered)
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
		it.Event = new(StakingbindingsRecovered)
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
func (it *StakingbindingsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingbindingsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingbindingsRecovered represents a Recovered event raised by the Stakingbindings contract.
type StakingbindingsRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) FilterRecovered(opts *bind.FilterOpts) (*StakingbindingsRecoveredIterator, error) {

	logs, sub, err := _Stakingbindings.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &StakingbindingsRecoveredIterator{contract: _Stakingbindings.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *StakingbindingsRecovered) (event.Subscription, error) {

	logs, sub, err := _Stakingbindings.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingbindingsRecovered)
				if err := _Stakingbindings.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) ParseRecovered(log types.Log) (*StakingbindingsRecovered, error) {
	event := new(StakingbindingsRecovered)
	if err := _Stakingbindings.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingbindingsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Stakingbindings contract.
type StakingbindingsRewardAddedIterator struct {
	Event *StakingbindingsRewardAdded // Event containing the contract specifics and raw log

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
func (it *StakingbindingsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingbindingsRewardAdded)
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
		it.Event = new(StakingbindingsRewardAdded)
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
func (it *StakingbindingsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingbindingsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingbindingsRewardAdded represents a RewardAdded event raised by the Stakingbindings contract.
type StakingbindingsRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Stakingbindings *StakingbindingsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*StakingbindingsRewardAddedIterator, error) {

	logs, sub, err := _Stakingbindings.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &StakingbindingsRewardAddedIterator{contract: _Stakingbindings.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Stakingbindings *StakingbindingsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *StakingbindingsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Stakingbindings.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingbindingsRewardAdded)
				if err := _Stakingbindings.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Stakingbindings *StakingbindingsFilterer) ParseRewardAdded(log types.Log) (*StakingbindingsRewardAdded, error) {
	event := new(StakingbindingsRewardAdded)
	if err := _Stakingbindings.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingbindingsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the Stakingbindings contract.
type StakingbindingsRewardPaidIterator struct {
	Event *StakingbindingsRewardPaid // Event containing the contract specifics and raw log

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
func (it *StakingbindingsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingbindingsRewardPaid)
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
		it.Event = new(StakingbindingsRewardPaid)
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
func (it *StakingbindingsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingbindingsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingbindingsRewardPaid represents a RewardPaid event raised by the Stakingbindings contract.
type StakingbindingsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Stakingbindings *StakingbindingsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*StakingbindingsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingbindings.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingbindingsRewardPaidIterator{contract: _Stakingbindings.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Stakingbindings *StakingbindingsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *StakingbindingsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingbindings.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingbindingsRewardPaid)
				if err := _Stakingbindings.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Stakingbindings *StakingbindingsFilterer) ParseRewardPaid(log types.Log) (*StakingbindingsRewardPaid, error) {
	event := new(StakingbindingsRewardPaid)
	if err := _Stakingbindings.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingbindingsRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the Stakingbindings contract.
type StakingbindingsRewardsDurationUpdatedIterator struct {
	Event *StakingbindingsRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *StakingbindingsRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingbindingsRewardsDurationUpdated)
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
		it.Event = new(StakingbindingsRewardsDurationUpdated)
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
func (it *StakingbindingsRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingbindingsRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingbindingsRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the Stakingbindings contract.
type StakingbindingsRewardsDurationUpdated struct {
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_Stakingbindings *StakingbindingsFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*StakingbindingsRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _Stakingbindings.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingbindingsRewardsDurationUpdatedIterator{contract: _Stakingbindings.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_Stakingbindings *StakingbindingsFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *StakingbindingsRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Stakingbindings.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingbindingsRewardsDurationUpdated)
				if err := _Stakingbindings.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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

// ParseRewardsDurationUpdated is a log parse operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_Stakingbindings *StakingbindingsFilterer) ParseRewardsDurationUpdated(log types.Log) (*StakingbindingsRewardsDurationUpdated, error) {
	event := new(StakingbindingsRewardsDurationUpdated)
	if err := _Stakingbindings.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingbindingsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Stakingbindings contract.
type StakingbindingsStakedIterator struct {
	Event *StakingbindingsStaked // Event containing the contract specifics and raw log

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
func (it *StakingbindingsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingbindingsStaked)
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
		it.Event = new(StakingbindingsStaked)
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
func (it *StakingbindingsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingbindingsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingbindingsStaked represents a Staked event raised by the Stakingbindings contract.
type StakingbindingsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*StakingbindingsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingbindings.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingbindingsStakedIterator{contract: _Stakingbindings.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingbindingsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingbindings.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingbindingsStaked)
				if err := _Stakingbindings.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) ParseStaked(log types.Log) (*StakingbindingsStaked, error) {
	event := new(StakingbindingsStaked)
	if err := _Stakingbindings.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingbindingsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Stakingbindings contract.
type StakingbindingsWithdrawnIterator struct {
	Event *StakingbindingsWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingbindingsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingbindingsWithdrawn)
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
		it.Event = new(StakingbindingsWithdrawn)
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
func (it *StakingbindingsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingbindingsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingbindingsWithdrawn represents a Withdrawn event raised by the Stakingbindings contract.
type StakingbindingsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*StakingbindingsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingbindings.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingbindingsWithdrawnIterator{contract: _Stakingbindings.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingbindingsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingbindings.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingbindingsWithdrawn)
				if err := _Stakingbindings.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Stakingbindings *StakingbindingsFilterer) ParseWithdrawn(log types.Log) (*StakingbindingsWithdrawn, error) {
	event := new(StakingbindingsWithdrawn)
	if err := _Stakingbindings.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
