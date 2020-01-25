package storage

// MemoryStorage for storing data in memory
type MemoryStorage interface {
	AddItem(item interface{}) error
	DeleteItem(id string) error
	ChangeItem(id string, args map[string]string) error
	GetItem(id string) (interface{}, error)
	GetItems() []interface{}
}
