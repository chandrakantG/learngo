package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) }}

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 10; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; i < 15; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}

/**
A: Cannot compile
B: Can compile, runs normally, and memory is stable
C: Can compile, memory may surge during runtime => answer
D: Can compile, memory surges during runtime, but will be recovered after a while
*/
