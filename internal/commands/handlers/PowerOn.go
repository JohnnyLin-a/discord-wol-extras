package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/johnnylin-a/discord-wol-extras/pkg/data"
	"github.com/johnnylin-a/go-wol-lib/pkg/wol"
)

var PowerOn = CommandHandler{
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		content := "Failed to power on..."
		err := wol.Wake(data.ConfigDatabase.WolMAC, data.ConfigDatabase.WolBroadcastIP+":9")
		if err == nil {
			content = "Powered on."
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	},
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "poweron",
		Description: "Power on the server",
	},
}
