package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var Ping = CommandHandler{
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "I am alive. You are <@" + i.Member.User.ID + ">, who joined on " + i.Member.JoinedAt.Format(time.RFC822) + ".\nThis bot was created by .azuri",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	},
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping test for this bot",
	},
}
