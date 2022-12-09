package main

import (
	"fmt"
	"gitlab.com/mjedari/number-server/configs"
	"gitlab.com/mjedari/number-server/src/cache"
	"gitlab.com/mjedari/number-server/src/handler"
	"gitlab.com/mjedari/number-server/src/reporter"
	"gitlab.com/mjedari/number-server/src/storage"
	"net"
	"os"
	"sync"
)

func initConfig() (*configs.Configuration, int) {
	config := configs.GetConfigs()
	maxConnection := config.Connection.Max
	return config, maxConnection
}

func main() {
	// initiate configuration
	config, maxConnection := initConfig()

	// initiate cache
	cache := cache.NewCache()

	// listening on tcp
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Server is listening on %v:%v \n", config.Server.Host, config.Server.Port)

	reporter := reporter.NewReporter()

	storage := storage.NewStorage(config.Storage.Path)
	defer storage.GetFile().Close()

	if err = cache.SyncWithStorage(storage); err != nil {
		panic("can't sync cache with storage")
	}

	go reporter.Report()

	w := sync.WaitGroup{}
	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		handler := handler.NewHandler(c)
		if handler.ConnectionLimitExceeded(maxConnection) {
			continue
		}

		//w.Add(1)
		go handler.HandleRequest(cache, storage, reporter, &w)
	}
	//w.Wait()
}
