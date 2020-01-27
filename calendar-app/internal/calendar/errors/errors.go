package errors

import "errors"

var (
	// ErrEventNotFound error occupies when looking event is not exist
	ErrEventNotFound = errors.New("event is not found")
	// ErrSameTime error occupies when try adding an event with the same time
	ErrSameTime = errors.New("time is occupied by another event")
	// ErrSameID error occupies when try adding an event with id which already in storage
	ErrSameID = errors.New("event with this id is already exists")
)
