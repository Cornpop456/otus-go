package models

import "time"

// Event represents event type for calendar
type Event struct {
	ID          string
	Name        string
	Description string
	EventDate   *Date
	RawDate     *time.Time
}
