package main

import "fmt"

func main() {
	s := []float64{10.2, 10.4}

	sum := 0.0
	length := float64(len(s))

	for _, f := range s {
		sum += f
	}

	fmt.Println(sum / length)
}
