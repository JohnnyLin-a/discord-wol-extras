package commands

import (
	"github.com/johnnylin-a/discord-wol-extras/internal/commands/handlers"
)

var CommandList = map[string]handlers.CommandHandler{}

func init() {
	CommandList[handlers.Ping.ApplicationCommand.Name] = handlers.Ping
}
