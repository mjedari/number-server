package reporter

import (
	"fmt"
	"gitlab.com/mjedari/number-server/src/cache"
	"time"
)

const ReportInterval = 10 * time.Second

type Reporter struct {
	Ticker           *time.Ticker
	quit             chan struct{}
	UniqueNumbers    uint
	DuplicateNumbers uint
}

func NewReporter() *Reporter {
	ticker := time.NewTicker(ReportInterval)
	quit := make(chan struct{})
	return &Reporter{Ticker: ticker, quit: quit}
}

func (r *Reporter) Report() {
	fmt.Println("Reporting service has just started...")
	for {
		select {
		case <-r.Ticker.C:
			fmt.Printf("Received %d unique number, %d duplicate. Unique total: %d\n", r.UniqueNumbers, r.DuplicateNumbers, len(cache.NumberCache.List))
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
