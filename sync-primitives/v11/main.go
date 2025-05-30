package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var ints = make([]int, 0, 1000)
	// ch := make(chan bool)
	go func() {
		for i := range 1000 {
			ints = append(ints, i)
		}
		wg.Done()
		// ch <- true
	}()
	go func() {
		// <-ch
		for i := range 1000 {
			ints = append(ints, i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(len(ints))
}

// data race occur so use channel to sync slice inserctions
// use  go run -race main.go
