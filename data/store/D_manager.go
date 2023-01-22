package store

type Store interface {
	SetKey(key string, value string) error
	GetKey(key string) (string, error)
}

type DataManager interface {
	Retrieve(input interface{}) (out interface{}, err error)
}
