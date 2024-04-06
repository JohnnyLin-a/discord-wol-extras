package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/johnnylin-a/discord-wol-extras/internal/commands"
	"github.com/johnnylin-a/discord-wol-extras/internal/commands/handlers"
	"github.com/johnnylin-a/discord-wol-extras/pkg/data"
)

var s *discordgo.Session

func init() {
	cleanenv.ReadEnv(&data.ConfigDatabase)
	cleanenv.ReadConfig(".env", &data.ConfigDatabase)
}

func main() {
	// connect discord session
	var err error
	s, err = discordgo.New("Bot " + data.ConfigDatabase.DiscordToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	stopProcessingInteractions := s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Println("Interaction from", i.GuildID, i.Member.User.ID)
		log.Println("local values:", data.ConfigDatabase.DiscordGuildID)
		if data.ConfigDatabase.DiscordGuildID == i.GuildID {
			if h, ok := commands.CommandList[i.ApplicationCommandData().Name]; ok {
				h.Handler(s, i)
			} else {
				handlers.Placeholder.Handler(s, i)
			}
		}
	})

	s.ChannelMessageSend(data.ConfigDatabase.DiscordMsgChannelID, "Bot Online!")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Service started.")
	log.Println("Press Ctrl+C to exit")
	<-stop
	log.Println("Gracefully stopping...")
	stopProcessingInteractions()
	log.Println("Done cleanup! Shutting down...")

}
