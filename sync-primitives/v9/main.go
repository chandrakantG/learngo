package main

import "fmt"

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func main() {
	go f()
	c <- 0
	fmt.Println(a)
}

/**
A: Cannot compile
B: Output 1 => goroutine assign value 1 to a
C: Output 0
D: panic
*/
