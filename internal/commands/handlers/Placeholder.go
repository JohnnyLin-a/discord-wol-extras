package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var Placeholder = CommandHandler{
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Command under construction",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	},
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "Placeholder",
		Description: "Placeholder for nonexistant or under-construction commands",
	},
}
