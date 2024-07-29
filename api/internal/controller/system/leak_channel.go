package system

import (
	"fmt"
	"runtime"
	"time"
)

func leakChannel() {
	ch := make(chan []byte, 100)
	go func() {
		for {
			ch <- make([]byte, 1024*1024*10) // Allocate 10 MB
			time.Sleep(100 * time.Millisecond)
		}
	}()
	// The channel is never closed, leading to a leak
	_ = ch
	fmt.Println("Created leaking channel")
}

func (i impl) MonitorChannel() {
	runtime.GC() // Run garbage collector to clean up unused memory

	for i := 0; i < 5; i++ {
		leakChannel()
		time.Sleep(500 * time.Millisecond)
	}

	for {
		time.Sleep(5 * time.Second)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
			m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
	}
}
