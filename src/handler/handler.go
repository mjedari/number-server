package handler

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/mjedari/number-server/src/cache"
	"gitlab.com/mjedari/number-server/src/reporter"
	"gitlab.com/mjedari/number-server/src/storage"
	"net"
	"regexp"
	"sync"
	"sync/atomic"
)

type IHandler interface {
	HandleRequest(cache cache.ICache, storage storage.IStorage, report reporter.IReporter, w *sync.WaitGroup)
	ConnectionLimitExceeded(maxConnection int) bool
}

type Handler struct {
	Conn net.Conn
	Live uint32
}

func NewHandler(conn net.Conn) IHandler {
	return &Handler{Conn: conn}
}

func (h *Handler) HandleRequest(cache cache.ICache, storage storage.IStorage, report reporter.IReporter, w *sync.WaitGroup) {
	atomic.AddUint32(&h.Live, 1)

	buf := make([]byte, 9)
	_, err := h.Conn.Read(buf)
	if err != nil {
		message := "internal error"
		logrus.Error("Error reading:", err.Error())
		h.terminate(w, message)
		return
	}

	if err := validateNumber(buf); err != nil {
		message := "Invalid number received!\n"
		h.terminate(w, message)
		return
	}
	status := cache.Write(buf)
	if !status {
		// can't store number into the cache map
		report.IncDuplicateNumber(1)
		message := fmt.Sprintf("Number %v is duplicated.\n", string(buf))
		h.terminate(w, message)
		return
	}

	report.IncUniqueNumber(1)
	err = storage.PersistNumber(buf)
	message := fmt.Sprintf("Number %v has been stored.\n", string(buf))
	if err != nil {
		message = err.Error()
	}

	h.terminate(w, message)
}

func (h *Handler) closeConnection(message string) error {
	_, err := h.Conn.Write([]byte(message))
	if err != nil {
		logrus.Errorln("failed in sending response", err)
		return err
	}
	if err = h.Conn.Close(); err != nil {
		logrus.Errorln("closing connection has problem: ", err)
		return err
	}

	return nil
}

func (h *Handler) terminate(w *sync.WaitGroup, message string) {
	err := h.closeConnection(message)
	if err != nil {
		fmt.Println("Got error in closing connection")
	}

	atomic.AddUint32(&h.Live, ^uint32(0))
	//w.Done()
}

func (h *Handler) ConnectionLimitExceeded(maxConnection int) bool {
	if h.Live <= uint32(maxConnection) {
		return false
	}

	message := fmt.Sprintf("`\nMaximum connection exceeded! %d liveConnections is in use Please try later...\n", maxConnection)
	h.closeConnection(message)

	return false
}

func validateNumber(number []byte) error {
	reg, _ := regexp.Compile(`\d{9}`)
	if ok := reg.Match(number); !ok {
		return errors.New("invalid number received")
	}
	return nil
}
