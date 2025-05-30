package main

import "fmt"

func main() {
	fmt.Println(ReverseStr("abc"))
	fmt.Println(ReverseStr("cba"))
}

func ReverseStr(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
