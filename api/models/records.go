package models

import (
	"errors"
	"log"
	"os"

	"github.com/aniruddha2000/goEtcd/filesystem"
	"github.com/spf13/afero"
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

// Return In Memory struct
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
	FS afero.Fs
}

type DiskFS struct {
	FS filesystem.Fs
}

// Return the Disk structure file system
func NewDisk() *DiskFS {
	diskFs := filesystem.NewOsFs()
	ok, err := filesystem.DirExists(diskFs, "storage")
	// afero.DirExists()
	if err != nil {
		log.Fatalf("Dir exists: %v", err)
	}
	if !ok {
		err := diskFs.Mkdir("storage", os.ModePerm)
		if err != nil {
			log.Fatalf("Create dir: %v", err)
		}
	}
	return &DiskFS{FS: diskFs}
}

// Store key, value in the file system
func (d *DiskFS) Store(key, val string) {
	file, err := d.FS.Create("storage/" + key)
	if err != nil {
		log.Fatalf("Create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(val))
	if err != nil {
		log.Fatalf("Writing file: %v", err)
	}
}

func (d *DiskFS) List() map[string]string {
	m := make(map[string]string, 2)
	dir, err := filesystem.ReadDir(d.FS, "storage")
	if err != nil {
		log.Fatalf("Error reading the directory: %v", err)
	}

	for _, fileName := range dir {
		content, err := filesystem.ReadFile(d.FS, "storage/"+fileName.Name())
		if err != nil {
			log.Fatalf("Error reading the file: %v", err)
		}
		m[fileName.Name()] = string(content)
	}
	return m
}

func (d *DiskFS) Get(key string) (string, error) {
	ok, err := filesystem.Exists(d.FS, "storage/"+key)
	if err != nil {
		log.Fatalf("File exist: %v", err)
	}

	if ok {
		file, err := filesystem.ReadFile(d.FS, "storage/"+key)
		if err != nil {
			log.Fatalf("Error reading the file: %v", err)
		}
		return string(file), nil
	}
	return "", errors.New("key not found")
}

func (d *DiskFS) Delete(key string) error {
	ok, err := filesystem.Exists(d.FS, "storage/"+key)
	if err != nil {
		log.Fatalf("File exist: %v", err)
	}
	if ok {
		err = d.FS.Remove("storage/" + key)
		if err != nil {
			log.Fatalf("Delete file err: %v", err)
		}
		return nil
	}
	return errors.New("key not found")
}
