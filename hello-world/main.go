package main

import (
	"fmt"
	"os"
)

func main() {
	// assign a new value to the string variable below
	name := os.Args[1]
	fmt.Println("Hello great", name, "!")

	// changes the name, declares the age with 85
	name, age := "gandalf", 85

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}

// go run hello-world/main.go chandrakant
