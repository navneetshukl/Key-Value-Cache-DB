package storage

import (
	"errors"
	"time"
)

func NewKV() *KV {
	return &KV{Store: make(map[string]Value)}
}

func (k *KV) Set(key, value string) {
	k.mutex.Lock()
	val := Value{
		Value: value,
		Time:  time.Now(),
	}
	k.Store[key] = val
	k.mutex.Unlock()

}

func (k *KV) Get(key string) (string, error) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	value, ok := k.Store[key]
	if !ok {
		return "", errors.New("error in finding the key")
	}
	return value.Value, nil
}

func (k *KV) CheckTTL() {
	for {

		for key, value := range k.Store {
			if time.Since(value.Time) > 20*time.Second {
				delete(k.Store, key)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}
