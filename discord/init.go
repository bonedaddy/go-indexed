package discord

import "github.com/bwmarrin/discordgo"

func init() {
	helpEmbed = &discordgo.MessageEmbed{
		Title:       "Indexed Finance Bot Help Menu",
		Description: "NDXBot is an unofficial Discord Bot for making read-only calls to Indexed Finance smart contracts.\n\nAll commands must be invoked with !ndx <cmd>\n\nAnything with <..> after command name expects an argument\n\nAnything with [..] after command name is an optional argument\n\nExample (stake-earned)\n!ndx stake-earned defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13\n\nExample (pool-balance)\n!ndx pool-balance defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13\n\nAll available commands are listed below",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			// URL: "https://pbs.twimg.com/profile_images/1342395531318976518/kIv5abLc_400x400.jpg",
			URL: "https://gateway.temporal.cloud/ipfs/QmSky8KsZ6q9zz6kmj3TrbNTTvwhGtGmVbYB9iXWLPD5VC",
		},
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "pool-balance <pool-name> <account-address>",
				Value: "returns the index pool balance for a given pool held by account-address",
			},
			&discordgo.MessageEmbedField{
				Name:  "pool-tokens <pool-name>",
				Value: "returns the current tokens basketed in the pool",
			},
			&discordgo.MessageEmbedField{
				Name:  "stake-earned <stake-type> <account-address>",
				Value: "returns the amount of staking rewards earned, only supported stake-type params are defi5 and univ2-eth-defi5",
			},
			&discordgo.MessageEmbedField{
				Name:  "notify <command> <args...> [-interval]",
				Value: "enable notifying user of certain conditions, run `!ndx notify help` for more information",
			},
			&discordgo.MessageEmbedField{
				Name:  "uniswap <command> <args...>",
				Value: "enable querying uniswap pairs, run `!ndx uniswap help` for more information",
			},
		},
	}
	notifyHelpEmbed = &discordgo.MessageEmbed{
		Title:       "Indexed Finance Notify Command Help",
		Description: "The notify command allow you to schedule notifications about arbitrary conditions\n\nEach command accepts a final optional parameter -interval=SECONDS which specifies the frequency in seconds to notify the user which defaults to 60 seconds if not specified\n",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			// URL: "https://pbs.twimg.com/profile_images/1342395531318976518/kIv5abLc_400x400.jpg",
			URL: "https://gateway.temporal.cloud/ipfs/QmSky8KsZ6q9zz6kmj3TrbNTTvwhGtGmVbYB9iXWLPD5VC",
		},
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "help",
				Value: "display this help menu",
			},
			&discordgo.MessageEmbedField{
				Name:  "stake-earned <pool-name> <account-address>",
				Value: "enable notifying the user about their earned staked rewards\n\nexample: `!ndx notify stake-earned defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13 -interval=60`",
			},
		},
	}
	uniswapHelpEmbed = &discordgo.MessageEmbed{
		Title:       "Indexed Finance Uniswap Command Help",
		Description: "The uniswap command allows you to query indexed finance uniswap pairs",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			// URL: "https://pbs.twimg.com/profile_images/1342395531318976518/kIv5abLc_400x400.jpg",
			URL: "https://gateway.temporal.cloud/ipfs/QmSky8KsZ6q9zz6kmj3TrbNTTvwhGtGmVbYB9iXWLPD5VC",
		},
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "help",
				Value: "display this help menu",
			},
			&discordgo.MessageEmbedField{
				Name:  "exchange-amount <pair> <amount>",
				Value: "retrieve the exchange amount for the given pair. `!ndx uniswap exchange-amount defi5-eth 1` will return the amount of ETH received for swapping 1 DEFI5 to ETH",
			},
		},
	}
}
