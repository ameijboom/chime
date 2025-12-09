package server

import (
	"context"
	"strconv"

	"github.com/ameijboom/chime/generated"
	"github.com/ameijboom/chime/internal/bot"
	"github.com/ameijboom/chime/internal/bot/handlers"
	"github.com/ameijboom/chime/internal/repository"
)

type Server struct {
	generated.UnimplementedEventsServer
	repo          *repository.Repository
	bot           *bot.Bot
	eventsHandler *handlers.EventsHandler
}

func NewServer(repo *repository.Repository, discordBot *bot.Bot) *Server {
	return &Server{
		repo:          repo,
		bot:           discordBot,
		eventsHandler: &handlers.EventsHandler{},
	}
}

func (s *Server) CreateScheduledEvent(ctx context.Context, req *generated.CreateScheduledEventReqeuest) (*generated.CreateScheduledEventResponse, error) {
	// Call Discord bot handler to create the scheduled event
	discordEvent, err := s.eventsHandler.CreateScheduledEvent(
		s.bot.GetSession(),
		s.bot.GetGuildID(),
		req,
	)
	if err != nil {
		return nil, err
	}

	// Parse the Discord event ID as uint64
	guildEventID, err := strconv.ParseUint(discordEvent.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	// Convert protobuf types to repository types
	var partyType *repository.PartyType
	if req.Type != nil {
		pt := protoToPartyType(*req.Type)
		partyType = &pt
	}

	// Create the event in the database
	event := &repository.Event{
		GuildEventID: guildEventID,
		Status:       protoToEventStatus(req.Status),
		PartyType:    partyType,
	}

	if err := s.repo.CreateEvent(event); err != nil {
		return nil, err
	}

	return &generated.CreateScheduledEventResponse{
		ID: discordEvent.ID,
	}, nil
}

func protoToEventStatus(status generated.ScheduledEventStatus) repository.EventStatus {
	switch status {
	case generated.ScheduledEventStatus_ONGOING:
		return repository.Ongoing
	case generated.ScheduledEventStatus_ENDED:
		return repository.Ended
	case generated.ScheduledEventStatus_ONHOLD:
		return repository.OnHold
	default:
		return repository.Ongoing
	}
}

func protoToPartyType(partyType generated.ScheduledEventPartyType) repository.PartyType {
	switch partyType {
	case generated.ScheduledEventPartyType_LIGHT:
		return repository.Light
	case generated.ScheduledEventPartyType_FULL:
		return repository.Full
	default:
		return repository.Full
	}
}
