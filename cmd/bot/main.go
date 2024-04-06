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
	// Prepare discord session
	var err error
	s, err = discordgo.New("Bot " + data.ConfigDatabase.DiscordToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	stopProcessingInteractions := s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if data.ConfigDatabase.DiscordGuildID == i.GuildID {
			if h, ok := commands.CommandList[i.ApplicationCommandData().Name]; ok {
				h.Handler(s, i)
			} else {
				handlers.Placeholder.Handler(s, i)
			}
		}
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		s.ChannelMessageSend(data.ConfigDatabase.DiscordMsgChannelID, "Bot Online!")
	})

	// Connect to discord APIs
	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Service started.")
	log.Println("Press Ctrl+C to exit")
	<-stop
	log.Println("Gracefully stopping...")
	stopProcessingInteractions()
	log.Println("Done cleanup! Shutting down...")

}
