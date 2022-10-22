package storage

import (
	"os"
	"sync"
)

type Storage struct {
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

func (s *Storage) PersistNumber(err error, buf []byte, w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	{
		if _, err = s.File.WriteString(string(buf) + "\n"); err != nil {
			panic(err)
		}
	}
	m.Unlock()
}
