package kvstore

import (
	"sync"
)

//KVStore is a interface for key value store that can store value of any type
type KVStore interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
	Delete(string)
}

type kvData struct {
	sync.Mutex
	data map[string]interface{}
}

func (store *kvData) Get(key string) (interface{}, bool) {
	store.Lock()
	value, err := store.data[key]
	store.Unlock()
	return value, err
}

func (store *kvData) Set(key string, value interface{}) {
	store.Lock()
	store.data[key] = value
	store.Unlock()
}

func (store *kvData) Delete(key string) {
	store.Lock()
	delete(store.data, key)
	store.Unlock()
}

//NewKVStore function to create new KVStore
func NewKVStore() KVStore {
	return &kvData{data: make(map[string]interface{})}
}
