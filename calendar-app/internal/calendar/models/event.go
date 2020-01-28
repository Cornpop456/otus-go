package models

import "time"

// Event represents event type for calendar
type Event struct {
	ID          string
	Name        string
	Description string
	EventDate   *Date
	RawDate     time.Time
}

// Date for calendar
type Date struct {
	Year     string
	Month    string
	Day      string
	Time     string
	Timezone string
}

func (d *Date) String() string {
	return d.Year + ":" + d.Month + ":" + d.Day + ":" + d.Time + ":" + d.Timezone
}

// ParseDateFromString returns *Date from date string
func ParseDateFromString(dateString string) (*Date, error) {
	date, err := time.Parse(time.RFC822, dateString)

	if err != nil {
		return nil, err
	}

	d := &Date{}

	d.Year = date.Format("2006")
	d.Month = date.Format("January")
	d.Day = date.Format("Monday")
	d.Time = date.Format("15:04")
	d.Timezone = date.Format("MST")

	return d, nil
}
