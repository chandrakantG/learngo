package main

import (
	"fmt"
	"time"
)

// Semaphore to limit concurrency
var sema = make(chan struct{}, 3)

func main() {
	userCount := 10
	done := make(chan struct{}, userCount)

	for i := 0; i < userCount; i++ {
		go Read(i, done)
	}

	for i := 0; i < userCount; i++ {
		<-done
	}

	fmt.Println("All goroutines complete.")
	close(done)
}

func Read(i int, done chan<- struct{}) {
	sema <- struct{}{}        // acquire semaphore
	defer func() { <-sema }() // release semaphore

	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second) // simulate work

	done <- struct{}{}
}
