package commands

import (
	"github.com/bwmarrin/discordgo"
)

var Registry []Command

func ProcessCommandRegistry(s *discordgo.Session, guildID string) ([]*discordgo.ApplicationCommand, error) {
	registered := make([]*discordgo.ApplicationCommand, len(Registry))
	handlers := make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

	for index, cmd := range Registry {
		definition := cmd.Definition()

		reg, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, definition)
		if err != nil {
			return nil, err
		}

		registered[index] = reg
		handlers[definition.Name] = cmd.Handler()
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := handlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	return registered, nil
}

type Command interface {
	Definition() *discordgo.ApplicationCommand
	Handler() func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// Register adds a command to the registry
func Register(cmd Command) {
	Registry = append(Registry, cmd)
}
