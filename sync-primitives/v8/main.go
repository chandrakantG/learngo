package main

import "fmt"

func main() {
	var ch chan int // nil channel,do not use this syntax
	// ch := make(chan int) // not nil channel,use this syntax
	var count int
	go func() {
		fmt.Println("goroutine:1")
		ch <- 1
	}()
	go func() {
		fmt.Println("goroutine:2")
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}

/**
A: Cannot compile
B: Output 1
C: Output 0
D: panic => can not close nil channel
*/
