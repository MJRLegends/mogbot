package memory

import (
	"log"
	"runtime"
	"time"
)

func Ticker() {
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				PrintMemUsage()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	log.Printf(
		"Alloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v",
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys),
		m.NumGC,
	)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
