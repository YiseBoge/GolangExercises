package main

import "fmt"

func one(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	one(3, 4, 5, 6, 6, 1)
}
