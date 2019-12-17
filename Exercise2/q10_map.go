package main

import "fmt"

func increment(x int) int {
	return x + 1
}

func exaggerate(x string) string {
	return x + "!"
}

func mapInt(f func(n int) int, nums []int) []int {
	var results []int

	for _, num := range nums {
		results = append(results, f(num))
	}
	return results
}

func mapString(f func(n string) string, strings []string) []string {
	var results []string

	for _, str := range strings {
		results = append(results, f(str))
	}
	return results
}

func main() {
	fmt.Println(mapInt(increment, []int{0, 1, 2, 3, 4}))
	fmt.Println(mapString(exaggerate, []string{"hello", "what", "good job", "ayiiiii", "adios"}))
}
