package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
	"github.com/johnnylin-a/discord-wol-extras/pkg/data"
)

var Status = CommandHandler{
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		conn, err := rcon.Dial(data.ConfigDatabase.RCONHost+":"+data.ConfigDatabase.RCONPort, data.ConfigDatabase.RCONPassword)
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error did not dial RCON successfully",
				},
			})
			return
		}
		defer conn.Close()
		response, err := conn.Execute("ShowPlayers")
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error did not send server command successfully",
				},
			})
			return
		}
		if !strings.HasPrefix(response, "name,playeruid,steamid") {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error RCON server sent wrong response.\nGot:\n" + response,
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "RCON response successful",
			},
		})

	},
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "status",
		Description: "Check if RCON server is alive",
	},
}
