package models

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Key value strore struct
type InMemory struct {
	Data map[string]string `json:"data"`
}

type StorageSetter interface {
	Store(string, string)
}

type StorageGetter interface {
	List() map[string]string
	Get(string) (string, error)
}

type StorageDestroyer interface {
	Delete(string) error
}

type Storage interface {
	StorageSetter
	StorageGetter
	StorageDestroyer
}

func NewCache() *InMemory {
	return &InMemory{Data: make(map[string]string, 2)}
}

func (r *InMemory) Store(key, val string) {
	r.Data[key] = val
}

func (r *InMemory) List() map[string]string {
	return r.Data
}

func (r *InMemory) Get(key string) (string, error) {
	val, ok := r.Data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return val, nil
}

func (r *InMemory) Delete(key string) error {
	_, ok := r.Data[key]
	if !ok {
		return errors.New("key not found")
	}
	delete(r.Data, key)
	return nil
}

type Disk struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func NewDisk() *Disk {
	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		err = os.Mkdir("storage", os.ModePerm)
		if err != nil {
			log.Fatalf("Create directory: %v", err)
		}
	}
	return &Disk{}
}

func (d *Disk) Store(key, val string) {
	fileName := filepath.Join("storage", key)

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(val))
	if err != nil {
		log.Fatalf("Writing file: %v", err)
	}
}

func (d *Disk) List() map[string]string {
	m := make(map[string]string, 2)
	err := filepath.Walk("storage", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Travarse storage directory: %v", err)
		}

		if !info.IsDir() {
			key := strings.Split(path, "/")[1]
			val, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("%v Read file err: %v", path, err)
			}
			m[key] = string(val)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return m
}

func (d *Disk) Get(key string) (string, error) {
	path := filepath.Join("storage", key)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		val, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("%v Read file err: %v", path, err)
		}
		return string(val), nil
	}
	return "", errors.New("key not found")
}

func (d *Disk) Delete(key string) error {
	path := filepath.Join("storage", key)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := os.Remove(path)
		if err != nil {
			log.Fatalf("%v Delete file err: %v", path, err)
		}
		return nil
	}
	return errors.New("key not found")
}
