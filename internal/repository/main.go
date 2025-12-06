package repository

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ErrEventFull         = errors.New("event is full")
	ErrUserAlreadyMember = errors.New("user is already a member of this event")
)

type Repository struct {
	db *gorm.DB
}

func NewDatabase(databaseUrl string) (*Repository, error) {
	db, err := gorm.Open(sqlite.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db,
	}, nil
}

func (r *Repository) RunMigrations() error {
	return r.db.AutoMigrate(Models)
}

func (r *Repository) GetEventByMessageID(messageID uint64) (*Event, error) {
	var event Event
	if err := r.db.Preload("Members").Where("message_id = ?", messageID).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *Repository) AddMemberToEvent(member *Attendee, messageID uint64) error {
	event, err := r.GetEventByMessageID(messageID)
	if err != nil {
		return err
	}

	if len(event.Members) >= event.MaxSize() {
		return ErrEventFull
	}

	member.EventID = event.ID
	event.Members = append(event.Members, *member)

	return r.db.Save(&event).Error
}

func (r *Repository) CreateEvent(e *Event) error {
	return r.db.Create(e).Error
}

func (r *Repository) UpdateEvent(e *Event) error {
	return r.db.Save(e).Error
}

func (r *Repository) UpdateEventStatus(messageID uint64, status EventStatus) error {
	return r.db.Model(&Event{}).Where("message_id = ?", messageID).Update("status", status).Error
}

func (r *Repository) DeleteEvent(messageID uint64) error {
	return r.db.Where("message_id = ?", messageID).Delete(&Event{}).Error
}

func (r *Repository) GetEventByGuildEventID(guildEventID uint64) (*Event, error) {
	var event Event
	if err := r.db.Preload("Members").Where("guild_event_id = ?", guildEventID).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *Repository) GetAllEvents() ([]*Event, error) {
	var events []*Event
	if err := r.db.Preload("Members").Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (r *Repository) GetEventsByStatus(status EventStatus) ([]*Event, error) {
	var events []*Event
	if err := r.db.Preload("Members").Where("status = ?", status).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (r *Repository) RemoveMemberFromEvent(messageID uint64, userId uint64) error {
	event, err := r.GetEventByMessageID(messageID)
	if err != nil {
		return err
	}

	return r.db.Where("event_id = ? AND user_id = ?", event.ID, userId).Delete(&Attendee{}).Error
}

func (r *Repository) GetMemberRole(messageID uint64, userId uint64) (*JobRole, error) {
	event, err := r.GetEventByMessageID(messageID)
	if err != nil {
		return nil, err
	}

	var member Attendee
	if err := r.db.Where("event_id = ? AND user_id = ?", event.ID, userId).First(&member).Error; err != nil {
		return nil, err
	}

	return member.Role, nil
}

func (r *Repository) UpdateMemberRole(messageID uint64, userId uint64, role *JobRole) error {
	event, err := r.GetEventByMessageID(messageID)
	if err != nil {
		return err
	}

	return r.db.Model(&Attendee{}).Where("event_id = ? AND user_id = ?", event.ID, userId).Update("role", role).Error
}

func (r *Repository) IsMemberInEvent(messageID uint64, userId uint64) (bool, error) {
	event, err := r.GetEventByMessageID(messageID)
	if err != nil {
		return false, err
	}

	var count int64
	if err := r.db.Model(&Attendee{}).Where("event_id = ? AND user_id = ?", event.ID, userId).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
