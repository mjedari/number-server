package main

import (
	"errors"
	"fmt"
	"gitlab.com/mjedari/number-server/cache"
	"gitlab.com/mjedari/number-server/reporter"
	"gitlab.com/mjedari/number-server/storage"
	"net"
	"os"
	"regexp"
	"sync"
)

const (
	MaxConnection = 5
)

var liveConnections = 0

type Connection struct {
	Node net.Conn
}

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	reporter := reporter.NewReporter()

	fmt.Println("Server started listening ...")
	cache.NumberCache = make(cache.Cache, 2097152)

	storage := storage.NewStorage("numbers.log")
	defer storage.File.Close()

	if err = cache.NumberCache.SyncCacheWithStorage(storage); err != nil {
		panic("can't sync cache with storage")
	}

	go reporter.Report()

	m := sync.Mutex{}
	w := sync.WaitGroup{}

	for {
		fmt.Println("start fo loop")
		c, err := listener.Accept()
		fmt.Println("start listening")
		conn := Connection{Node: c}
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		if liveConnections >= MaxConnection {
			conn.Node.Write([]byte(fmt.Sprintf("`\nMaximum connection exceeded! %d liveConnections is in use Please try later...\n", MaxConnection)))
			conn.Node.Close()
			continue
		}

		w.Add(1)
		go conn.handleRequest(storage, reporter, &w, &m)
		w.Wait()
	}
}

func (c *Connection) handleRequest(storage *storage.Storage, report *reporter.Reporter, w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	liveConnections += 1
	m.Unlock()

	buf := make([]byte, 9)
	_, err := c.Node.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		c.terminateHandler(w, m)
		return
	}

	if err := validateNumber(buf); err != nil {
		c.Node.Write([]byte("Invalid number received!\n"))
		c.terminateHandler(w, m)
		return
	}

	if err := cache.NumberCache.Write(buf, w, m); err != nil {
		// can't store number into the cache map
		report.IncDuplicateNumber(1)
		c.Node.Write([]byte(fmt.Sprintf("Number %v is duplicated.\n", string(buf))))
		c.terminateHandler(w, m)
		return
	}

	report.IncUniqueNumber(1)

	go storage.PersistNumber(err, buf, w, m)
	c.Node.Write([]byte(fmt.Sprintf("Number %v has been stored.\n", string(buf))))

	c.terminateHandler(w, m)
}

func (c *Connection) terminateHandler(w *sync.WaitGroup, m *sync.Mutex) {
	if err := c.Node.Close(); err != nil {
		fmt.Println("closing connection has problem: ", err)

	}
	fmt.Println("connection closed")
	m.Lock()
	liveConnections -= 1
	m.Unlock()
	w.Done()
}

func validateNumber(number []byte) error {
	reg, _ := regexp.Compile(`\d{9}`)
	if ok := reg.Match(number); !ok {
		return errors.New("invalid number received")
	}
	return nil
}
