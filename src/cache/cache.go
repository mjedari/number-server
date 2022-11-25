package cache

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"gitlab.com/mjedari/number-server/src/storage"
	"sync"
)

type Cache struct {
	mu sync.Mutex

	List map[string]bool
}

func NewCache() *Cache {
	NumberCache = Cache{
		mu:   sync.Mutex{},
		List: map[string]bool{},
	}
	return &NumberCache
}

var NumberCache Cache

func (c *Cache) SyncWithStorage(storage *storage.Storage) error {
	// fetch all numbers
	// cache all number into memory
	scanner := bufio.NewScanner(storage.File)

	for scanner.Scan() {
		NumberCache.List[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		logrus.Error("cache package:", err)
	}
	return nil
}

func (c *Cache) Write(buf []byte) bool {
	status := true
	receivedNumber := string(buf)
	c.mu.Lock()
	{
		if ok := NumberCache.List[receivedNumber]; ok {
			status = false
		}

		NumberCache.List[receivedNumber] = true
	}
	c.mu.Unlock()
	return status
}
