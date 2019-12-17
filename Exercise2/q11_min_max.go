package main

import "fmt"

func min(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(min(x))
	fmt.Println(max(x))
}
