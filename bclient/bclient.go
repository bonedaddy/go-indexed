package bclient

import (
	mcapscontroller "github.com/bonedaddy/go-indexed/bindings/marketcap_sqrt_controller"
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	ec *ethclient.Client
}

// NewInfuraClient returns an eth client connected to infura
func NewInfuraClient(token string, websockets bool) (*Client, error) {
	var url string
	if websockets {
		url = INFURA_WS_URL + token
	} else {
		url = INFURA_HTTP_URL + token
	}
	return NewClient(url)
}

func NewClient(url string) (*Client, error) {
	ec, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{ec}, nil
}

// DEFI5 returns a DEFI5 contract binding
func (c *Client) DEFI5() (IndexPool, error) {
	return poolbindings.NewPoolbindings(DEFI5TokenAddress, c.ec)
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

// StakingAt returns a staking rewards bindings at the given address
func (c *Client) StakingAt(addr common.Address) (*stakingbindings.Stakingbindings, error) {
	return stakingbindings.NewStakingbindings(DEFI5StakingAddress, c.ec)
}
