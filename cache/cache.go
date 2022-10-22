package cache

import (
	"bufio"
	"errors"
	"fmt"
	"gitlab.com/mjedari/number-server/storage"
	"sync"
)

type Cache map[string]bool

// It's better to use this approach
//type Cache struct {
//	mu sync.Mutex
//
//	data map[string]bool
//}

var NumberCache Cache

func (c *Cache) SyncCacheWithStorage(storage *storage.Storage) error {
	// fetch all numbers
	// cache all number into memory
	scanner := bufio.NewScanner(storage.File)

	for scanner.Scan() {
		NumberCache[scanner.Text()] = true
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err)
	}
	return nil
}

func (c *Cache) Write(buf []byte, w *sync.WaitGroup, m *sync.Mutex) error {
	receivedNumber := string(buf)
	m.Lock()
	{
		if err := checkNumber(receivedNumber); err != nil {
			return err
		}

		NumberCache[receivedNumber] = true
	}
	m.Unlock()
	return nil
}

func checkNumber(key string) error {
	if ok := NumberCache[key]; ok {
		return errors.New("duplicated entity")
	}
	return nil
}
