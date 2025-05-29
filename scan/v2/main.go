package main

import "fmt"

func main() {
	var sx, sy, ex, ey int
	fmt.Println("enter 4 values:")
	_, err := fmt.Scanf("%d %d %d %d", &sx, &sy, &ex, &ey)
	if err != nil {
		fmt.Println(err)
	}
	ans := max(abs(sx-ex), abs(sy-ey))

	fmt.Println(ans)

}

// max returns the larger of two integers.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// return absolute value
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
