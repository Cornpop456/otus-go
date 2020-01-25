package memstorage

import (
	"sync"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/models"
)

// EventsLocalStorage implements Storage interface
type EventsLocalStorage struct {
	events []*models.Event
	mux    sync.Mutex
}

// NewEventsLocalStorage returns EventsLocalStorage
func NewEventsLocalStorage() *EventsLocalStorage {
	return &EventsLocalStorage{
		events: make([]*models.Event, 0, 20),
	}
}

// AddItem adds item to storage implementation
func (s *EventsLocalStorage) AddItem(item interface{}) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	event, ok := item.(*models.Event)
	if !ok {
		return ErrEventNotFound
	}

	for _, v := range s.events {
		if event.EventDate.String() == v.EventDate.String() {
			return ErrSameTime
		}
	}

	s.events = append(s.events, event)
	return nil
}

// DeleteItem deletes item from storage
func (s *EventsLocalStorage) DeleteItem(id string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	for i, v := range s.events {
		if id == v.ID {
			s.events[i] = s.events[len(s.events)-1]
			s.events[len(s.events)-1] = nil
			s.events = s.events[:len(s.events)-1]
			return nil
		}
	}
	return ErrEventNotFound
}

// ChangeItem changes item in storage
func (s *EventsLocalStorage) ChangeItem(id string, args map[string]string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, v := range s.events {
		if newDateString, ok := args["Date"]; ok {
			rawDate, err := time.Parse(time.RFC822, newDateString)
			if err != nil {
				return err
			}

			date := &models.Date{
				Year:  rawDate.Format("2006"),
				Month: rawDate.Format("Jan"),
				Day:   rawDate.Format("Mon"),
				Time:  rawDate.Format("15:04:05"),
			}

			for _, v := range s.events {
				if date.String() == v.EventDate.String() {
					return ErrSameTime
				}

				v.EventDate = date
				v.RawDate = &rawDate
			}
		}
		if id == v.ID {
			if newName, ok := args["Name"]; ok {

				v.Name = newName
			}
			if newDesc, ok := args["Description"]; ok {
				v.Description = newDesc
			}
			return nil
		}
	}
	return ErrEventNotFound
}

// GetItem gets item from storage
func (s *EventsLocalStorage) GetItem(id string) (interface{}, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, v := range s.events {
		if id == v.ID {
			return v, nil
		}
	}
	return nil, ErrEventNotFound
}

// GetItems gets all items from storage
func (s *EventsLocalStorage) GetItems() []interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()

	res := make([]interface{}, 0, len(s.events))
	for _, v := range s.events {
		res = append(res, v)
	}
	return res
}
