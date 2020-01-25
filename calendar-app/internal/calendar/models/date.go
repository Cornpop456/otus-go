package models

// Date for calendar
type Date struct {
	Year     string
	Month    string
	Day      string
	Time     string
	Timezone string
}

func (d Date) String() string {
	return d.Year + ":" + d.Month + ":" + d.Day + ":" + d.Time + ":" + d.Timezone
}
