package storage

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type Storage struct {
	mu sync.Mutex

	File *os.File
}

func NewStorage(name string) *Storage {

	file, err := createOrOpenFile(name)
	if err != nil {
		panic("error opening storage file")
	}

	return &Storage{File: file}
}

func createOrOpenFile(fileName string) (*os.File, error) {
	if _, err := os.Stat(fileName); err != nil {
		return os.Create(fileName)
	}

	return os.OpenFile(fileName, os.O_RDWR, 0644)
}

func (s *Storage) PersistNumber(buf []byte) error {
	var status error
	s.mu.Lock()
	{
		if _, err := s.File.WriteString(fmt.Sprintln(string(buf))); err != nil {
			status = errors.New("can not store number")
		}
	}
	s.mu.Unlock()
	return status
}
