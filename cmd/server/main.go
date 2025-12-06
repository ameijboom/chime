package main

import (
	"flag"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID        = flag.String("guild", "", "Guild ID for registering slash commands.")
	Token          = flag.String("token", "", "Discord Access Token.")
	RemoveCommands = flag.Bool("remove-commands", false, "Removes commands upon shutdown")
)

var s *discordgo.Session

func init() {
	flag.Parse()
}

func init() {
	var err error

	s, err = discordgo.New("Bot " + *Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {

}
