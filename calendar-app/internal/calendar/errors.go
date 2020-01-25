package calendar

import "errors"

var (
	// ErrEventNotFound error occupies when looking event is not exist
	ErrEventNotFound = errors.New("event is not found")
	// ErrSameTime error occupies when try adding an event with the same time
	ErrSameTime = errors.New("time is occupied by another event")
)
