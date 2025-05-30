package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu    sync.RWMutex
	count int
)

func main() {
	go A()
	time.Sleep(1 * time.Millisecond)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	fmt.Println("func A")
	B()
}

func B() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("func B")
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("func C")
}

/**
A: Can not compile
B: output 1 => can call read lock inside read lock or can call read lock inside lock
C: program hangs
D: panic
*/
