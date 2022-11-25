package storage

import (
	"fmt"
	"os"
	"strconv"
	"sync"
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
	// arrange
	number := []byte("123456789")
	storage := NewStorage("tmp.log")
	defer teardown(storage, t)

	w := sync.WaitGroup{}
	w.Add(20000)
	for i := 1; i <= 20000; i++ {
		go func(number []byte, w *sync.WaitGroup) {
			err := storage.PersistNumber([]byte(strconv.Itoa(i)))
			fmt.Printf("this %v :\n %v", i, err)
			w.Done()
		}(number, &w)
	}

	w.Wait()

}

func teardown(s *Storage, t *testing.T) {
	err := os.Remove(s.File.Name())
	if err != nil {
		t.Error("failed to remove temp storage file")
	}
}
