package filesystem

import (
	"os"
)

func DirExists(fs Fs, name string) (bool, error) {
	file, err := fs.Stat(name)
	if err == nil && file.IsDir() {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Exists(fs Fs, name string) (bool, error) {
	_, err := fs.Stat(name)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadDir(fs Fs, dirName string) ([]os.FileInfo, error) {
	dir, err := fs.Open(dirName)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	list, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func ReadFile(fs Fs, name string) ([]byte, error) {
	file, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}
