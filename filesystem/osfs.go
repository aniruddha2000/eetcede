package filesystem

import (
	"os"
)

type OsFs struct {
	Fs
}

func NewOsFs() Fs {
	return &OsFs{}
}

func (OsFs) Create(name string) (File, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (OsFs) Open(name string) (File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return file, err
}

func (OsFs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (OsFs) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (OsFs) Remove(name string) error {
	return os.Remove(name)
}
