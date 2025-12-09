package commands

import (
	"github.com/ameijboom/chime/internal/bot/handlers"
	"github.com/bwmarrin/discordgo"
)

type EventsCommand struct{}

var handler handlers.EventsHandler

func (c *EventsCommand) Definition() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "events",
		Description: "Manage FC Events in Discord",
	}
}

func (c *EventsCommand) Handler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// data := repository.Event{
	// 	GuildEventID: 0,
	// 	MessageID:    0,
	// 	Status:       repository.Ongoing,
	// 	PartyType:    nil,
	// 	Members:      make([]repository.Attendee, 0),
	// }

	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// start, err := time.Parse(time.RFC3339, "2026-01-01T12:00:00+01:00")
		// if err != nil {
		// 	log.Fatalf("Encountered error while parsing start time: %v", err)
		// }
		// end, err := time.Parse(time.RFC3339, "2026-01-01T13:00:00+01:00")
		// if err != nil {
		// 	log.Fatalf("Encountered error while parsing end time: %v", err)
		// }

		// event, err := handler.CreateScheduledEvent(s, i.GuildID, handlers.NewEvent(&data, "Suki's Event created with Chime!", "This is a test event that can be safely ignored! I really hope this doesn't actually send a notification to everyone in the server now that I think about it...", start, end))
		// if err != nil {
		// 	log.Fatalf("Failed to create Guild Event: %v", err)
		// }

		// s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// 	Type: discordgo.InteractionResponseChannelMessageWithSource,
		// 	Data: &discordgo.InteractionResponseData{
		// 		Content: "I've scheduled the following Scheduled Event!\n" + "[" + event.ID + "]->\n" + event.Name + "\n" + event.Description,
		// 	},
		// })
	}
}

func init() {
	Register(&EventsCommand{})
}
