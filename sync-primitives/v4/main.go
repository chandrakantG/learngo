package main

import (
	"fmt"
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 { // double check not necessary
		o.done = 1
		f()
	}
}

func main() {
	o := &Once{}
	o.Do(f)
	o.Do(f)
	o.Do(f)
	o.Do(f)

	// a := &Once{}
	// a.Do(f)
	// a.Do(f)
	// a.Do(f)
}

func f() {
	fmt.Println("func f()")
}

/**
A: Cannot compile
B: Can compile, singleton is implemented correctly => this is answer
C: Can compile, but there are concurrency issues, function f may be executed multiple times
D: Can compile, but the program will panic
*/
