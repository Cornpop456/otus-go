package calendar

import (
	"fmt"
	"sync"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/interfaces/storage"
	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/models"
)

// Calendar represents basic calendar
type Calendar struct {
	eventsStorage storage.EventsStorage
	eventsNumber  int
	mux           sync.Mutex
}

// New returns new calendar
func New(storage storage.EventsStorage) *Calendar {
	return &Calendar{eventsStorage: storage, eventsNumber: 0}
}

// AddEvent adds new event to calendar
func (c *Calendar) AddEvent(name string, desc string, date time.Time) (string, error) {
	dateModel, err := models.ParseDateFromString(date.Format(time.RFC822))

	if err != nil {
		return "", err
	}

	event := models.Event{
		Name:        name,
		Description: desc,
		EventDate:   dateModel,
		RawDate:     date,
	}

	uuid, err := c.eventsStorage.AddItem(event)

	if err != nil {
		return "", err
	}

	c.mux.Lock()
	c.eventsNumber++
	c.mux.Unlock()

	fmt.Println("New event was added to calendar")
	return uuid, nil
}

// DeleteEvent deletes event from calendar
func (c *Calendar) DeleteEvent(eventID string) error {
	if err := c.eventsStorage.DeleteItem(eventID); err != nil {
		return err
	}

	c.mux.Lock()
	c.eventsNumber--
	c.mux.Unlock()

	fmt.Println("Event was deleted from calendar")
	return nil
}

// ChangeEvent changes event in calendar
func (c *Calendar) ChangeEvent(eventID string, args map[string]string) error {
	if err := c.eventsStorage.ChangeItem(eventID, args); err != nil {
		return err
	}
	fmt.Println("Event was changed in calendar")
	return nil
}

// GetEvent gets single event
func (c *Calendar) GetEvent(eventID string) (*models.Event, error) {
	event, err := c.eventsStorage.GetItem(eventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// GetEvents get all events
func (c *Calendar) GetEvents() []*models.Event {
	return c.eventsStorage.GetItems()
}
