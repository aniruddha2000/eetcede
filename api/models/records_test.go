package models

import (
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestInMemory(t *testing.T) {
	t.Run("Store", func(t *testing.T) {
		cache := NewCache()
		cache.Store("Aniruddha", "Basak")

		if _, ok := cache.Data["Aniruddha"]; !ok {
			t.Errorf("func -> %v | Data not being inserted", t.Name())
		}
	})

	t.Run("List", func(t *testing.T) {
		cache := NewCache()
		cache.Store("Aniruddha", "Basak")
		cache.Store("Golang", "Etcd")

		want := map[string]string{
			"Aniruddha": "Basak",
			"Golang":    "Etcd",
		}
		got := cache.List()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("func -> %v | Data not Listed properly", t.Name())
		}
	})

	t.Run("Get", func(t *testing.T) {
		cache := NewCache()
		cache.Store("Aniruddha", "Basak")

		_, err := cache.Get("Golang")
		if err == nil {
			t.Errorf("func -> %v | Should get key not found error", t.Name())
		}

		got, _ := cache.Get("Aniruddha")
		if !reflect.DeepEqual(got, "Basak") {
			t.Errorf("func -> %v | Value not getting", t.Name())
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cache := NewCache()
		cache.Store("Aniruddha", "Basak")

		cache.Delete("Aniruddha")

		_, err := cache.Get("Aniruddha")
		if err == nil {
			t.Errorf("func -> %v | Should get key not found error", t.Name())
		}
	})
}

func TestDisk(t *testing.T) {
	t.Run("Store", func(t *testing.T) {
		disk := GetMemDisk()
		key, val := "file", "system"
		disk.Store(key, val)

		if ok, _ := afero.Exists(disk.FS, key); !ok {
			t.Error("Failed to create file")
		}

		file, err := afero.ReadFile(disk.FS, key)
		if err != nil {
			t.Errorf("Couldn't read file : %v", err)
		}

		if !reflect.DeepEqual(string(file), val) {
			t.Error("Couldn't write into the file")
		}
	})

	t.Run("List", func(t *testing.T) {
		disk := GetMemDisk()
		want := map[string]string{
			"file": "system",
			"data": "structure",
		}

		for k, v := range want {
			disk.Store(k, v)
		}

		got := disk.List()

		if !reflect.DeepEqual(want, got) {
			t.Error("Data not Listed properly")
		}
	})

	t.Run("Get", func(t *testing.T) {
		key, val := "file", "system"
		disk := GetMemDisk()
		disk.Store(key, val)

		_, err := disk.Get("Golang")
		if err == nil {
			t.Error("Should get key not found error")
		}

		got, _ := disk.Get(key)

		if !reflect.DeepEqual(val, got) {
			t.Error("not getting the value")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		key, val := "file", "system"
		disk := GetMemDisk()
		disk.Store(key, val)

		disk.Delete(key)

		_, err := disk.Get(key)
		if err == nil {
			t.Error("File should have been deleted")
		}
	})
}

func GetMemDisk() *Disk {
	test_fs := afero.NewMemMapFs()
	return &Disk{FS: test_fs}
}
