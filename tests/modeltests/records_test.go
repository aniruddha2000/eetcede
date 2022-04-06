package modeltests

import (
	"reflect"
	"testing"

	"github.com/aniruddha2000/goEtcd/api/models"
)

func TestInMemory(t *testing.T) {
	t.Run("Store", func(t *testing.T) {
		cache := models.NewCache()
		cache.Store("Aniruddha", "Basak")

		if _, ok := cache.Data["Aniruddha"]; !ok {
			t.Errorf("func -> %v | Data not being inserted", t.Name())
		}
	})

	t.Run("List", func(t *testing.T) {
		cache := models.NewCache()
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
		cache := models.NewCache()
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
		cache := models.NewCache()
		cache.Store("Aniruddha", "Basak")

		cache.Delete("Aniruddha")

		_, err := cache.Get("Aniruddha")
		if err == nil {
			t.Errorf("func -> %v | Should get key not found error", t.Name())
		}
	})
}
