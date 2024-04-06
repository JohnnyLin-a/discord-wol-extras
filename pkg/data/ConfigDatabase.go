package data

type TConfigDatabase struct {
	DiscordToken        string `env:"DISCORD_TOKEN"`
	DiscordGuildID      string `env:"DISCORD_GUILD_ID"`
	DiscordMsgChannelID string `env:"DISCORD_MSG_CHANNEL_ID"`
}

var ConfigDatabase TConfigDatabase
