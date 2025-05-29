package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println("Initial:", people)
	sort.Sort(ByAge(people))
	fmt.Println("Sorted:", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("Reversed:", people)

	s := []int{5, 2, 7, 3, 1}
	sort.Ints(s)
	fmt.Println("s:", s)
	fmt.Println(sort.SearchInts(s, 5))
}
