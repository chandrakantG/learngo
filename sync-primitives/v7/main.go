package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int // nil channel,do not use this syntax
	// ch := make(chan int) // not nil channel,use this syntax
	go func() {
		fmt.Println("go routine:1 start")
		ch = make(chan int, 1)
		ch <- 1
		fmt.Println("go routine:1 end")
	}()
	go func(ch chan int) {
		fmt.Println("go routine:2 start")
		// time.Sleep(time.Second)
		<-ch
		fmt.Println("go routine:2 end")
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

/**
A: Cannot compile
B: After a while, it always outputs #goroutines: 1
C: After a while, it always outputs #goroutines: 2 => second goroutine hang because of nil channel at goroutine start running
D: panic
*/
