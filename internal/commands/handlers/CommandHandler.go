package handlers

import "github.com/bwmarrin/discordgo"

type CommandHandler struct {
	Handler            func(s *discordgo.Session, i *discordgo.InteractionCreate)
	ApplicationCommand *discordgo.ApplicationCommand
}
