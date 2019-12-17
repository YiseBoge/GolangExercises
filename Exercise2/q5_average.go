package main

import "fmt"

func average(s []float64) float64 {
	sum := 0.0
	length := float64(len(s))

	for _, f := range s {
		sum += f
	}

	return sum / length
}

func main() {
	s := []float64{10.2, 10.4}
	fmt.Println(average(s))
}
