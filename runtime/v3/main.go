package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	userCount := 10
	ch := make(chan bool, 2)

	for i := range userCount {
		wg.Add(1)
		go Read(ch, i)
	}

	wg.Wait()
}

func Read(ch chan bool, i int) {
	defer wg.Done()

	ch <- true
	fmt.Printf("read: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}
