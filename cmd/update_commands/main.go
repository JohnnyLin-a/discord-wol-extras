package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/johnnylin-a/discord-wol-extras/internal/commands"
	"github.com/johnnylin-a/discord-wol-extras/pkg/data"
)

var s *discordgo.Session

func init() {
	cleanenv.ReadEnv(&data.ConfigDatabase)
	cleanenv.ReadConfig(".env", &data.ConfigDatabase)

	var err error
	s, err = discordgo.New("Bot " + data.ConfigDatabase.DiscordToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	for _, v := range commands.CommandList {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, data.ConfigDatabase.DiscordGuildID, v.ApplicationCommand)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.ApplicationCommand.Name, err)
		}
		log.Println("Added command " + v.ApplicationCommand.Name)
	}
	s.Close()
}
