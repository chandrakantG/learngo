package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	// wg.Done()
	wg.Wait()
}

/**
A: can not compile => WaitGroup must not be modified (via Add) while Wait() is still blocking in another goroutine.
B: no output, exit normally
C: program hangs
D: panic
*/

// A WaitGroup is not safe to reuse until the previous Wait() has returned. Go’s sync.WaitGroup enforces this to prevent race conditions and logic errors.
