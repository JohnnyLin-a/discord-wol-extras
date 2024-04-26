package commands

import (
	"github.com/johnnylin-a/discord-wol-extras/internal/commands/handlers"
)

var CommandList = map[string]handlers.CommandHandler{}

func init() {
	enabledCommands := []handlers.CommandHandler{
		handlers.Ping,
		handlers.PowerOn,
		handlers.PowerOffWithRCONSafety,
		handlers.RestartRCON,
		handlers.Status,
	}
	for _, v := range enabledCommands {
		CommandList[v.ApplicationCommand.Name] = v
	}
}
