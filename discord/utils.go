package discord

import (
	"errors"

	"github.com/bonedaddy/go-indexed/bclient"
	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"
)

func (c *Client) getIndexPool(name string) (bclient.IndexPool, error) {
	switch name {
	case "defi5":
		return c.bc.DEFI5()
	case "cc10":
		return c.bc.CC10()
	default:
		return nil, errors.New("invalid pool name")
	}
}

func (c *Client) getStakingRewards(name string) (*stakingbindings.Stakingbindings, error) {
	switch name {
	case "defi5":
		return c.bc.StakingAt("defi5")
	case "univ2-eth-defi5":
		return c.bc.StakingAt("univ2-eth-defi5")
	default:
		return nil, errors.New("invalid staking rewards name")
	}
}
