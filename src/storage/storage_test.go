package storage

import (
	"os"
	"testing"
)

func TestNewStorage(t *testing.T) {
	storage := NewStorage("tmp.log")
	defer teardown(storage, t)

	if storage == nil {
		t.Error("failed to construct an storage instance")
	}
	// how to clean up the file that I've created?
}

func TestStorage_PersistNumber(t *testing.T) {
	// arrange
	number := []byte("123456789")
	storage := NewStorage("tmp.log")
	defer teardown(storage, t)

	// act
	err := storage.PersistNumber(number)

	// assert
	if err != nil {
		t.Error("failed to persist number to storage")
	}
}

func TestStorage_ConcurrentPersistNumber(t *testing.T) {
	//
}

func teardown(s IStorage, t *testing.T) {
	err := os.Remove(s.GetFile().Name())
	if err != nil {
		t.Error("failed to remove temp storage file")
	}
}
