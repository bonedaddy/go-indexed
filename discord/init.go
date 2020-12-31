package discord

import "github.com/bwmarrin/discordgo"

func init() {
	helpEmbed = &discordgo.MessageEmbed{
		Title:       "Indexed Finance Bot Help Menu",
		Description: "All commands must be invoked with !ndx <cmd>\n\nAnything with <..> after command name expects an argument\n\nAnything with [..] after command name is an optional argument\n\nExample (stake-earned)\n!ndx stake-earned defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13\n\nExample (pool-balance)\n!ndx pool-balance defi5 0x5a361A1dfd52538A158e352d21B5b622360a7C13\n\nAll available commands are listed below",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://pbs.twimg.com/profile_images/1342395531318976518/kIv5abLc_400x400.jpg",
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
				Value: "enable notifying user of certain conditions, only supported command is stake-earned. -interval specifies the notification frequency in seconds, and will notify up to a maximum of 10 tiems",
			},
		},
	}
}
