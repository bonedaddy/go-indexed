package discord

import (
	"errors"
	"fmt"
	"strings"

	stakingbindings "github.com/bonedaddy/go-indexed/bindings/staking_rewards"
	"github.com/bwmarrin/discordgo"
)

func (c *Client) getStakingRewards(name string) (*stakingbindings.Stakingbindings, error) {
	switch name {
	case "defi5":
		return c.bc.StakingAt("defi5")
	case "cc10":
		return c.bc.StakingAt("cc10")
	case "univ2-eth-defi5":
		return c.bc.StakingAt("univ2-eth-defi5")
	case "univ2-eth-cc10":
		return c.bc.StakingAt("univ2-eth-cc10")
	default:
		return nil, errors.New("invalid staking rewards name")
	}
}

// BaseEmbed returns a base message embed type to be customized
func BaseEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://pbs.twimg.com/profile_images/1342395531318976518/kIv5abLc_400x400.jpg",
		},
		Color: 0x00ff00,
	}
}

// ParseValue converts numbers like 10000000 into 10M
func ParseValue(num float64) string {
	parts := strings.Split(fmt.Sprintf("%0.2f", num), ".")
	if len(parts) == 0 {
		return ""
	}
	strPart := parts[0]
	/*
		500,000 (hundred thousands 6 numbers)
		1,000,000 (millions 7 numbers)
		10,000,000 (10 millions 8 numbers)
		100,000,000 (100 millions 9 numbers)
		1,000,000,000 (billions 10 numbers)
	*/
	switch len(strPart) {
	case 6:
		return string(strPart[0]) + string(strPart[1]) + string(strPart[2]) + "K"
	case 7:
		return string(strPart[0]) + "M"
	case 8:
		return string(strPart[0]) + string(strPart[1]) + "M"
	case 9:
		return string(strPart[0]) + string(strPart[1]) + string(strPart[2]) + "M"
	case 10:
		return string(strPart[0]) + "B"
	default:
		return strPart
	}
}
