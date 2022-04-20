package filesystem

import "os"

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
