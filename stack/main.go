package main

import "fmt"

type MyStack struct {
	data []int
}

func main() {
	s := new(MyStack)
	fmt.Println(s)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func (s *MyStack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *MyStack) Pop() int {
	if len(s.data) <= 0 {
		fmt.Println("Stack is Empty")
		return -1
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}
