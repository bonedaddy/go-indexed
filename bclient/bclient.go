package bclient

import (
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

// StakingAt returns a staking rewards bindings at the given address
func (c *Client) StakingAt(addr common.Address) (*stakingbindings.Stakingbindings, error) {
	return stakingbindings.NewStakingbindings(DEFI5StakingAddress, c.ec)
}
