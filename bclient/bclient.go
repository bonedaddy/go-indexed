package bclient

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bonedaddy/go-indexed/bindings/erc20"
	governoralpha "github.com/bonedaddy/go-indexed/bindings/governor_alpha"
	mcapscontroller "github.com/bonedaddy/go-indexed/bindings/marketcap_sqrt_controller"
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"
	uv2oraclebindings "github.com/bonedaddy/go-indexed/bindings/uniswapv2_oracle"
	"github.com/bonedaddy/go-indexed/uniswap"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client wraps ethclient and provides helper functions for interacting with the indexed finance smart contracts
type Client struct {
	ec *ethclient.Client
	uc *uniswap.Client
}

// NewInfuraClient returns an eth client connected to infura
func NewInfuraClient(token string, websockets bool) (*Client, error) {
	var url string
	if websockets {
		url = InfuraWSURL + token
	} else {
		url = InfuraHTTPURL + token
	}
	return NewClient(url)
}

// NewClient returns an eth client connected to an RPC
func NewClient(url string) (*Client, error) {
	ec, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{ec, uniswap.NewClient(ec)}, nil
}

// CurrentBlock returns the current block known by the ethereum client
func (c *Client) CurrentBlock() (uint64, error) {
	return c.ec.BlockNumber(context.Background())
}

// NDX returns an NDX contract binding
func (c *Client) NDX() (*erc20.Erc20, error) {
	return erc20.NewErc20(NDXTokenAddress, c.ec)
}

// DEFI5 returns a DEFI5 contract binding
func (c *Client) DEFI5() (IndexPool, error) {
	return poolbindings.NewPoolbindings(DEFI5TokenAddress, c.ec)
}

// CC10 returns a CC10 contract binding
func (c *Client) CC10() (IndexPool, error) {
	return poolbindings.NewPoolbindings(CC10TokenAddress, c.ec)
}

// ORCL5 returns a ORCL5 contract binding
func (c *Client) ORCL5() (IndexPool, error) {
	return poolbindings.NewPoolbindings(ORCL5TokenAddress, c.ec)
}

// MCAPControllerAt returns the marketcap square root controller bindings for an IndexPool
func (c *Client) MCAPControllerAt(ip IndexPool) (*mcapscontroller.Mcapscontroller, error) {
	cntrl, err := ip.GetController(nil)
	if err != nil {
		return nil, err
	}
	return mcapscontroller.NewMcapscontroller(cntrl, c.ec)
}

// OracleFor returns the uniswap v2 oracle address for tne given index pool
func (c *Client) OracleFor(ip IndexPool) (common.Address, error) {
	mcaps, err := c.MCAPControllerAt(ip)
	if err != nil {
		return common.Address{}, err
	}
	return mcaps.Oracle(nil)
}

// OracleAt returns the uniswapv2 oracle contract binding that the current idnex pool is using
func (c *Client) OracleAt(ip IndexPool) (*uv2oraclebindings.Uv2oraclebindings, error) {
	addr, err := c.OracleFor(ip)
	if err != nil {
		return nil, err
	}
	return uv2oraclebindings.NewUv2oraclebindings(addr, c.ec)
}

// StakingAt returns a staking rewards bindings at the given address
func (c *Client) StakingAt(contractType string) (*stakingbindings.Stakingbindings, error) {
	switch contractType {
	case "defi5":
		return stakingbindings.NewStakingbindings(DEFI5StakingAddress, c.ec)
	case "univ2-eth-defi5":
		return stakingbindings.NewStakingbindings(DEFI5UNILPStakingAddress, c.ec)
	case "cc10":
		return stakingbindings.NewStakingbindings(CC10StakingAddress, c.ec)
	case "univ2-eth-cc10":
		return stakingbindings.NewStakingbindings(CC10UNILPStakingAddress, c.ec)
	default:
		return nil, errors.New("unsupported staking contract")
	}

}

// PoolTokensFor returns the pools tokens baseketed in the pool, and their ERC20 name
func (c *Client) PoolTokensFor(ip IndexPool) (map[string]common.Address, error) {
	tokens, err := ip.GetCurrentTokens(nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var out = make(map[string]common.Address)
	for _, token := range tokens {
		ec, err := erc20.NewErc20(token, c.ec)
		if err != nil {
			log.Println("erc: ", err)
			return nil, err
		}
		// get the token name
		// ignore error as some tokens such as maker cause this problem
		// https://github.com/ethereum/go-ethereum/issues/21754#issuecomment-716231021
		symbol, _ := ec.Symbol(nil)
		if symbol == "" {
			symbol = guessTokenSymbol(token.String())
		}
		out[symbol] = token
	}
	return out, nil
}

// Uniswap returns a uniswap client helper
func (c *Client) Uniswap() *uniswap.Client { return c.uc }

// EthClient returns the underlying ethclient connection
func (c *Client) EthClient() *ethclient.Client { return c.ec }

// GovernorAlpha returns the GovernorAlpha contracts binding at the active governance address
func (c *Client) GovernorAlpha() (*governoralpha.Governoralpha, error) {
	return governoralpha.NewGovernoralpha(GovernorAlpha, c.ec)
}

// Close terminates the blockchain connection
func (c *Client) Close() {
	c.ec.Close()
}

func guessTokenName(address string) string {
	if address == "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2" {
		return "Maker"
	}
	return fmt.Sprintf("unknown-%v-%v", time.Now().UnixNano(), len(address))
}

func guessTokenSymbol(address string) string {
	if address == "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2" {
		return "MKR"
	}
	return fmt.Sprintf("unknown-%v-%v", time.Now().UnixNano(), len(address))
}
