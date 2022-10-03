package main

import (
	"bufio"
	"errors"
	"fmt"
	"sync"
)

type Cache map[string]bool

func (c *Cache) syncCacheWithStorage(storage *Storage) error {
	// fetch all numbers
	// cache all number into memory
	scanner := bufio.NewScanner(storage.file)

	for scanner.Scan() {
		numberCache[scanner.Text()] = true
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err)
	}
	return nil
}

func (c *Cache) write(buf []byte, w *sync.WaitGroup, m *sync.Mutex) error {
	receivedNumber := string(buf)
	m.Lock()
	{
		if err := checkNumber(receivedNumber); err != nil {
			return err
		}

		numberCache[receivedNumber] = true
	}
	m.Unlock()
	return nil
}

func checkNumber(key string) error {
	if ok := numberCache[key]; ok {
		return errors.New("duplicated entity")
	}
	return nil
}
