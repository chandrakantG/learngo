package main

import "fmt"

type S1 struct{}
type S2 struct{}
type S3 struct{}
type S4 struct {
	M *int
}
type S5 struct {
	M *int
}
type S6 struct {
	M *int
}

func main() {
	var x S1
	_ = identity1(x)

	var p S2
	// y := &p
	_ = identity2(&p)

	var q S3
	_ = *ref3(q)

	var i int
	refStruct4(i)
	fmt.Println(i)

	var j int
	refStruct5(&j)

	var a S6
	var b int
	refStruct6(&b, &a)
}

// go run -gcflags '-m -l' main.go

func identity1(x S1) S1 {
	return x
}

func identity2(z *S2) *S2 {
	return z
}

// Escape analysis is how Go decides whether a variable is allocated on the stack or the heap.
// If a value “leaks” to the result or another function, it can't safely live only on the stack, so Go puts it on the heap.
// Heap allocations are slower and involve garbage collection, so it affects performance.

func ref3(z S3) *S3 {
	return &z
}

func refStruct4(y int) (z S4) {
	z.M = &y
	return z
}

func refStruct5(y *int) (z S5) {
	z.M = y
	return z
}

func refStruct6(y *int, z *S6) {
	z.M = y
}

// how and when a value "leaks" (escapes) in Go is key to writing efficient, allocation-aware code.
// In Go, a value "leaks" (aka escapes to the heap) when the compiler determines that it must outlive the current function's stack frame — usually because it's returned, captured, or shared.
// If a value is used outside the function where it's created, it leaks
