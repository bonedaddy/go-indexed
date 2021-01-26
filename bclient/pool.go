package bclient

import (
	"log"
	"math/big"
	"strings"

	"github.com/bonedaddy/go-indexed/bindings/erc20"
	"github.com/bonedaddy/go-indexed/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
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
	TotalSupply(opts *bind.CallOpts) (*big.Int, error)
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

// GetTotalValueLocked returns the total value locked into the contracts
func (c *Client) GetTotalValueLocked(ip IndexPool) (float64, error) {
	tokens, err := c.PoolTokensFor(ip)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get pool tokens")
	}

	uc := c.Uniswap()

	type tokenValue struct {
		tokenUsdValue float64        // the value of a single token
		tokenAddress  common.Address // the address of the token contract
	}

	values := make(map[string]*tokenValue)

	for symbol, addr := range tokens {
		// hard coded list for alternate price lookups of X->ETH->DAI
		switch strings.ToLower(symbol) {
		default:
			erc, err := erc20.NewErc20(addr, c.ec)
			if err != nil {
				return 0, errors.Wrap(err, "failed to get erc20 contract")
			}
			dec, err := erc.Decimals(nil)
			if err != nil {
				return 0, errors.Wrap(err, "failed to get decimals")
			}

			tokenEthPrice, err := uc.GetExchangeAmount(utils.ToWei("1.0", int(dec)), addr, WETHTokenAddress)
			if err != nil {
				return 0, errors.Wrap(err, "failed to get exchange amount")
			}
			tokenEthPriceDec := utils.ToDecimal(tokenEthPrice, int(dec))

			ethDaiPrice, err := c.EthDaiPrice()
			if err != nil {
				return 0, errors.Wrap(err, "failed to get dai price")
			}
			ethDaiPriceDec := utils.ToDecimal(utils.ToWei(ethDaiPrice.Int64(), 18), 18)

			// derive price of token by getting the amount of ETH you would get from 1 token
			// and convert that into dai
			tdF, _ := tokenEthPriceDec.Float64()
			edF, _ := ethDaiPriceDec.Float64()

			values[symbol] = &tokenValue{
				tokenUsdValue: edF * tdF,
				tokenAddress:  addr,
			}
			log.Printf("%s USD value: %0.2f", symbol, values[symbol].tokenUsdValue)
		}
	}
	var totalValueUSD float64
	for _, object := range values {
		erc, err := erc20.NewErc20(object.tokenAddress, c.ec)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get erc20 contract")
		}
		dec, err := erc.Decimals(nil)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get decimals")
		}

		poolBalance, err := ip.GetBalance(nil, object.tokenAddress)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get pool token balance")
		}

		poolBalanceF, _ := utils.ToDecimal(poolBalance, int(dec)).Float64()
		poolUSDValue := object.tokenUsdValue * poolBalanceF

		totalValueUSD += poolUSDValue

	}

	return totalValueUSD, nil
}
