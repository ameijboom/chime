package handlers

import (
	"github.com/ameijboom/chime/generated"
	"github.com/bwmarrin/discordgo"
)

type EventsHandler struct{}

func (h *EventsHandler) CreateScheduledEvent(s *discordgo.Session, g string, e *generated.CreateScheduledEventReqeuest) (*discordgo.GuildScheduledEvent, error) {
	start, end := e.StartTime.AsTime(), e.EndTime.AsTime()

	event, err := s.GuildScheduledEventCreate(g, &discordgo.GuildScheduledEventParams{
		Name:               e.Title,
		Description:        e.Description,
		ScheduledStartTime: &start,
		ScheduledEndTime:   &end,
		EntityType:         discordgo.GuildScheduledEventEntityTypeVoice,
		ChannelID:          "1436078067712790611",
		PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
	})
	if err != nil {
		return nil, err
	}

	return event, nil
}
