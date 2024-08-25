package storage

import (
	"sync"
	"time"
)

type Value struct {
	Value string
	Time  time.Time
}

type KV struct {
	Store map[string]Value
	mutex sync.Mutex
}

type KeyValueDB interface {
	Set(key, value string)
	Get(key string) (string, error)
	CheckTTL()
}
