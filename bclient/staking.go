package bclient

import (
	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"
	"github.com/bonedaddy/go-indexed/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// StakeBalanceOf returns the balance of staking tokens owned by account
func StakeBalanceOf(
	sp *stakingbindings.Stakingbindings,
	ip IndexPoolRead,
	addr common.Address) (decimal.Decimal, error) {
	bal, err := sp.BalanceOf(nil, addr)
	if err != nil {
		return decimal.Decimal{}, err
	}
	dec, err := ip.Decimals(nil)
	if err != nil {
		return decimal.Decimal{}, nil
	}
	return utils.ToDecimal(bal, int(dec)), nil
}

// StakeEarned returns the amount of staking tokens earned
func StakeEarned(
	sp *stakingbindings.Stakingbindings,
	ip IndexPoolRead,
	addr common.Address) (decimal.Decimal, error) {
	earned, err := sp.Earned(nil, addr)
	if err != nil {
		return decimal.Decimal{}, nil
	}
	dec, err := ip.Decimals(nil)
	if err != nil {
		return decimal.Decimal{}, nil
	}
	return utils.ToDecimal(earned, int(dec)), nil
}
