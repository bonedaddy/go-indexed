package bclient

import (
	"math/big"

	"github.com/bonedaddy/go-indexed/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// IndexPoolRead are read-only IndexPool contract calls
// See https://docs.indexed.finance/indexed-finance-docs/smart-contracts/pool#indexpool for more information
type IndexPoolRead interface {
	IsPublicSwap(opts *bind.CallOpts) (bool, error)
	GetController(opts *bind.CallOpts) (common.Address, error)
	GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetNumTokens(opts *bind.CallOpts) (*big.Int, error)
	GetMaxPoolTokens(opts *bind.CallOpts) (*big.Int, error)
	GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error)
	GetSwapFee(opts *bind.CallOpts) (*big.Int, error)
	GetCurrentDesiredTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error)
	GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error)
	GetUsedBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error)
	GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error)
	ExtrapolatePoolValueFromToken(opts *bind.CallOpts) (common.Address, *big.Int, error)
	BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error)
	Decimals(opts *bind.CallOpts) (uint8, error)
}

// IndexPool provides helper functions around the IndexPool contract
type IndexPool interface {
	IndexPoolRead
}

// BalanceOfDecimal returns the IndexPool balance in decimal (non wei) format
func BalanceOfDecimal(ip IndexPool, addr common.Address) (decimal.Decimal, error) {
	bal, err := ip.BalanceOf(nil, addr)
	if err != nil {
		return decimal.Decimal{}, err
	}
	dec, err := ip.Decimals(nil)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return utils.ToDecimal(bal, int(dec)), nil
}
