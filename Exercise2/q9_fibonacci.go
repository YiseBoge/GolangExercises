package main

import "fmt"

func fibonacci(n int) []int {
	if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	}

	res := []int{0, 1}

	for i := 1; i < n; i++ {
		length := len(res)
		res = append(res, res[length-1]+res[length-2])
	}

	return res
}
func main() {
	fmt.Println(fibonacci(8))
}
