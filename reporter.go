package main

import (
	"fmt"
	"time"
)

type Reporter struct {
	Ticker           *time.Ticker
	quit             chan struct{}
	UniqueNumbers    uint
	DuplicateNumbers uint
}

func NewReporter() *Reporter {
	ticker := time.NewTicker(ReporterInterval)
	quit := make(chan struct{})
	return &Reporter{Ticker: ticker, quit: quit}
}

func (r *Reporter) Report() {
	for {
		select {
		case <-r.Ticker.C:
			fmt.Printf("Received %d unique number, %d duplicate. Unique total: %d\n", r.UniqueNumbers, r.DuplicateNumbers, len(numberCache))
			r.resetNumbers()
		case <-r.quit:
			r.Ticker.Stop()
			return
		}
	}
}

func (r *Reporter) IncUniqueNumber(num uint) {
	r.UniqueNumbers += num
}

func (r *Reporter) IncDuplicateNumber(num uint) {
	r.DuplicateNumbers += num
}

func (r *Reporter) resetNumbers() {
	r.UniqueNumbers = 0
	r.DuplicateNumbers = 0
}
