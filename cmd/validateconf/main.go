package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/johnnylin-a/discord-wol-extras/pkg/data"
	"github.com/johnnylin-a/discord-wol-extras/pkg/utils"
)

func init() {
	cleanenv.ReadEnv(&data.ConfigDatabase)
	cleanenv.ReadConfig(".env", &data.ConfigDatabase)
}

func main() {
	err := utils.ValidateConf(data.ConfigDatabase)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("OK")
}
