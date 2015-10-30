package storage

import "errors"

type status struct {
	Key   string
	URL   string
	Count int
}

var (
	// ErrNotFound when key is not found in the store
	ErrNotFound = errors.New("Not Found")
)

// Database is an interface used for the backend for storing url and counts to a database
type Database interface {
	Get(key string) (status, error)
	Set(url string) (status, error)
	Incr(Key string)
}
