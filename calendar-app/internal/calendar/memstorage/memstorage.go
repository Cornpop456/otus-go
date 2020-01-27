package memstorage

import (
	"sync"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/errors"
	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/models"
	"github.com/Cornpop456/otus-go/calendar-app/internal/pkg/utils"
)

// EventsLocalStorage implements Storage interface
type EventsLocalStorage struct {
	events map[string]*models.Event
	mux    sync.Mutex
}

// NewEventsLocalStorage returns EventsLocalStorage
func NewEventsLocalStorage() *EventsLocalStorage {
	return &EventsLocalStorage{
		events: make(map[string]*models.Event),
	}
}

// AddItem adds item to storage implementation
func (s *EventsLocalStorage) AddItem(event models.Event) (string, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	uuid, err := utils.GenetateUUID()

	if err != nil {
		return "", err
	}

	if _, ok := s.events[uuid]; ok {
		return "", errors.ErrSameID
	}

	event.ID = uuid

	for _, v := range s.events {
		if event.RawDate.String() == v.RawDate.String() {
			return "", errors.ErrSameTime
		}
	}

	s.events[uuid] = &event
	return uuid, nil
}

// DeleteItem deletes item from storage
func (s *EventsLocalStorage) DeleteItem(id string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.events[id]; !ok {
		return errors.ErrEventNotFound
	}

	delete(s.events, id)
	return nil
}

// ChangeItem changes item in storage
func (s *EventsLocalStorage) ChangeItem(id string, args map[string]string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	event, ok := s.events[id]

	if !ok {
		return errors.ErrEventNotFound
	}

	if newDateString, ok := args["Date"]; ok {
		rawDate, err := time.Parse(time.RFC822, newDateString)

		if err != nil {
			return err
		}

		emptyDateModel := models.Date{}
		dateModel, err := emptyDateModel.Parse(newDateString)

		if err != nil {
			return err
		}

		for _, val := range s.events {
			if rawDate.String() == val.RawDate.String() {
				return errors.ErrSameTime
			}
		}

		event.EventDate = dateModel
		event.RawDate = rawDate
	}
	if newName, ok := args["Name"]; ok {
		event.Name = newName
	}
	if newDesc, ok := args["Description"]; ok {
		event.Description = newDesc
	}
	return nil
}

// GetItem gets item from storage
func (s *EventsLocalStorage) GetItem(id string) (*models.Event, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	event, ok := s.events[id]

	if !ok {
		return nil, errors.ErrEventNotFound
	}

	return event, nil
}

// GetItems gets all items from storage
func (s *EventsLocalStorage) GetItems() []*models.Event {
	s.mux.Lock()
	defer s.mux.Unlock()

	res := make([]*models.Event, 0, len(s.events))
	for _, v := range s.events {
		res = append(res, v)
	}

	return res
}
