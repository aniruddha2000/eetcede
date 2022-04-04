package models

import "errors"

// Key value strore struct
type Record struct {
	Data map[string]string `json:"data"`
}

type StorageSetter interface {
	Store(string, string)
}

type StorageGetter interface {
	FindAllRecords() map[string]string
	FindRecord(string) (string, error)
}

type StorageDestroyer interface {
	DeleteRecordByKey(string) error
}

type Storage interface {
	StorageSetter
	StorageGetter
	StorageDestroyer
}

func NewRecord() *Record {
	return &Record{Data: make(map[string]string, 2)}
}

func (r *Record) Store(key, val string) {
	r.Data[key] = val
}

func (r *Record) FindAllRecords() map[string]string {
	return r.Data
}

func (r *Record) FindRecord(key string) (string, error) {
	val, ok := r.Data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return val, nil
}

func (r *Record) DeleteRecordByKey(key string) error {
	_, ok := r.Data[key]
	if !ok {
		return errors.New("key not found")
	}
	delete(r.Data, key)
	return nil
}
