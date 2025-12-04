package database

import (
	"gorm.io/gorm"
)

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
	Id        uint64
	MessageId uint64

	Status    EventStatus
	PartyType *PartyType

	Members []EventMember `gorm:"foreignKey:EventID"`
}

type EventMember struct {
	gorm.Model
	EventID uint64
	UserID  uint64
	Role    *JobRole
}
