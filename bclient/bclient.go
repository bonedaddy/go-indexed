package bclient

import (
	"context"
	"errors"

	mcapscontroller "github.com/bonedaddy/go-indexed/bindings/marketcap_sqrt_controller"
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"
	uv2oraclebindings "github.com/bonedaddy/go-indexed/bindings/uniswapv2_oracle"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client wraps ethclient and provides helper functions for interacting with the indexed finance smart contracts
type Client struct {
	ec *ethclient.Client
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
	return &Client{ec}, nil
}

// CurrentBlock returns the current block known by the ethereum client
func (c *Client) CurrentBlock() (uint64, error) {
	return c.ec.BlockNumber(context.Background())
}

// DEFI5 returns a DEFI5 contract binding
func (c *Client) DEFI5() (IndexPool, error) {
	return poolbindings.NewPoolbindings(DEFI5TokenAddress, c.ec)
}

// CC10 returns a CC10 contract binding
func (c *Client) CC10() (IndexPool, error) {
	return poolbindings.NewPoolbindings(CC10TokenAddress, c.ec)
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
	default:
		return nil, errors.New("unsupported staking contract")
	}

}
