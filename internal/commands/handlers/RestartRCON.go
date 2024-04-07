package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
	"github.com/johnnylin-a/discord-wol-extras/pkg/data"
)

var RestartRCON = CommandHandler{
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		conn, err := rcon.Dial(data.ConfigDatabase.RCONHost+":"+data.ConfigDatabase.RCONPort, data.ConfigDatabase.RCONPassword)
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Failed to connect to RCON Server",
				},
			})
			return
		}
		defer conn.Close()
		response, err := conn.Execute("shutdown 10") // TODO MAYBE, customize this for non-PalWorld servers
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Failed to send shutdown command",
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
	},
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "restartserver",
		Description: "Restart and update server without rebooting machine",
	},
}
