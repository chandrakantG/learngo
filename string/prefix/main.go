package main

import "fmt"

func main() {
	fmt.Println(HasPrefix("hello", "hello"))
	fmt.Println(HasPrefix("hello", "hi"))
	fmt.Println(HasSuffix("hello", "llo"))
	fmt.Println(HasSuffix("hello", "lo"))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
