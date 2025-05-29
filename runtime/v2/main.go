package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	userCount := 10

	for i := range userCount {
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
}

func read(i int) {
	defer wg.Done()
	fmt.Println("read:", i)
}
