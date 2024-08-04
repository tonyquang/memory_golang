package system

import (
	"log"
	"time"
)

func leakChannel() {
	ch := make(chan []byte, 100)
	go func() {
		ch <- make([]byte, 1024*1024*10) // Allocate 10 MB
		time.Sleep(100 * time.Millisecond)
	}()
	// The channel is never closed, leading to a leak
	_ = ch
	log.Println("Created leaking channel")
}

func (i impl) MonitorChannel() {
	for i := 0; i < 3; i++ {
		leakChannel()
		time.Sleep(500 * time.Millisecond)
	}
}
