package system

import (
	"fmt"
	"runtime"
	"time"
)

func leakMap() {
	m := make(map[int][]byte)
	for i := 0; i < 1000000; i++ {
		m[i] = make([]byte, 1024*1024) // Allocate 1 MB
	}
	_ = m
	fmt.Println("Created large map")
}

func (i impl) MonitorMap() {
	runtime.GC() // Run garbage collector to clean up unused memory

	for {
		leakMap()
		time.Sleep(1 * time.Second)
		if time.Now().Second()%5 == 0 {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
				m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
		}
	}
}
