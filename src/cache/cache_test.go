package cache

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/mjedari/number-server/src/storage"
	"os"
	"testing"
)

func TestNewCache(t *testing.T) {
	// How to test?
	// Do we need to test constructions?
	cache := NewCache()
	if cache == nil {
		t.Error("creating a cache instance has been failed ")
	}
}

func TestCache_SyncWithStorage(t *testing.T) {
	// arrange
	cache := NewCache()
	storage := storage.NewStorage("numbers.log")
	defer teardown(storage, t)

	// act
	err := cache.SyncWithStorage(storage)

	// assert: it works
	if err != nil {
		logrus.Error("failed to sync with storage")
	}
}

func TestCache_Write(t *testing.T) {
	// arrange
	cache := NewCache()
	number := "123456"
	buf := []byte(number)

	// act
	cache.Write(buf)

	// assert
	ok := cache.GetList()[number]
	if !ok {
		logrus.Errorf("fail to store number %v", number)
	}
}

func teardown(s storage.IStorage, t *testing.T) {
	err := os.Remove(s.GetFile().Name())
	if err != nil {
		t.Error("failed to remove temp storage file")
	}
}
