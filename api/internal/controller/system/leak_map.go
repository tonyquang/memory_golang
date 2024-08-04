package system

import (
	"log"
	"time"
)

func leakMap() {
	m := make(map[int][]byte)
	for i := 0; i < 30; i++ {
		m[i] = make([]byte, 1024*1024) // Allocate 1 MB
	}
	_ = m
	log.Println("Created large map")
}

func (i impl) MonitorMap() {
	leakMap()
	time.Sleep(1 * time.Second)
}
