package calendar

import (
	"fmt"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/interfaces/storage"
	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/models"
	"github.com/Cornpop456/otus-go/calendar-app/utils"
)

// Calendar represents basic calendar
type Calendar struct {
	eventsStorage storage.Storage
	eventsNumber  int
}

// New returns new calendar
func New(storage storage.Storage) *Calendar {
	return &Calendar{storage, 0}
}

// AddEvent adds new event to calendar
func (c *Calendar) AddEvent(name string, desc string, date time.Time) error {
	uuid, err := utils.GenetateUUID()

	if err != nil {
		return err
	}

	dateStruct := &models.Date{
		Year:     date.Format("2006"),
		Month:    date.Format("January"),
		Day:      date.Format("Monday"),
		Time:     date.Format("15:04"),
		Timezone: date.Format("MST"),
	}

	event := &models.Event{
		ID:          uuid,
		Name:        name,
		Description: desc,
		EventDate:   dateStruct,
		RawDate:     &date,
	}

	if err := c.eventsStorage.AddItem(event); err != nil {
		return err
	}
	c.eventsNumber++
	fmt.Println("New event was added to calendar")
	return nil
}

// DeleteEvent deletes event from calendar
func (c *Calendar) DeleteEvent(eventID string) error {
	if err := c.eventsStorage.DeleteItem(eventID); err != nil {
		return err
	}
	c.eventsNumber--
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
func (c Calendar) GetEvent(eventID string) (*models.Event, error) {
	event, err := c.eventsStorage.GetItem(eventID)
	if err != nil {
		return nil, err
	}
	return event.(*models.Event), nil
}

// GetEvents get all events
func (c Calendar) GetEvents() []*models.Event {
	res := make([]*models.Event, 0, c.eventsNumber)
	for _, v := range c.eventsStorage.GetItems() {
		event := v.(*models.Event)
		res = append(res, event)
	}
	return res
}
