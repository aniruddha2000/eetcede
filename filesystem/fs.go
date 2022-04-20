package filesystem

import (
	"io"
	"os"
)

type FileSystem struct {
	Fs
}

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
	Sync() error
	WriteString(s string) (ret int, err error)
}

type Fs interface {
	Create(name string) (File, error)
	Mkdir(name string, perm os.FileMode) error
	Open(name string) (File, error)
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	Stat(name string) (os.FileInfo, error)
	Remove(name string) error
}
