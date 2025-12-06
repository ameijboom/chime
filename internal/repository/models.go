package repository

import (
	"gorm.io/gorm"
)

var Models = []any{
	&Event{},
	&Attendee{},
}

type EventStatus string
type PartyType string
type JobRole string

const (
	Ongoing EventStatus = "ongoing"
	Ended   EventStatus = "ended"
	OnHold  EventStatus = "onhold"

	Light PartyType = "light"
	Full  PartyType = "full"

	Tank   JobRole = "tank"
	Healer JobRole = "healer"
	Dps    JobRole = "dps"
)

type Event struct {
	gorm.Model
	GuildEventID uint64
	MessageID    uint64

	Status    EventStatus
	PartyType *PartyType

	Members []Attendee `gorm:"foreignKey:EventID"`
}

type Attendee struct {
	gorm.Model
	EventID uint
	UserID  uint64
	Role    *JobRole
}
