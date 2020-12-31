package bclient

import (
	"errors"
	"math/big"
	"strings"

	"github.com/bonedaddy/go-indexed/uniswap"
	"github.com/ethereum/go-ethereum/common"
)

// NdxEthPairAddress returns the UniswapV2 pair of NDX-ETH
func (c *Client) NdxEthPairAddress() common.Address {
	return uniswap.GeneratePairAddress(NDXTokenAddress, WETHTokenAddress)
}

// Defi5EthPairAddress returns the UniswapV2 pair of DEFI5-ETH
func (c *Client) Defi5EthPairAddress() common.Address {
	return uniswap.GeneratePairAddress(DEFI5TokenAddress, WETHTokenAddress)
}

// Cc10EthPairAddress returns the UniswapV2 pair of CC10-ETH
func (c *Client) Cc10EthPairAddress() common.Address {
	return uniswap.GeneratePairAddress(CC10TokenAddress, WETHTokenAddress)
}

// ExchangeAmount returns the exchange amount for a variety of pairs
func (c *Client) ExchangeAmount(amount *big.Int, pair string) (*big.Int, error) {
	switch strings.ToLower(pair) {
	case "ndx-eth":
		return c.uc.GetExchangeAmount(amount, NDXTokenAddress, WETHTokenAddress)
	case "eth-ndx":
		return c.uc.GetExchangeAmount(amount, WETHTokenAddress, NDXTokenAddress)
	case "cc10-eth":
		return c.uc.GetExchangeAmount(amount, CC10TokenAddress, WETHTokenAddress)
	case "eth-cc10":
		return c.uc.GetExchangeAmount(amount, WETHTokenAddress, CC10TokenAddress)
	case "defi5-eth":
		return c.uc.GetExchangeAmount(amount, DEFI5TokenAddress, WETHTokenAddress)
	case "eth-defi5":
		return c.uc.GetExchangeAmount(amount, WETHTokenAddress, DEFI5TokenAddress)
	default:
		return nil, errors.New("unsupported pair")
	}
}
