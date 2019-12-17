package main

import "fmt"

func order(x int, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

func main() {
	fmt.Println(order(4, 3))
}
