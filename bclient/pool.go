package bclient

import (
	"math/big"
	"strings"

	"github.com/bonedaddy/go-indexed/bindings/multicall"
	"github.com/bonedaddy/go-indexed/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// ERC20I denotes ERC20 interface functions
type ERC20I interface {
	BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error)
	Decimals(opts *bind.CallOpts) (uint8, error)
	TotalSupply(opts *bind.CallOpts) (*big.Int, error)
}

// IndexPoolRead are read-only IndexPool contract calls
// See https://docs.indexed.finance/indexed-finance-docs/smart-contracts/pool#indexpool for more information
type IndexPoolRead interface {
	ERC20I
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
func (c *Client) GetTotalValueLocked(ip IndexPool, mc *multicall.Multicall, logger *zap.Logger, poolAddress common.Address) (float64, error) {
	tokens, err := c.PoolTokensForMC(mc, poolAddress)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get pool tokens")
	}

	uc := c.Uniswap()

	type tokenValue struct {
		tokenUsdValue float64        // the value of a single token
		tokenAddress  common.Address // the address of the token contract
	}
	decimals := make(map[common.Address]uint8)
	values := make(map[string]*tokenValue)

	var tokenAddrs []common.Address
	for _, addr := range tokens {
		tokenAddrs = append(tokenAddrs, addr)
	}

	addrs, decs, err := mc.GetDecimals(nil, tokenAddrs)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get token decimals from multicall")
	}

	if len(addrs) != len(decs) {
		return 0, errors.New("mismatching addr and dec lengths")
	}

	for i := 0; i < len(addrs); i++ {
		decimals[addrs[i]] = decs[i]
	}

	for symbol, addr := range tokens {
		// hard coded list for alternate price lookups of X->ETH->DAI
		switch strings.ToLower(symbol) {
		default:
			tokenEthPrice, err := uc.GetExchangeAmount(utils.ToWei("1.0", int(decimals[addr])), addr, WETHTokenAddress)
			if err != nil {
				return 0, errors.Wrap(err, "failed to get exchange amount")
			}
			tokenEthPriceDec := utils.ToDecimal(tokenEthPrice, int(decimals[addr]))

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
			logger.Debug("usd value retrieve", zap.String("asset", symbol), zap.Float64("usd.value", values[symbol].tokenUsdValue))
		}
	}

	balances := make(map[common.Address]*big.Int)

	addrs, bals, err := mc.GetBalances(nil, poolAddress, tokenAddrs)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get pool token balances from multicall")
	}

	if len(addrs) != len(bals) {
		return 0, errors.New("mismatching addr and bal lengths")
	}

	for i := 0; i < len(addrs); i++ {
		balances[addrs[i]] = bals[i]
	}

	var totalValueUSD float64
	for _, object := range values {

		poolBalanceF, _ := utils.ToDecimal(balances[object.tokenAddress], int(decimals[object.tokenAddress])).Float64()
		poolUSDValue := object.tokenUsdValue * poolBalanceF

		totalValueUSD += poolUSDValue

	}

	return totalValueUSD, nil
}
