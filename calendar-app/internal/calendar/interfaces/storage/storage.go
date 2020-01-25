package storage

// Storage interface for storing data
type Storage interface {
	AddItem(item interface{}) error
	DeleteItem(id string) error
	ChangeItem(id string, args map[string]string) error
	GetItem(id string) (interface{}, error)
	GetItems() []interface{}
}
