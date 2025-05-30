package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	chain string
)

func main() {
	chain = "chain"
	A()
	fmt.Println(chain)
}

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

/**
A: Can not compile => can not call lock inside lock(deadlock)
B: output main --> A --> B --> C
C: output main
D: panic
*/
