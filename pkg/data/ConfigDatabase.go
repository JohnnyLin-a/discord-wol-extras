package data

type TConfigDatabase struct {
	DiscordToken                string `env:"DISCORD_TOKEN"`
	DiscordGuildID              string `env:"DISCORD_GUILD_ID"`
	DiscordMsgChannelID         string `env:"DISCORD_MSG_CHANNEL_ID"`
	TruenasHost                 string `env:"TRUENAS_HOST"`
	TruenasAPIKey               string `env:"TRUENAS_API_KEY"`
	TruenasIsHTTPS              string `env:"TRUENAS_IS_HTTPS"`
	TruenasIsMismatchCertForTLS string `env:"TRUENAS_IS_MISMATCH_CERT_FOR_TLS"`
	WolMAC                      string `env:"WOL_MAC"`
	WolBroadcastIP              string `env:"WOL_BROADCAST_IP"`
}

var ConfigDatabase TConfigDatabase
