package models

import "errors"

// Key value strore struct
type InMemory struct {
	Data map[string]string `json:"data"`
}

type StorageSetter interface {
	Store(string, string)
}

type StorageGetter interface {
	ListRecord() map[string]string
	GetValue(string) (string, error)
}

type StorageDestroyer interface {
	DeleteRecord(string) error
}

type Storage interface {
	StorageSetter
	StorageGetter
	StorageDestroyer
}

func NewRecord() *InMemory {
	return &InMemory{Data: make(map[string]string, 2)}
}

func (r *InMemory) Store(key, val string) {
	r.Data[key] = val
}

func (r *InMemory) ListRecord() map[string]string {
	return r.Data
}

func (r *InMemory) GetValue(key string) (string, error) {
	val, ok := r.Data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return val, nil
}

func (r *InMemory) DeleteRecord(key string) error {
	_, ok := r.Data[key]
	if !ok {
		return errors.New("key not found")
	}
	delete(r.Data, key)
	return nil
}
