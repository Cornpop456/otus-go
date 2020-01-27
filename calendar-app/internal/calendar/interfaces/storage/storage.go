package storage

import "github.com/Cornpop456/otus-go/calendar-app/internal/calendar/models"

// EventsStorage interface for storing events
type EventsStorage interface {
	AddItem(event models.Event) (string, error)
	DeleteItem(id string) error
	ChangeItem(id string, args map[string]string) error
	GetItem(id string) (*models.Event, error)
	GetItems() []*models.Event
}
