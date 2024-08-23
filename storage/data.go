package storage

import (
	"errors"
	"sync"
)

type KV struct {
	Store map[string]string
	mutex sync.Mutex
}

func NewKV() *KV {
	return &KV{Store: make(map[string]string)}
}

type KeyValueDB interface {
	Set(key, value string)
	Get(key string) (string, error)
}

func (k *KV) Set(key, value string) {
	k.mutex.Lock()
	k.Store[key] = value
	k.mutex.Unlock()

}

func (k *KV) Get(key string) (string, error) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	value, ok := k.Store[key]
	if !ok {
		return "", errors.New("error in finding the key")
	}
	return value, nil
}
