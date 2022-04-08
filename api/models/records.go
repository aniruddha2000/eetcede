package models

import (
	"errors"
	"log"
	"os"
	"path/filepath"

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
	FS  afero.Fs
	Key string `json:"key"`
	Val string `json:"val"`
}

// type OpenFileFS interface {
// 	fs.FS
// 	OpenFile(name string, flag int, perm os.FileMode) (fs.File, error)
// }

// func OpenFile(fsys fs.FS, name string, flag int, perm os.FileMode) (fs.File, error) {
// 	if fsys, ok := fsys.(OpenFileFS); ok {
// 		return fsys.OpenFile(name, flag, perm)
// 	}
// 	if flag == os.O_RDONLY {
// 		return fsys.Open(name)
// 	}
// 	return nil, fmt.Errorf("open %s: Operation not supported", name)
// }

// func Create(fsys fs.FS, name string) (fs.File, error) {
// 	return OpenFile(fsys, name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
// }

func NewDisk() *Disk {
	diskFs := afero.NewOsFs()
	ok, err := afero.DirExists(diskFs, "storage")
	if err != nil {
		log.Fatalf("Dir exists: %v", err)
	}
	if !ok {
		err := diskFs.Mkdir("storage", os.ModePerm)
		if err != nil {
			log.Fatalf("Create dir: %v", err)
		}
	}
	return &Disk{FS: diskFs}
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

	// file, err := d.FS.Open(key)
	// if err != nil {
	// 	log.Fatalf("Couldn't open file in filesystem : %v", err)
	// }
	// file.Read()
}

func (d *Disk) List() map[string]string {
	m := make(map[string]string, 2)
	dir, err := afero.ReadDir(d.FS, "storage")
	if err != nil {
		log.Fatalf("Error reading the directory: %v", err)
	}

	for _, fileName := range dir {
		content, err := afero.ReadFile(d.FS, "storage/"+fileName.Name())
		if err != nil {
			log.Fatalf("Error reading the file: %v", err)
		}
		m[fileName.Name()] = string(content)
	}
	return m
}

func (d *Disk) Get(key string) (string, error) {
	_, err := d.FS.Open("storage/" + key)
	if err != nil {
		return "", err
	}

	file, err := afero.ReadFile(d.FS, "storage/"+key)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}
	return string(file), nil
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
