package main

import (
	"fmt"
	"time"
)

func main() {
	userCount := 10
	ch := make(chan bool, 2)
	for i := range userCount {
		ch <- true
		go read(ch, i)
	}
	time.Sleep(time.Second)
}

func read(ch chan bool, i int) {
	fmt.Println("read: ", i)
	fmt.Println("ch", <-ch)
}
