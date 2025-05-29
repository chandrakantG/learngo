package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter no:")
	scanner.Scan()
	number := scanner.Text()
	fmt.Println("number:", number)

	fmt.Println("please enter string with space:")
	scanner.Scan()
	arr := strings.Split(strings.Trim(scanner.Text(), " "), " ")

	fmt.Println(arr)
	fmt.Println("please enter 5 numbers:")
	var n, sx, sy, ex, ey int
	_, err := fmt.Scanf("%d %d %d %d %d", &n, &sx, &sy, &ex, &ey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, sx, sy, ex, ey)
}
