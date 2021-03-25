package discord

import (
	"strings"

	"github.com/bonedaddy/dgc"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bwmarrin/discordgo"
)

// lists various indexed contracts

func (c *Client) listContracts(ctx *dgc.Ctx) {
	embed := BaseEmbed()
	embed.Title = "Contracts Overview"
	embed.Description = "While these will be correct please verify the addresses before using them for financial transactions"
	arguments := ctx.Arguments
	filter := strings.ToLower(arguments.Get(0).Raw())
	switch filter {
	case "list-filters":
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "filters",
				Value: "all, staking, indexes, governance, uniswap-pairs",
			},
		}
	case "uniswap-pairs":
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "DEFI5",
				Value: bclient.DEFI5LP.String(),
			},
			{
				Name:  "CC10",
				Value: bclient.CC10LP.String(),
			},
			{
				Name:  "ORCL5",
				Value: bclient.ORCL5LP.String(),
			},
			{
				Name:  "DEGEN",
				Value: bclient.DEGENLP.String(),
			},
			{
				Name:  "NDX",
				Value: bclient.NDXLP.String(),
			},
		}
	case "governance":
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Governor Alpha",
				Value: bclient.GovernorAlpha.String(),
			},
		}
	case "indexes":
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "DEFI5",
				Value: bclient.DEFI5TokenAddress.String(),
			},
			{
				Name:  "CC10",
				Value: bclient.CC10TokenAddress.String(),
			},
			{
				Name:  "ORCL5",
				Value: bclient.ORCL5TokenAddress.String(),
			},
			{
				Name:  "DEGEN",
				Value: bclient.DEGEN10TokenAddress.String(),
			},
		}
	case "staking":
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "DEFI5 Staking",
				Value: bclient.DEFI5StakingAddress.String(),
			},
			{
				Name:  "CC10 Staking",
				Value: bclient.CC10StakingAddress.String(),
			},
			{
				Name:  "UNIV2:DEFI5-ETH LP Staking",
				Value: bclient.DEFI5UNILPStakingAddress.String(),
			},
			{
				Name:  "UNIV2:CC10-ETH LP Staking",
				Value: bclient.CC10UNILPStakingAddress.String(),
			},
			{
				Name:  "UNIV2:DEGEN-ETH LP Staking",
				Value: bclient.DEGENUNILPStakingAddress.String(),
			},
		}
	case "all", "":
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "NDX Token",
				Value: bclient.NDXTokenAddress.String(),
			},
			{
				Name:  "DEFI5 Token",
				Value: bclient.DEFI5TokenAddress.String(),
			},
			{
				Name:  "CC10 Token",
				Value: bclient.CC10TokenAddress.String(),
			},
			{
				Name:  "ORCL5 Token",
				Value: bclient.ORCL5TokenAddress.String(),
			},
			{
				Name:  "DEGEN Token",
				Value: bclient.DEGEN10TokenAddress.String(),
			},
			{
				Name:  "DEFI5 Staking",
				Value: bclient.DEFI5StakingAddress.String(),
			},
			{
				Name:  "CC10 Staking",
				Value: bclient.CC10StakingAddress.String(),
			},
			{
				Name:  "UNIV2:DEFI5-ETH LP Staking",
				Value: bclient.DEFI5UNILPStakingAddress.String(),
			},
			{
				Name:  "UNIV2:CC10-ETH LP Staking",
				Value: bclient.CC10UNILPStakingAddress.String(),
			},
			{
				Name:  "UNIV2:DEGEN-ETH LP Staking",
				Value: bclient.DEGENUNILPStakingAddress.String(),
			},
			{
				Name:  "Governor Alpha",
				Value: bclient.GovernorAlpha.String(),
			},
		}
	default:
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "filters",
				Value: "all, staking, indexes, governance, uniswap-pairs",
			},
		}
		embed.Description = "invalid filter given, supported filters are listed below"
	}
	ctx.RespondEmbed(embed)
}
